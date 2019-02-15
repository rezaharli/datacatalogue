package models

import "time"

type SysUser struct {
	ID       int
	Username int
	Password string
	Name     string
	IsActive bool
	Role     SysRole

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSysUserModel() *SysUser {
	return new(SysUser)
}

func (m *SysUser) TableName() string {
	return "sys_user"
}
