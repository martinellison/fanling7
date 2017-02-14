package pagestore

import (
	"persist"
	"store"
	"testing"
)

// `TestPersistentPageStore1` tests
func TestPersistentPageStore1(t *testing.T) {
	s := MakePersistentPageStore("test", &MockFreezer{})
	store.TestStore1(s, t)
}

// `TestPersistentPageStore2` tests
func TestPersistentPageStore2(t *testing.T) {
	s := MakePersistentPageStore("test", &MockFreezer{})
	store.TestStore2(s, t)
}

// `TestPersistentPageStore3` tests
func TestPersistentPageStore3(t *testing.T) {
	s := MakePersistentPageStore("test", &MockFreezer{})
	store.TestStore3(s, t)
}

// `TestPersistentPageStore4` tests
func TestPersistentPageStore4(t *testing.T) {
	s := MakePersistentPageStore("test", &MockFreezer{})
	store.TestStore4(s, &MockFreezable{}, t)
}

// `TestPageStore1` tests
func TestPageStore1(t *testing.T) {
	s := MakePageStore("test", &MockFreezer{})
	store.TestStore1(s, t)
}

// `TestPageStore2` tests
func TestPageStore2(t *testing.T) {
	s := MakePageStore("test", &MockFreezer{})
	store.TestStore2(s, t)
}

// `TestPageStore3` tests
func TestPageStore3(t *testing.T) {
	s := MakePageStore("test", &MockFreezer{})
	store.TestStore3(s, t)
}

// `TestPageStore4` tests
func TestPageStore4(t *testing.T) {
	s := MakePageStore("test", &MockFreezer{})
	store.TestStore4(s, &MockFreezable{}, t)
}

type MockFreezer struct{}
type MockFreezable struct{}

func (mf *MockFreezable) Meta() persist.PageMeta       { return persist.PageMeta{Type: "test", Ident: "abc"} }
func (mf *MockFreezable) Detail() []byte               { return []byte("xyz") }
func (mf *MockFreezable) Ident() string                { return "abc" }
func (mf *MockFreezable) IndexKey(indexNum int) string { return "jkl" }

func (mf *MockFreezer) Freeze(meta persist.PageMeta, detail []byte) (f Freezable) {
	return &MockFreezable{}
}
