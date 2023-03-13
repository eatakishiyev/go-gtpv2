package constants

type ApnRestrictions int

const (
	NoExistingContextOrRestrictions ApnRestrictions = iota
	Public1
	Public2
	Private1
	Private2
)
