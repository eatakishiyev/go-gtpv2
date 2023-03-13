package information_element

import (
	"encoding/binary"
	"go-gtp-v2/information-element/constants"
)

type Ambr struct {
	AmbrUplink   uint32
	AmbrDownlink uint32
	IeHeader     InformationElementHeader
}

func (ambr *Ambr) Decode(data []byte) {
	ambr.AmbrUplink = binary.BigEndian.Uint32(data[0:4])
	ambr.AmbrDownlink = binary.BigEndian.Uint32(data[4:])
}

func (ambr *Ambr) Encode() []byte {
	var encodedData []byte
	tmpBuff := make([]byte, 4)
	binary.BigEndian.PutUint32(tmpBuff, ambr.AmbrUplink)
	encodedData = append(encodedData, tmpBuff...)
	binary.BigEndian.PutUint32(tmpBuff, ambr.AmbrDownlink)
	encodedData = append(encodedData, tmpBuff...)
	return encodedData
}

func (ambr *Ambr) GetType() constants.IEType {
	return constants.AggregateMaximumBitRate
}

func (ambr *Ambr) GetHeader() *InformationElementHeader {
	return &ambr.IeHeader
}
