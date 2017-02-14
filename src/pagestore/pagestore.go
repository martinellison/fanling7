// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
package pagestore

import (
	"itemset"
	"store"
)

type PageStore struct {
	items     itemset.ItemSet
	backStore *PersistentPageStore
	flagged   []map[string]store.Storable
}

func MakePageStore(path string, freezer Freezer) *PageStore {
	s := &PageStore{backStore: MakePersistentPageStore(path, freezer), flagged: make([]map[string]store.Storable, store.MaxFlag)}
	s.items.Init()
	return s
}
func (ps *PageStore) Add(storable store.Storable)     { ps.items.Add(storable); ps.backStore.Add(storable) }
func (ps *PageStore) Changed(storable store.Storable) { ps.backStore.Changed(storable) }
func (ps *PageStore) Get(ident string) (storable store.Storable, found bool) {
	storableId, found := ps.items.GetItemOrNot(ident)
	if found {
		storable = storableId.(store.Storable)
		return
	}
	storable, found = ps.backStore.Get(ident)
	return
}
func (ps *PageStore) GetByIndex(indexNum int) []store.Storable {
	ss := make([]store.Storable, store.MaxIndex)
	return ss
}
func (ps *PageStore) Flag(flagNum int, storable store.Storable) {
	m := ps.flagged[flagNum]
	if m == nil {
		m = make(map[string]store.Storable, 0)
		ps.flagged[flagNum] = m
	}
	m[storable.Ident()] = storable
}
func (ps *PageStore) GetByFlagAndClear(flagNum int) map[string]store.Storable {
	m := ps.flagged[flagNum]
	if m == nil {
		m = make(map[string]store.Storable, store.MaxFlag)
		ps.flagged[flagNum] = m
	}
	ps.flagged[flagNum] = m
	return m
}
func (ps *PageStore) Close() { ps.backStore.Close() }

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
