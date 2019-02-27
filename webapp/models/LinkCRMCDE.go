package models

type LinkCRMCDE struct {
	ID     int
	CRM_ID int
	CDE_ID int
}

func NewLinkCRMCDEModel() *LinkCRMCDE {
	return new(LinkCRMCDE)
}

func (m *LinkCRMCDE) TableName() string {
	return "tbl_link_crm_cde"
}
