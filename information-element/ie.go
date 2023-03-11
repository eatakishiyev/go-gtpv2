package information_element

type InformationElement interface {
	Decode(data []byte)
	Encode() []byte
	InformationElementType() IEType
}
