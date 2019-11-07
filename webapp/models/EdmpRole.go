package models

import "time"

type EdmpRole struct {
	ID                int
	Role_Name         string
	Role_Type         string
	Role_Description  string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewEdmpRoleModel() *EdmpRole {
	return new(EdmpRole)
}

func (m *EdmpRole) TableName() string {
	return "TBL_EDMP_ROLE"
}
