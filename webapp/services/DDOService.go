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
		args = append(args, search, searchDDM.GetString("SubDataDomain"), searchDDM.GetString("SubDataDomainOwner"))
	} else {
		args = append(args, searchDDM.GetString("DataDomain"), searchDDM.GetString("SubDataDomain"), searchDDM.GetString("SubDataDomainOwner"))
	}

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err = h.BuildQueryFromFile(filePath, "left-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	q = `SELECT res.*, 
			COUNT(DISTINCT data_domain) OVER () COUNT_data_domain,
			COUNT(DISTINCT sub_domains) OVER () COUNT_sub_domains,
			COUNT(DISTINCT sub_domain_owner) OVER () COUNT_sub_domain_owner,
			COUNT(DISTINCT bank_id) OVER () COUNT_bank_id
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
	args = append(args, searchDDM.GetString("BusinessTerm"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err = h.BuildQueryFromFile(filePath, "right-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	q = `SELECT res.*, 
			COUNT(DISTINCT business_term) OVER () COUNT_business_term,
			COUNT(DISTINCT bt_description) OVER () COUNT_bt_description,
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

	otherArgs := make([]string, 0)
	if payload.GetString("BusinessTerm") != "" {
		otherArgs = append(otherArgs, "")
	} else {
		otherArgs = append(otherArgs, payload.GetString("Right"))
	}

	otherArgs = append(otherArgs,
		payload.GetString("BusinessTerm"),
		payload.GetString("BusinessAlias"),
		payload.GetString("DownstreamProcessName"),
		payload.GetString("SystemName"),
		payload.GetString("ItamId"),
		payload.GetString("TableName"),
		payload.GetString("ColumnName"),
	)

	///////// FILTER
	q = `SELECT * FROM (
		` + q + `
	) 
	WHERE (
			tbtid = '` + otherArgs[0] + `'
		) OR (
			business_term = '` + otherArgs[1] + `' `
	if otherArgs[2] != "" {
		q += `AND business_alias = '` + otherArgs[2] + `' `
	}
	if otherArgs[3] != "" {
		q += `AND downstream_process_name = '` + otherArgs[3] + `' `
	}
	if otherArgs[4] != "" {
		q += `AND system_name_dd = '` + otherArgs[4] + `' `
	}
	if otherArgs[5] != "" {
		q += `AND itam_id_dd = '` + otherArgs[5] + `' `
	}
	if otherArgs[6] != "" {
		q += `AND table_name_dd = '` + otherArgs[6] + `' `
	}
	if otherArgs[7] != "" {
		q += `AND column_name_dd = '` + otherArgs[7] + `' `
	}
	q += `) `

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

func (s *DDOService) GetddSource(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	tabs := ""

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

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

	///////// FILTER
	q = `SELECT DISTINCT business_term, business_alias, downstream_process_name, system_name_dd, itam_id_dd, table_name_dd, column_name_dd 
		FROM (
		` + q + `
	) `

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
