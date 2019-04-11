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
	*Base
}

func NewDSCService() *DSCService {
	ret := new(DSCService)
	return ret
}

func (s *DSCService) GetAllSystem(tabs, loggedinid, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "dsc.sql")
	gridArgs.QueryName = "dsc-view"
	gridArgs.PageNumber = pageNumber
	gridArgs.RowsPerPage = rowsPerPage

	if loggedinid != "" {
		gridArgs.MainArgs = append(gridArgs.MainArgs, "MY", loggedinid)
	} else {
		gridArgs.MainArgs = append(gridArgs.MainArgs, "ALL", "0000000")
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

	gridArgs.DropdownFilter = append(gridArgs.DropdownFilter,
		filterSystemName,
		searchDDM.GetString("ItamID"),
	)

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("SYSTEM_NAME"),
			colFilterM.GetString("ITAM_ID"),
			colFilterM.GetString("DATASET_CUSTODIAN"),
			colFilterM.GetString("BANK_ID"),
		)
	}

	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DSCService) GetHomepageCounts(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	system := payload.GetString("System")
	args = append(args, system, system, system)

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dsc.sql")
	q, err := h.BuildQueryFromFile(filePath, "dsc-view-homepage", args...)
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

func (s *DSCService) GetCDETable(system string, colFilter interface{}, pageNumber, rowsPerPage int) ([]toolkit.M, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "dsc.sql")
	gridArgs.QueryName = "dsc-view-cde"
	gridArgs.PageNumber = pageNumber
	gridArgs.RowsPerPage = rowsPerPage

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("CDE"),
			colFilterM.GetString("DESCRIPTION"),
			colFilterM.GetString("TABLE_NAME"),
			colFilterM.GetString("COLUMN_NAME"),
			colFilterM.GetString("DSP_NAME"),
			colFilterM.GetString("PROCESS_OWNER"),
		)
	}

	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DSCService) GetTableName(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
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

	gridArgs.DropdownFilter = append(gridArgs.DropdownFilter, searchDDM.GetString("TableName"))
	gridArgs.DropdownFilter = append(gridArgs.DropdownFilter, searchDDM.GetString("ColumnName"))

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "", "")
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
			colFilterM.GetString("TABLE_NAME"),
			colFilterM.GetString("COLUMN_NAME"),
			colFilterM.GetString("BUSINESS_ALIAS_NAME"),
			cdeYesNo,
		)
	}

	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DSCService) GetInterfacesRightTable(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	gridArgs.QueryName = "right-grid"
	gridArgs.PageNumber = pageNumber
	gridArgs.RowsPerPage = rowsPerPage

	gridArgs.MainArgs = append(gridArgs.MainArgs, toolkit.ToString(systemID))

	///////// --------------------------------------------------DROPDOWN FILTER
	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	gridArgs.DropdownFilter = append(gridArgs.DropdownFilter, searchDDM.GetString("TableName"))
	gridArgs.DropdownFilter = append(gridArgs.DropdownFilter, searchDDM.GetString("ColumnName"))

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
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

	gridArgs.GroupCol = "LIST_OF_CDE"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
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
