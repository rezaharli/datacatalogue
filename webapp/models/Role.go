package models

import "time"

type Role struct {
	ID                int
	Role_Name         string
	Role_Type         string
	Role_Description  string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewRoleModel() *Role {
	return new(Role)
}

func (m *Role) TableName() string {
	return "Tbl_Role"
}
