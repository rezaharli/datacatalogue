package models

import "time"

type DSProcesses struct {
	ID                int
	Name              string
	Description       string
	Process_ID        string
	Owner_ID          int
	Owner_Name        string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewDSProcessesModel() *DSProcesses {
	return new(DSProcesses)
}

func (m *DSProcesses) TableName() string {
	return "Tbl_DS_Processes"
}
