package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/nyan2d/note/model"
)

type NavigationPage struct {
	SearchEntry  *widget.Entry
	FolderSelect *widget.Select
	NotesList    *widget.List
	CreateButton *widget.Button
	topBox       *fyne.Container
	container    *fyne.Container

	selectedNoteID int
}

func NewNavigationPage(folders *model.Folders, notes *model.Notes) *NavigationPage {
	x := &NavigationPage{}

    x.selectedNoteID = -1
	x.SearchEntry = widget.NewEntry()
	x.SearchEntry.SetPlaceHolder("Type to search")
	x.FolderSelect = widget.NewSelect(folders.StringSlice(), func(string) {})
	x.NotesList = widget.NewListWithData(notes, x.notesListCreate, x.notesListUpdate)
	x.CreateButton = widget.NewButton("Create New", func() {})
	x.topBox = container.NewVBox(x.SearchEntry, x.FolderSelect)
	x.container = container.NewBorder(x.topBox, x.CreateButton, nil, nil, x.NotesList)

	return x
}

func (np *NavigationPage) CanvasObject() fyne.CanvasObject {
	return np.container
}

func (np *NavigationPage) notesListCreate() fyne.CanvasObject {
	return widget.NewLabel("")
}

func (np *NavigationPage) notesListUpdate(item binding.DataItem, obj fyne.CanvasObject) {
	obj.(*widget.Label).Bind(item.(binding.String))
}
