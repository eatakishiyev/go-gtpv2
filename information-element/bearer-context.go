package information_element

import (
	"bytes"
	"go-gtp-v2/information-element/constants"
	"io"
)

type BearerContext struct {
	InformationElements []InformationElement
	IeHeader            InformationElementHeader
}

func (b *BearerContext) Decode(data []byte) {
	reader := bytes.NewReader(data)
	for {
		informationElement, err := CreateInformationElement(reader)
		if err == io.EOF {
			break
		}
		b.InformationElements = append(b.InformationElements, *informationElement)
	}
}

func (b *BearerContext) Encode() []byte {
	var encoded []byte
	for _, element := range b.InformationElements {
		tmp := EncodeInformationElement(element)
		encoded = append(encoded, tmp...)
	}
	return encoded
}

func (b *BearerContext) GetType() constants.IEType {
	return constants.BearerContext
}

func (b *BearerContext) GetHeader() *InformationElementHeader {
	return &b.IeHeader
}
