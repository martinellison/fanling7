// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
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

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
