package models

type DSProcesses struct {
	ID         int
	Name       string
	Owner_ID   int
	Owner_Name string
}

func NewDSProcessesModel() *DSProcesses {
	return new(DSProcesses)
}

func (m *DSProcesses) TableName() string {
	return "Tbl_DS_Processes"
}
