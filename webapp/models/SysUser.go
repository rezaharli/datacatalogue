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

type UserUsage struct {
	ID          int
	Username    string
	Fullname    string
	Role        string
	Module      string
	Action      string
	Description string
	Time        string
	ResourceURL string
}

func NewUserUsageModel() *UserUsage {
	return new(UserUsage)
}

func (m *UserUsage) TableName() string {
	return "tbl_usage"
}
