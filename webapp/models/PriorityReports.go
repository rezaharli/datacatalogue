package models

import "time"

type PriorityReports struct {
	ID                int
	Name              string
	Owner_ID          int
	Lead_ID           int
	Sub_Risk_Type_ID  int
	Rationale         string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewPriorityReportsModel() *PriorityReports {
	return new(PriorityReports)
}

func (m *PriorityReports) TableName() string {
	return "Tbl_Priority_Reports"
}
