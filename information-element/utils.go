package information_element

import (
	"encoding/binary"
)

func EncodeInformationElement(element InformationElement) []byte {
	elementData := element.Encode()
	encodedData := make([]byte, 4)
	encodedData[0] = byte(element.GetType())
	binary.BigEndian.PutUint16(encodedData[1:], uint16(len(elementData)))
	flags := element.GetHeader().CrFlag
	flags = flags | (element.GetHeader().CrFlag & 0b00001111)
	encodedData[3] = flags
	encodedData = append(encodedData, elementData...)
	return encodedData
}

func Encode5ByteNumber(number uint64) []byte {
	return []byte{uint8(number >> 32), uint8(number >> 24), uint8(number >> 16), uint8(number >> 8), uint8(number)}
}

func Decode5ByteNumber(data []byte) uint64 {
	return uint64(data[0])<<32 | uint64(data[1])<<24 | uint64(data[2])<<16 | uint64(data[3])<<8 | uint64(data[4])

}
