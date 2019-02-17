package models

type System struct {
	ID             int
	System_Name    string
	ITAM_ID        int
	MD_Resource_ID int
}

func NewSystemModel() *System {
	return new(System)
}

func (m *System) TableName() string {
	return "Tbl_System"
}
