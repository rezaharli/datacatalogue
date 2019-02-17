package models

type MDResource struct {
	ID          int
	Name        string
	Type        string
	Description string
}

func NewMDResource() *MDResource {
	return new(MDResource)
}

func (m *MDResource) TableName() string {
	return "Tbl_MD_Resource"
}
