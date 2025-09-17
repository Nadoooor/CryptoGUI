package Hashing

import (
	HistoryFs "CryptoGUI/History"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/getevo/hash"
)

func Hashing(History *fyne.Container, INPUT *widget.Entry, OUTPUT *widget.Entry, CIPHER *widget.Entry) *fyne.Container {
	current := "blank"
	background, _ := fyne.LoadResourceFromPath("Hashing.png")
	back := canvas.NewImageFromResource(background)
	back.FillMode = canvas.ImageFillStretch
	back.Resize(fyne.NewSize(300, 600))

	in := widget.NewEntry()
	in.Move(fyne.NewPos(35, 105))
	in.Resize(fyne.NewSize(232, 40))

	view := widget.NewMultiLineEntry()

	view.Move(fyne.NewPos(35, 275))
	view.Resize(fyne.NewSize(232, 175))
	ciphers := []string{"CRC32", "CRC64", "FNV-1a 32bit", "FNV-1a 64bit", "MD5", "SHA1", "SHA256", "SHA384", "SHA512"}
	drop := widget.NewSelect(ciphers, func(s string) {
		current = s

	})
	drop.PlaceHolder = "Choose your Hash Type!!"
	drop.Move(fyne.NewPos(35, 190))
	drop.Resize(fyne.NewSize(229, 40))
	view.Wrapping = fyne.TextWrapWord
	view.SetMinRowsVisible(3)
	encode := widget.NewButton("Hash It!!", func() {
		switch current {
		case "CRC32":
			view.SetText(hash.CRC32String(in.Text))
		case "CRC64":
			view.SetText(hash.CRC64String(in.Text))
		case "FNV-1a 32bit":
			view.SetText(hash.FNV32String(in.Text))
		case "FNV-1a 64bit":
			view.SetText(hash.FNV64String(in.Text))
		case "MD5":
			view.SetText(hash.MD5(in.Text))
		case "SHA1":
			view.SetText(hash.SHA1(in.Text))
		case "SHA256":
			view.SetText(hash.SHA256(in.Text))
		case "SHA384":
			view.SetText(hash.SHA384(in.Text))
		case "SHA512":
			view.SetText(hash.SHA512(in.Text))
		default:
			view.SetText("Choose your Hash Type!!")

		}
		HistoryFs.HistoryFH(in, view, current, History, INPUT, OUTPUT, CIPHER)
		History.Refresh()
	})

	encode.Resize(fyne.NewSize(230, 47))
	encode.Move(fyne.NewPos(35, 495))

	return container.NewWithoutLayout(back, in, view, drop, encode)
}
