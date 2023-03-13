package information_element

import (
	"bytes"
	"go-gtp-v2/information-element/constants"
	"io"
	"strings"
)

type Apn struct {
	Value    string
	IeHeader InformationElementHeader
}

func (apn *Apn) Decode(data []byte) {
	reader := bytes.NewReader(data)
	var apnParts []string
	for {
		length, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		apnData := make([]byte, length)
		reader.Read(apnData)
		apnParts = append(apnParts, string(apnData))
	}

	apn.Value = strings.Join(apnParts, ".")
}

func (apn *Apn) Encode() []byte {
	var encodedData []byte
	apnParts := strings.Split(apn.Value, ".")
	for _, apnPart := range apnParts {
		encodedData = append(encodedData, byte(len(apnPart)))
		encodedData = append(encodedData, apnPart...)
	}
	return encodedData
}

func (apn *Apn) GetType() constants.IEType {
	return constants.Apn
}

func (apn *Apn) GetHeader() *InformationElementHeader {
	return &apn.IeHeader
}
