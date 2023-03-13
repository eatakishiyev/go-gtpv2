package information_element

import (
	"go-gtp-v2/information-element/constants"
	"math/big"
)

type BearerLevelQos struct {
	PreEmptionCapability         byte
	PriorityLevel                byte
	PreEmptionVulnerability      byte
	Qci                          uint8
	MaxBitRateForUplink          uint64
	MaxBitRateForDownlink        uint64
	GuaranteedBitRateForUplink   uint64
	GuaranteedBitRateForDownlink uint64
	IeHeader                     InformationElementHeader
}

func (b *BearerLevelQos) Decode(data []byte) {
	octet := data[0] & 0b11111111
	b.PreEmptionCapability = (octet & 0b01000000) >> 6
	b.PriorityLevel = (octet & 0b00111100) >> 2
	b.PreEmptionVulnerability = octet & 0b00000001

	b.Qci = data[1] & 0b11111111
	b.MaxBitRateForUplink = big.NewInt(0).SetBytes(data[2:7]).Uint64()
	b.MaxBitRateForDownlink = big.NewInt(0).SetBytes(data[7:12]).Uint64()
	b.GuaranteedBitRateForUplink = big.NewInt(0).SetBytes(data[12:17]).Uint64()
	b.GuaranteedBitRateForDownlink = big.NewInt(0).SetBytes(data[17:22]).Uint64()
}

func (b *BearerLevelQos) Encode() []byte {
	encoded := make([]byte, 2)
	encoded[0] = 0
	encoded[0] = b.PreEmptionCapability << 6
	encoded[0] = (b.PriorityLevel << 2) | encoded[0]
	encoded[0] = encoded[0] | (b.PreEmptionVulnerability)

	encoded[1] = b.Qci
	encoded = append(encoded, Encode5ByteNumber(b.MaxBitRateForUplink)...)
	encoded = append(encoded, Encode5ByteNumber(b.MaxBitRateForDownlink)...)
	encoded = append(encoded, Encode5ByteNumber(b.GuaranteedBitRateForUplink)...)
	encoded = append(encoded, Encode5ByteNumber(b.GuaranteedBitRateForDownlink)...)
	return encoded
}

func (b *BearerLevelQos) GetType() constants.IEType {
	return constants.BearerLevelQos
}

func (b *BearerLevelQos) GetHeader() *InformationElementHeader {
	return &b.IeHeader
}
