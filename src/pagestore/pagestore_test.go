package pagestore

import (
	"filepersist"
	"store"
	"testing"
)

func makeTestPersistentPageStore() *PersistentPageStore {
	p := &filepersist.FilePersistor{}
	p.Open("?")
	s := &PersistentPageStore{Persistor: p}
	return s
}

// `TestPersistentPageStore1` tests
func TestPersistentPageStore1(t *testing.T) {
	s := makeTestPersistentPageStore()
	store.TestStore1(s, t)
}

// `TestPersistentPageStore2` tests
func TestPersistentPageStore2(t *testing.T) {
	s := makeTestPersistentPageStore()
	store.TestStore2(s, t)
}

// `TestPersistentPageStore3` tests
func TestPersistentPageStore3(t *testing.T) {
	s := makeTestPersistentPageStore()
	store.TestStore3(s, t)
}

// `TestPersistentPageStore4` tests
func TestPersistentPageStore4(t *testing.T) {
	s := makeTestPersistentPageStore()
	store.TestStore4(s, t)
}

func makeTestPageStore() *PageStore {
	s := &PageStore{backStore: makeTestPersistentPageStore()}
	s.items.Init()
	return s
}

// `TestPageStore1` tests
func TestPageStore1(t *testing.T) {
	s := makeTestPageStore()
	store.TestStore1(s, t)
}

// `TestPageStore2` tests
func TestPageStore2(t *testing.T) {
	s := makeTestPageStore()
	store.TestStore2(s, t)
}

// `TestPageStore3` tests
func TestPageStore3(t *testing.T) {
	s := makeTestPageStore()
	store.TestStore3(s, t)
}

// `TestPageStore4` tests
func TestPageStore4(t *testing.T) {
	s := makeTestPageStore()
	store.TestStore4(s, t)
}
