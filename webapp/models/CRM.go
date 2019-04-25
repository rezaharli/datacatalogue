package models

import "time"

type CRM struct {
	ID                int
	Name              string
	Prority_Report_ID int
	CRM_Rationale     string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewCRMModel() *CRM {
	return new(CRM)
}

func (m *CRM) TableName() string {
	return "Tbl_CRM"
}
