package Guesser

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Guesser() *fyne.Container {
	IN := widget.NewMultiLineEntry()
	IN.Wrapping = fyne.TextWrapWord
	IN.SetMinRowsVisible(3)
	Guessing_about := widget.NewMultiLineEntry()
	Guessing_about.Wrapping = fyne.TextWrapWord
	Guessing_about.SetMinRowsVisible(3)

	background2, _ := fyne.LoadResourceFromPath("CryptoGUIHistory.png")
	back2 := canvas.NewImageFromResource(background2)
	back2.FillMode = canvas.ImageFillStretch
	back2.Resize(fyne.NewSize(300, 600))

	return container.NewWithoutLayout(back2)
}
