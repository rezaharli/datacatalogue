package models

import "time"

type ReferenceLink struct {
	ID                int
	Object_Type       string
	Object_ID         int
	URL               string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
}

func NewReferenceLinkModel() *ReferenceLink {
	return new(ReferenceLink)
}

func (m *ReferenceLink) TableName() string {
	return "Tbl_Reference_Link"
}
