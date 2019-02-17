package models

type SubCategory struct {
	ID          int
	Name        string
	Type        string
	Category_ID Category
}
