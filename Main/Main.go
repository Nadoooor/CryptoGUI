package main

import (
	"CryptoGUI/Base"
	"CryptoGUI/Binary"
	"CryptoGUI/Brute"
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
	// the app backbone
	a := app.New()
	w := a.NewWindow("CryptoGUI")
	// some vars and ui components
	current := "blank"
	historyload := JOSN.Load()
	HistoryWin, INPUT, OUTPUT, CIPHER, History := HistoryFs.HistoryWin(historyload)
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
	// the drop down sections it was kinda wierd but i got it
	drop.PlaceHolder = "Select Cipher"
	drop.Move(fyne.NewPos(35, 190))
	drop.Resize(fyne.NewSize(229, 40))
	view.Wrapping = fyne.TextWrapWord
	view.SetMinRowsVisible(3)
	// making the input entry to submit when i press enter
	in.OnSubmitted = func(s string) {
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
	}
	// the decode button
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
		// adding the information to the history window
		HistoryFs.HistoryF(in, view, current, History, INPUT, OUTPUT, CIPHER)

	})
	// moving and resizing the decode button
	decode.Resize(fyne.NewSize(95, 47))
	decode.Move(fyne.NewPos(35, 495))
	// the encode button
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
		// adding the encoding information to the history window this time
		HistoryFs.HistoryF(in, view, current, History, INPUT, OUTPUT, CIPHER)
	})
	// moving and resizing the encode button
	encode.Resize(fyne.NewSize(95, 47))
	encode.Move(fyne.NewPos(170, 495))
	// adding all the history buttons and info to the windows from the json file
	for _, Each := range historyload {
		History.Add(widget.NewButton(Each.DateNtime, func() {
			INPUT.SetText(Each.INPUT)
			OUTPUT.SetText(Each.OUTPUT)
			CIPHER.SetText(Each.CIPHER)
		}))

	}
	History.Refresh()
	// the first container for the ciphers encoding and decoding section maybe i should move this all to a separite package
	container1 := container.NewWithoutLayout(back, in, view, drop, decode, encode)
	// adding all the tabs to the app so i can choose whatever i want.
	tabs := container.NewAppTabs(
		container.NewTabItem("Tool", container1),
		container.NewTabItem("Brute", Brute.Bruter(w)),
		container.NewTabItem("History", HistoryWin),
		container.NewTabItem("Hashing", Hashing.Hashing(History, INPUT, OUTPUT, CIPHER)),
	)
	// setting all the code into fyne gui to show
	w.SetContent(tabs)
	w.SetFixedSize(true)
	// just resizing the window itself and making it fixed size
	w.Resize(fyne.NewSize(300, 630))

	w.ShowAndRun()
}
