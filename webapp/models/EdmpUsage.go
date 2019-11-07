package models

import "time"

type EdmpUsage struct {
	ID          int
	USERNAME    int
	FULLNAME    string
	ROLE        string
	MODULE      string
	DESCRIPTION string
	TIME        time.Time
	RESOURCEURL string
}

func NewEdmpUsageModel() *EdmpUsage {
	return new(EdmpUsage)
}

func (m *EdmpUsage) TableName() string {
	return "TBL_EDMP_USAGE"
}
