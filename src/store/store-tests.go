package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorable struct{}

func (ms *MockStorable) Ident() (ident string)        { return "abc" }
func (ms *MockStorable) IndexKey(indexNum int) string { return "xyz" }

func TestStore1(s Store, t *testing.T) {
	assert := assert.New(t)
	_, found := s.Get("fred")
	assert.False(found)
}
func TestStore2(s Store, t *testing.T) {
	assert := assert.New(t)
	ss := s.GetByIndex(0)
	assert.Equal(0, len(ss))
}
func TestStore3(s Store, t *testing.T) {
	assert := assert.New(t)
	ss := s.GetByFlagAndClear(0)
	assert.Equal(0, len(ss))
}
func TestStore4(s Store, ms Storable, t *testing.T) {
	assert := assert.New(t)
	s.Add(ms)
	storable, found := s.Get("abc")
	assert.True(found)
	assert.Equal(ms, storable)
}
