package models

type DSProcessDetail struct {
	ID                        int
	Process_ID                int
	Business_Term_ID          int
	Segment_ID                int
	Imm_Prec_System_ID        int
	Ultimate_Source_System_ID int
}

func NewDSProcessesDetailModel() *DSProcessDetail {
	return new(DSProcessDetail)
}

func (m *DSProcessDetail) TableName() string {
	return "Tbl_DS_Process_Detail"
}
