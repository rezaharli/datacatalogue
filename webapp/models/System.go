package models

import "time"

type System struct {
	ID                int
	System_Name       string
	Display_Name      string
	ITAM_ID           int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewSystemModel() *System {
	return new(System)
}

func (m *System) TableName() string {
	return "Tbl_System"
}
