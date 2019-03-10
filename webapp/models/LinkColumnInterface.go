package models

type LinkColumnInterface struct {
	ID                  int
	Column_ID           int
	Imm_Prec_System_ID  int
	Imm_Prec_System_SLA int
	Imm_Prec_System_OLA int
	Imm_Succ_System_ID  int
	Imm_Succ_System_SLA int
	Imm_Succ_System_OLA int
}

func NewLinkColumnInterfaceModel() *LinkColumnInterface {
	return new(LinkColumnInterface)
}

func (m *LinkColumnInterface) TableName() string {
	return "Tbl_Link_Column_Interface"
}
