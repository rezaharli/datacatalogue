package models

type Policy struct {
	ID                 int
	Info_Asset_Name    string
	Description        string
	Confidentiality    int
	Integrity          int
	Availability       int
	Overall_CIA_Rating int
}

func NewPolicyModel() *Policy {
	return new(Policy)
}

func (m *Policy) TableName() string {
	return "Tbl_Policy"
}
