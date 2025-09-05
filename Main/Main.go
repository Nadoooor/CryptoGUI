package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("CryptoGUI")
	current := "blank"
	background, _ := fyne.LoadResourceFromPath("CryptoGUI.png")
	back := canvas.NewImageFromResource(background)
	back.FillMode = canvas.ImageFillStretch
	back.Resize(fyne.NewSize(300, 600))

	in := widget.NewEntry()
	in.Move(fyne.NewPos(35, 105))
	in.Resize(fyne.NewSize(229, 40))

	view := widget.NewEntry()
	view.Disable()
	view.Move(fyne.NewPos(35, 280))
	view.Resize(fyne.NewSize(229, 170))
	ciphers := []string{"Base64", "Base32", "Hex", "Binary", "Caeser"}
	drop := widget.NewSelect(ciphers, func(s string) {
		current = s
	})
	drop.PlaceHolder = "Select Cipher"
	drop.Move(fyne.NewPos(35, 190))
	drop.Resize(fyne.NewSize(229, 40))

	decode := widget.NewButton("Decode", func() {
		view.SetText(current)
	})
	decode.Resize(fyne.NewSize(95, 42))
	decode.Move(fyne.NewPos(35, 500))
	encode := widget.NewButton("Encode", func() {})
	encode.Resize(fyne.NewSize(95, 42))
	encode.Move(fyne.NewPos(170, 500))

	container1 := container.NewWithoutLayout(back, in, view, drop, decode, encode)
	w.SetContent(container1)
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(300, 600))

	w.ShowAndRun()
}
