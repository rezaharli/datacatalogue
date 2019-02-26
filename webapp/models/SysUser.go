package models

type SysUser struct {
	ID       int
	Username int
	Password string
	Email    string
	Name     string
	Status   int
	Role     string

	CreatedAt string
	UpdatedAt string
}

func NewSysUserModel() *SysUser {
	return new(SysUser)
}

func (m *SysUser) TableName() string {
	return "tbl_users"
}
