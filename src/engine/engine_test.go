package engine

import (
	"persist"
	"testing"

	"ui"
)

type MockPage struct {
}

func (mf *MockPage) Meta() persist.PageMeta                         { return persist.PageMeta{Type: "test", Ident: "abc"} }
func (mf *MockPage) Detail() []byte                                 { return []byte("xyz") }
func (mf *MockPage) Ident() string                                  { return "abc" }
func (mf *MockPage) IndexKey(indexNum int) string                   { return "jkl" }
func (mf *MockPage) ApplyAction(arg2 string, arg3 int, r ui.Result) {}
func (mf *MockPage) GetPageYAMLDetail() string                      { return "" }
func (mf *MockPage) SetDetailAndProcess(arg2 string, r ui.Result)   {}
func (mf *MockPage) CanEdit() (_swig_ret bool)                      { return false }
func (mf *MockPage) Actions() ui.StringVector                       { sv := ui.NewStringVector(); return sv }

// `TestPersistentPageStore1` tests
func TestEngine1(t *testing.T) {
	e := NewEngine()
	ui.TestUI1(e, t)
}
