package itemset

import (
	"errors"
)

// `AddList`
type AddList interface {
	Add(string)
}

// `Identified`
type Identified interface {
	Ident() string
}

// `ItemSet`
type ItemSet struct {
	items     map[string]Identified
	itemList  []Identified
	iterating bool
	index     int
}

// `Init` initialises the ident set.
func (is *ItemSet) Init() { is.items = make(map[string]Identified, 0) }

// `Clear` clears the ident set
func (is *ItemSet) Clear() {
	is.items = make(map[string]Identified, 0)
	is.itemList = nil
	is.iterating = false
}

// `Add` adds an item.
func (is *ItemSet) Add(i Identified) error {
	id := i.Ident()
	if _, alreadyExists := is.items[id]; alreadyExists {
		return errors.New("duplicate item: " + id)
	}
	is.items[id] = i
	if is.iterating {
		is.itemList = append(is.itemList, i)
	}
	return nil
}

// `GetItems` gets a list of the names of all items.
func (is *ItemSet) GetItems(al AddList) {
	for t := range is.items {
		al.Add(t)
	}
}

// `Items` gets a list of all items.
func (is *ItemSet) Items() (items []Identified) {
	items = make([]Identified, 0)
	for _, i := range is.items {
		items = append(items, i)
	}
	return
}

// `ItemExists` tests whether there is an item with the given identifier.
func (is *ItemSet) ItemExists(ident string) bool {
	_, found := is.items[ident]
	return found
}

// `GetItem` gets the item with a given ident.
func (is *ItemSet) GetItem(ident string) (item Identified, err error) {
	item, found := is.items[ident]
	if !found {
		err = errors.New("Item not found: " + ident)
	}
	return
}

// `GetItemOrNot` gets the item with a given ident or returns found == false.
func (is *ItemSet) GetItemOrNot(ident string) (item Identified, found bool) {
	item, found = is.items[ident]
	return
}

// `GetFirst` gets the first item in a item set.
func (is *ItemSet) GetFirst() (item Identified, more bool) {
	if is.items == nil || len(is.items) == 0 {
		is.iterating = false
		return nil, false
	}
	is.itemList = make([]Identified, 0)
	for _, i := range is.items {
		is.itemList = append(is.itemList, i)
	}
	is.iterating = true
	is.index = 0
	return is.itemList[0], true
}

// `GetNext` gets the next item in a item set.
func (is *ItemSet) GetNext() (item Identified, more bool) {
	is.index++
	if is.index >= len(is.itemList) {
		is.iterating = false
		is.itemList = nil
		return nil, false
	}
	return is.itemList[is.index], true
}

// `Size` returns the number of items in a item set.
func (is *ItemSet) Size() int { return len(is.items) }
