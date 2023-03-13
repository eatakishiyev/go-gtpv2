package information_element

import "go-gtp-v2/information-element/constants"

type ApnRestriction struct {
	Restriction constants.ApnRestrictions
	IeHeader    InformationElementHeader
}

func (apnRestriction *ApnRestriction) Decode(data []byte) {
	apnRestriction.Restriction = constants.ApnRestrictions(data[0])
}

func (apnRestriction *ApnRestriction) Encode() []byte {
	return []byte{byte(apnRestriction.Restriction)}
}

func (apnRestriction *ApnRestriction) GetType() constants.IEType {
	return constants.ApnRestriction
}

func (apnRestriction *ApnRestriction) GetHeader() *InformationElementHeader {
	return &apnRestriction.IeHeader
}
