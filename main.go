package main

import (
	"fmt"
)

func main() {
	challengeOne()
	challengeTwo()
	challengeThree()
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
	fmt.Printf("746865206b696420646f6e277420706c6179\n\n")
}

func challengeThree() {
	encodedPhrase := decodeFromHex([]byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	maxScore := 0
	decodedMaxScorePhrase := make([]byte, len(encodedPhrase))
	for i := 0; i <= 255; i++ {
		decodedPhrase := xorBuffer(encodedPhrase, uint8(i))
		score := scorePhrase(decodedPhrase)
		if score > maxScore {
			maxScore = score
			decodedMaxScorePhrase = decodedPhrase
		}
	}
	fmt.Println("--- Challenge Three ---")
	fmt.Printf("Xor cipher decoded phrase: \n%s\n\n", decodedMaxScorePhrase)
}
