package main

import (
	"fmt"
)

func main() {
	challengeOne()
	challengeTwo()
}

func challengeOne() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hexDecoded := decodeFromHex([]byte(hexString))
	base64Encoded := encodeToBase64([]byte(hexDecoded))
	fmt.Println("--- Challenge One ---")
	fmt.Printf("String to Base64, response and expected respectively: \n%s\n", base64Encoded)
	fmt.Printf("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t\n\n")
}

func challengeTwo() {
	buffer1 := decodeFromHex([]byte("1c0111001f010100061a024b53535009181c"))
	buffer2 := decodeFromHex([]byte("686974207468652062756c6c277320657965"))
	xordBuffers := xorSameLengthBuffers(buffer1, buffer2)
	fmt.Println("--- Challenge Two ---")
	fmt.Printf("Xord buffers, response and expected respectively: \n%s\n", encodeToHex(xordBuffers))
	fmt.Println("746865206b696420646f6e277420706c6179\n\n")
}
