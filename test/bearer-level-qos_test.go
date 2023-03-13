package test

import (
	"bytes"
	information_element "go-gtp-v2/information-element"
	"testing"
)

func TestBearerLevelQosEncode(t *testing.T) {
	bearerLevelQos := information_element.BearerLevelQos{
		PreEmptionCapability:         0,
		PriorityLevel:                2,
		PreEmptionVulnerability:      0,
		Qci:                          8,
		MaxBitRateForUplink:          1,
		MaxBitRateForDownlink:        1,
		GuaranteedBitRateForUplink:   1,
		GuaranteedBitRateForDownlink: 1,
	}
	want := []byte{0x08, 0x08, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x01}
	got := bearerLevelQos.Encode()
	if bytes.Equal(want, got) == false {
		t.Fatalf("Test failed. Want %d but got %d", want, got)
	}
}

func TestBearerLevelQosDecode(t *testing.T) {
	want := information_element.BearerLevelQos{
		PreEmptionCapability:         0,
		PriorityLevel:                2,
		PreEmptionVulnerability:      0,
		Qci:                          8,
		MaxBitRateForUplink:          1,
		MaxBitRateForDownlink:        1,
		GuaranteedBitRateForUplink:   1,
		GuaranteedBitRateForDownlink: 1,
	}
	example := []byte{0x08, 0x08, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x01}
	bearerLevelQos := information_element.BearerLevelQos{}
	bearerLevelQos.Decode(example)

	if want != bearerLevelQos {
		t.Fatalf("Test failed. Want %o , got %o", want, bearerLevelQos)
	}
}
