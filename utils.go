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
