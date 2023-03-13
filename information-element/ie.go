package information_element

import "go-gtp-v2/information-element/constants"

type InformationElement interface {
	Decode(data []byte)
	Encode() []byte
	GetType() constants.IEType
	GetHeader() *InformationElementHeader
}
