package services

import (
	"github.com/eaciit/toolkit"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type DSCService struct {
}

func NewDSCService() *DSCService {
	ret := new(DSCService)
	return ret
}

func (s *DSCService) GetAllSystem(search string, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
			ts.id,
			ts.system_name,
			ts.itam_id,
			tp.first_name,
			tp.bank_id
		FROM 
			Tbl_System ts
			INNER JOIN Tbl_MD_Resource tmr ON tmr.system_id = ts.id
			INNER JOIN Tbl_MD_Table tmt ON tmt.resource_id = tmr.id
			LEFT JOIN tbl_business_term tbt ON tmt.business_term_id = tbt.id
			LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
			LEFT JOIN tbl_link_subcategory_people tlsp on tlsp.subcategory_id = tsc.id
			LEFT JOIN tbl_people tp on tlsp.people_id = tp.id
			LEFT JOIN Tbl_MD_Column tmc ON tmc.table_id = tmt.id
			LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.business_term_id = tbt.id
			LEFT JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id
			LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
			LEFT JOIN tbl_policy tpol on tbt.policy_id = tpol.id`

	if search != "" {
		q += `
			WHERE
				upper(ts.system_name) LIKE upper('%` + search + `%')
				OR upper(ts.itam_id) LIKE upper('%` + search + `%')
				OR upper(tp.first_name) LIKE upper('%` + search + `%')
				OR upper(tp.bank_id) LIKE upper('%` + search + `%')`
	}

	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewSystemModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  pageNumber,
		RowsPerPage: rowsPerPage,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetTableName(systemID int, search string, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
			tmt.id,
			ts.id as tsid,
			tmt.name as table_name,
			tmc.name as column_name,
			tmc.alias_name,
			tmc.cde
		FROM 
			Tbl_System ts
			INNER JOIN Tbl_MD_Resource tmr ON tmr.system_id = ts.id
			INNER JOIN Tbl_MD_Table tmt ON tmt.resource_id = tmr.id
			LEFT JOIN tbl_business_term tbt ON tmt.business_term_id = tbt.id
			LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
			LEFT JOIN tbl_link_subcategory_people tlsp on tlsp.subcategory_id = tsc.id
			LEFT JOIN tbl_people tp on tlsp.people_id = tp.id
			LEFT JOIN Tbl_MD_Column tmc ON tmc.table_id = tmt.id
			LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.business_term_id = tbt.id
			LEFT JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id
			LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
			LEFT JOIN tbl_policy tpol on tbt.policy_id = tpol.id
		WHERE
			ts.id = ` + toolkit.ToString(systemID)

	if search != "" {
		q += `
			AND
				upper(tmt.name) LIKE upper('%` + search + `%')
				OR upper(tmc.name) LIKE upper('%` + search + `%')
				OR upper(tmc.alias_name) LIKE upper('%` + search + `%')
				OR upper(tmc.cde) LIKE upper('%` + search + `%')`
	}

	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewSystemModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  pageNumber,
		RowsPerPage: rowsPerPage,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetInterfacesRightTable(systemID int, search string, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
			tmt.id,
			ts.id as tsid,
			tmt.name as table_name,
			tmc.imm_prec_system_id,
			tmc.imm_succ_system_id,
			tdp.owner_id
		FROM 
			Tbl_System ts
			INNER JOIN Tbl_MD_Resource tmr ON tmr.system_id = ts.id
			INNER JOIN Tbl_MD_Table tmt ON tmt.resource_id = tmr.id
			LEFT JOIN tbl_business_term tbt ON tmt.business_term_id = tbt.id
			LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
			LEFT JOIN tbl_link_subcategory_people tlsp on tlsp.subcategory_id = tsc.id
			LEFT JOIN tbl_people tp on tlsp.people_id = tp.id
			LEFT JOIN Tbl_MD_Column tmc ON tmc.table_id = tmt.id
			LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.business_term_id = tbt.id
			LEFT JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id
			LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
			LEFT JOIN tbl_policy tpol on tbt.policy_id = tpol.id
		WHERE
			ts.id = ` + toolkit.ToString(systemID)

	if search != "" {
		q += `
			AND
				upper(tmt.name) LIKE upper('%` + search + `%')
				OR upper(tmc.imm_prec_system_id) LIKE upper('%` + search + `%')
				OR upper(tmc.imm_succ_system_id) LIKE upper('%` + search + `%')
				OR upper(tdp.owner_id) LIKE upper('%` + search + `%')`
	}

	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewSystemModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  pageNumber,
		RowsPerPage: rowsPerPage,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetDetails(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	toolkit.Println(payload.GetString("left"), payload.GetString("right"))
	toolkit.Println(payload.GetString("TableName"))
	toolkit.Println(payload.GetString("ColumnName"))
	toolkit.Println(payload.GetString("ScreenLabel"))

	q := `SELECT 
			ts.id,
			ts.system_name,
			ts.itam_id,
			tp.first_name,
			tp.bank_id,
			tmc.alias_name,
			tmt.name as table_name,
			tmc.name as column_name,
			tbt.bt_name,
			tbt.description as business_description,
			tmc.cde,
			tmc.status,
			tmc.data_type,
			tmc.data_format,
			tmc.data_length,
			tmc.example,
			tmc.derived,
			tmc.Derivation_Logic,
			tmc.Sourced_from_Upstream,
			tmc.System_Checks,
			tc.Name as domain,
			tsc.name as subdomain,
			tpol.info_asset_name,
			tpol.description as info_asset_description,
			tpol.confidentiality,
			tpol.integrity,
			tpol.availability,
			tpol.overall_cia_rating,
			tmt.record_category,
			tmc.pii_flag,
			tmc.imm_prec_system_id,
			tmc.imm_succ_system_id,
			tmc.threshold
		FROM
			Tbl_System ts
			INNER JOIN Tbl_MD_Resource tmr ON tmr.system_id = ts.id
			INNER JOIN Tbl_MD_Table tmt ON tmt.resource_id = tmr.id
			LEFT JOIN tbl_business_term tbt ON tmt.business_term_id = tbt.id
			LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
			LEFT JOIN tbl_link_subcategory_people tlsp on tlsp.subcategory_id = tsc.id
			LEFT JOIN tbl_people tp on tlsp.people_id = tp.id
			LEFT JOIN Tbl_MD_Column tmc ON tmc.table_id = tmt.id
			LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.business_term_id = tbt.id
			LEFT JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id
			LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
			LEFT JOIN tbl_policy tpol on tbt.policy_id = tpol.id
		WHERE ROWNUM = 1 
		AND ts.id = ` + payload.GetString("left") + ` AND tmt.id = ` + payload.GetString("right") + ` `

	if payload.GetString("TableName") != "" {
		q += `AND tmt.name = '` + payload.GetString("TableName") + `' `
	}

	if payload.GetString("ColumnName") != "" {
		q += ` AND tmc.name = '` + payload.GetString("ColumnName") + `' `
	}

	if payload.GetString("ScreenLabel") != "" {
		q += ` AND tmc.alias_name = '` + payload.GetString("ScreenLabel") + `' `
	}

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

func (s *DSCService) GetddSource(leftParam string) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
			tmt.name as table_name,
			tmc.name as column_name,
			tmc.alias_name
		FROM
			Tbl_System ts
			INNER JOIN Tbl_MD_Resource tmr ON tmr.system_id = ts.id
			INNER JOIN Tbl_MD_Table tmt ON tmt.resource_id = tmr.id
			LEFT JOIN tbl_business_term tbt ON tmt.business_term_id = tbt.id
			LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
			LEFT JOIN tbl_link_subcategory_people tlsp on tlsp.subcategory_id = tsc.id
			LEFT JOIN tbl_people tp on tlsp.people_id = tp.id
			LEFT JOIN Tbl_MD_Column tmc ON tmc.table_id = tmt.id
			LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.business_term_id = tbt.id
			LEFT JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id
			LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
			LEFT JOIN tbl_policy tpol on tbt.policy_id = tpol.id
		AND ts.id = ` + leftParam
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
