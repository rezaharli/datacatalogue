package models

type SubCategory struct {
	ID          int
	Name        string
	Type        string
	Category_ID int
}

func NewSubCategoryModel() *SubCategory {
	return new(SubCategory)
}

func (m *SubCategory) TableName() string {
	return "Tbl_Subcategory"
}
