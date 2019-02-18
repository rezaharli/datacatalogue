package models

type BusinessTerm struct {
	ID                      int
	BT_Name                 string
	Parent_ID               int
	Description             string
	CDE                     bool
	CDE_Rationale           string
	Status                  bool
	Policy_ID               int
	DQ_Standards            string
	Threshold               int
	Golden_Source_ID        int
	Target_Golden_Source_ID int
}

func NewBusinessTermModel() *BusinessTerm {
	return new(BusinessTerm)
}

func (m *BusinessTerm) TableName() string {
	return "Tbl_Business_Term"
}
