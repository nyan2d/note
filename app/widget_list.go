package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type List struct {
	*widget.List
	selectedID widget.ListItemID
	OnSelected func(id widget.ListItemID)
}

func NewNotAListWithData(data binding.DataList, createItem func() fyne.CanvasObject, updateItem func(binding.DataItem, fyne.CanvasObject)) *List {
	notalist := &List{
		List: widget.NewListWithData(data, createItem, updateItem),
	}

	notalist.List.OnSelected = func(id widget.ListItemID) {
		notalist.selectedID = id
		if notalist.OnSelected != nil {
			notalist.OnSelected(id)
		}
	}

	return notalist
}

func (nl *List) SelectedID() widget.ListItemID {
	return nl.selectedID
}
