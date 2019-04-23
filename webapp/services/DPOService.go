package services

import (
	"path/filepath"
	"strings"

	"github.com/eaciit/clit"
	"github.com/eaciit/toolkit"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type DPOService struct {
	*Base
}

func NewDPOService() *DPOService {
	ret := new(DPOService)
	return ret
}

func (s *DPOService) GetLeftTable(tabs, loggedinid, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	gridArgs.QueryName = "left-grid"
	gridArgs.PageNumber = pageNumber
	gridArgs.RowsPerPage = rowsPerPage

	if loggedinid != "" {
		gridArgs.MainArgs = append(gridArgs.MainArgs, loggedinid)
	}

	///////// --------------------------------------------------DROPDOWN FILTER
	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	filterProcessName := ""
	if search != "" {
		filterProcessName = search
	} else {
		filterProcessName = searchDDM.GetString("ProcessName")
	}

	gridArgs.DropdownFilter = append(gridArgs.DropdownFilter,
		filterProcessName,
		searchDDM.GetString("ProcessOwner"),
	)

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("DOWNSTREAM_PROCESS"),
			colFilterM.GetString("PROCESS_OWNER"),
			colFilterM.GetString("BANK_ID"),
		)
	}

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DPOService) GetRightTable(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
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

	gridArgs.DropdownFilter = append(gridArgs.DropdownFilter, searchDDM.GetString("CDEName"))

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "", "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("CDE_NAME"),
			colFilterM.GetString("SEGMENT"),
			colFilterM.GetString("IMM_PREC_SYSTEM"),
			colFilterM.GetString("ULT_SOURCE_SYSTEM"),
			colFilterM.GetString("BUSINESS_DESCRIPTION"),
			colFilterM.GetString("CDE_RATIONALE"),
		)
	}

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DPOService) GetDetails(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	tabs := ""

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")), payload.GetString("Right"))

	if strings.Contains(payload.GetString("Which"), "my") == true {
		tabs = "dpomy"
	} else {
		tabs = "dpoall"
	}

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "details", []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	otherArgs := make([]string, 0)
	otherArgs = append(otherArgs,
		payload.GetString("ImmSystemName"),
		payload.GetString("ImmItamID"),
		payload.GetString("ImmTableName"),
		payload.GetString("ImmColumnName"),
		payload.GetString("ImmScreenLabel"),
		payload.GetString("UltSystemName"),
		payload.GetString("UltItamID"),
		payload.GetString("UltTableName"),
		payload.GetString("UltColumnName"),
		payload.GetString("UltScreenLabel"),
	)

	checkNotEmpty := func(s []string) bool {
		for _, v := range s {
			if v != "" {
				return true
			}
		}
		return false
	}

	q = `SELECT rownum, a.* FROM (
		` + q + `
	) a `

	if checkNotEmpty(otherArgs) == true {
		q += `WHERE (( ID IS NOT NULL `
		if otherArgs[0] != "" {
			q += `AND IMM_SYSTEM = '` + otherArgs[0] + `' `
		}
		if otherArgs[1] != "" {
			q += `AND IMM_ITAM_ID = '` + otherArgs[1] + `' `
		}
		if otherArgs[2] != "" {
			q += `AND IMM_TABLE_NAME = '` + otherArgs[2] + `' `
		}
		if otherArgs[3] != "" {
			q += `AND IMM_COLUMN_NAME = '` + otherArgs[3] + `' `
		}
		if otherArgs[4] != "" {
			q += `AND IMM_SCREEN_LABEL_NAME = '` + otherArgs[4] + `' `
		}
		if otherArgs[5] != "" {
			q += `AND ULT_SYSTEM = '` + otherArgs[5] + `' `
		}
		if otherArgs[6] != "" {
			q += `AND ULT_ITAM_ID = '` + otherArgs[6] + `' `
		}
		if otherArgs[7] != "" {
			q += `AND ULT_TABLE_NAME = '` + otherArgs[7] + `' `
		}
		if otherArgs[8] != "" {
			q += `AND ULT_COLUMN_NAME = '` + otherArgs[8] + `' `
		}
		if otherArgs[9] != "" {
			q += `AND ULT_SCREEN_LABEL_NAME = '` + otherArgs[9] + `' `
		}
		q += `)) AND rownum = 1 `
	} else {
		q += `WHERE rownum = 1 `
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

func (s *DPOService) GetddSource(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	tabs := ""

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")), payload.GetString("Right"))

	if strings.Contains(payload.GetString("Which"), "my") == true {
		tabs = "dpomy"
	} else {
		tabs = "dpoall"
	}

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "details", []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	///////// FILTER
	q = `SELECT DISTINCT 
			Imm_System, 
			Imm_Itam_ID, 
			Imm_Table_Name, 
			Imm_Column_Name, 
			Imm_Screen_Label_name, 
			Ult_System, 
			Ult_Itam_ID, 
			Ult_Table_Name, 
			Ult_Column_Name, 
			Ult_Screen_Label_name
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
