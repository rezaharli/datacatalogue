package models

type LinkSubcategoryPeople struct {
	ID             int
	Subcategory_ID int
	People_ID      int
}

func NewLinkSubcategoryPeopleModel() *LinkSubcategoryPeople {
	return new(LinkSubcategoryPeople)
}

func (m *LinkSubcategoryPeople) TableName() string {
	return "tbl_link_subcategory_people"
}
