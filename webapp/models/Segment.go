package models

import "time"

type Segment struct {
	ID                int
	Name              string
	Subdomain_ID      int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewSegmentModel() *Segment {
	return new(Segment)
}

func (m *Segment) TableName() string {
	return "Tbl_Segment"
}
