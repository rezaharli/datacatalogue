package services

import (
	"path/filepath"

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

func (s *DPOService) GetLeftTable(tabs, loggedinid, search string, searchDD, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	gridArgs.QueryName = "dpo-view"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	if loggedinid != "" {
		gridArgs.MainArgs = append(gridArgs.MainArgs, "MY", loggedinid)
	} else {
		gridArgs.MainArgs = append(gridArgs.MainArgs, "ALL", "0000000")
	}

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"DSP_NAME", "DSP_OWNER",
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

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DPOService) GetHomepageCounts(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	system := payload.GetString("System")
	args = append(args, system, system)

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	q, err := h.BuildQueryFromFile(filePath, "dpo-view-homepage", []string{}, args...)
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

func (s *DPOService) GetDataelementsTable(system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	gridArgs.QueryName = "dpo-dataelements"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"SYSTEM_NAME", "ITAM_ID", "ALIAS_NAME", "CDE", "TABLE_NAME", "COLUMN_NAME", "ULT_SYSTEM_NAME", "DATA_SLA",
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

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *DPOService) GetDatalineageTable(system string, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	gridArgs.QueryName = "dpo-datalineage"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	gridArgs.MainArgs = append(gridArgs.MainArgs, system)

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"PR_NAME", "SYSTEM_NAME", "CDE_COUNT",
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

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Left"), payload.GetString("Right"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
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

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Left"), payload.GetString("Right"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
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

func (s *DPOService) GetDetailsImmediatePrecedingSystem(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Left"), payload.GetString("Right"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	q, err := h.BuildQueryFromFile(filePath, "details-immediate-preceding-system", []string{}, args...)
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
			q += `AND SYSTEM_NAME = '` + otherArgs[0] + `' `
		}
		if otherArgs[1] != "" {
			q += `AND ITAM_ID = '` + otherArgs[1] + `' `
		}
		if otherArgs[2] != "" {
			q += `AND TABLE_NAME = '` + otherArgs[2] + `' `
		}
		if otherArgs[3] != "" {
			q += `AND COLUMN_NAME = '` + otherArgs[3] + `' `
		}
		if otherArgs[4] != "" {
			q += `AND DATA_ELEMENT = '` + otherArgs[4] + `' `
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

func (s *DPOService) GetddSourceImmediatePrecedingSystem(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Left"), payload.GetString("Right"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	q, err := h.BuildQueryFromFile(filePath, "details-immediate-preceding-system", []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	///////// FILTER
	q = `SELECT DISTINCT 
			SYSTEM_NAME, 
			ITAM_ID, 
			TABLE_NAME, 
			COLUMN_NAME, 
			DATA_ELEMENT
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

func (s *DPOService) GetDetailsUltimateSourceSystem(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Left"), payload.GetString("Right"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	q, err := h.BuildQueryFromFile(filePath, "details-ultimate-source-system", []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	otherArgs := make([]string, 0)
	otherArgs = append(otherArgs,
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
			q += `AND SYSTEM_NAME = '` + otherArgs[0] + `' `
		}
		if otherArgs[1] != "" {
			q += `AND ITAM_ID = '` + otherArgs[1] + `' `
		}
		if otherArgs[2] != "" {
			q += `AND TABLE_NAME = '` + otherArgs[2] + `' `
		}
		if otherArgs[3] != "" {
			q += `AND COLUMN_NAME = '` + otherArgs[3] + `' `
		}
		if otherArgs[4] != "" {
			q += `AND DATA_ELEMENT = '` + otherArgs[4] + `' `
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

func (s *DPOService) GetddSourceUltimateSourceSystem(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Left"), payload.GetString("Right"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	q, err := h.BuildQueryFromFile(filePath, "details-ultimate-source-system", []string{}, args...)
	if err != nil {
		return nil, 0, err
	}

	///////// FILTER
	q = `SELECT DISTINCT 
			SYSTEM_NAME, 
			ITAM_ID, 
			TABLE_NAME, 
			COLUMN_NAME, 
			DATA_ELEMENT
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

func (s *DPOService) GetDetailsDomainView(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Right"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	q, err := h.BuildQueryFromFile(filePath, "details-domain-view", []string{}, args...)
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

func (s *DPOService) GetddSourceDomainView(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Left"), payload.GetString("Right"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	q, err := h.BuildQueryFromFile(filePath, "details-domain-view", []string{}, args...)
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

func (s *DPOService) GetDetailsDataStandards(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Left"), payload.GetString("Right"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	q, err := h.BuildQueryFromFile(filePath, "details-data-standards", []string{}, args...)
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

func (s *DPOService) GetddSourceDataStandards(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, payload.GetString("Left"), payload.GetString("Right"))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", "dpo.sql")
	q, err := h.BuildQueryFromFile(filePath, "details-data-standards", []string{}, args...)
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
