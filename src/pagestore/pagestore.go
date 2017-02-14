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
