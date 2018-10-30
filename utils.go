package main

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func encodeToBase64(src []byte) []byte {
	dest := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dest, src)
	return dest
}

func decodeFromHex(src []byte) []byte {
	dest := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dest, src)
	if err != nil {
		log.Fatal(err)
	}
	return dest
}

func encodeToHex(src []byte) []byte {
	dest := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dest, src)
	return dest
}

func xorSameLengthBuffers(buffer1 []byte, buffer2 []byte) []byte {
	resultBuffer := make([]byte, len(buffer1))
	for index := range buffer1 {
		resultBuffer[index] = buffer1[index] ^ buffer2[index]
	}
	return resultBuffer
}

// byte is the same as uint8 (0 - 255)
func xorBuffer(buffer []byte, value uint8) []byte {
	xoredBuffer := make([]byte, len(buffer))
	for pos := range buffer {
		xoredBuffer[pos] = buffer[pos] ^ value
	}
	return xoredBuffer
}

func scorePhrase(buffer []byte) int {
	sum := 0
	for _, letter := range buffer {
		sum += scoreByLetter[string(letter)]
	}
	return sum
}

var scoreByLetter = map[string]int{
	" ": 27,
	"E": 26,
	"T": 25,
	"A": 24,
	"O": 23,
	"I": 22,
	"N": 21,
	"S": 20,
	"R": 19,
	"H": 18,
	"D": 17,
	"L": 16,
	"U": 15,
	"C": 14,
	"M": 13,
	"F": 12,
	"Y": 11,
	"W": 10,
	"G": 9,
	"P": 8,
	"B": 7,
	"V": 6,
	"K": 5,
	"X": 4,
	"Q": 3,
	"J": 2,
	"Z": 1,
	"e": 26,
	"t": 25,
	"a": 24,
	"o": 23,
	"i": 22,
	"n": 21,
	"s": 20,
	"r": 19,
	"h": 18,
	"d": 17,
	"l": 16,
	"u": 15,
	"c": 14,
	"m": 13,
	"f": 12,
	"y": 11,
	"w": 10,
	"g": 9,
	"p": 8,
	"b": 7,
	"v": 6,
	"k": 5,
	"x": 4,
	"q": 3,
	"j": 2,
	"z": 1,
}
