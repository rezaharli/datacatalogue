package models

import "time"

type SubCategory struct {
	ID                int
	Name              string
	Type              string
	Category_ID       int
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewSubCategoryModel() *SubCategory {
	return new(SubCategory)
}

func (m *SubCategory) TableName() string {
	return "Tbl_Subcategory"
}
