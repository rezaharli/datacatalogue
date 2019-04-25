package models

import "time"

type LinkColumnBusinessTerm struct {
	ID                int
	Column_ID         int
	Business_Term_ID  int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
}

func NewLinkColumnBusinessTermModel() *LinkColumnBusinessTerm {
	return new(LinkColumnBusinessTerm)
}

func (m *LinkColumnBusinessTerm) TableName() string {
	return "Tbl_Link_Column_Business_Term"
}
