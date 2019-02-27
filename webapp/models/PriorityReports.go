package models

type PriorityReports struct {
	ID               int
	Name             string
	Owner_ID         int
	Lead_ID          int
	Sub_Risk_Type_ID int
	Rationale        string
}

func NewPriorityReportsModel() *PriorityReports {
	return new(PriorityReports)
}

func (m *PriorityReports) TableName() string {
	return "Tbl_Priority_Reports"
}
