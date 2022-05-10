package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/nyan2d/note/model"
)

type EditPage struct {
	Content      *widget.Entry
	FolderSelect *widget.Select
	SaveButton   *widget.Button
	CancelButton *widget.Button
	buttons      *fyne.Container
	bottomBox    *fyne.Container
	container    *fyne.Container
}

func NewEditPage(folders model.Folders) *EditPage {
	x := &EditPage{}

	x.Content = widget.NewMultiLineEntry()
	x.FolderSelect = widget.NewSelect(folders.StringSlice(), func(string) {})
	x.SaveButton = widget.NewButton("Save", func() {})
	x.CancelButton = widget.NewButton("Cancel", func() {})

	x.buttons = container.NewGridWithColumns(2, x.SaveButton, x.CancelButton)
	x.bottomBox = container.NewVBox(x.FolderSelect, x.buttons)
	x.container = container.NewBorder(nil, x.bottomBox, nil, nil, x.Content)

	return x
}

func (ep *EditPage) Reset() {
	ep.Content.SetText("")
	ep.FolderSelect.SetSelectedIndex(-1)
}

func (ep *EditPage) CanvasObject() fyne.CanvasObject {
	return ep.container
}
