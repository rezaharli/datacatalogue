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
	*Base
}

func NewDDOService() *DDOService {
	ret := new(DDOService)
	return ret
}

func (s *DDOService) GetLeftTable(tabs, loggedinid, search string, searchDD, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "ddo.sql"
	queryName := "ddo-view"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	if loggedinid != "" {
		gridArgs.MainArgs = append(gridArgs.MainArgs, "MY", loggedinid)
	} else {
		gridArgs.MainArgs = append(gridArgs.MainArgs, "ALL", "0000000")
	}

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"DATA_DOMAIN", "SUB_DOMAINS", "SUB_DOMAIN_OWNER", "BANK_ID",
	)

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colFilterM.GetString(colname))
		}

		filterTypes := colFilterM.Get("filterTypes")
		if filterTypes != nil {
			gridArgs.ColumnFilterType = colFilterM.Get("filterTypes").(map[string]interface{})
		}
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DDOService) GetHomepageCounts(payload toolkit.M) (interface{}, int, error) {
	fileName := "ddo.sql"
	queryName := "ddo-view-homepage"

	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	system := payload.GetString("System")
	args = append(args, system, system, system)

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	funcLog(funcName(), fileName, queryName)

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

func (s *DDOService) GetBusinesstermTable(system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "ddo.sql"
	queryName := "ddo-businessterm"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"BT_NAME", "BT_DESCRIPTION",
	)

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colFilterM.GetString(colname))
		}

		filterTypes := colFilterM.Get("filterTypes")
		if filterTypes != nil {
			gridArgs.ColumnFilterType = colFilterM.Get("filterTypes").(map[string]interface{})
		}
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DDOService) GetSystemsTable(system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "ddo.sql"
	queryName := "ddo-systems"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"SYSTEM_NAME", "BT_COUNT",
	)

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colFilterM.GetString(colname))
		}

		filterTypes := colFilterM.Get("filterTypes")
		if filterTypes != nil {
			gridArgs.ColumnFilterType = colFilterM.Get("filterTypes").(map[string]interface{})
		}
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DDOService) GetSystemsBusinesstermTable(subdomain, system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "ddo.sql"
	queryName := "ddo-systems-businessterm"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, subdomain, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"BT_NAME", "TABLE_NAME", "COLUMN_NAME",
	)

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colFilterM.GetString(colname))
		}

		filterTypes := colFilterM.Get("filterTypes")
		if filterTypes != nil {
			gridArgs.ColumnFilterType = colFilterM.Get("filterTypes").(map[string]interface{})
		}
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DDOService) GetDownstreamTable(system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "ddo.sql"
	queryName := "ddo-downstream"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"DP_NAME", "PROCESS_OWNER", "BT_COUNT",
	)

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colFilterM.GetString(colname))
		}

		filterTypes := colFilterM.Get("filterTypes")
		if filterTypes != nil {
			gridArgs.ColumnFilterType = colFilterM.Get("filterTypes").(map[string]interface{})
		}
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DDOService) GetDownstreamBusinesstermTable(subdomain, system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "ddo.sql"
	queryName := "ddo-downstream-businessterm"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system, subdomain)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"BT_NAME", "SYSTEM_CONSUMED", "TABLE_NAME", "COLUMN_NAME", "GOLDEN_SOURCE", "ALIAS_NAME", "GS_SYSTEM_NAME", "GS_TABLE_NAME", "GS_COLUMN_NAME",
	)

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colFilterM.GetString(colname))
		}

		filterTypes := colFilterM.Get("filterTypes")
		if filterTypes != nil {
			gridArgs.ColumnFilterType = colFilterM.Get("filterTypes").(map[string]interface{})
		}
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	gridArgs.GroupCol = "BT_NAME"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DDOService) GetRightTable(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	fileName := tabs + ".sql"
	queryName := "right-grid"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pageNumber
	gridArgs.RowsPerPage = rowsPerPage

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, toolkit.ToString(systemID))
	gridArgs.MainArgs = append(gridArgs.MainArgs, toolkit.ToString(systemID))

	///////// --------------------------------------------------DROPDOWN FILTER
	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	gridArgs.DropdownFilter = append(gridArgs.DropdownFilter,
		searchDDM.GetString("BusinessTerm"),
	)

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "")
	} else {
		cdeYesNo := colFilterM.GetString("CDE_YES_NO")

		if cdeYesNo != "" {
			if strings.EqualFold(cdeYesNo, "yes") {
				cdeYesNo = "1"
			} else {
				cdeYesNo = "0"
			}
		}

		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("BUSINESS_TERM"),
			colFilterM.GetString("BT_DESCRIPTION"),
			cdeYesNo,
		)
	}

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DDOService) GetDetails(payload toolkit.M) (interface{}, int, error) {
	fileName := "ddo.sql"
	queryName := "details"

	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Subdomain"), payload.GetString("BTname"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	funcLog(funcName(), fileName, queryName)

	otherArgs := make([]string, 0)
	otherArgs = append(otherArgs,
		payload.GetString("SystemName"),
		payload.GetString("TableName"),
		payload.GetString("BusinessAliasName"),
	)

	checkNotEmpty := func(s []string) bool {
		for _, v := range s {
			if v != "" {
				return true
			}
		}
		return false
	}

	if checkNotEmpty(otherArgs) == true {
		///////// FILTER
		q = `SELECT filtered.* FROM (
			` + q + `
		) filtered `

		q += `WHERE ( ID IS NOT NULL `
		if otherArgs[0] != "" {
			q += `AND SYSTEM_NAME = '` + otherArgs[0] + `' `
			q += `AND GS_SYSTEM_NAME = '` + otherArgs[0] + `' `
		}
		if otherArgs[1] != "" {
			q += `AND TABLE_NAME = '` + otherArgs[1] + `' `
		}
		if otherArgs[2] != "" {
			q += `AND ALIAS_NAME = '` + otherArgs[2] + `' `
		}
		q += `) `
	}

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewCategoryModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  1,
		RowsPerPage: 1,
		GroupCol:    "system_name",
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DDOService) GetddSource(payload toolkit.M) (interface{}, int, error) {
	fileName := "ddo.sql"
	queryName := "details"

	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Subdomain"), payload.GetString("BTname"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	funcLog(funcName(), fileName, queryName)

	///////// FILTER
	q = `SELECT DISTINCT 
			SYSTEM_NAME, 
			TABLE_NAME, 
			ALIAS_NAME
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

func (s *DDOService) GetDetailsBusinessMetadataFromDomain(payload toolkit.M) (interface{}, int, error) {
	fileName := "ddodetails.sql"
	queryName := "details-business-metadata-from-domain"

	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	funcLog(funcName(), fileName, queryName)

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
		TableName:   m.NewCategoryModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  1,
		RowsPerPage: 1,
		GroupCol:    "business_alias",
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DDOService) GetddSourceBusinessMetadataFromDomain(payload toolkit.M) (interface{}, int, error) {
	fileName := "ddodetails.sql"
	queryName := "details-business-metadata-from-domain"

	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	funcLog(funcName(), fileName, queryName)

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
	fileName := "ddodetails.sql"
	queryName := "details-downstream-usage-of-business-term"

	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	funcLog(funcName(), fileName, queryName)

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
	fileName := "ddodetails.sql"
	queryName := "details-downstream-usage-of-business-term"

	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	funcLog(funcName(), fileName, queryName)

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
	fileName := "ddodetails.sql"
	queryName := "details-business-term-residing"

	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	funcLog(funcName(), fileName, queryName)

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
	q = `SELECT rownum, a.* FROM (
		` + q + `
	) a
	WHERE ((
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
	q += `)) AND rownum = 1 `

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
	fileName := "ddodetails.sql"
	queryName := "details-business-term-residing"

	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	funcLog(funcName(), fileName, queryName)

	otherArgs := make([]string, 0)
	if payload.GetString("BusinessTerm") != "" {
		otherArgs = append(otherArgs, "")
	} else {
		otherArgs = append(otherArgs, payload.GetString("Right"))
	}

	otherArgs = append(otherArgs, payload.GetString("BusinessTerm"))

	///////// FILTER
	q = `SELECT DISTINCT SYSTEM_NAME, ITAM_ID, TABLE_NAME, COLUMN_NAME
		FROM (
		` + q + `
	)
	WHERE (
		tbtid = '` + otherArgs[0] + `'
	) OR (
		business_term = '` + otherArgs[1] + `'
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
