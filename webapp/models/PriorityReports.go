package models

type PriorityReports struct {
	ID               int
	Name             string
	Owner_ID         int
	Lead_ID          int
	Sub_Risk_Type_ID int
	Rationale        string
}
