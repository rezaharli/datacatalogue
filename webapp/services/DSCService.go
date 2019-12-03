package services

import (
	"path/filepath"

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

func (s *DSCService) GetEdmpDDDropdowns(payload toolkit.M) (interface{}, int, error) {
	fileName := "edmp.sql"
	queryName := "edmp-dd-dropdowns"

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

func (s *DSCService) GetEdmpDDTechnicalTable(system string, globalFilter, colFilter interface{}, pagination toolkit.M, defaultSort []string) (res []toolkit.M, total int, err error) {
	fileName := "edmp.sql"
	queryName := "edmp-dd-technical"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")
	gridArgs.DefaultSort = defaultSort

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"COUNTRY", "BUSINESS_SEGMENT", "EDM_SOURCE_SYSTEM_NAME", "CLUSTER_NAME", "TIER", "ITAM",
		"DATA_XRAY", "EDM_GOLDEN_SYSTEM_NAME", "DATABASE_NAME", "TABLE_NAME", "COLUMN_NAME", "DATA_TYPE", "COLUMN_LENGTH", "NULLABLE",
		"PRIMARY_KEY", "CERTIFIED", "PII", "DATA_LINEAGE", "BUSINESS_ALIAS_NAME", "BUSINESS_ALIAS_DESCRIPTION", "CDE", "DATA_LENGTH",
		"EXAMPLE", "DERIVED", "DERIVATION_LOGIC", "SOURCED_FROM_UPSTREAM", "SYSTEM_CHECKS",
	)

	globalFilterM, err := toolkit.ToM(globalFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.GlobalFilter = append(gridArgs.GlobalFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.GlobalFilter = append(gridArgs.GlobalFilter, globalFilterM.Get(colname))

			filterTypes := globalFilterM.Get("filterTypes")
			if filterTypes != nil {
				gridArgs.GlobalFilterType = globalFilterM.Get("filterTypes").(map[string]interface{})
			}
		}
	}

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colFilterM.Get(colname))

			filterTypes := colFilterM.Get("filterTypes")
			if filterTypes != nil {
				gridArgs.ColumnFilterType = colFilterM.Get("filterTypes").(map[string]interface{})
			}
		}
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	gridArgs.GroupCol = "TABLE_NAME"
	result, total, err := s.Base.ExecuteGridQueryFromFile(gridArgs)

	res = result
	return
}

func (s *DSCService) GetEdmpDDBusinessTable(system string, globalFilter, colFilter interface{}, pagination toolkit.M, defaultSort []string) (res []toolkit.M, total int, err error) {
	fileName := "edmp.sql"
	queryName := "edmp-dd-business"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")
	gridArgs.DefaultSort = defaultSort

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"COUNTRY", "BUSINESS_SEGMENT", "EDM_SOURCE_SYSTEM_NAME", "CLUSTER_NAME", "TIER", "ITAM",
		"EDM_GOLDEN_SYSTEM_NAME", "DATABASE_NAME", "TABLE_NAME", "COLUMN_NAME",
		"TABLE_DESCRIPTION", "COLUMN_DESCRIPTION", "BUSINESS_TERM", "BUSINESS_DESCRIPTION", "DETERMINES_CLIENT_LOCATION",
		"DETERMINES_ACCOUNT", "BUSINESS_SEGMENT", "PRODUCT_CATEGORY", "BUSINESS_ALIAS_NAME", "BUSINESS_ALIAS_DESCRIPTION", "CDE",
		"DOMAIN", "SUBDOMAIN", "DOMAIN_OWNER", "BUSINESS_TERM_DESCRIPTION",
	)

	globalFilterM, err := toolkit.ToM(globalFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.GlobalFilter = append(gridArgs.GlobalFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.GlobalFilter = append(gridArgs.GlobalFilter, globalFilterM.Get(colname))

			filterTypes := globalFilterM.Get("filterTypes")
			if filterTypes != nil {
				gridArgs.GlobalFilterType = globalFilterM.Get("filterTypes").(map[string]interface{})
			}
		}
	}

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colFilterM.Get(colname))

			filterTypes := colFilterM.Get("filterTypes")
			if filterTypes != nil {
				gridArgs.ColumnFilterType = colFilterM.Get("filterTypes").(map[string]interface{})
			}
		}
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	gridArgs.GroupCol = "TABLE_NAME"
	result, total, err := s.Base.ExecuteGridQueryFromFile(gridArgs)

	res = result
	return
}
