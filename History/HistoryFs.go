package HistoryFs

import (
	"CryptoGUI/JOSN"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HistoryF(in *widget.Entry, view *widget.Entry, current string, HistoryC *fyne.Container, INPUT *widget.Entry, OUTPUT *widget.Entry, CIPHER *widget.Entry) {
	var historyload []JOSN.History = JOSN.Load()
	if in.Text != "" && view.Text != "Choose a Chipher!!" {
		time2 := time.Now().Format("2006-01-02 PM 15:04:05")
		first := in.Text
		second := view.Text
		Third := current
		HistoryC.Add(widget.NewButton(time2, func() {
			INPUT.SetText(first)
			OUTPUT.SetText(second)
			CIPHER.SetText(Third)
		}))
		HistoryC.Refresh()
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

func HistoryWin(history []JOSN.History) (*fyne.Container, *widget.Entry, *widget.Entry, *widget.Entry, *fyne.Container) {
	INPUT := widget.NewMultiLineEntry()
	INPUT.Wrapping = fyne.TextWrapWord
	INPUT.SetMinRowsVisible(3)
	OUTPUT := widget.NewMultiLineEntry()
	OUTPUT.Wrapping = fyne.TextWrapWord
	OUTPUT.SetMinRowsVisible(3)
	CIPHER := widget.NewMultiLineEntry()
	CIPHER.Wrapping = fyne.TextWrapWord
	CIPHER.SetMinRowsVisible(3)
	History := container.NewVBox()
	Clear := widget.NewButton("Clear", func() {
		history = nil
		History.RemoveAll()
		JOSN.Save(history)
		History.Refresh()
		INPUT.SetText("")
		OUTPUT.SetText("")
		CIPHER.SetText("")

	})
	Clear.Resize(fyne.NewSize(50, 30))
	Clear.Move(fyne.NewPos(210, 68))

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

	return container.NewWithoutLayout(back2, scroll, INPUT, OUTPUT, CIPHER, Clear), INPUT, OUTPUT, CIPHER, History
}
