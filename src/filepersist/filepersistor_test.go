// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
package filepersist

import (
	"os"
	"persist"
	"testing"

	"github.com/stretchr/testify/assert"
)

// `TestFilePersist1` tests
func TestFilePersist1(t *testing.T) {
	assert := assert.New(t)
	path := os.TempDir() + "/fp"
	os.RemoveAll(path)
	os.MkdirAll(path, 0777)
	var fp FilePersistor
	fp.Open(path)
	l0 := fp.GetPageList()
	assert.Equal(0, len(l0))
	fp.Close()
}

// `TestFilePersist2` tests
func TestFilePersist2(t *testing.T) {
	assert := assert.New(t)
	path := os.TempDir() + "/fp"
	os.RemoveAll(path)
	os.MkdirAll(path, 0777)
	var fp FilePersistor
	fp.Open(path)
	fp.WritePage(persist.PageMeta{Type: "fred", Ident: "nurke"}, []byte("abc"))
	l1 := fp.GetPageList()
	assert.Equal(1, len(l1))
	assert.Equal("nurke", l1[0])
	meta, detail, found := fp.ReadPage("nurke")
	assert.True(found)
	assert.Equal("nurke", meta.Ident)
	assert.Equal("fred", meta.Type)
	assert.Equal("abc", string(detail))
	fp.Close()
}

// `TestFilePersist3` tests
func TestFilePersist3(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() {
		path := os.TempDir() + "/badpath"
		os.RemoveAll(path)
		var fp FilePersistor
		fp.Open(path)
		_ = fp.GetPageList()
		fp.Close()
	})
}

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
