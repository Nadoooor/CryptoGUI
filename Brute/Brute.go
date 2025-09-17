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
	// initalizing the cancel button to cancel the bruteforcing prossess
	ctx, cancel2 := context.WithCancel(context.Background())
	*cancelFunc = cancel2

	// the go routine that will brute force the hash but in the background
	go func() {
		// when i click on the Brute button it should disappear and make the cancel button show
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
		// the scanner that will get the file and read it line by line
		scanner := bufio.NewScanner(file)

		found := false
		// initializing the counters that will count how many words was checked
		var counter int
		var counter2 int
		// the for loop that will brute the hash itself
		for scanner.Scan() {
			select {
			// just in case i press the cancel button it stops the process
			case <-ctx.Done():
				fyne.Do(
					func() {
						// and show this text to tell the user that the process canceled
						OUT.SetText("Cancelled")
					})
				return
			default:
				// adding one to the counters
				counter++
				counter2++
				// assinging the word from the file in the line to this var so i can use it
				line := scanner.Text()
				// just taking the word and hashing it to get the hash value and compare it.
				hashed := hashing(line)
				// if condition to check eather the hashes are matched or not.
				if hashed == thehash {
					fyne.Do(
						func() {
							// if matched it will print in the OUTPUT the cracked password or code
							OUT.SetText("CRACKED: " + line)
							// and will reset the button so the brute button show again
							BruteIT.Show()
							cancle.Hide()
						},
					)
					// set found bool to true
					found = true
					return

				}
				if counter == 500 {
					// adding the counter to the OUTPUT so the user can see how many words are used till now.
					fyne.Do(func() {
						OUT.SetText("Finshed: " + strconv.Itoa(counter2))
					})
					counter = 0
				}

			}
		}
		if !found {
			// so this massage won't appear unless the bruter didn't find the word in the list
			fyne.Do(func() {
				OUT.SetText("Couldn't Crack the Hash :(")
				BruteIT.Show()
				cancle.Hide()
			})
		}

	}()
}

// the window itself
func Bruter(myWindow fyne.Window) *fyne.Container {
	// just some vars and UI components
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
	// here is the file detector when you drag the word list to the window to use
	myWindow.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		for _, uri := range uris {
			// it gets the path of the file to use and print in the OUTPUT that the wordlist was set
			path = uri.Path()
			OUT.SetText("WORDLIST SET TO: " + path)
		}

	})
	// moving the INPUT entry to its place and resizing it
	IN.Resize(fyne.NewSize(231, 40))
	IN.Move(fyne.NewPos(35, 107))
	// the same thing for the OUTPUT
	OUT.Resize(fyne.NewSize(231, 162))
	OUT.Move(fyne.NewPos(35, 325))
	// making the drop down (this time it was really)
	Hashes := []string{"CRC32", "CRC64", "FNV-1a 32bit", "FNV-1a 64bit", "MD5", "SHA1", "SHA256", "SHA384", "SHA512"}
	HASHES := widget.NewSelect(Hashes, func(s string) {
		current = s
	})
	HASHES.PlaceHolder = "Choose your Hash Type!!"
	// moving and resizing the dropdown
	HASHES.Resize(fyne.NewSize(229, 40))
	HASHES.Move(fyne.NewPos(35, 233))
	// making the brute button
	BruteIT = widget.NewButton("Brute It!!", func() {
		// the switch case that will see what i have chosen in the dropdown and use it when i press the button to hash the wordlist
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
	// moving and resizing the button (i've said this a lot will stop ok)
	BruteIT.Resize(fyne.NewSize(231, 40))
	BruteIT.Move(fyne.NewPos(35, 520))
	// the cancel button
	cancle = widget.NewButton("Cancle", func() {
		cancle.Hide()
		BruteIT.Show()
		cancelFunc()
	})
	// movin... won't say it again sorry :-)
	cancle.Resize(fyne.NewSize(231, 40))
	cancle.Move(fyne.NewPos(35, 520))
	background2, _ := fyne.LoadResourceFromPath("CryptoGUIBrute.png")
	back2 := canvas.NewImageFromResource(background2)
	back2.FillMode = canvas.ImageFillStretch
	back2.Resize(fyne.NewSize(300, 600))
	// just return the container that the tab in the main func will use
	return container.NewWithoutLayout(back2, IN, OUT, HASHES, cancle, BruteIT)
}
