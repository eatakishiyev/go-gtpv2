package test

import (
	"bytes"
	"go-gtp-v2/information-element"
	"testing"
)

const apn = "nokia.bkc.mnc002.mcc400.gprs"

func TestApnEncode(t *testing.T) {
	var testApn = information_element.Apn{
		Value: apn,
	}

	want := []byte{0x05, 0x6e, 0x6f, 0x6b, 0x69, 0x61, 0x03, 0x62, 0x6b, 0x63, 0x06, 0x6d, 0x6e, 0x63, 0x30, 0x30, 0x32, 0x06, 0x6d,
		0x63, 0x63, 0x34, 0x30, 0x30, 0x04, 0x67, 0x70, 0x72, 0x73}
	got := testApn.Encode()
	if !bytes.Equal(want, got) {
		t.Fatalf("apn parameter encoded incorrectly, want %d, got %d", want, got)
	}
}

func TestApnDecode(t *testing.T) {
	var testApn = new(information_element.Apn)
	testApn.Decode([]byte{0x05, 0x6e, 0x6f, 0x6b, 0x69, 0x61, 0x03, 0x62, 0x6b, 0x63, 0x06, 0x6d, 0x6e, 0x63, 0x30, 0x30, 0x32, 0x06, 0x6d,
		0x63, 0x63, 0x34, 0x30, 0x30, 0x04, 0x67, 0x70, 0x72, 0x73})
	if testApn.Value != apn {
		t.Fatalf("apn decoded incorrectly, want %s, got %s", apn, testApn.Value)
	}
}
