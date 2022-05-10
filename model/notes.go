package model

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/nyan2d/note/util"
	"github.com/nyan2d/smartbolt"
)

type Notes struct {
	notes       *smartbolt.Bucket[int, Note]
	cachedItems []*Note
	items       []*Note
	filter      func(data string) bool

	listeners []binding.DataListener
}

func NewNotes(database *smartbolt.Smartbolt) *Notes {
	n := &Notes{
		notes: smartbolt.OpenBucket[int, Note](database, "notes"),
	}

	n.RecacheAllItems()

	return n
}

func (n *Notes) RecacheAllItems() {
	notes, _ := n.notes.GetAll()
	n.cachedItems = []*Note{}
	for _, v := range notes {
        note := v
		n.cachedItems = append(n.cachedItems, &note)
	}
	n.updateFilteredItems()
	n.notifyDataChanged()
}

func (n *Notes) ApplyFilter(filter string) {
	n.filter = func(data string) bool {
		return util.IsFuzzyMatch(filter, data)
	}
	n.updateFilteredItems()
	n.notifyDataChanged()
}

func (n *Notes) RemoveFilter() {
	n.filter = nil
	n.notifyDataChanged()
}

func (n *Notes) Append(item Note) {
	itemPtr := &item
	n.cachedItems = append(n.cachedItems, itemPtr)
	n.updateFilteredItems()
	n.notifyDataChanged()
}

func (n *Notes) AddListener(listener binding.DataListener) {
	n.listeners = append(n.listeners, listener)
}

func (n *Notes) RemoveListener(listener binding.DataListener) {
	for i, listen := range n.listeners {
		if listen != listener {
			continue
		}

		if i == len(n.listeners)-1 {
			n.listeners = n.listeners[:len(n.listeners)-1]
		} else {
			n.listeners = append(n.listeners[:i], n.listeners[i+1:]...)
		}
	}
}

func (n *Notes) GetItem(index int) (binding.DataItem, error) {
	element := n.items[index]
	return binding.BindString(&element.Title), nil
}

func (n *Notes) GetElement(index int) Note {
	return *n.items[index]
}

func (m *Notes) Length() int {
	return len(m.items)
}

func (n *Notes) notifyDataChanged() {
	for _, listener := range n.listeners {
		listener.DataChanged()
	}
}

func (n *Notes) updateFilteredItems() {
	n.items = []*Note{}
	for _, v := range n.cachedItems {
		if n.filter == nil {
			n.items = append(n.items, v)
		} else {
			if n.filter(v.Title) {
				n.items = append(n.items, v)
			}
		}
	}
}
