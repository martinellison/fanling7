package pagestore

import (
	"filepersist"
	"persist"
	"store"
)

type PersistentPageStore struct {
	persist.Persistor
	Freezer
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
	s := &PersistentPageStore{Persistor: p, Freezer: freezer}
	return s
}

func (pps *PersistentPageStore) Add(storable store.Storable) {
	ps := storable.(Freezable)
	pps.Persistor.WritePage(ps.Meta(), ps.Detail())
}
func (pps *PersistentPageStore) Changed(storable store.Storable) {}
func (pps *PersistentPageStore) Get(ident string) (storable store.Storable, found bool) {
	meta, detail, found := pps.Persistor.ReadPage(ident)
	if found {
		storable = pps.Freezer.Freeze(meta, detail)
	}
	return
}
func (pps *PersistentPageStore) GetByIndex(indexNum int) []store.Storable                { return nil }
func (pps *PersistentPageStore) Flag(flagNum int, storable store.Storable)               {}
func (pps *PersistentPageStore) GetByFlagAndClear(flagNum int) map[string]store.Storable { return nil }
func (pps *PersistentPageStore) Close()                                                  { pps.Persistor.Close() }

type Freezable interface {
	store.Storable
	Meta() persist.PageMeta
	Detail() []byte
}
type Freezer interface {
	Freeze(meta persist.PageMeta, detail []byte) (f Freezable)
}
