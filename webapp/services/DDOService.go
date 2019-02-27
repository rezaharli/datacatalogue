package services

import (
	"github.com/eaciit/toolkit"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type DDOService struct {
}

func NewDDOService() *DDOService {
	ret := new(DDOService)
	return ret
}

func (s *DDOService) GetLeftTable(sortKey, sortOrder string, skip, take int, filter toolkit.M) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := "SELECT tc.id, tc.name as domain, tsc.name as sub_domain FROM tbl_category tc JOIN tbl_subcategory tsc ON tsc.category_id = tc.id"
	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewCategoryModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	// err := h.NewDBcmd().GetAll(h.GetAllParam{
	// 	Filter:      filter,
	// 	SortKey:     sortKey,
	// 	SortOrder:   sortOrder,
	// 	Skip:        skip,
	// 	Take:        take,
	// 	TableName:   m.NewCategoryModel().TableName(),
	// 	ResultRows:  &resultRows,
	// 	ResultTotal: &resultTotal,
	// })

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DDOService) GetRightTable(systemID int) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT tsc.category_id as tscid, tbt.id, tbt.BT_Name, tbt.Description, tbt.CDE 
		FROM
			tbl_category tc
			JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
			JOIN tbl_link_subcategory_people tlscp ON tlscp.subcategory_id = tsc.id
			JOIN Tbl_People	tp on tlscp.people_id = tp.id
			join tbl_business_term tbt on tbt.parent_id = tsc.id
			join tbl_md_column tmc on tmc.business_term_id = tbt.id
			join tbl_md_table tmt on tmc.table_id = tmt.id
			join tbl_md_resource tmr on tmt.resource_id = tmr.id
			join tbl_system ts on tmr.system_id = ts.id
			join tbl_policy tpo on tbt.policy_id = tpo.id
			join tbl_ds_process_detail tdpd on tdpd.business_term_id = tbt.id
			join tbl_ds_processes tdp on tdpd.process_id = tdp.id
		WHERE tsc.category_id = ` + toolkit.ToString(systemID)
	toolkit.Println(q)
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

func (s *DDOService) GetDetails(leftParam int, rightParam int) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT 
			tc.name as data_domain,
			tsc.name as sub_domain,
			tlscp.people_id as sub_domain_owner,
			tp.bank_id,
			tbt.bt_name,
			tbt.description as business_term_description,
			tmc.alias_name,
			tbt.cde,
			tbt.dq_standards,
			tbt.policy_guidance,
			tbt.mandatory,
			tbt.golden_source_system_id,
			ts.itam_id,
			tmt.name as golden_source_table_name,
			tmc.name as golden_source_column_name,
			tpo.info_asset_name,
			tpo.description as info_asset_desc,
			tpo.confidentiality,
			tpo.integrity,
			tpo.availability,
			tpo.overall_cia_rating,
			tmc.record_category,
			tmc.pii_flag,
			tdp.name as downstream_process_name,
			tdp.owner_name as downstream_process_owner,
			ts.system_name,
			tmc.name as column_name,
			tmc.ddo_threshold
		FROM
			tbl_category tc
			JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
			JOIN tbl_link_subcategory_people tlscp ON tlscp.subcategory_id = tsc.id
			JOIN Tbl_People	tp on tlscp.people_id = tp.id
			join tbl_business_term tbt on tbt.parent_id = tsc.id
			join tbl_md_column tmc on tmc.business_term_id = tbt.id
			join tbl_md_table tmt on tmc.table_id = tmt.id
			join tbl_md_resource tmr on tmt.resource_id = tmr.id
			join tbl_system ts on tmr.system_id = ts.id
			join tbl_policy tpo on tbt.policy_id = tpo.id
			join tbl_ds_process_detail tdpd on tdpd.business_term_id = tbt.id
			join tbl_ds_processes tdp on tdpd.process_id = tdp.id
		WHERE
			tc.id = ` + toolkit.ToString(leftParam) + ` and tbt.id = ` + toolkit.ToString(rightParam)
	toolkit.Println(q)
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
