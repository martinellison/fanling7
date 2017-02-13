package pagestore

import (
	"itemset"
	"persist"
	"store"
)

type PageStore struct {
	items     itemset.ItemSet
	backStore *PersistentPageStore
	flagged   []map[string]store.Storable
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
	ss := make([]store.Storable, 0)
	return ss
}
func (ps *PageStore) Flag(flagNum int, storable store.Storable) {
	if ps.flagged == nil {
		ps.flagged = make([]map[string]store.Storable, 0)
	}
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
		m = make(map[string]store.Storable, 0)
		ps.flagged[flagNum] = m
	}
	ps.flagged[flagNum] = m
	return m
}
func (ps *PageStore) Close() { ps.backStore.Close() }

type PersistentPageStore struct {
	persist.Persistor
}

func (ps *PersistentPageStore) Add(storable store.Storable)                             {}
func (ps *PersistentPageStore) Changed(storable store.Storable)                         {}
func (ps *PersistentPageStore) Get(ident string) (storable store.Storable, found bool)  { return }
func (ps *PersistentPageStore) GetByIndex(indexNum int) []store.Storable                { return nil }
func (ps *PersistentPageStore) Flag(flagNum int, storable store.Storable)               {}
func (ps *PersistentPageStore) GetByFlagAndClear(flagNum int) map[string]store.Storable { return nil }
func (ps *PersistentPageStore) Close()                                                  {}
