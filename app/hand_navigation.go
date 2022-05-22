package app

import (
	"fyne.io/fyne/v2/widget"
)

func (a *App) hNavCreate() {
	a.editPage.Reset()
	a.SwitchPage(a.editPage)
}

func (a *App) hNavSearchTextChanged(s string) {
	a.navigationPage.NotesList.UnselectAll()
	a.SwitchPage(a.homePage)
	if len(s) > 0 {
		a.Notes.ApplyFilter(s)
	} else {
		a.Notes.RemoveFilter()
	}
}

func (a *App) hNavListElementSelected(id widget.ListItemID) {
	note := a.Notes.GetElement(id)
	a.previewPage.SetContent(note)
	a.SwitchPage(a.previewPage)
	a.navigationPage.selectedNoteID = note.ID
}
