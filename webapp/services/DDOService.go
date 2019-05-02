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
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "ddo.sql")
	gridArgs.QueryName = "ddo-view"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	if loggedinid != "" {
		gridArgs.MainArgs = append(gridArgs.MainArgs, "MY", loggedinid)
	} else {
		gridArgs.MainArgs = append(gridArgs.MainArgs, "ALL", "0000000")
	}

	///////// --------------------------------------------------DROPDOWN FILTER
	// searchDDM, err := toolkit.ToM(searchDD)
	// if err != nil {
	// 	return nil, 0, err
	// }

	// filterSubDomains := ""
	// if search != "" {
	// 	filterSubDomains = search
	// } else {
	// 	filterSubDomains = searchDDM.GetString("SubDataDomain")
	// }

	// gridArgs.DropdownFilter = append(gridArgs.DropdownFilter,
	// 	filterSubDomains,
	// 	searchDDM.GetString("DataDomain"),
	// 	searchDDM.GetString("SubDataDomainOwner"),
	// )

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("DATA_DOMAIN"),
			colFilterM.GetString("SUB_DOMAINS"),
			colFilterM.GetString("SUB_DOMAIN_OWNER"),
			colFilterM.GetString("BANK_ID"),
		)
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
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	system := payload.GetString("System")
	args = append(args, system, system, system)

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "ddo.sql")
	q, err := h.BuildQueryFromFile(filePath, "ddo-view-homepage", []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

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
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "ddo.sql")
	gridArgs.QueryName = "ddo-businessterm"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("BUSINESS_TERM"),
			colFilterM.GetString("DESCRIPTION"),
		)
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
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "ddo.sql")
	gridArgs.QueryName = "ddo-systems"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("SYSTEM_NAME"),
			colFilterM.GetString("BT_COUNT"),
		)
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
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "ddo.sql")
	gridArgs.QueryName = "ddo-systems-businessterm"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	gridArgs.MainArgs = append(gridArgs.MainArgs, subdomain, system)

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("BT_NAME"),
			colFilterM.GetString("TABLE_NAME"),
			colFilterM.GetString("COLUMN_NAME"),
		)
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DDOService) GetRightTable(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	gridArgs.QueryName = "right-grid"
	gridArgs.PageNumber = pageNumber
	gridArgs.RowsPerPage = rowsPerPage

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
	q, err := h.BuildQueryFromFile(filePath, "details-business-metadata-from-domain", []string{}, args...)
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
	q, err := h.BuildQueryFromFile(filePath, "details-business-metadata-from-domain", []string{}, args...)
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
	q, err := h.BuildQueryFromFile(filePath, "details-downstream-usage-of-business-term", []string{}, args...)
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
	q, err := h.BuildQueryFromFile(filePath, "details-downstream-usage-of-business-term", []string{}, args...)
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
	q, err := h.BuildQueryFromFile(filePath, "details-business-term-residing", []string{}, args...)
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
	q, err := h.BuildQueryFromFile(filePath, "details-business-term-residing", []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

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
