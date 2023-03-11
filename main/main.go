package main

import (
	"bytes"
	"fmt"
	"go-gtp-v2/information-element"
)

func main() {
	var testApn = information_element.ApnIE{
		Value: "nokia.bkc.mnc002.mcc400.gprs",
	}

	exampleData := []byte{0x05, 0x6e, 0x6f, 0x6b, 0x69, 0x61, 0x03, 0x62, 0x6b, 0x63, 0x06, 0x6d, 0x6e, 0x63, 0x30, 0x30, 0x32, 0x06, 0x6d,
		0x63, 0x63, 0x34, 0x30, 0x30, 0x04, 0x67, 0x70, 0x72, 0x73}
	encoded := testApn.Encode()
	if !bytes.Equal(exampleData, encoded) {
		fmt.Printf("apn parameter encoded incorrectly, expecting %d, got %d", exampleData, encoded)
	}
}