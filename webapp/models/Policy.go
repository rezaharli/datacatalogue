package models

import "time"

type Policy struct {
	ID                 int
	Info_Asset_Name    string
	Description        string
	Confidentiality    int
	Integrity          int
	Availability       int
	Overall_CIA_Rating int
	Created_DateTime   time.Time
	Modified_DateTime  time.Time
	Status             int
}

func NewPolicyModel() *Policy {
	return new(Policy)
}

func (m *Policy) TableName() string {
	return "Tbl_Policy"
}
