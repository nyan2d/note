package app

import "fyne.io/fyne/v2"

type MainMenu struct {
	menu     *fyne.MainMenu
	fileMenu *fyne.Menu
	editMenu *fyne.Menu
	helpMenu *fyne.Menu

	FileQuitMenuItem    *fyne.MenuItem
	EditFoldersMenuItem *fyne.MenuItem
	EditEditMenuItem    *fyne.MenuItem
	EditRemoveMenuItem  *fyne.MenuItem
	HelpGithubMenuItem  *fyne.MenuItem
	HelpAboutMenuItem   *fyne.MenuItem
}

func empty() {}

func NewMainMenu() *MainMenu {
	x := &MainMenu{}

	x.FileQuitMenuItem = fyne.NewMenuItem("Quit", empty)
	x.EditFoldersMenuItem = fyne.NewMenuItem("Folders", empty)
	x.EditEditMenuItem = fyne.NewMenuItem("Edit", empty)
	x.EditRemoveMenuItem = fyne.NewMenuItem("Remove", empty)
	x.HelpGithubMenuItem = fyne.NewMenuItem("Github", empty)
	x.HelpAboutMenuItem = fyne.NewMenuItem("About", empty)

	x.fileMenu = fyne.NewMenu("File", x.FileQuitMenuItem)
	x.editMenu = fyne.NewMenu("Edit", x.EditFoldersMenuItem, x.EditEditMenuItem, x.EditRemoveMenuItem)
	x.helpMenu = fyne.NewMenu("Help", x.HelpGithubMenuItem, x.HelpAboutMenuItem)

	x.menu = fyne.NewMainMenu(x.fileMenu, x.editMenu, x.helpMenu)

	return x
}

func (mm *MainMenu) Menu() *fyne.MainMenu {
	return mm.menu
}
