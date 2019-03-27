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

func (s *RFOService) GetLeftTable(tabs, loggedinid, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
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

	filterPriorityReportName := ""
	if search != "" {
		filterPriorityReportName = search
	} else {
		filterPriorityReportName = searchDDM.GetString("PriorityReportName")
	}
	gridArgs.DropdownFilter = append(gridArgs.DropdownFilter,
		filterPriorityReportName,
		searchDDM.GetString("RiskReportingLead"),
	)

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("PRIORITY_REPORT"),
			colFilterM.GetString("RR_LEAD"),
			colFilterM.GetString("BANK_ID"),
		)
	}

	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *RFOService) GetRightTable(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
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
