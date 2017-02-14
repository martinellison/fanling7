// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
package filepersist

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"persist"

	"github.com/lanior/panic"
	"github.com/y0ssar1an/q"
	yaml "gopkg.in/yaml.v2"
)

const pageSuffix = ".page"

type FilePersistor struct{ dir string }

func (p *FilePersistor) Open(path string) { q.Q("opening file persistor", p.dir); p.dir = path }

func (p *FilePersistor) Close() {}
func (p *FilePersistor) GetPageList() (paths []string) {
	lenDir := len(p.dir) + 1
	lenSuffix := len(pageSuffix)
	q.Q("searching directory", p.dir)
	paths = make([]string, 0)
	panic.PanicIf(filepath.Walk(p.dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			suffixPos := len(path) - lenSuffix
			if !info.IsDir() && path[0] != '.' && path[suffixPos:] == pageSuffix {
				paths = append(paths, path[lenDir:suffixPos])
			}
			return nil
		}))
	return
}

var separator []byte

func init() {
	separator = []byte("--- # \n")
}

func (p *FilePersistor) ReadPage(ident string) (meta persist.PageMeta, detail []byte, found bool) {
	path := p.pagePath(ident)
	meta, detail, found = p.ReadPageFromPath(path)
	return
}
func (p *FilePersistor) ReadPageFromPath(path string) (meta persist.PageMeta, detail []byte, found bool) {
	pageFile, err := os.Open(path)
	if err != nil && os.IsNotExist(err) {
		return
	}
	panic.PanicIf(err)
	found = true
	pageText, err := ioutil.ReadAll(pageFile)
	panic.PanicIf(err)
	docs := bytes.SplitN(pageText, separator, 3)
	panic.PanicIf(len(docs) < 3, "cannot split input from %s (only %d parts): %s", path, len(docs), pageText)
	panic.PanicIf(yaml.Unmarshal(docs[1], &meta))
	detail = docs[2]
	return
}
func (p *FilePersistor) WritePage(meta persist.PageMeta, detail []byte) {
	pageText := separator
	header, err := yaml.Marshal(meta)
	panic.PanicIf(err)
	pageText = append(pageText, header...)
	pageText = append(pageText, separator...)
	pageText = append(pageText, detail...)
	path := p.pagePath(meta.Ident)
	panic.PanicIf(ioutil.WriteFile(path, pageText, 0644))
	//shouldPath := e.pagePath(newPage.Ident())
	//if path != shouldPath {
	//fmt.Printf("git mv \"%s\" \"%s\"\n", path, shouldPath)
	//}
}

// `pagePath` returns the file path for a page.
func (p *FilePersistor) pagePath(ident string) string { return filepath.Join(p.dir, ident+pageSuffix) }

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
