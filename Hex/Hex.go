package Hex

import "encoding/hex"

func Ftexttohex(text string) string {
	encoded := hex.EncodeToString([]byte(text))
	if encoded == "" {
		return "Input Is Empty!"

	} else {
		return encoded
	}
}

func Fhextotext(text string) string {
	decoded, err := hex.DecodeString(text)
	if text == "" {
		return "Input Is Empty!"
	} else if err != nil {
		return "An Error Occured " + err.Error()
	} else {
		return string(decoded)
	}
}
