package test

import (
	"bytes"
	information_element "go-gtp-v2/information-element"
	"testing"
)

func TestCauseEncode(t *testing.T) {
	want := []byte{0x10, 0x00}
	cause := information_element.Cause{
		CauseValue: 16,
		Flag:       0,
	}
	got := cause.Encode()
	if bytes.Equal(want, got) == false {
		t.Fatalf("Test failed. Want %d, got %d", want, got)
	}
}

func TestCauseDecode(t *testing.T) {
	example := []byte{0x10, 0x00}
	want := information_element.Cause{
		CauseValue: 16,
		Flag:       0,
	}
	got := information_element.Cause{}
	got.Decode(example)
	if got != want {
		t.Fatalf("Test failed. Want %d, got %d", want, got)
	}
}
