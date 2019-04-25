package models

import "time"

type MDColumnDetails struct {
	ID                    int
	Column_ID             int
	Alias_Name            string
	Description           string
	CDE                   int
	CDE_Rationale         string
	Mandatory             int
	Sourced_from_Upstream int
	Derived               int
	Derivation_Logic      string
	PII_Flag              int
	Status                int
	Data_SLA_Signed       int
	Example               string
	System_Checks         string
	DQ_Standards          string
	Threshold             int
	Record_Category       string
	Created_DateTime      time.Time
	Modified_DateTime     time.Time
}

func NewMDColumnDetailsModel() *MDColumnDetails {
	return new(MDColumnDetails)
}

func (m *MDColumnDetails) TableName() string {
	return "Tbl_MD_Column_Details"
}
