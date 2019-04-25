package models

import "time"

type LinkColumnGoldenSource struct {
	ID                      int
	Column_ID               int
	Golden_Source_Column_ID int
	Created_DateTime        time.Time
	Modified_DateTime       time.Time
}

func NewLinkColumnGoldenSourceModel() *LinkColumnGoldenSource {
	return new(LinkColumnGoldenSource)
}

func (m *LinkColumnGoldenSource) TableName() string {
	return "Tbl_Link_Column_Golden_Source"
}
