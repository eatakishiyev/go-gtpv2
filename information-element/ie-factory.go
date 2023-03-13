package information_element

import (
	"bytes"
	"go-gtp-v2/information-element/constants"
)

func CreateInformationElement(reader *bytes.Reader) (*InformationElement, error) {
	firstByte, _ := reader.ReadByte()
	ieType := constants.IEType(firstByte)

	buf := make([]byte, 2)
	reader.Read(buf)
	//length := binary.BigEndian.Uint16(buf)

	var informationElement InformationElement
	switch ieType {
	case constants.BearerContext:
		informationElement = new(BearerContext)
	}

	return &informationElement, nil
}
