package models

import "time"

type Edmp struct {
	ID                          int
	COUNTRY                     string
	BUSINESS_SEGMENT            string
	EDM_SOURCE_SYSTEM_NAME      string
	CLUSTER_NAME                string
	TIER                        string
	ITAM                        int
	EDM_GOLDEN_SYSTEM_NAME      string
	DATABASE_NAME               string
	TABLE_NAME                  string
	TABLE_DESCRIPTION           string
	COLUMN_NAME                 string
	COLUMN_DESCRIPTION          string
	DATA_TYPE                   string
	DATA_LENGTH                 string
	NULLABLE                    string
	PRIMARY_KEY                 string
	PII                         string
	CERTIFIED                   string
	DATA_XRAY                   string
	DATA_LINEAGE                string
	BUSINESS_TERM               string
	BUSINESS_DESCRIPTION        string
	CDE                         string
	DETERMINES_CLIENT_LOCATION  string
	DETERMINES_ACCOUNT          string
	PRODUCT_CATEGORY            string
	CONSUMING_APPLICATION       string
	CONSUMING_APPLICATION_ITAM  int
	CONSUMING_APPLICATION_OWNER string
	CONSUMER_DESCRIPTION        string
	TECH_CONTACT                int
	BUSINESS_OWNERSHIP          string
	ACCESS_ROLE                 string
	ROLE_DESCRIPTION            string
	CONSUMING_TECH_METADATA     string
	CREATED_DATETIME            time.Time
	MODIFIED_DATETIME           time.Time
}

func NewEdmpModel() *Edmp {
	return new(Edmp)
}

func (m *Edmp) TableName() string {
	return "TBL_EDMP"
}
