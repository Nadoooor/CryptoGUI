package HistoryFs

import (
	"CryptoGUI/JOSN"
	"time"

	"fyne.io/fyne/v2"
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
