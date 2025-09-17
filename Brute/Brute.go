package Brute

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/getevo/hash"
)

func Bruter(myWindow fyne.Window) *fyne.Container {
	var current string
	var path string
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

	BruteIT := widget.NewButton("Brute It!!", func() {
		go func() {
			var thehash string = IN.Text
			file, err := os.Open(path)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)

			switch current {
			case "CRC32":

				for scanner.Scan() {
					line := scanner.Text()
					hashed := hash.CRC32String(line)
					if hashed == thehash {
						OUT.SetText("CRACKED: " + line)
						break
					} else {
						OUT.SetText("Couldn't Crack :(")
					}
				}
			case "CRC64":

				for scanner.Scan() {
					line := scanner.Text()
					hashed := hash.CRC64String(line)
					if hashed == thehash {
						OUT.SetText("CRACKED: " + line)
						break
					} else {
						OUT.SetText("Couldn't Crack :(")
					}
				}

			case "FNV-1a 32bit":

				for scanner.Scan() {
					line := scanner.Text()
					hashed := hash.FNV32String(line)

					if hashed == thehash {
						OUT.SetText("CRACKED: " + line)
						break
					} else {
						OUT.SetText("Couldn't Crack :(")
					}
				}
			case "FNV-1a 64bit":

				for scanner.Scan() {
					line := scanner.Text()
					hashed := hash.FNV64String(line)
					if hashed == thehash {
						OUT.SetText("CRACKED: " + line)
						break
					} else {
						OUT.SetText("Couldn't Crack :(")
					}
				}
			case "MD5":
				var counter int
				var counter2 int
				for scanner.Scan() {
					counter++
					counter2++
					line := scanner.Text()
					hashed := hash.MD5(line)
					if counter != 5 {
						if hashed == thehash {
							OUT.SetText("CRACKED: " + line)
							break
						} else {
							OUT.SetText("Couldn't Crack :(")
						}
					} else {
						OUT.SetText("Finshed: " + strconv.Itoa(counter2))
						time.Sleep(2 * time.Second)
						counter = 0
					}
				}
			case "SHA1":

				for scanner.Scan() {
					line := scanner.Text()
					hashed := hash.SHA1(line)
					if hashed == thehash {
						OUT.SetText("CRACKED: " + line)
						break
					} else {
						OUT.SetText("Couldn't Crack :(")
					}
				}
			case "SHA256":

				for scanner.Scan() {
					line := scanner.Text()
					hashed := hash.SHA256(line)
					if hashed == thehash {
						OUT.SetText("CRACKED: " + line)
						break
					} else {
						OUT.SetText("Couldn't Crack :(")
					}
				}
			case "SHA384":

				for scanner.Scan() {
					line := scanner.Text()
					hashed := hash.SHA384(line)
					if hashed == thehash {
						OUT.SetText("CRACKED: " + line)
						break
					} else {
						OUT.SetText("Couldn't Crack :(")
					}
				}
			case "SHA512":

				for scanner.Scan() {
					line := scanner.Text()
					hashed := hash.SHA512(line)
					if hashed == thehash {
						OUT.SetText("CRACKED: " + line)
						break
					} else {
						OUT.SetText("Couldn't Crack :(")
					}
				}
			default:
				OUT.SetText("Choose your Hash Type!!")

			}
		}()
	})
	BruteIT.Resize(fyne.NewSize(231, 40))
	BruteIT.Move(fyne.NewPos(35, 520))

	background2, _ := fyne.LoadResourceFromPath("CryptoGUIBrute.png")
	back2 := canvas.NewImageFromResource(background2)
	back2.FillMode = canvas.ImageFillStretch
	back2.Resize(fyne.NewSize(300, 600))

	return container.NewWithoutLayout(back2, IN, OUT, HASHES, BruteIT)
}
