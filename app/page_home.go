package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type HomePage struct {
	captionLabel *widget.Label
	container    *fyne.Container
}

func NewHomePage() *HomePage {
	x := &HomePage{}

	captionText := "Riddle me this\nRiddle me that\nWho's afraid of the big black?"
	x.captionLabel = widget.NewLabel(captionText)
	x.container = container.NewCenter(x.captionLabel)

	return x
}

func (hp *HomePage) CanvasObject() fyne.CanvasObject {
	return hp.container
}
