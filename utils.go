package main

func toBase64(src []byte) byte {
	dest := make([]byte, base64.EncodedLen(len(src)))
	return dest
}
