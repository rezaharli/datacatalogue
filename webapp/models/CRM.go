package models

type CRM struct {
	ID                int
	Name              string
	Prority_Report_ID int
	CRM_Rationale     string
}

func NewCRMModel() *CRM {
	return new(CRM)
}

func (m *CRM) TableName() string {
	return "Tbl_CRM"
}
