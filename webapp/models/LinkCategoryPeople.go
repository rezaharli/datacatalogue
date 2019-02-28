package models

type LinkCategoryPeople struct {
	ID          int
	Category_ID int
	People_ID   int
}

func NewLinkCategoryPeopleModel() *LinkCategoryPeople {
	return new(LinkCategoryPeople)
}

func (m *LinkCategoryPeople) TableName() string {
	return "tbl_link_category_people"
}
