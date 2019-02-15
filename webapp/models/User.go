package models

import "time"

type User struct {
	ID       int
	Username int
	Password string
	Name     string
	IsActive bool
	Role     string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserModel() *User {
	return new(User)
}

func (m *User) TableName() string {
	return "sys_user"
}
