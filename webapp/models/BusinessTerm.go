package models

type BusinessTerm struct {
	ID                      int
	BT_Name                 string
	Parent_ID               SubCategory
	Description             string
	CDE                     bool
	CDE_Rationale           string
	Status                  bool
	Policy_ID               Policy
	DQ_Standards            string
	Threshold               int
	Golden_Source_ID        System
	Target_Golden_Source_ID System
}
