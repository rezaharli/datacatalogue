package models

import "time"

type BusinessTerm struct {
	ID                         int
	BT_Name                    string
	Parent_ID                  int
	Description                string
	Mandatory                  int
	Policy_ID                  int
	Policy_Guidance            string
	DQ_Standards               string
	Threshold                  int
	Remarks                    string
	Golden_Source_System_ID    int
	Golden_Source_ITAM_ID      int
	Golden_Source_TableName_ID int
	Golden_Source_Column_ID    int
	Target_Golden_Source_ID    int
	DDO_DQ_Standards           string
	DDO_Threshold              int
	Created_DateTime           time.Time
	Modified_DateTime          time.Time
	Status                     int
}

func NewBusinessTermModel() *BusinessTerm {
	return new(BusinessTerm)
}

func (m *BusinessTerm) TableName() string {
	return "Tbl_Business_Term"
}
