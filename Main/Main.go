package main

import (
	"CryptoGUI/Base"
	"CryptoGUI/Binary"
	"CryptoGUI/Hashing"
	"CryptoGUI/Hex"
	HistoryFs "CryptoGUI/History"
	"CryptoGUI/JOSN"

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
	historyload := JOSN.Load()

	background, _ := fyne.LoadResourceFromPath("CryptoGUI.png")
	back := canvas.NewImageFromResource(background)
	back.FillMode = canvas.ImageFillStretch
	back.Resize(fyne.NewSize(300, 600))

	in := widget.NewEntry()
	in.Move(fyne.NewPos(35, 105))
	in.Resize(fyne.NewSize(232, 40))

	view := widget.NewMultiLineEntry()

	view.Move(fyne.NewPos(35, 275))
	view.Resize(fyne.NewSize(232, 175))
	ciphers := []string{"Base64", "Base32", "Hex", "Binary"}
	drop := widget.NewSelect(ciphers, func(s string) {
		current = s
	})
	drop.PlaceHolder = "Select Cipher"
	drop.Move(fyne.NewPos(35, 190))
	drop.Resize(fyne.NewSize(229, 40))
	view.Wrapping = fyne.TextWrapWord
	view.SetMinRowsVisible(3)
	History := container.NewVBox()
	INPUT := widget.NewMultiLineEntry()
	INPUT.Wrapping = fyne.TextWrapWord
	INPUT.SetMinRowsVisible(3)
	OUTPUT := widget.NewMultiLineEntry()
	OUTPUT.Wrapping = fyne.TextWrapWord
	OUTPUT.SetMinRowsVisible(3)
	CIPHER := widget.NewMultiLineEntry()
	CIPHER.Wrapping = fyne.TextWrapWord
	CIPHER.SetMinRowsVisible(3)
	decode := widget.NewButton("Decode", func() {
		switch current {

		case "Base64":
			view.SetText(Base.F64totext(in.Text))
		case "Base32":
			view.SetText(Base.F32totext(in.Text))
		case "Hex":
			view.SetText(Hex.Fhextotext(in.Text))
		case "Binary":
			view.SetText(Binary.FBINtotext(in.Text))
		default:
			view.SetText("Choose a Chipher!!")
		}
		HistoryFs.HistoryF(in, view, current, History, INPUT, OUTPUT, CIPHER)

	})
	decode.Resize(fyne.NewSize(95, 47))
	decode.Move(fyne.NewPos(35, 495))
	encode := widget.NewButton("Encode", func() {
		switch current {

		case "Base64":
			view.SetText(Base.Ftextto64(in.Text))
		case "Base32":
			view.SetText(Base.Ftextto32(in.Text))
		case "Hex":
			view.SetText(Hex.Ftexttohex(in.Text))
		case "Binary":
			view.SetText(Binary.FtexttoBIN(in.Text))
		default:
			view.SetText("Choose a Chipher!!")
		}
		HistoryFs.HistoryF(in, view, current, History, INPUT, OUTPUT, CIPHER)
	})
	encode.Resize(fyne.NewSize(95, 47))
	encode.Move(fyne.NewPos(170, 495))
	for _, Each := range historyload {
		History.Add(widget.NewButton(Each.DateNtime, func() {
			INPUT.SetText(Each.INPUT)
			OUTPUT.SetText(Each.OUTPUT)
			CIPHER.SetText(Each.CIPHER)
		}))

	}
	History.Refresh()
	container1 := container.NewWithoutLayout(back, in, view, drop, decode, encode)

	background2, _ := fyne.LoadResourceFromPath("CryptoGUIHistory.png")
	back2 := canvas.NewImageFromResource(background2)
	back2.FillMode = canvas.ImageFillStretch
	back2.Resize(fyne.NewSize(300, 600))

	scroll := container.NewScroll(History)

	scroll.Move(fyne.NewPos(35, 105))
	scroll.Resize(fyne.NewSize(230, 120))
	INPUT.Move(fyne.NewPos(35, 275))
	INPUT.Resize(fyne.NewSize(231, 70))
	OUTPUT.Move(fyne.NewPos(35, 380))
	OUTPUT.Resize(fyne.NewSize(231, 67))
	CIPHER.Move(fyne.NewPos(35, 485))
	CIPHER.Resize(fyne.NewSize(231, 67))

	container2 := container.NewWithoutLayout(back2, scroll, INPUT, OUTPUT, CIPHER)
	tabs := container.NewAppTabs(
		container.NewTabItem("Tool", container1),
		container.NewTabItem("History", container2),
		container.NewTabItem("Hashing", Hashing.Hashing(History, INPUT, OUTPUT, CIPHER)),
	)
	w.SetContent(tabs)
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(300, 630))

	w.ShowAndRun()
}
