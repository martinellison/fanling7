// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
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

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
