package information_element

import "encoding/binary"

type Ambr struct {
	AmbrUplink   uint32
	AmbrDownlink uint32
}

func (instance *Ambr) Decode(data []byte) {
	instance.AmbrUplink = binary.BigEndian.Uint32(data[0:4])
	instance.AmbrDownlink = binary.BigEndian.Uint32(data[4:])
}

func (instance *Ambr) Encode() []byte {
	var encodedData []byte
	tmpBuff := make([]byte, 4)
	binary.BigEndian.PutUint32(tmpBuff, instance.AmbrUplink)
	encodedData = append(encodedData, tmpBuff...)
	binary.BigEndian.PutUint32(tmpBuff, instance.AmbrDownlink)
	encodedData = append(encodedData, tmpBuff...)
	return encodedData
}

func (instance *Ambr) InformationElementType() IEType {
	return AggregateMaximumBitRate
}
