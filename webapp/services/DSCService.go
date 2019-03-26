package services

import (
	"path/filepath"
	"strings"

	"github.com/eaciit/clit"
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

func (s *DSCService) GetAllSystem(tabs, loggedinid, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	if loggedinid != "" {
		args = append(args, loggedinid)
	}

	///////// --------------------------------------------------DROPDOWN FILTER
	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	filterSystemName := ""
	if search != "" {
		filterSystemName = search
	} else {
		filterSystemName = searchDDM.GetString("SystemName")
	}

	args = append(args, filterSystemName)
	args = append(args, searchDDM.GetString("ItamID"))

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		args = append(args, "", "", "", "")
	} else {
		args = append(args,
			colFilterM.GetString("SYSTEM_NAME"),
			colFilterM.GetString("ITAM_ID"),
			colFilterM.GetString("DATASET_CUSTODIAN"),
			colFilterM.GetString("BANK_ID"),
		)
	}

	///////// --------------------------------------------------BUILD QUERY FROM ARGS
	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err = h.BuildQueryFromFile(filePath, "left-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	///////// --------------------------------------------------EXECUTE
	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
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

func (s *DSCService) GetTableName(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)
	args = append(args, toolkit.ToString(systemID))
	args = append(args, toolkit.ToString(systemID))

	///////// --------------------------------------------------DROPDOWN FILTER
	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	args = append(args, searchDDM.GetString("TableName"))
	args = append(args, searchDDM.GetString("ColumnName"))

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		args = append(args, "", "", "", "")
	} else {
		args = append(args,
			colFilterM.GetString("TABLE_NAME"),
			colFilterM.GetString("COLUMN_NAME"),
			colFilterM.GetString("BUSINESS_ALIAS_NAME"),
			colFilterM.GetString("CDE_YES_NO"),
		)
	}

	///////// --------------------------------------------------BUILD QUERY FROM ARGS
	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err = h.BuildQueryFromFile(filePath, "right-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	///////// --------------------------------------------------EXECUTE
	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
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

func (s *DSCService) GetInterfacesRightTable(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)
	args = append(args, toolkit.ToString(systemID))

	///////// --------------------------------------------------DROPDOWN FILTER
	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	args = append(args, searchDDM.GetString("TableName"))
	args = append(args, searchDDM.GetString("ColumnName"))

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		args = append(args, "", "", "", "", "", "", "", "", "")
	} else {
		args = append(args,
			colFilterM.GetString("LIST_OF_CDE"),
			colFilterM.GetString("IMM_PREC_SYSTEM_NAME"),
			colFilterM.GetString("IMM_PREC_SYSTEM_SLA"),
			colFilterM.GetString("IMM_PREC_SYSTEM_OLA"),
			colFilterM.GetString("IMM_SUCC_SYSTEM_NAME"),
			colFilterM.GetString("IMM_SUCC_SYSTEM_SLA"),
			colFilterM.GetString("IMM_SUCC_SYSTEM_OLA"),
			colFilterM.GetString("LIST_DOWNSTREAM_PROCESS"),
			colFilterM.GetString("DOWNSTREAM_PROCESS_OWNER"),
		)
	}

	///////// --------------------------------------------------BUILD QUERY FROM ARGS
	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err = h.BuildQueryFromFile(filePath, "right-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	///////// --------------------------------------------------EXECUTE
	filePath = filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err = h.BuildQueryFromFile(filePath, "right-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewSystemModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  pageNumber,
		RowsPerPage: rowsPerPage,
		GroupCol:    "list_of_cde",
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetDetails(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	tabs := ""

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

	if strings.Contains(payload.GetString("Which"), "interfaces") == true {
		tabs = "dscinterfaces"
	} else if strings.Contains(payload.GetString("Which"), "my") == true {
		tabs = "dscmy"
	} else {
		tabs = "dscall"
	}

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "details", args...)
	if err != nil {
		return nil, 0, err
	}

	otherArgs := make([]string, 0)
	if payload.GetString("TableName") != "" && payload.GetString("ColumnName") != "" {
		otherArgs = append(otherArgs, "", "")
	} else {
		otherArgs = append(otherArgs, payload.GetString("Right"), payload.GetString("Column"))
	}

	otherArgs = append(otherArgs, payload.GetString("TableName"), payload.GetString("ColumnName"), payload.GetString("ScreenLabel"))

	///////// FILTER
	q = `SELECT * FROM (
		` + q + `
	) 
	WHERE (
			tmtid = '` + otherArgs[0] + `'
			AND tmcid = '` + otherArgs[1] + `'
		) OR (
			table_name = '` + otherArgs[2] + `'
			AND column_name = '` + otherArgs[3] + `' `
	if otherArgs[4] != "" {
		q += `AND business_alias_name = '` + otherArgs[4] + `' `
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

func (s *DSCService) GetddSource(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	tabs := ""

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

	if strings.Contains(payload.GetString("Which"), "interfaces") == true {
		tabs = "dscinterfaces"
	} else if strings.Contains(payload.GetString("Which"), "my") == true {
		tabs = "dscmy"
	} else {
		tabs = "dscall"
	}

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "details", args...)
	if err != nil {
		return nil, 0, err
	}

	///////// FILTER
	q = `SELECT DISTINCT table_name, column_name, business_alias_name 
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
