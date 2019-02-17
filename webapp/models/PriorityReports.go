package models

type PriorityReports struct {
	ID               int
	Name             string
	Owner_ID         People
	Sub_Risk_Type_ID SubCategory
	Rationale        string
}
