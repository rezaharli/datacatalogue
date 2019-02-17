package models

type MDTable struct {
	ID                 int
	Resource_ID        int
	Name               string
	UUID               string
	Type               string
	Description        string
	Business_Term_ID   int
	Status             bool
	Imm_Prec_System_ID int
	Imm_Succ_System_ID int
	Golden_Source      bool
	Data_SLA_Signed    bool
	DSC_DQ_Standards   string
	DSC_Threshold      int
	DPO_DQ_Standards   string
	DPO_Threshold      int
	DDO_DQ_Standards   string
	DDO_Threshold      int
}

func NewMDTableModel() *MDTable {
	return new(MDTable)
}

func (m *MDTable) TableName() string {
	return "Tbl_MD_Table"
}
