package models

import "time"

type LinkRolePeople struct {
	ID                int
	Role_ID           int
	People_ID         int
	Object_Type       string
	Object_ID         int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
}

func NewLinkRolePeopleModel() *LinkRolePeople {
	return new(LinkRolePeople)
}

func (m *LinkRolePeople) TableName() string {
	return "tbl_link_role_people"
}
