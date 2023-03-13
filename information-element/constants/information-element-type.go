package constants

type IEType int

const (
	Imsi                         IEType = 1
	Cause                               = 2
	RestartCounter                      = 3
	Apn                                 = 71
	AggregateMaximumBitRate             = 72
	EpsBearerId                         = 73
	MobileEquipmentIdentity             = 75
	Msisdn                              = 76
	Indication                          = 77
	ProtocolConfigurationOptions        = 78
	PdnAddressAllocation                = 79
	BearerLevelQos                      = 80
	RatType                             = 82
	ServingNetwork                      = 83
	UserLocationInfo                    = 86
	FTeid                               = 87
	BearerContext                       = 93
	ChargingCharacteristics             = 95
	PdnType                             = 99
	TimeZone                            = 114
	ApnRestriction                      = 127
	SelectionMode                       = 128
	UlTimeStamp                         = 170
)
