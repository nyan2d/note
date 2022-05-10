package app

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/nyan2d/note/model"
)

type PreviewPage struct {
	TitleLabel      *widget.Label
	DateLabel       *widget.Label
	ContentRichText *widget.RichText
	topBox          *fyne.Container
	container       *fyne.Container
}

func NewPreviewPage() *PreviewPage {
	x := &PreviewPage{}

	x.TitleLabel = widget.NewLabel("title")
	x.DateLabel = widget.NewLabel("date")
	x.ContentRichText = widget.NewRichTextFromMarkdown("")
    x.ContentRichText.Wrapping = fyne.TextWrapWord

	x.topBox = container.NewGridWithColumns(2, x.TitleLabel, x.DateLabel)
	x.container = container.NewBorder(x.topBox, nil, nil, nil, x.ContentRichText)

	return x
}

func (pp *PreviewPage) SetContent(note model.Note) {
    pp.TitleLabel.SetText(note.Title)
    pp.DateLabel.SetText(note.CreatedAt.Format(time.RFC822))
    pp.ContentRichText.ParseMarkdown(note.Content)
}

func (pp *PreviewPage) CanvasObject() fyne.CanvasObject {
    return pp.container
}
