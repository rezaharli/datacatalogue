package models

import "time"

type MDColumn struct {
	ID                int
	Table_ID          int
	Name              string
	UUID              string
	Type              string
	Data_Type         string
	Data_Format       string
	Data_Length       string
	Primary_Key       int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
}

func NewMDColumnModel() *MDColumn {
	return new(MDColumn)
}

func (m *MDColumn) TableName() string {
	return "Tbl_MD_Column"
}
