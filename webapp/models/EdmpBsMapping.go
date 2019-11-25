package models

type EdmpBsMapping struct {
	BUSINESS_SEGMENT_OLD string
	BUSINESS_SEGMENT_NEW string
}

func NewEdmpBsMappingModel() *EdmpBsMapping {
	return new(EdmpBsMapping)
}

func (m *EdmpBsMapping) TableName() string {
	return "TBL_EDMP_BS_MAPPING"
}
