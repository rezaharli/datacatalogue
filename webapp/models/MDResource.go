package models

import "time"

type MDResource struct {
	ID                int
	Name              string
	Type              string
	Description       string
	System_ID         int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewMDResource() *MDResource {
	return new(MDResource)
}

func (m *MDResource) TableName() string {
	return "Tbl_MD_Resource"
}
