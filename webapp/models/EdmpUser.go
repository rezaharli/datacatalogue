package models

type EdmpUser struct {
	ID        int
	USERNAME  int
	PASSWORD  string
	EMAIL     string
	NAME      string
	STATUS    int
	ROLE      string
	CREATEDAT string
	UPDATEDAT string
}

func NewEdmpUserModel() *EdmpUser {
	return new(EdmpUser)
}

func (m *EdmpUser) TableName() string {
	return "TBL_EDMP_USERS"
}
