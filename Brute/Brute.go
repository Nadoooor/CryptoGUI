package Brute

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
	"github.com/getevo/hash"
)

func compair(cancelFunc *context.CancelFunc, IN *widget.Entry, OUT *widget.Entry, path string, BruteIT *widget.Button, cancle *widget.Button, hashing func(v any) string) {
	ctx, cancel2 := context.WithCancel(context.Background())
	*cancelFunc = cancel2
	go func() {
		fyne.Do(func() {

			BruteIT.Hide()
			cancle.Show()
		})
		var thehash string = IN.Text
		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		found := false
		var counter int
		var counter2 int
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				fyne.Do(
					func() {
						OUT.SetText("Cancelled")
					})
				return
			default:
				counter++
				counter2++
				line := scanner.Text()

				hashed := hashing(line)

				if hashed == thehash {
					fyne.Do(
						func() {
							OUT.SetText("CRACKED: " + line)
							BruteIT.Show()
							cancle.Hide()
						},
					)
					found = true
					return

				}
				if counter == 500 {
					fyne.Do(func() {
						OUT.SetText("Finshed: " + strconv.Itoa(counter2))
					})
					counter = 0
				}

			}
		}
		if !found {
			fyne.Do(func() {
				OUT.SetText("Couldn't Crack the Hash :(")
				BruteIT.Show()
				cancle.Hide()
			})
		}

	}()
}

func Bruter(myWindow fyne.Window) *fyne.Container {

	var current string
	var path string
	var BruteIT *widget.Button
	var cancle *widget.Button
	var cancelFunc context.CancelFunc
	IN := widget.NewMultiLineEntry()
	IN.Wrapping = fyne.TextWrapWord
	IN.SetMinRowsVisible(3)
	OUT := widget.NewMultiLineEntry()
	OUT.Wrapping = fyne.TextWrapWord
	OUT.SetMinRowsVisible(3)
	myWindow.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		for _, uri := range uris {
			path = uri.Path()
			OUT.SetText("FILE SET TO: " + path)
		}

	})

	IN.Resize(fyne.NewSize(231, 40))
	IN.Move(fyne.NewPos(35, 107))
	OUT.Resize(fyne.NewSize(231, 162))
	OUT.Move(fyne.NewPos(35, 325))
	Hashes := []string{"CRC32", "CRC64", "FNV-1a 32bit", "FNV-1a 64bit", "MD5", "SHA1", "SHA256", "SHA384", "SHA512"}
	HASHES := widget.NewSelect(Hashes, func(s string) {
		current = s
	})
	HASHES.PlaceHolder = "Choose your Hash Type!!"
	HASHES.Resize(fyne.NewSize(229, 40))
	HASHES.Move(fyne.NewPos(35, 233))

	BruteIT = widget.NewButton("Brute It!!", func() {

		switch current {
		case "CRC32":
			compair(&cancelFunc, IN, OUT, path, BruteIT, cancle, hash.CRC32String)
		case "CRC64":
			compair(&cancelFunc, IN, OUT, path, BruteIT, cancle, hash.CRC64String)
		case "FNV-1a 32bit":
			compair(&cancelFunc, IN, OUT, path, BruteIT, cancle, hash.FNV32String)
		case "FNV-1a 64bit":
			compair(&cancelFunc, IN, OUT, path, BruteIT, cancle, hash.FNV64String)
		case "MD5":
			compair(&cancelFunc, IN, OUT, path, BruteIT, cancle, hash.MD5)
		case "SHA1":
			compair(&cancelFunc, IN, OUT, path, BruteIT, cancle, hash.SHA1)
		case "SHA256":
			compair(&cancelFunc, IN, OUT, path, BruteIT, cancle, hash.SHA256)
		case "SHA384":
			compair(&cancelFunc, IN, OUT, path, BruteIT, cancle, hash.SHA384)
		case "SHA512":
			compair(&cancelFunc, IN, OUT, path, BruteIT, cancle, hash.SHA512)
		default:
			OUT.SetText("Choose your Hash Type!!")

		}

	})
	BruteIT.Resize(fyne.NewSize(231, 40))
	BruteIT.Move(fyne.NewPos(35, 520))
	cancle = widget.NewButton("Cancle", func() {
		cancle.Hide()
		BruteIT.Show()
		cancelFunc()
	})
	cancle.Resize(fyne.NewSize(231, 40))
	cancle.Move(fyne.NewPos(35, 520))
	background2, _ := fyne.LoadResourceFromPath("CryptoGUIBrute.png")
	back2 := canvas.NewImageFromResource(background2)
	back2.FillMode = canvas.ImageFillStretch
	back2.Resize(fyne.NewSize(300, 600))

	return container.NewWithoutLayout(back2, IN, OUT, HASHES, cancle, BruteIT)
}
