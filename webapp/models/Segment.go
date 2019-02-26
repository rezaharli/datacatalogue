package models

type Segment struct {
	ID           int
	Name         string
	Subdomain_ID int
}

func NewSegmentModel() *Segment {
	return new(Segment)
}

func (m *Segment) TableName() string {
	return "Tbl_Segment"
}
