// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
package persist

// A `PageMeta` describes a Page independently of the page type.
type PageMeta struct {
	Type  string
	Ident string
}

type Persistor interface {
	Open(path string)
	Close()
	GetPageList() (paths []string)
	ReadPage(ident string) (meta PageMeta, detail []byte, found bool)
	ReadPageFromPath(path string) (meta PageMeta, detail []byte, found bool)
	WritePage(meta PageMeta, detail []byte)
}

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
