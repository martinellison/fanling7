// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
package store

type Store interface {
	Add(storable Storable)
	Changed(storable Storable)
	Get(ident string) (storable Storable, found bool)
	GetByIndex(indexNum int) []Storable
	Flag(flagNum int, storable Storable)
	GetByFlagAndClear(flagNum int) map[string]Storable
	Close()
}

type Storable interface {
	Ident() (ident string)
	IndexKey(indexNum int) string
}

const MaxFlag = 1
const MaxIndex = 1

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
