package app

import (
	"strings"
	"time"

	"github.com/nyan2d/note/model"
	"github.com/nyan2d/note/util"
	"github.com/nyan2d/smartbolt"
)

func (a *App) hEditpageCancel() {
	a.editPage.Reset()
	a.SwitchPage(a.homePage)
}

func (a *App) hEditpageSave() {
	content := a.editPage.Content.Text
	folderID := a.editPage.FolderSelect.SelectedIndex() // TODO:

	if strings.TrimSpace(content) == "" {
		a.ShowAlert("1")
		return
	}

	if folderID == -1 {
		a.ShowAlert("2")
		return
	}

	nbucket := smartbolt.OpenBucket[int, model.Note](a.db, "notes")
	note := model.Note{
		ID:        nbucket.NextID(),
		Folder:    folderID,
		Title:     util.GetTitle(content),
		Content:   content,
		CreatedAt: time.Now(),
		EditedAt:  time.Now(),
	}
	nbucket.Put(note.ID, note)
    a.notes.Append(note) // TODO: ???
}
