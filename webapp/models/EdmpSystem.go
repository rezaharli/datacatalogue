package models

import "time"

type EdmpSystem struct {
	ID                int
	System_Name       string
	Display_Name      string
	ITAM_ID           int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewEdmpSystemModel() *EdmpSystem {
	return new(EdmpSystem)
}

func (m *EdmpSystem) TableName() string {
	return "TBL_EDMP_SYSTEM"
}
