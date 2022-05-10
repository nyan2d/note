package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type FoldersWindow struct {
	window       fyne.Window
	FoldersList  *widget.List
	AddButton    *widget.Button
	RemoveButton *widget.Button
	RenameButton *widget.Button
	CloseButton  *widget.Button
	buttons      *fyne.Container
	container    *fyne.Container
}

func NewFoldersWindow(application fyne.App) *FoldersWindow {
	x := &FoldersWindow{}

	x.window = application.NewWindow("Folders")

	hahadata := []string{
		"Первое",
		"Второе",
		"Ещё одно",
	}

	x.FoldersList = widget.NewList(
		func() int {
			return len(hahadata)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(hahadata[id])
		},
	)

	x.AddButton = widget.NewButton("Add", empty)
	x.RemoveButton = widget.NewButton("Remove", empty)
	x.RenameButton = widget.NewButton("Rename", empty)
	x.CloseButton = widget.NewButton("Close", func() { x.window.Hide() })

	x.buttons = container.NewGridWithColumns(4, x.AddButton, x.RemoveButton, x.RenameButton, x.CloseButton)
	x.container = container.NewBorder(nil, x.buttons, nil, nil, x.FoldersList)
	x.window.SetContent(x.container)

	return x
}

func (fw *FoldersWindow) Show() {
	fw.window.Show()
}
