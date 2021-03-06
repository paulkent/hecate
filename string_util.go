package main

import (
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

func isASCII(val byte) bool {
	return (val >= 0x20 && val < 0x7f)
}

func isCode(val byte) bool {
	return val == 0x09 || val == 0x0A || val == 0x0D
}

func isPrintable(val byte) bool {
	return isASCII(val) || isCode(val)
}

func drawStringAtPoint(str string, x int, y int, fg termbox.Attribute, bg termbox.Attribute) int {
	x_pos := x
	for _, runeValue := range str {
		termbox.SetCell(x_pos, y, runeValue, fg, bg)
		x_pos++
	}
	return x_pos - x
}

func removeRuneAtIndex(value []byte, index int) []byte {
	var runes []rune
	new_string := make([]byte, utf8.UTFMax*(len(value)+1))
	pos := 0
	for _, runeValue := range string(value) {
		if pos != index {
			runes = append(runes, runeValue)
		}
		pos++
	}
	pos = 0
	for _, runeValue := range runes {
		pos += utf8.EncodeRune(new_string[pos:], runeValue)
	}
	return new_string[0:pos]
}

func insertRuneAtIndex(value []byte, index int, newRuneValue rune) []byte {
	var runes []rune
	new_string := make([]byte, utf8.UTFMax*(len(value)+1))
	pos := 0
	for _, runeValue := range string(value) {
		if pos == index {
			runes = append(runes, newRuneValue)
		}
		runes = append(runes, runeValue)
		pos++
	}
	if index == pos {
		runes = append(runes, newRuneValue)
	}
	pos = 0
	for _, runeValue := range runes {
		pos += utf8.EncodeRune(new_string[pos:], runeValue)
	}
	return new_string[0:pos]
}

func binaryStringForInteger8(value uint8) string {
	new_string := make([]byte, 1)
	new_string[0] = value
	return string(new_string[0:1])
}

func binaryStringForInteger16(value uint16, big_endian bool) string {
	new_string := make([]byte, 2)
	if big_endian {
		new_string[0] = uint8(value >> 8)
		new_string[1] = uint8(value & 0xFF)
	} else {
		new_string[0] = uint8(value & 0xFF)
		new_string[1] = uint8(value >> 8)
	}
	return string(new_string[0:2])
}

func binaryStringForInteger32(value uint32, big_endian bool) string {
	new_string := make([]byte, 4)
	if big_endian {
		new_string[0] = uint8(value >> 24)
		new_string[1] = uint8((value >> 16) & 0xFF)
		new_string[2] = uint8((value >> 8) & 0xFF)
		new_string[3] = uint8(value & 0xFF)
	} else {
		new_string[0] = uint8(value & 0xFF)
		new_string[1] = uint8((value >> 8) & 0xFF)
		new_string[2] = uint8((value >> 16) & 0xFF)
		new_string[3] = uint8(value >> 24)
	}
	return string(new_string[0:4])
}

func binaryStringForInteger64(value uint64, big_endian bool) string {
	new_string := make([]byte, 8)
	if big_endian {
		new_string[0] = uint8((value >> 56) & 0xFF)
		new_string[1] = uint8((value >> 48) & 0xFF)
		new_string[2] = uint8((value >> 40) & 0xFF)
		new_string[3] = uint8((value >> 32) & 0xFF)
		new_string[4] = uint8((value >> 24) & 0xFF)
		new_string[5] = uint8((value >> 16) & 0xFF)
		new_string[6] = uint8((value >> 8) & 0xFF)
		new_string[7] = uint8(value & 0xFF)
	} else {
		new_string[0] = uint8(value & 0xFF)
		new_string[1] = uint8((value >> 8) & 0xFF)
		new_string[2] = uint8((value >> 16) & 0xFF)
		new_string[3] = uint8((value >> 24) & 0xFF)
		new_string[4] = uint8((value >> 32) & 0xFF)
		new_string[5] = uint8((value >> 40) & 0xFF)
		new_string[6] = uint8((value >> 48) & 0xFF)
		new_string[7] = uint8((value >> 56) & 0xFF)
	}
	return string(new_string[0:8])
}
