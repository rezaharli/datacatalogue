package services

import (
	"github.com/eaciit/toolkit"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type RFOService struct {
}

func NewRFOService() *RFOService {
	ret := new(RFOService)
	return ret
}

func (s *RFOService) GetLeftTable(sortKey, sortOrder string, skip, take int, filter toolkit.M) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT 
			tpr.id, tpr.name, tpr.owner_id 
		FROM 
			Tbl_Priority_Reports tpr
			JOIN tbl_subcategory tsc ON tpr.sub_risk_type_id = tsc.id
			JOIN tbl_category tc ON tsc.category_id = tc.id
			LEFT JOIN tbl_crm tcrm ON tcrm.prority_report_id = tpr.id
			LEFT JOIN tbl_link_crm_cde tlcc ON tlcc.crm_id = tcrm.id
			LEFT JOIN tbl_business_term tbt ON tlcc.cde_id = tbt.id`
	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewCategoryModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *RFOService) GetRightTable(systemID int) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT 
		tpr.id,
		tc.name as Principal_Risk_Type,
		tsc.name as Risk_Sub_Type,
		tpr.rationale as Priority_Report_Rationale,
		tcrm.name as CRM_Name,
		tcrm.crm_rationale as CRM_Rationale,
		tlcc.cde_id as Associated_CDEs,
		tbt.cde_rationale
	FROM 
		Tbl_Priority_Reports tpr
		JOIN tbl_subcategory tsc ON tpr.sub_risk_type_id = tsc.id
		JOIN tbl_category tc ON tsc.category_id = tc.id
		LEFT JOIN tbl_crm tcrm ON tcrm.prority_report_id = tpr.id
		LEFT JOIN tbl_link_crm_cde tlcc ON tlcc.crm_id = tcrm.id
		LEFT JOIN tbl_business_term tbt ON tlcc.cde_id = tbt.id
	WHERE
		tpr.id = ` + toolkit.ToString(systemID)
	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewCategoryModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}
