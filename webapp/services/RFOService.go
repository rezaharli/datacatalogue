package services

import (
	"path/filepath"

	"github.com/eaciit/clit"
	"github.com/eaciit/toolkit"
)

type RFOService struct {
	*Base
}

func NewRFOService() *RFOService {
	ret := new(RFOService)
	return ret
}

func (s *RFOService) GetLeftTable(tabs, loggedinid, search string, searchDD, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "rfo.sql"
	queryName := "rfo-view"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = "rfo-view"
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
		"PRINCIPAL_RISK_TYPES", "RISK_SUB_TYPE", "RISK_FRAMEWORK_OWNER", "RISK_REPORTING_LEAD", "PR_COUNT", "CRM_COUNT", "CDE_COUNT",
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
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *RFOService) GetRightTable(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	fileName := tabs + ".sql"
	queryName := "right-grid"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = "right-grid"
	gridArgs.PageNumber = pageNumber
	gridArgs.RowsPerPage = rowsPerPage

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, toolkit.ToString(systemID))

	///////// --------------------------------------------------DROPDOWN FILTER
	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	gridArgs.DropdownFilter = append(gridArgs.DropdownFilter,
		searchDDM.GetString("PrincipalRiskType"),
		searchDDM.GetString("RiskSubType"),
	)

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "", "", "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("PRINCIPAL_RISK"),
			colFilterM.GetString("RISK_SUB"),
			colFilterM.GetString("PR_RATIONALE"),
			colFilterM.GetString("CRM_NAME"),
			colFilterM.GetString("CRM_RATIONALE"),
			colFilterM.GetString("ASSOC_CDES"),
			colFilterM.GetString("CDE_RATIONALE"),
		)
	}

	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *RFOService) GetPriorityTable(system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "rfo.sql"
	queryName := "rfo-priority"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = "rfo-priority"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"PRIORITY_REPORT", "PRIORITY_REPORT_RATIONALE", "CRM_NAME", "CRM_RATIONALE", "CDE_NAME", "CDE_RATIONALE",
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
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}
