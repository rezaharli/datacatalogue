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

func (s *DSCService) GetAllSystem(tabs, loggedinid, search string, searchDD, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "dsc.sql"
	queryName := "dsc-view"

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
		"SYSTEM_NAME", "ITAM_ID", "DATASET_CUSTODIAN", "BANK_ID",
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

	gridArgs.GroupCol = "SYSTEM_NAME"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DSCService) GetHomepageCounts(payload toolkit.M) (interface{}, int, error) {
	fileName := "dsc.sql"
	queryName := "dsc-view-homepage"

	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	system := payload.GetString("System")
	args = append(args, system, system, system, system)

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

func (s *DSCService) GetCDETable(system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "dsc.sql"
	queryName := "dsc-view-cde"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"CDE", "DESCRIPTION", "UPSTREAM_SYSTEM", "TABLE_NAME", "COLUMN_NAME", "DSP_NAME", "PROCESS_OWNER",
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

	gridArgs.GroupCol = "CDE"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DSCService) GetCDPTable(system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "dsc.sql"
	queryName := "dsc-view-cdp"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"DSP_NAME", "PROCESS_OWNER", "CDE_COUNT",
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

func (s *DSCService) GetCDPCDETable(system, dspName string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "dsc.sql"
	queryName := "dsc-view-cdp-cde"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system, dspName)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"CDE", "DESCRIPTION", "TABLE_NAME", "COLUMN_NAME",
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

func (s *DSCService) GetIARCTable(system string, colFilter interface{}, pagination toolkit.M) (res []toolkit.M, total int, err error) {
	fileName := "dsc.sql"
	queryName := "dsc-view-policy-ia"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"INFORMATION_ASSET_NAMES", "INFORMATION_ASSET_DESCRIPTION", "CONFIDENTIALITY", "INTEGRITY", "AVAILABILITY", "OVERALL_CIA_RATING",
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

	gridArgs.GroupCol = "-"
	result, total, err := s.Base.ExecuteGridQueryFromFile(gridArgs)

	res = result
	return
}

func (s *DSCService) GetIARCPersonalDataTable(system string, colFilter interface{}, pagination toolkit.M) (res []toolkit.M, total int, err error) {
	fileName := "dsc.sql"
	queryName := "dsc-view-policy"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"TABLE_NAME", "COLUMN_NAME", "BUSINESS_ALIAS_NAME", "BUSINESS_ALIAS_DESCRIPTION", "CDE_YES_NO", "PII_FLAG",
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

	gridArgs.GroupCol = "-"
	result, total, err := s.Base.ExecuteGridQueryFromFile(gridArgs)

	res = result
	return
}

func (s *DSCService) GetInterfacesTable(system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "dsc.sql"
	queryName := "dsc-view-interfaces"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"IMM_INTERFACE", "CDE_COUNT", "PROCESS_OWNER",
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

	gridArgs.GroupCol = "IMM_INTERFACE"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DSCService) GetInterfacesCDETable(system, dspName string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	fileName := "dsc.sql"
	queryName := "dsc-view-interfaces-cde"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system, dspName)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"CDE", "DESCRIPTION", "TABLE_NAME", "COLUMN_NAME", "IMM_PREC_SYSTEM", "IMM_SUCC_SYSTEM",
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

	gridArgs.GroupCol = "CDE"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DSCService) GetDDTable(system string, colFilter interface{}, pagination toolkit.M) (res []toolkit.M, total int, err error) {
	fileName := "dsc.sql"
	queryName := "dsc-view-dd"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"SYSTEM_NAME", "ITAM_ID", "TABLE_NAME", "COLUMN_NAME", "BUSINESS_ALIAS_NAME", "BUSINESS_ALIAS_DESCRIPTION", "CDE_YES_NO", "DATA_TYPE",
		"DATA_LENGTH", "EXAMPLE", "DERIVED_YES_NO", "DERIVATION_LOGIC", "SOURCED_FROM_UPSTREAM_YES_NO", "SYSTEM_CHECKS",
		"DOMAIN", "SUBDOMAIN", "DOMAIN_OWNER", "BUSINESS_TERM", "BUSINESS_TERM_DESCRIPTION",
		"INFORMATION_ASSET_NAMES", "INFORMATION_ASSET_DESCRIPTION", "CONFIDENTIALITY", "INTEGRITY", "AVAILABILITY", "OVERALL_CIA_RATING", "RECORD_CATEGORIES", "PII_FLAG",
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

	gridArgs.GroupCol = "-"
	result, total, err := s.Base.ExecuteGridQueryFromFile(gridArgs)

	res = result
	return
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

func (s *DSCService) GetEdmpDDTechnicalTable(system string, colFilter interface{}, pagination toolkit.M) (res []toolkit.M, total int, err error) {
	fileName := "edmp.sql"
	queryName := "edmp-dd-technical"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"COUNTRY", "BUSINESS_SEGMENT", "EDM_SOURCE_SYSTEM_NAME", "CLUSTER_NAME", "TIER", "ITAM",
		"EDM_GOLDEN_SYSTEM_NAME", "DATABASE_NAME", "TABLE_NAME", "COLUMN_NAME",
		"DATA_TYPE", "COLUMN_LENGTH", "NULLABLE", "PRIMARY_KEY", "CERTIFIED", "DATA_XRAY", "DATA_LINEAGE", "BUSINESS_ALIAS_NAME",
		"BUSINESS_ALIAS_DESCRIPTION", "CDE", "DATA_LENGTH", "EXAMPLE", "DERIVED", "DERIVATION_LOGIC", "SOURCED_FROM_UPSTREAM", "SYSTEM_CHECKS",
	)

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

func (s *DSCService) GetEdmpDDBusinessTable(system string, colFilter interface{}, pagination toolkit.M) (res []toolkit.M, total int, err error) {
	fileName := "edmp.sql"
	queryName := "edmp-dd-business"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

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

func (s *DSCService) GetEdmpDDConsumptionTable(system string, colFilter interface{}, pagination toolkit.M) (res []toolkit.M, total int, err error) {
	fileName := "edmp.sql"
	queryName := "edmp-dd-consumption"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"COUNTRY", "BUSINESS_SEGMENT", "EDM_SOURCE_SYSTEM_NAME", "CLUSTER_NAME", "TIER", "ITAM",
		"EDM_GOLDEN_SYSTEM_NAME", "DATABASE_NAME", "TABLE_NAME", "COLUMN_NAME",
		"CONSUMING_APPLICATION", "CONSUMING_APPLICATION_ITAM", "CONSUMING_APPLICATION_OWNER", "CONSUMER_DESCRIPTION",
		"TECH_CONTACT", "BUSINESS_OWNERSHIP", "ACCESS_ROLE", "ROLE_DESCRIPTION", "CONSUMING_TECH_METADATA",
	)

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

func (s *DSCService) GetEdmpIarcPersonalTable(system string, colFilter interface{}, pagination toolkit.M) (res []toolkit.M, total int, err error) {
	fileName := "edmp.sql"
	queryName := "edmp-iarc-personal"

	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	gridArgs.QueryName = queryName
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	funcLog(funcName(), fileName, queryName)

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"EDM_SOURCE_SYSTEM_NAME", "DATABASE_NAME", "TABLE_NAME", "COLUMN_NAME", "BUSINESS_ALIAS_NAME", "BUSINESS_ALIAS_DESCRIPTION", "CDE", "PII",
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

	gridArgs.GroupCol = "-"
	result, total, err := s.Base.ExecuteGridQueryFromFile(gridArgs)

	res = result
	return
}

func (s *DSCService) GetDetailsLeftPanel(payload toolkit.M) (interface{}, int, error) {
	fileName := "dsc.sql"
	queryName := "details-left-panel"

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
	if payload.GetString("TableName") != "" && payload.GetString("ColumnName") != "" {
		otherArgs = append(otherArgs, "", "")
	} else {
		otherArgs = append(otherArgs, payload.GetString("Right"), payload.GetString("Column"))
	}

	otherArgs = append(otherArgs,
		payload.GetString("TableName"),
		payload.GetString("ColumnName"),
	)

	///////// FILTER
	q = `SELECT rownum, a.* FROM (
		` + q + `
	) a
	WHERE ((
			tmtid = '` + otherArgs[0] + `'
			AND tmcid = '` + otherArgs[1] + `'
		) OR (
			table_name = '` + otherArgs[2] + `'
			AND column_name = '` + otherArgs[3] + `' `
	q += `)) `
	// q += `AND rownum = 1 `

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

func (s *DSCService) GetDetailsRightPanel(payload toolkit.M) (interface{}, int, error) {
	fileName := "dsc.sql"
	queryName := "details"

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
	if payload.GetString("TableName") != "" && payload.GetString("ColumnName") != "" {
		otherArgs = append(otherArgs, "", "")
	} else {
		otherArgs = append(otherArgs, payload.GetString("Right"), payload.GetString("Column"))
	}

	otherArgs = append(otherArgs,
		payload.GetString("TableName"),
		payload.GetString("ColumnName"),
		payload.GetString("ScreenLabel"),
		payload.GetString("BusinessTerm"),
		payload.GetString("Prec"),
		payload.GetString("PrecIncoming"),
		payload.GetString("Succ"),
		payload.GetString("SuccIncoming"),
	)

	///////// FILTER
	q = `SELECT rownum, a.* FROM (
		` + q + `
	) a
	WHERE ((
			tmtid = '` + otherArgs[0] + `'
			AND tmcid = '` + otherArgs[1] + `'
		) OR (
			table_name = '` + otherArgs[2] + `'
			AND column_name = '` + otherArgs[3] + `' `
	if otherArgs[4] != "" {
		q += `AND business_alias_name = '` + otherArgs[4] + `' `
	}
	if otherArgs[5] != "" {
		q += `AND business_term = '` + otherArgs[5] + `' `
	}
	if otherArgs[6] != "" {
		q += `AND IMM_PRECEEDING_SYSTEM = '` + otherArgs[6] + `' `
	}
	if otherArgs[7] != "" {
		q += `AND imm_prec_incoming = '` + otherArgs[7] + `' `
	}
	if otherArgs[8] != "" {
		q += `AND IMM_SUCCEEDING_SYSTEM = '` + otherArgs[8] + `' `
	}
	if otherArgs[9] != "" {
		q += `AND imm_succ_incoming = '` + otherArgs[9] + `' `
	}
	q += `)) `
	// q += `AND rownum = 1 `

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

func (s *DSCService) GetddSourceRightPanel(payload toolkit.M) (interface{}, int, error) {
	fileName := "dsc.sql"
	queryName := "details"

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
	q = `SELECT DISTINCT 
			table_name, 
			column_name, 
			business_alias_name, 
			business_term, 
			IMM_PRECEEDING_SYSTEM, 
			imm_prec_incoming, 
			IMM_SUCCEEDING_SYSTEM, 
			imm_succ_incoming
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
