package models

import "time"

type EdmpPeople struct {
	ID                int
	First_Name        string
	Last_Name         string
	Bank_ID           string
	Email_ID          string
	Function          string
	Org_Unit          string
	Status            string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
}

func NewEdmpPeopleModel() *EdmpPeople {
	return new(EdmpPeople)
}

func (m *EdmpPeople) TableName() string {
	return "TBL_EDMP_PEOPLE"
}
