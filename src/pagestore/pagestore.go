package pagestore

import (
	"itemset"
	"persist"
	"store"
)

type PageStore struct {
	items     itemset.ItemSet
	backStore *PersistentPageStore
}

func (ps *PageStore) Add(storable store.Storable)                            {}
func (ps *PageStore) Changed(storable store.Storable)                        {}
func (ps *PageStore) Get(ident string) (storable store.Storable, found bool) { return nil, false }
func (ps *PageStore) GetByIndex(indexNum int) []store.Storable               { return nil }
func (ps *PageStore) Flag(flagNum int, storable store.Storable)              {}
func (ps *PageStore) GetByFlagAndClear(flagNum int) []store.Storable         { return nil }
func (ps *PageStore) Close()                                                 {}

type PersistentPageStore struct {
	persist.Persistor
}

func (ps *PersistentPageStore) Add(storable store.Storable)                            {}
func (ps *PersistentPageStore) Changed(storable store.Storable)                        {}
func (ps *PersistentPageStore) Get(ident string) (storable store.Storable, found bool) { return }
func (ps *PersistentPageStore) GetByIndex(indexNum int) []store.Storable               { return nil }
func (ps *PersistentPageStore) Flag(flagNum int, storable store.Storable)              {}
func (ps *PersistentPageStore) GetByFlagAndClear(flagNum int) []store.Storable         { return nil }
func (ps *PersistentPageStore) Close()                                                 {}
