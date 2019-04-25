package models

import "time"

type LinkDSProcessSegment struct {
	ID                int
	Process_ID        string
	Segment_Name      string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
}

func NewLinkDSProcessSegmentModel() *LinkDSProcessSegment {
	return new(LinkDSProcessSegment)
}

func (m *LinkDSProcessSegment) TableName() string {
	return "Tbl_Link_DS_Process_Segment"
}
