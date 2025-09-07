package Binary

import (
	"fmt"
	"strconv"
	"strings"
)

func FtexttoBIN(text string) string {
	encoded := ""
	for _, char := range []byte(text) {
		encoded += fmt.Sprintf("%08b", char)
		encoded = encoded + " "
	}
	return encoded
}

func FBINtotext(text string) string {
	encoded := strings.Fields(text)
	var builder strings.Builder
	for _, Here := range encoded {
		val, err := strconv.ParseUint(Here, 2, 8)
		if err != nil {
			builder.WriteString("�")
		}
		builder.WriteByte(byte(val))
	}
	return builder.String()
}
