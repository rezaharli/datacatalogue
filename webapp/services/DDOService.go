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

func (s *DDOService) GetLeftTable(tabs, loggedinid, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	if loggedinid != "" {
		args = append(args, loggedinid)
	}

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "left-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	filterSubDomains := ""
	if search != "" {
		filterSubDomains = search
	} else {
		filterSubDomains = searchDDM.GetString("SubDataDomain")
	}
	filterDataDomain := searchDDM.GetString("DataDomain")
	filterSubDomainOwner := searchDDM.GetString("SubDataDomainOwner")

	///////// DROPDOWN FILTER
	q = `SELECT * FROM (
		` + q + `
	) WHERE (
		upper(sub_domains) LIKE upper('%` + filterSubDomains + `%')
		AND upper(data_domain) LIKE upper('%` + filterDataDomain + `%')
		AND upper(sub_domain_owner) LIKE upper('%` + filterSubDomainOwner + `%')
	) `

	colFilterM, err := toolkit.ToM(colFilter)
	cf := make([]string, 0)
	if err != nil {
		cf = append(cf, "", "", "", "")
	} else {
		cf = append(cf, colFilterM.GetString("SUB_DOMAINS"), colFilterM.GetString("DATA_DOMAIN"), colFilterM.GetString("SUB_DOMAIN_OWNER"), colFilterM.GetString("BANK_ID"))
	}

	///////// COLUMN FILTER
	if cf[0] != "" || cf[1] != "" || cf[2] != "" || cf[3] != "" {
		q += `AND (
			upper(sub_domains) LIKE upper('%` + cf[0] + `%')
			AND upper(data_domain) LIKE upper('%` + cf[1] + `%')
			AND upper(sub_domain_owner) LIKE upper('%` + cf[2] + `%')
			AND upper(bank_id) LIKE upper('%` + cf[3] + `%')
		) `
	}

	///////// COUNT
	q = `SELECT res.*, 
			COUNT(DISTINCT sub_domains) OVER () COUNT_sub_domains,
			COUNT(DISTINCT data_domain) OVER () COUNT_data_domain,
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

func (s *DDOService) GetRightTable(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(systemID))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "right-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	filterBusinessTerm := searchDDM.GetString("BusinessTerm")

	///////// DROPDOWN FILTER
	q = `SELECT * FROM (
		` + q + `
	) WHERE (
		upper(business_term) LIKE upper('%` + filterBusinessTerm + `%')
	) `

	colFilterM, err := toolkit.ToM(colFilter)
	cf := make([]string, 0)
	if err != nil {
		cf = append(cf, "", "", "")
	} else {
		cf = append(cf, colFilterM.GetString("BUSINESS_TERM"), colFilterM.GetString("BT_DESCRIPTION"), colFilterM.GetString("BUSINESS_ALIAS_NAME"), colFilterM.GetString("CDE_YES_NO"))
	}

	///////// COLUMN FILTER
	if cf[0] != "" || cf[1] != "" || cf[2] != "" {
		q += `AND (
			upper(business_term) LIKE upper('%` + cf[0] + `%')
			AND upper(bt_description) LIKE upper('%` + cf[1] + `%')
			AND upper(cde_yes_no) LIKE upper('%` + cf[2] + `%')
		) `
	}

	///////// COUNT
	q = `SELECT res.*, 
			COUNT(DISTINCT business_term) OVER () COUNT_business_term,
			COUNT(DISTINCT bt_description) OVER () COUNT_bt_description,
			(SELECT COUNT (cde_yes_no) FROM ( ` + q + `) res2 WHERE cde_yes_no = 1) COUNT_cde_yes_no
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

func (s *DDOService) GetDetailsBusinessMetadataFromDomain(payload toolkit.M) (interface{}, int, error) {
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
	q, err := h.BuildQueryFromFile(filePath, "details-business-metadata-from-domain", args...)
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
		payload.GetString("GSTableName"),
		payload.GetString("GSColumnName"),
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
		q += `AND GS_TABLE_NAME = '` + otherArgs[3] + `' `
	}
	if otherArgs[3] != "" {
		q += `AND GS_COLUMN_NAME = '` + otherArgs[3] + `' `
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

func (s *DDOService) GetddSourceBusinessMetadataFromDomain(payload toolkit.M) (interface{}, int, error) {
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
	q, err := h.BuildQueryFromFile(filePath, "details-business-metadata-from-domain", args...)
	if err != nil {
		return nil, 0, err
	}

	///////// FILTER
	q = `SELECT DISTINCT business_term, business_alias, GS_TABLE_NAME, GS_COLUMN_NAME
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

func (s *DDOService) GetDetailsDownstreamUsageOfBusinessTerm(payload toolkit.M) (interface{}, int, error) {
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
	q, err := h.BuildQueryFromFile(filePath, "details-downstream-usage-of-business-term", args...)
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
		payload.GetString("DownstreamProcessName"),
	)

	///////// FILTER
	q = `SELECT rownum, a.* FROM (
		` + q + `
	) a
	WHERE ((
			tbtid = '` + otherArgs[0] + `'
		) OR (
			business_term = '` + otherArgs[1] + `' `
	if otherArgs[2] != "" {
		q += `AND DOWNSTREAM_PROCESS_NAME = '` + otherArgs[2] + `' `
	}
	q += `)) AND ROWNUM = 1 `

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

func (s *DDOService) GetddSourceDownstreamUsageOfBusinessTerm(payload toolkit.M) (interface{}, int, error) {
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
	q, err := h.BuildQueryFromFile(filePath, "details-downstream-usage-of-business-term", args...)
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
	)

	///////// FILTER
	q = `SELECT DISTINCT DOWNSTREAM_PROCESS_NAME
		FROM (
		` + q + `
	) 
	WHERE (
			tbtid = '` + otherArgs[0] + `'
		) OR (
			business_term = '` + otherArgs[1] + `' `
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

func (s *DDOService) GetDetailsBTResiding(payload toolkit.M) (interface{}, int, error) {
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
	q, err := h.BuildQueryFromFile(filePath, "details-business-term-residing", args...)
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
		q += `AND SYSTEM_NAME = '` + otherArgs[2] + `' `
	}
	if otherArgs[3] != "" {
		q += `AND ITAM_ID = '` + otherArgs[3] + `' `
	}
	if otherArgs[4] != "" {
		q += `AND TABLE_NAME = '` + otherArgs[4] + `' `
	}
	if otherArgs[5] != "" {
		q += `AND COLUMN_NAME = '` + otherArgs[5] + `' `
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

func (s *DDOService) GetddSourceBTResiding(payload toolkit.M) (interface{}, int, error) {
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
	q, err := h.BuildQueryFromFile(filePath, "details-business-term-residing", args...)
	if err != nil {
		return nil, 0, err
	}

	///////// FILTER
	q = `SELECT DISTINCT SYSTEM_NAME, ITAM_ID, TABLE_NAME, COLUMN_NAME
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
