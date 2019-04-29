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
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "rfo.sql")
	gridArgs.QueryName = "rfo-view"
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

	// filterPriorityReportName := ""
	// if search != "" {
	// 	filterPriorityReportName = search
	// } else {
	// 	filterPriorityReportName = searchDDM.GetString("PriorityReportName")
	// }
	// gridArgs.DropdownFilter = append(gridArgs.DropdownFilter,
	// 	filterPriorityReportName,
	// 	searchDDM.GetString("RiskReportingLead"),
	// )

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "", "", "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("PRINCIPAL_RISK_TYPES"),
			colFilterM.GetString("RISK_SUB_TYPE"),
			colFilterM.GetString("RISK_FRAMEWORK_OWNER"),
			colFilterM.GetString("RISK_REPORTING_LEAD"),
			colFilterM.GetString("PR_COUNT"),
			colFilterM.GetString("CRM_COUNT"),
			colFilterM.GetString("CDE_COUNT"),
		)
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
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

func (s *RFOService) GetPriorityTable(system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "rfo.sql")
	gridArgs.QueryName = "rfo-priority"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "", "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("PRIORITY_REPORT"),
			colFilterM.GetString("PRIORITY_REPORT_RATIONALE"),
			colFilterM.GetString("CRM_NAME"),
			colFilterM.GetString("CRM_RATIONALE"),
			colFilterM.GetString("CDE_NAME"),
			colFilterM.GetString("CDE_RATIONALE"),
		)
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}
