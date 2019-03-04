package models

type LinkRolePeople struct {
	ID          int
	Role_ID     int
	People_ID   int
	Object_Type string
	Object_ID   int
}

func NewLinkRolePeopleModel() *LinkRolePeople {
	return new(LinkRolePeople)
}

func (m *LinkRolePeople) TableName() string {
	return "tbl_link_role_people"
}
