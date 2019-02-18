package models

type MDColumn struct {
	ID                    int
	Table_ID              int
	Name                  string
	UUID                  string
	Type                  string
	Description           string
	Business_Term_ID      int
	Data_Type             string
	Data_Format           string
	Data_Length           int
	Example               string
	Derived               bool
	Derivation_Logic      string
	Mandatory             bool
	Status                bool
	Alias_Name            string
	CDE                   bool
	Sourced_from_Upstream bool
	System_Checks         string
	Imm_Prec_System_ID    int
	Imm_Succ_System_ID    int
	Data_SLA_Signed       bool
	Golden_Source         bool
	DQ_Standards          string
	Threshold             int
	DPO_DQ_Standards      string
	DPO_Threshold         int
	DDO_DQ_Standards      string
	DDO_Threshold         int
}

func NewMDColumnModel() *MDColumn {
	return new(MDColumn)
}

func (m *MDColumn) TableName() string {
	return "Tbl_MD_Column"
}
