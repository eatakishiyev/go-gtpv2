package information_element

import (
	"go-gtp-v2/information-element/constants"
)

type Cause struct {
	CauseValue uint32
	Flag       uint32
	IeHeader   InformationElementHeader
}

func (c *Cause) Decode(data []byte) {
	c.CauseValue = uint32(data[0] & 0b11111111)
	c.Flag = uint32(data[1] & 0b11111111)
}

func (c *Cause) Encode() []byte {
	data := make([]byte, 2)
	data[0] = byte(c.CauseValue)
	data[1] = byte(c.Flag)

	return data
}

func (c *Cause) GetType() constants.IEType {
	return constants.Cause
}

func (c *Cause) GetHeader() *InformationElementHeader {
	return &c.IeHeader
}
