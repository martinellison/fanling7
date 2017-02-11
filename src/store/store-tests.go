package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore1(s Store, t *testing.T) {
	assert := assert.New(t)
	storable, found := s.Get("fred")
	assert.False(found)
}
func TestStore2(s Store, t *testing.T) {
	assert := assert.New(t)
	ss := s.GetByIndex(0)
	assert.Equal(0, len(ss))
}
func TestStore3(s Store, t *testing.T) {
	assert := assert.New(t)
	ss := GetByFlagAndClear(0)
	assert.Equal(0, len(ss))
}
