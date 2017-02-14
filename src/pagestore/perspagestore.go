// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
package pagestore

import (
	"filepersist"
	"persist"
	"store"
)

type PersistentPageStore struct {
	persist.Persistor
	Freezer
	flagged []map[string]store.Storable
}

//Open(path string)
//Close()
//GetPageList() (paths []string)
//ReadPage(ident string) (meta PageMeta, detail []byte, found bool)
//ReadPageFromPath(path string) (meta PageMeta, detail []byte, found bool)
//WritePage(meta PageMeta, detail []byte)

func MakePersistentPageStore(path string, freezer Freezer) *PersistentPageStore {
	p := &filepersist.FilePersistor{}
	p.Open(path)
	s := &PersistentPageStore{Persistor: p, Freezer: freezer, flagged: make([]map[string]store.Storable, store.MaxFlag)}
	return s
}

func (pps *PersistentPageStore) Add(storable store.Storable) {
	ps := storable.(Freezable)
	pps.Persistor.WritePage(ps.Meta(), ps.Detail())
}
func (pps *PersistentPageStore) Changed(storable store.Storable) {
	ps := storable.(Freezable)
	pps.Persistor.WritePage(ps.Meta(), ps.Detail())
}
func (pps *PersistentPageStore) Get(ident string) (storable store.Storable, found bool) {
	meta, detail, found := pps.Persistor.ReadPage(ident)
	if found {
		storable = pps.Freezer.Freeze(meta, detail)
	}
	return
}
func (pps *PersistentPageStore) GetByIndex(indexNum int) []store.Storable {
	//  panic("not coded") // TODO: need index over all pages
	return []store.Storable{}
}
func (pps *PersistentPageStore) Flag(flagNum int, storable store.Storable) {
	m := pps.flagged[flagNum]
	if m == nil {
		m = make(map[string]store.Storable, 0)
		pps.flagged[flagNum] = m
	}
	m[storable.Ident()] = storable
}
func (pps *PersistentPageStore) GetByFlagAndClear(flagNum int) map[string]store.Storable {
	m := pps.flagged[flagNum]
	if m == nil {
		m = make(map[string]store.Storable, store.MaxFlag)
		pps.flagged[flagNum] = m
	}
	pps.flagged[flagNum] = m
	return m
	return nil
}
func (pps *PersistentPageStore) Close() { pps.Persistor.Close() }

type Freezable interface {
	store.Storable
	Meta() persist.PageMeta
	Detail() []byte
}
type Freezer interface {
	Freeze(meta persist.PageMeta, detail []byte) (f Freezable)
}

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
