package models

import "time"

type CDE struct {
	ID                int
	Name              string
	CRM_ID            int
	CDE_Rationale     string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewCDEModel() *CDE {
	return new(CDE)
}

func (m *CDE) TableName() string {
	return "Tbl_CDE"
}
