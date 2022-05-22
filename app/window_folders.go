package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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

func NewFoldersWindow(app *App) *FoldersWindow {
	x := &FoldersWindow{}

	x.window = app.App.NewWindow("Folders")

	x.FoldersList = widget.NewListWithData(
		app.Folders,
		func() fyne.CanvasObject {
			return widget.NewLabel("-")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			item := di.(binding.String)
			co.(*widget.Label).Bind(item)
		},
	)

	x.AddButton = widget.NewButton("Add", app.hWfoldersAdd)
	x.RemoveButton = widget.NewButton("Remove", app.hWfoldersRemove)
	x.RenameButton = widget.NewButton("Rename", app.hWfoldersRename)
	x.CloseButton = widget.NewButton("Close", app.hWfoldersClose)

	x.buttons = container.NewGridWithColumns(4, x.AddButton, x.RemoveButton, x.RenameButton, x.CloseButton)
	x.container = container.NewBorder(nil, x.buttons, nil, nil, x.FoldersList)
	x.window.SetContent(x.container)

	return x
}

func (fw *FoldersWindow) Show() {
	fw.window.Show()
}

func (fw *FoldersWindow) Hide() {
	fw.window.Hide()
}
