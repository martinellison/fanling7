package pagestore

import (
	"store"
	"testing"

	"github.com/stretchr/testify/assert"
)

// `TestFilePersist1` tests
func TestPageStore1(t *testing.T) {
	assert := assert.New(t)
	s := &PageStore{}
	store.TestStore1(s, t)
}
