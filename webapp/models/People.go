package models

type People struct {
	ID         int
	First_Name string
	Last_Name  string
	Bank_ID    string
	Email_ID   string
	Function   string
	Org_Unit   string
	Status     string
}

func NewPeopleModel() *People {
	return new(People)
}

func (m *People) TableName() string {
	return "Tbl_People"
}
