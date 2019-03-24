package services

import (
	"path/filepath"

	"github.com/eaciit/clit"
	"github.com/eaciit/toolkit"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type RFOService struct {
}

func NewRFOService() *RFOService {
	ret := new(RFOService)
	return ret
}

func (s *RFOService) GetLeftTable(tabs, loggedinid, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
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

	filterPriorityReportName := ""
	if search != "" {
		filterPriorityReportName = search
	} else {
		filterPriorityReportName = searchDDM.GetString("PriorityReportName")
	}
	filterRiskReportingLead := searchDDM.GetString("RiskReportingLead")

	///////// DROPDOWN FILTER
	q = `SELECT * FROM (
		` + q + `
	) WHERE (
		upper(PRIORITY_REPORT) LIKE upper('%` + filterPriorityReportName + `%')
		AND upper(RR_LEAD) LIKE upper('%` + filterRiskReportingLead + `%')
	) `

	colFilterM, err := toolkit.ToM(colFilter)
	cf := make([]string, 0)
	if err != nil {
		cf = append(cf, "", "", "")
	} else {
		cf = append(cf, colFilterM.GetString("PRIORITY_REPORT"), colFilterM.GetString("RR_LEAD"), colFilterM.GetString("BANK_ID"))
	}

	///////// COLUMN FILTER
	if cf[0] != "" || cf[1] != "" || cf[2] != "" {
		q += `AND (
			upper(PRIORITY_REPORT) LIKE upper('%` + cf[0] + `%')
			AND upper(RR_LEAD) LIKE upper('%` + cf[1] + `%')
			AND upper(BANK_ID) LIKE upper('%` + cf[2] + `%')
		) `
	}

	///////// COUNT
	q = `SELECT res.*, 
			COUNT(DISTINCT PRIORITY_REPORT) OVER () COUNT_PRIORITY_REPORT,
			COUNT(DISTINCT RR_LEAD) OVER () COUNT_RR_LEAD,
			COUNT(DISTINCT BANK_ID) OVER () COUNT_BANK_ID
		FROM (
			` + q + `
		) res `

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewCategoryModel().TableName(),
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

func (s *RFOService) GetRightTable(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
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

	filterPrincipalRiskType := searchDDM.GetString("PrincipalRiskType")
	filterRiskSubType := searchDDM.GetString("RiskSubType")

	///////// DROPDOWN FILTER
	q = `SELECT * FROM (
		` + q + `
	) WHERE (
		upper(PRINCIPAL_RISK) LIKE upper('%` + filterPrincipalRiskType + `%')
		AND upper(RISK_SUB) LIKE upper('%` + filterRiskSubType + `%')
	) `

	colFilterM, err := toolkit.ToM(colFilter)
	cf := make([]string, 0)
	if err != nil {
		cf = append(cf, "", "", "", "", "", "", "")
	} else {
		cf = append(cf,
			colFilterM.GetString("PRINCIPAL_RISK"),
			colFilterM.GetString("RISK_SUB"),
			colFilterM.GetString("PR_RATIONALE"),
			colFilterM.GetString("CRM_NAME"),
			colFilterM.GetString("CRM_RATIONALE"),
			colFilterM.GetString("ASSOC_CDES"),
			colFilterM.GetString("CDE_RATIONALE"),
		)
	}

	///////// COLUMN FILTER
	if cf[0] != "" || cf[1] != "" || cf[2] != "" || cf[3] != "" || cf[4] != "" || cf[5] != "" || cf[6] != "" {
		q += `AND (
			upper(PRINCIPAL_RISK) LIKE upper('%` + cf[0] + `%')
			AND upper(RISK_SUB) LIKE upper('%` + cf[1] + `%')
			AND upper(PR_RATIONALE) LIKE upper('%` + cf[2] + `%')
			AND upper(CRM_NAME) LIKE upper('%` + cf[3] + `%')
			AND upper(CRM_RATIONALE) LIKE upper('%` + cf[4] + `%')
			AND upper(ASSOC_CDES) LIKE upper('%` + cf[5] + `%')
			AND upper(CDE_RATIONALE) LIKE upper('%` + cf[6] + `%')
		) `
	}

	///////// COUNT
	q = `SELECT res.*, 
			COUNT(DISTINCT PRINCIPAL_RISK) OVER () COUNT_PRINCIPAL_RISK,
			COUNT(DISTINCT RISK_SUB) OVER () COUNT_RISK_SUB,
			COUNT(DISTINCT PR_RATIONALE) OVER () COUNT_PR_RATIONALE,
			COUNT(DISTINCT CRM_NAME) OVER () COUNT_CRM_NAME,
			COUNT(DISTINCT CRM_RATIONALE) OVER () COUNT_CRM_RATIONALE,
			COUNT(DISTINCT ASSOC_CDES) OVER () COUNT_ASSOC_CDES,
			COUNT(DISTINCT CDE_RATIONALE) OVER () COUNT_CDE_RATIONALE
		FROM (
			` + q + `
		) res `

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewCategoryModel().TableName(),
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
