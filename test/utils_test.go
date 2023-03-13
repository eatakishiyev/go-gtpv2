package test

import (
	"bytes"
	information_element "go-gtp-v2/information-element"
	"go-gtp-v2/information-element/constants"
	"testing"
)

func TestInformationElementEncode(t *testing.T) {
	restriction := information_element.ApnRestriction{
		IeHeader: struct {
			CrFlag   byte
			Instance byte
		}{CrFlag: 0, Instance: 0},
		Restriction: constants.NoExistingContextOrRestrictions,
	}
	want := []byte{0x7f, 0x00, 0x01, 0x00, 0x00}
	got := information_element.EncodeInformationElement(&restriction)
	if bytes.Equal(want, got) == false {
		t.Fatalf("Test failed. Want %q , got %q", want, got)
	}
}
