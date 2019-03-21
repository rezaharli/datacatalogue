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
	Data_Length           string
	Example               string
	Derived               int
	Derivation_Logic      string
	Status                int
	Alias_Name            string
	CDE                   int
	Sourced_from_Upstream int
	System_Checks         string
	Imm_Prec_System_ID    int
	Imm_Prec_System_SLA   int
	Imm_Prec_System_OLA   int
	Imm_Succ_System_ID    int
	Imm_Succ_System_SLA   int
	Imm_Succ_System_OLA   int
	Data_SLA_Signed       int
	Golden_Source         int
	DQ_Standards          string
	Threshold             int
	DPO_DQ_Standards      string
	DPO_Threshold         int
	DDO_DQ_Standards      string
	DDO_Threshold         int
	PII_Flag              int
	Record_Category       string
}

func NewMDColumnModel() *MDColumn {
	return new(MDColumn)
}

func (m *MDColumn) TableName() string {
	return "Tbl_MD_Column"
}
