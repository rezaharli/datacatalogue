package services

import (
	"path/filepath"
	"strings"

	"github.com/eaciit/clit"
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

func (s *DDOService) GetLeftTable(tabs, loggedinid, search string, searchDD interface{}, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	if loggedinid != "" {
		args = append(args, loggedinid)
	}

	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	if search != "" {
		args = append(args, search, searchDDM.GetString("SubDataDomain"))
	} else {
		args = append(args, searchDDM.GetString("DataDomain"), searchDDM.GetString("SubDataDomain"))
	}

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err = h.BuildQueryFromFile(filePath, "left-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	q = `SELECT res.*, 
			COUNT(DISTINCT data_domain) OVER () COUNT_data_domain,
			COUNT(DISTINCT sub_domains) OVER () COUNT_sub_domains
		FROM ( ` + q + `) res `

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewCategoryModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  pageNumber,
		RowsPerPage: rowsPerPage,
		GroupCol:    "-",
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DDOService) GetRightTable(tabs string, systemID int, search string, searchDD interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(systemID))
	args = append(args, searchDDM.GetString("SubDataDomainOwner"))
	args = append(args, searchDDM.GetString("BusinessTerm"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err = h.BuildQueryFromFile(filePath, "right-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	q = `SELECT res.*, 
			COUNT(DISTINCT business_term) OVER () COUNT_business_term,
			COUNT(DISTINCT business_term_description) OVER () COUNT_business_term_description,
			COUNT(DISTINCT cde_yes_no) OVER () COUNT_cde_yes_no
		FROM ( ` + q + `) res `

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewCategoryModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  pageNumber,
		RowsPerPage: rowsPerPage,
		GroupCol:    "-",
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DDOService) GetDetails(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	tabs := ""

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))
	args = append(args, payload.GetString("Right"))

	if strings.Contains(payload.GetString("Which"), "my") == true {
		tabs = "ddomy"
	} else {
		tabs = "ddoall"
	}

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "details", args...)
	if err != nil {
		return nil, 0, err
	}

	// q := `SELECT
	// 		tc.name as data_domain,
	// 		tsc.name as sub_domain,
	// 		tlscp.people_id as sub_domain_owner,
	// 		tp.bank_id,
	// 		tbt.bt_name,
	// 		tbt.description as business_term_description,
	// 		tmc.alias_name,
	// 		tbt.cde,
	// 		tbt.dq_standards,
	// 		tbt.policy_guidance,
	// 		tbt.mandatory,
	// 		tbt.golden_source_system_id,
	// 		ts.itam_id,
	// 		tmt.name as golden_source_table_name,
	// 		tmc.name as golden_source_column_name,
	// 		tpo.info_asset_name,
	// 		tpo.description as info_asset_desc,
	// 		tpo.confidentiality,
	// 		tpo.integrity,
	// 		tpo.availability,
	// 		tpo.overall_cia_rating,
	// 		tmc.record_category,
	// 		tmc.pii_flag,
	// 		tdp.name as downstream_process_name,
	// 		tdp.owner_name as downstream_process_owner,
	// 		ts.system_name,
	// 		tmc.name as column_name,
	// 		tmc.ddo_threshold
	// 	FROM
	// 		tbl_category tc
	// 		INNER JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
	// 		INNER JOIN tbl_business_term tbt ON tbt.parent_id = tsc.id
	// 		LEFT JOIN tbl_link_subcategory_people tlscp ON tlscp.subcategory_id = tsc.id
	// 		LEFT JOIN Tbl_People	tp on tlscp.people_id = tp.id
	// 		LEFT join tbl_md_column tmc on tmc.business_term_id = tbt.id
	// 		LEFT join tbl_md_table tmt on tmc.table_id = tmt.id
	// 		LEFT join tbl_md_resource tmr on tmt.resource_id = tmr.id
	// 		LEFT join tbl_system ts on tmr.system_id = ts.id
	// 		LEFT join tbl_policy tpo on tbt.policy_id = tpo.id
	// 		LEFT join tbl_ds_process_detail tdpd on tdpd.business_term_id = tbt.id
	// 		LEFT join tbl_ds_processes tdp on tdpd.process_id = tdp.id
	// 	WHERE
	// 		tc.id = ` + toolkit.ToString(leftParam) + ` and tbt.id = ` + toolkit.ToString(rightParam)

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewCategoryModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}
