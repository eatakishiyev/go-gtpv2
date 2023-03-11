package information_element

import (
	"bytes"
	"io"
	"strings"
)

type ApnIE struct {
	Value string
}

func (instance *ApnIE) Decode(data []byte) {
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

	instance.Value = strings.Join(apnParts, ".")
}

func (instance *ApnIE) Encode() []byte {
	var encodedData []byte
	apnParts := strings.Split(instance.Value, ".")
	for _, apnPart := range apnParts {
		encodedData = append(encodedData, byte(len(apnPart)))
		encodedData = append(encodedData, apnPart...)
	}
	return encodedData
}

func (instance *ApnIE) InformationElementType() IEType {
	return Apn
}
