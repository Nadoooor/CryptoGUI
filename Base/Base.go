package Base

import (
	"encoding/base32"
	"encoding/base64"
)

func F32totext(text string) string {
	decoded, err := base32.StdEncoding.DecodeString(text)
	if text == "" {
		return "Input Is Empty"
	} else if err != nil {
		return "An Error Occured " + err.Error()
	} else {
		return string(decoded)
	}
}

func F64totext(text string) string {
	decoded, err := base64.StdEncoding.DecodeString(text)
	if text == "" {
		return "Input Is Empty"
	} else if err != nil {
		return "An Error Occured " + err.Error()

	} else {
		return string(decoded)
	}
}

func Ftextto32(text string) string {
	encoded := base32.StdEncoding.EncodeToString([]byte(text))
	if text == "" {
		return "Input Is Empty"
	} else {
		return encoded
	}
}

func Ftextto64(text string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	if text == "" {
		return "Input Is Empty"

	} else {
		return encoded
	}
}
