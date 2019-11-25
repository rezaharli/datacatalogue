package models

type EdmpHeader struct {
	ID                     int
	COUNTRY                string
	BUSINESS_SEGMENT_NEW   string
	EDM_SOURCE_SYSTEM_NAME string
	CLUSTER_NAME           string
	TIER                   string
	ITAM                   int
}

func NewEdmpHeaderModel() *EdmpHeader {
	return new(EdmpHeader)
}

func (m *EdmpHeader) TableName() string {
	return "TBL_EDMP_HEADER"
}
