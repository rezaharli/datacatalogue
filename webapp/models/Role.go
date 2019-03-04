package models

type Role struct {
	ID               int
	Role_Name        string
	Role_Type        string
	Role_Description string
}

func NewRoleModel() *Role {
	return new(Role)
}

func (m *Role) TableName() string {
	return "Tbl_Role"
}
