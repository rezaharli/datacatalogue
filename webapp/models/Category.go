package models

type Category struct {
	ID   int
	Name string
	Type string
}

func NewCategoryModel() *Category {
	return new(Category)
}

func (m *Category) TableName() string {
	return "Tbl_Category"
}
