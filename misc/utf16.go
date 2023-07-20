package misc

import (
	"encoding/binary"
	"unicode/utf16"
)

func StringToUTF16LE(input string) []byte {
	utf16bytes := make([]byte, len(input)*2)

	utf16characters := utf16.Encode([]rune(input))

	for i, character := range utf16characters {
		binary.LittleEndian.PutUint16(utf16bytes[i*2:(i*2)+2], character)
	}

	return utf16bytes
}
