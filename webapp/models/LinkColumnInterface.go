package models

import "time"

type LinkColumnInterface struct {
	ID                        int
	Column_ID                 int
	System_SLA                int
	System_OLA                int
	Imm_Prec_System_ID        int
	Incoming_CDE_Name         string
	Incoming_Derived          int
	Incoming_Derivation_Logic string
	Imm_Prec_Column_ID        int
	Imm_Succ_System_ID        int
	Outgoing_CDE_Name         string
	Outgoing_Derived          int
	Outgoing_Derivation_Logic string
	Imm_Succ_Column_ID        int
	Created_DateTime          time.Time
	Modified_DateTime         time.Time
}

func NewLinkColumnInterfaceModel() *LinkColumnInterface {
	return new(LinkColumnInterface)
}

func (m *LinkColumnInterface) TableName() string {
	return "Tbl_Link_Column_Interface"
}
