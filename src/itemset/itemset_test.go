// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
package itemset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stringList struct {
	sa []string
}

// `Add` adds a sting to a string list and so makes `stringList` into an `AddList'.
func (sl *stringList) Add(s string) { sl.sa = append(sl.sa, s) }

type thing struct {
	id string
}

// `Ident` makes `thing` into an `identified`.
func (t thing) Ident() string { return t.id }

// `TestItemSet1` tests the case where there are no items in the `ItemSet`.
func TestItemSet1(t *testing.T) {
	assert := assert.New(t)
	var is ItemSet
	is.Init()
	assert.False(is.ItemExists("fred"))
	_, err := is.GetItem("fred")
	assert.NotNil(err)
}

// `TestItemSet2` tests the case where  there is one item in the `ItemSet`.
func TestItemSet2(t *testing.T) {
	assert := assert.New(t)
	var is ItemSet
	is.Init()
	err := is.Add(thing{id: "fred"})
	assert.Nil(err)
	assert.True(is.ItemExists("fred"))
}

// `TestItemSet3` tests the case where there is one item in the `ItemSet`.
func TestItemSet3(t *testing.T) {
	assert := assert.New(t)
	var is ItemSet
	is.Init()
	err := is.Add(thing{id: "fred"})
	assert.Nil(err)
	item, err := is.GetItem("fred")
	assert.Nil(err)
	assert.Equal("fred", item.Ident())
	item2, found := is.GetItemOrNot("fred")
	assert.True(found)
	assert.Equal("fred", item2.Ident())
}

// `TestItemSet4` tests `getItems` for the case where there is one item in the `ItemSet`.
func TestItemSet4(t *testing.T) {
	assert := assert.New(t)
	var is ItemSet
	is.Init()
	err := is.Add(thing{id: "fred"})
	assert.Nil(err)
	sl := &stringList{sa: make([]string, 0)}
	is.GetItems(sl)
	assert.Equal(1, len(sl.sa))
	assert.Equal("fred", sl.sa[0])
	items := is.Items()
	assert.Equal(1, len(items))
	assert.Equal("fred", items[0].Ident())
	assert.Equal(1, is.Size())
}

// `TestItemSet5` tests the case of trying to add the same item twice to the `ItemSet`.
func TestItemSet5(t *testing.T) {
	assert := assert.New(t)
	var is ItemSet
	is.Init()
	err := is.Add(thing{id: "fred"})
	assert.Nil(err)
	err = is.Add(thing{id: "fred"})
	assert.NotNil(err)
}

// `TestItemSet6` tests `getItems` `Clear`.
func TestItemSet6(t *testing.T) {
	assert := assert.New(t)
	var is ItemSet
	is.Init()
	err := is.Add(thing{id: "fred"})
	assert.Nil(err)
	assert.Equal(1, is.Size())
	is.Clear()
	assert.Equal(0, is.Size())
}

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
