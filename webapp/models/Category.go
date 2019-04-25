package models

import "time"

type Category struct {
	ID                int
	Name              string
	Type              string
	Created_DateTime  time.Time
	Modified_DateTime time.Time
	Status            int
}

func NewCategoryModel() *Category {
	return new(Category)
}

func (m *Category) TableName() string {
	return "Tbl_Category"
}
