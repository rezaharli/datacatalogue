package models

import "time"

type LinkPolicySystem struct {
	ID                int
	System_ID         int
	Policy_ID         int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
}

func NewLinkPolicySystemModel() *LinkPolicySystem {
	return new(LinkPolicySystem)
}

func (m *LinkPolicySystem) TableName() string {
	return "Tbl_Link_Policy_System"
}
