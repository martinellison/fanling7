package filepersist

import (
	"os"
	"persist"
	"testing"

	"github.com/stretchr/testify/assert"
)

// `TestFilePersist1` tests
func TestFilePersist1(t *testing.T) {
	assert := assert.New(t)
	path := os.TempDir() + "/fp"
	os.RemoveAll(path)
	os.MkdirAll(path, 0777)
	var fp FilePersistor
	fp.Open(path)
	l0 := fp.GetPageList()
	assert.Equal(0, len(l0))
	fp.Close()
}

// `TestFilePersist2` tests
func TestFilePersist2(t *testing.T) {
	assert := assert.New(t)
	path := os.TempDir() + "/fp"
	os.RemoveAll(path)
	os.MkdirAll(path, 0777)
	var fp FilePersistor
	fp.Open(path)
	fp.WritePage(persist.PageMeta{Type: "fred", Ident: "nurke"}, []byte("abc"))
	l1 := fp.GetPageList()
	assert.Equal(1, len(l1))
	assert.Equal("nurke", l1[0])
	meta, detail, found := fp.ReadPage("nurke")
	assert.True(found)
	assert.Equal("nurke", meta.Ident)
	assert.Equal("fred", meta.Type)
	assert.Equal("abc", string(detail))
	fp.Close()
}

// `TestFilePersist3` tests
func TestFilePersist3(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() {
		path := os.TempDir() + "/badpath"
		os.RemoveAll(path)
		var fp FilePersistor
		fp.Open(path)
		_ = fp.GetPageList()
		fp.Close()
	})
}
