package models

type BusinessTerm struct {
	ID                         int
	BT_Name                    string
	Parent_ID                  int
	Description                string
	CDE                        int
	CDE_Rationale              string
	Mandatory                  int
	Policy_ID                  int
	Policy_Guidance            string
	DQ_Standards               string
	Threshold                  int
	Golden_Source_System_ID    int
	Golden_Source_ITAM_ID      int
	Golden_Source_TableName_ID int
	Golden_Source_Column_ID    int
	Target_Golden_Source_ID    int
	DDO_DQ_Standards           string
	DDO_Threshold              int
}

func NewBusinessTermModel() *BusinessTerm {
	return new(BusinessTerm)
}

func (m *BusinessTerm) TableName() string {
	return "Tbl_Business_Term"
}
