package HistoryFs

import (
	"CryptoGUI/JOSN"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// the function that will save the history info of the ciphers
func HistoryF(in *widget.Entry, view *widget.Entry, current string, HistoryC *fyne.Container, INPUT *widget.Entry, OUTPUT *widget.Entry, CIPHER *widget.Entry) {
	// load the data from the JSON file
	var historyload []JOSN.History = JOSN.Load()
	// check if the input has anything to decode or encode
	if in.Text != "" && view.Text != "Choose a Chipher!!" {
		// here it will store the time and input and output and the cipher used in variables
		time2 := time.Now().Format("2006-01-02 PM 15:04:05")
		first := in.Text
		second := view.Text
		Third := current
		// then here it will add the button with the information to the list in the UI and the JSON file but it won't save to the JSON now.
		HistoryC.Add(widget.NewButton(time2, func() {
			INPUT.SetText(first)
			OUTPUT.SetText(second)
			CIPHER.SetText(Third)
		}))
		HistoryC.Refresh()
		// here it will create the row that has the data for a process and it will be add to the JSON file and to the UI.
		newrow := JOSN.History{
			DateNtime: time2,
			INPUT:     in.Text,
			OUTPUT:    view.Text,
			CIPHER:    current,
		}
		// here it will add append the row to the array
		historyload = append(historyload, newrow)
		// and finally it will add this all to the JOSN file and save it so when i open the app all the history will still be there
		JOSN.Save(historyload)

	}

}

// won't say anything more then this is the same process but for the hashing tab
func HistoryFH(in *widget.Entry, view *widget.Entry, current string, HistoryC *fyne.Container, INPUT *widget.Entry, OUTPUT *widget.Entry, CIPHER *widget.Entry) {
	var historyload []JOSN.History = JOSN.Load()
	if in.Text != "" && view.Text != "Choose your Hash Type!!" {
		time2 := time.Now().Format("2006-01-02 PM 15:04:05")
		first := in.Text
		second := view.Text
		Third := current
		HistoryC.Add(widget.NewButton(time2, func() {
			INPUT.SetText(first)
			OUTPUT.SetText(second)
			CIPHER.SetText(Third)
			HistoryC.Refresh()
		}))

		newrow := JOSN.History{
			DateNtime: time2,
			INPUT:     in.Text,
			OUTPUT:    view.Text,
			CIPHER:    current,
		}
		historyload = append(historyload, newrow)
		JOSN.Save(historyload)

	}

}

// here is the History tab itself
func HistoryWin(history []JOSN.History) (*fyne.Container, *widget.Entry, *widget.Entry, *widget.Entry, *fyne.Container) {
	// first just some UI components that will be used in the tab
	INPUT := widget.NewMultiLineEntry()
	INPUT.Wrapping = fyne.TextWrapWord
	INPUT.SetMinRowsVisible(3)
	OUTPUT := widget.NewMultiLineEntry()
	OUTPUT.Wrapping = fyne.TextWrapWord
	OUTPUT.SetMinRowsVisible(3)
	CIPHER := widget.NewMultiLineEntry()
	CIPHER.Wrapping = fyne.TextWrapWord
	CIPHER.SetMinRowsVisible(3)
	// here is the vertical container that will have all the buttons in it
	History := container.NewVBox()
	// the clear button to clear all the history from the UI and the File
	Clear := widget.NewButton("Clear", func() {
		history = nil
		History.RemoveAll()
		// just saving nothing to the JSON file to clear it
		JOSN.Save(history)
		History.Refresh()
		// reset all the info that is currently shown
		INPUT.SetText("")
		OUTPUT.SetText("")
		CIPHER.SetText("")

	})
	// hahahaha won't say it again :--)
	Clear.Resize(fyne.NewSize(50, 30))
	Clear.Move(fyne.NewPos(210, 68))
	// yay the background that i made it myself (all the design and backgrounds are originally made by me)
	background2, _ := fyne.LoadResourceFromPath("CryptoGUIHistory.png")
	back2 := canvas.NewImageFromResource(background2)
	// stretching the background to take all the screen
	back2.FillMode = canvas.ImageFillStretch
	back2.Resize(fyne.NewSize(300, 600))
	// the scrollable container that has all the history buttons on it so i can scroll through them
	scroll := container.NewScroll(History)
	// :---)
	scroll.Move(fyne.NewPos(35, 105))
	scroll.Resize(fyne.NewSize(230, 120))
	INPUT.Move(fyne.NewPos(35, 275))
	INPUT.Resize(fyne.NewSize(231, 70))
	OUTPUT.Move(fyne.NewPos(35, 380))
	OUTPUT.Resize(fyne.NewSize(231, 67))
	CIPHER.Move(fyne.NewPos(35, 485))
	CIPHER.Resize(fyne.NewSize(231, 67))
	// and finally returning the container that will be used in the tabs in the main func
	return container.NewWithoutLayout(back2, scroll, INPUT, OUTPUT, CIPHER, Clear), INPUT, OUTPUT, CIPHER, History
}
