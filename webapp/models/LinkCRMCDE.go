package models

import "time"

type LinkCRMCDE struct {
	ID                int
	CRM_ID            int
	CDE_ID            int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
}

func NewLinkCRMCDEModel() *LinkCRMCDE {
	return new(LinkCRMCDE)
}

func (m *LinkCRMCDE) TableName() string {
	return "tbl_link_crm_cde"
}
