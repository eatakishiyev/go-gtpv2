package test

import (
	"bytes"
	information_element "go-gtp-v2/information-element"
	"go-gtp-v2/information-element/constants"
	"testing"
)

func FuzzApnRestrictionEncode(f *testing.F) {
	f.Add(0, []byte{0x00})
	f.Add(1, []byte{0x01})
	f.Add(2, []byte{0x02})
	f.Add(3, []byte{0x03})
	f.Add(4, []byte{0x04})
	f.Fuzz(func(t *testing.T, restrictions int, value []byte) {
		apnRestriction := information_element.ApnRestriction{
			Restriction: constants.ApnRestrictions(restrictions),
		}

		if got, want := apnRestriction.Encode(), value; bytes.Equal(got, want) == false {
			t.Fatalf("Test failed. Got %q , want  %q", got, want)
		}
	})
}
