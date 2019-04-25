package models

import "time"

type DSProcessDetail struct {
	ID                        int
	Process_ID                int
	Column_ID                 int
	Segment_Name              string
	Imm_Prec_System_ID        int
	Ultimate_Source_System_ID int
	USS_CDE_NAME              string
	Created_DateTime          time.Time
	Modified_DateTime         time.Time
}

func NewDSProcessesDetailModel() *DSProcessDetail {
	return new(DSProcessDetail)
}

func (m *DSProcessDetail) TableName() string {
	return "Tbl_DS_Process_Detail"
}
