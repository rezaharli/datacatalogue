package models

import "time"

type MDTable struct {
	ID                int
	Resource_ID       int
	Schema_Name       string
	Name              string
	UUID              string
	Type              string
	Description       string
	Business_Term_ID  int
	Status            int
	Record_Category   string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
}

func NewMDTableModel() *MDTable {
	return new(MDTable)
}

func (m *MDTable) TableName() string {
	return "Tbl_MD_Table"
}
