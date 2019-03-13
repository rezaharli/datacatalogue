package services

import (
	"github.com/eaciit/toolkit"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type DPOService struct {
}

func NewDPOService() *DPOService {
	ret := new(DPOService)
	return ret
}

func (s *DPOService) GetLeftTable(loggedinid, search string, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
			tdp.id,
			tdp.Name as downstream_process,
			tdp.owner_id as process_owner,
			tp.bank_id
		FROM 
			tbl_ds_processes tdp
			LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = tdp.id
			LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id
		WHERE
			upper(tlrp.object_type) = upper('process') `

	if loggedinid != "" {
		a := toolkit.ToInt(loggedinid, "")
		toolkit.Println(a)
		q += `AND tp.bank_id = '` + toolkit.ToString(a) + `' `
	}

	if search != "" {
		q += `
			upper(tdp.Name) LIKE upper('%` + search + `%')
			OR upper(tdp.owner_id) LIKE upper('%` + search + `%')
			OR upper(tp.bank_id) LIKE upper('%` + search + `%') `
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

func (s *DPOService) GetRightTable(processID int, search string, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
			tbt.id,
			tdp.id as tdpid,
			tdpd.business_term_id as cde_name,
			ts.name as Segment,
			tdpd.imm_prec_system_id,
			tdpd.ultimate_source_system_id,
			tbt.description as business_description,
			tbt.cde_rationale
		FROM 
			tbl_ds_processes tdp
			LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = tdp.id
			LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id

            LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.process_id = tdp.id
			LEFT JOIN tbl_segment ts ON tdpd.segment_id = ts.id
			LEFT JOIN tbl_business_term tbt ON tdpd.business_term_id = tbt.id
			LEFT JOIN tbl_system tsy ON tdpd.imm_prec_system_id = tsy.id
			LEFT JOIN tbl_md_table tmt ON tmt.business_term_id = tbt.id
			LEFT JOIN tbl_md_column tmc ON tmc.table_id = tmt.id
			LEFT JOIN tbl_category tc ON ts.subdomain_id = tc.id
			LEFT JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
			LEFT JOIN tbl_link_category_people tlcp ON tlcp.category_id = tc.id
		WHERE tdp.id = ` + toolkit.ToString(processID)

	if search != "" {
		q += `
			AND
				upper(tdpd.business_term_id) LIKE upper('%` + search + `%')
				OR upper(ts.name) LIKE upper('%` + search + `%')
				OR upper(tdpd.imm_prec_system_id) LIKE upper('%` + search + `%')
				OR upper(tdpd.ultimate_source_system_id) LIKE upper('%` + search + `%')
				OR upper(tbt.description) LIKE upper('%` + search + `%')
				OR upper(tbt.cde_rationale) LIKE upper('%` + search + `%')`
	}

	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewSystemModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}
func (s *DPOService) GetDetails(leftParam int, rightParam int) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
			tdp.id,
			tdp.Name as downstream_process,
			tdp.owner_id as process_owner,
			tp.bank_id,
			tdpd.business_term_id,
			tdpd.business_term_id as cde_name,
			ts.name as Segment,
			tdpd.imm_prec_system_id,
			tdpd.ultimate_source_system_id,
			tbt.description as business_description,
			tbt.cde_rationale,
			tsy.system_name,
			tsy.itam_id,
			tmt.name as table_name,
			tmc.name as column_name,
			tmc.derived,
			tmc.derivation_logic,
			tmc.dq_standards as dq_requirements,
			tmc.threshold,
			tmc.data_sla_signed,
			tmc.golden_source,
			tbt.golden_source_system_id,
			tc.name as domain,
			tsc.name as subdomain,
			tlcp.people_id as domain_owner,
			tbt.bt_name as business_term,
			tmc.dpo_dq_standards,
			tbt.dq_standards,
			tmc.dpo_threshold
		FROM 
			tbl_ds_processes tdp
			INNER JOIN tbl_ds_process_detail tdpd ON tdpd.process_id = tdp.id
			LEFT JOIN tbl_people tp ON tdp.owner_id = tp.id
			LEFT JOIN tbl_segment ts ON tdpd.segment_id = ts.id
			LEFT JOIN tbl_business_term tbt ON tdpd.business_term_id = tbt.id
			LEFT JOIN tbl_system tsy ON tdpd.imm_prec_system_id = tsy.id
			LEFT JOIN tbl_md_table tmt ON tmt.business_term_id = tbt.id
			LEFT JOIN tbl_md_column tmc ON tmc.table_id = tmt.id
			LEFT JOIN tbl_category tc ON ts.subdomain_id = tc.id
			LEFT JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
			LEFT JOIN tbl_link_category_people tlcp ON tlcp.category_id = tc.id
		WHERE
			tdp.id = ` + toolkit.ToString(leftParam) + ` and tbt.id = ` + toolkit.ToString(rightParam)
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
