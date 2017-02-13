package pagestore

import (
	"store"
	"testing"
)

func makeTestPersistentPageStore() *PersistentPageStore { return &PersistentPageStore{} }

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
	return &PageStore{backStore: makeTestPersistentPageStore()}
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
