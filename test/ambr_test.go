package test

import (
	"bytes"
	"fmt"
	information_element "go-gtp-v2/information-element"
	"testing"
)

func TestAmbrDecodeTest(t *testing.T) {
	ambr := information_element.Ambr{}
	data := []byte{0x00, 0x11, 0x8c, 0x30, 0x00, 0x11, 0x8c, 0x30}
	fmt.Printf("first octets = %v\n", data[0:4])
	fmt.Printf("first octets = %v\n", data[4:])
	ambr.Decode(data)
	if ambr.AmbrDownlink != 1150000 || ambr.AmbrUplink != 1150000 {
		t.Fatalf("Got AmbrDownlink = %d, AmbrUplink = %d, want AmbrDownlink = 1150000,AmbrUplink = 1150000",
			ambr.AmbrDownlink, ambr.AmbrUplink)
	}
}

func TestAmbrEncodeTest(t *testing.T) {
	ambr := information_element.Ambr{
		AmbrUplink:   1150000,
		AmbrDownlink: 1150000,
	}

	var got = ambr.Encode()
	var want = []byte{0x00, 0x11, 0x8c, 0x30, 0x00, 0x11, 0x8c, 0x30}

	if !bytes.Equal(want, got) {
		t.Fatalf("Got %q, want %q", got, want)
	}
}
