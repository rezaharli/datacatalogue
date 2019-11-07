package models

import "time"

type EdmpLinkRolePeople struct {
	ID                int
	Role_ID           int
	People_ID         int
	Object_Type       string
	Object_ID         int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
}

func NewEdmpLinkRolePeopleModel() *EdmpLinkRolePeople {
	return new(EdmpLinkRolePeople)
}

func (m *EdmpLinkRolePeople) TableName() string {
	return "TBL_EDMP_LINK_ROLE_PEOPLE"
}
