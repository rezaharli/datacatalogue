package services

import (
	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/eaciit/clit"
	"github.com/eaciit/toolkit"
	"github.com/novalagung/gubrak"
)

type Base struct {
}

func NewBaseService() *Base {
	return new(Base)
}

type HeaderArgs struct {
	LoggedInID string
	Param1     string
	Param2     string

	Filename  string
	Queryname string
	FieldName string
	Headers   []interface{}

	Filter       string
	ScopeFilters interface{}
}

func (s *Base) GetHeaderOpts(headerArgs HeaderArgs) ([]toolkit.M, error) {
	fileName := headerArgs.Filename
	queryName := headerArgs.Queryname

	resultRows := make([]toolkit.M, 0)

	q := ""
	args := make([]interface{}, 0)

	if headerArgs.LoggedInID != "-" {
		if headerArgs.LoggedInID != "" {
			args = append(args, "MY", headerArgs.LoggedInID)
		} else {
			args = append(args, "ALL", "0000000")
		}
	}

	if headerArgs.Param1 != "" {
		args = append(args, toolkit.ToString(headerArgs.Param1))
	}
	if headerArgs.Param2 != "" {
		args = append(args, toolkit.ToString(headerArgs.Param2))
	}

	// args = append(args, toolkit.ToString(headerArgs.Filter))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, err
	}

	funcLog(funcName(), fileName, queryName)
	var fieldNames []string
	fieldNames = append(fieldNames, headerArgs.FieldName)

	additionalWhere := make(map[string]interface{}, 0)
	if strings.TrimSpace(headerArgs.Filter) != "" {
		additionalWhere[headerArgs.FieldName] = headerArgs.Filter
	}

	var columnFilterType map[string]interface{}
	for key, filter := range headerArgs.ScopeFilters.(map[string]interface{}) {
		if key == "filterTypes" {
			columnFilterType = filter.(map[string]interface{})
			continue
		}

		if filter == nil {
			continue
		}

		fieldNames = append(fieldNames, key)

		switch reflect.TypeOf(filter).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(filter)

			if s.Len() > 0 {
				additionalWhere[key] = filter
			}
		default:
			if toolkit.ToString(filter) != "" {
				additionalWhere[key] = filter
			}
		}
	}

	uniqueFieldNames, err := gubrak.Uniq(fieldNames)
	if err != nil {
		return nil, err
	}

	joinedFieldnames, err := gubrak.Join(uniqueFieldNames, ", ")
	if err != nil {
		return nil, err
	}

	///////// FILTER
	q = `SELECT DISTINCT ` + joinedFieldnames + `
		FROM (
		` + q + `
	) `

	err = h.NewDBcmd().ExecuteSQLQueryHeaderOpts(h.SqlQueryParam{
		TableName:        m.NewCategoryModel().TableName(),
		SqlQuery:         q,
		Results:          &resultRows,
		AdditionalWhere:  additionalWhere,
		PageNumber:       1,
		RowsPerPage:      -1,
		GroupCol:         headerArgs.FieldName,
		ColumnFilterType: columnFilterType,
	})

	if err != nil {
		return nil, err
	}

	return resultRows, nil
}

func (s *Base) GetRowCount(headerArgs HeaderArgs) ([]toolkit.M, error) {
	fileName := headerArgs.Filename
	queryName := headerArgs.Queryname

	resultRows := make([]toolkit.M, 0)

	q := ""
	args := make([]interface{}, 0)

	if headerArgs.LoggedInID != "-" {
		if headerArgs.LoggedInID != "" {
			args = append(args, "MY", headerArgs.LoggedInID)
		} else {
			args = append(args, "ALL", "0000000")
		}
	}

	if headerArgs.Param1 != "" {
		args = append(args, toolkit.ToString(headerArgs.Param1))
	}
	if headerArgs.Param2 != "" {
		args = append(args, toolkit.ToString(headerArgs.Param2))
	}

	// args = append(args, toolkit.ToString(headerArgs.Filter))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, err
	}

	var fieldNames []string
	fieldNames = append(fieldNames, headerArgs.FieldName)

	additionalWhere := make(map[string]interface{}, 0)
	if strings.TrimSpace(headerArgs.Filter) != "" {
		additionalWhere[headerArgs.FieldName] = headerArgs.Filter
	}

	var columnFilterType map[string]interface{}
	for key, filter := range headerArgs.ScopeFilters.(map[string]interface{}) {
		if key == "filterTypes" {
			columnFilterType = filter.(map[string]interface{})
			continue
		}

		if filter == nil {
			continue
		}

		fieldNames = append(fieldNames, key)

		switch reflect.TypeOf(filter).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(filter)

			if s.Len() > 0 {
				additionalWhere[key] = filter
			}
		default:
			if toolkit.ToString(filter) != "" {
				additionalWhere[key] = filter
			}
		}
	}

	err = h.NewDBcmd().ExecuteSQLQueryRowCount(h.SqlQueryParam{
		TableName:        m.NewCategoryModel().TableName(),
		SqlQuery:         q,
		Results:          &resultRows,
		AdditionalWhere:  additionalWhere,
		PageNumber:       1,
		RowsPerPage:      -1,
		ColumnFilterType: columnFilterType,
	})

	if err != nil {
		return nil, err
	}

	return resultRows, nil
}

func (s *Base) GetExportData(headerArgs HeaderArgs) ([]toolkit.M, error) {
	fileName := headerArgs.Filename
	queryName := headerArgs.Queryname

	resultRows := make([]toolkit.M, 0)

	q := ""
	args := make([]interface{}, 0)

	if headerArgs.LoggedInID != "-" {
		if headerArgs.LoggedInID != "" {
			args = append(args, "MY", headerArgs.LoggedInID)
		} else {
			args = append(args, "ALL", "0000000")
		}
	}

	if headerArgs.Param1 != "" {
		args = append(args, toolkit.ToString(headerArgs.Param1))
	}
	if headerArgs.Param2 != "" {
		args = append(args, toolkit.ToString(headerArgs.Param2))
	}

	// args = append(args, toolkit.ToString(headerArgs.Filter))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, err
	}

	funcLog(funcName(), fileName, queryName)
	additionalWhere := make(map[string]interface{}, 0)
	var columnFilterType map[string]interface{}
	for key, filter := range headerArgs.ScopeFilters.(map[string]interface{}) {
		if key == "filterTypes" {
			columnFilterType = filter.(map[string]interface{})
			continue
		}

		if filter == nil {
			continue
		}

		switch reflect.TypeOf(filter).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(filter)

			if s.Len() > 0 {
				additionalWhere[key] = filter
			}
		default:
			if toolkit.ToString(filter) != "" {
				additionalWhere[key] = filter
			}
		}
	}

	///////// FILTER
	q = `SELECT *
		FROM (
		` + q + `
	) `

	err = h.NewDBcmd().ExecuteSQLQueryHeaderOpts(h.SqlQueryParam{
		TableName:        m.NewCategoryModel().TableName(),
		SqlQuery:         q,
		Results:          &resultRows,
		AdditionalWhere:  additionalWhere,
		PageNumber:       1,
		RowsPerPage:      -1,
		GroupCol:         headerArgs.FieldName,
		ColumnFilterType: columnFilterType,
	})

	if err != nil {
		return nil, err
	}

	return resultRows, nil
}

type GridArgs struct {
	QueryFilePath string
	QueryName     string

	MainArgs         []interface{}
	ColumnFilter     []interface{}
	ColumnFilterType map[string]interface{}
	DropdownFilter   []interface{}
	Colnames         []string
	GroupCol         string

	PageNumber   int
	RowsPerPage  int
	OrderBy      string
	IsDescending bool

	DefaultSort []string
}

func (s *Base) ExecuteGridQueryFromFile(gridArgs GridArgs) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)

	resultTotal := 0

	args := make([]interface{}, 0)
	args = append(args, gridArgs.MainArgs...)
	args = append(args, gridArgs.DropdownFilter...)

	///////// --------------------------------------------------BUILD QUERY FROM ARGS
	q, err := h.BuildQueryFromFile(gridArgs.QueryFilePath, gridArgs.QueryName, gridArgs.Colnames, args...)
	if err != nil {
		return nil, 0, err
	}

	additionalWhere := make(map[string]interface{}, len(gridArgs.Colnames))
	for i, colname := range gridArgs.Colnames {

		if gridArgs.ColumnFilter[i] != nil {
			switch reflect.TypeOf(gridArgs.ColumnFilter[i]).Kind() {
			case reflect.Slice:
				s := reflect.ValueOf(gridArgs.ColumnFilter[i])

				if s.Len() > 0 {
					additionalWhere[colname] = gridArgs.ColumnFilter[i]
				}
			default:
				if toolkit.ToString(gridArgs.ColumnFilter[i]) != "" {
					additionalWhere[colname] = gridArgs.ColumnFilter[i]
				}
			}
		}
	}

	splitted := strings.Split(gridArgs.OrderBy, ".")
	orderBy := splitted[len(splitted)-1]

	///////// --------------------------------------------------EXECUTE
	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:        m.NewSystemModel().TableName(),
		SqlQuery:         q,
		Results:          &resultRows,
		ResultTotal:      resultTotal,
		PageNumber:       gridArgs.PageNumber,
		RowsPerPage:      gridArgs.RowsPerPage,
		OrderBy:          orderBy,
		IsDescending:     gridArgs.IsDescending,
		GroupCol:         gridArgs.GroupCol,
		AdditionalWhere:  additionalWhere,
		ColumnFilterType: gridArgs.ColumnFilterType,
		DefaultSort:      gridArgs.DefaultSort,
	})
	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func funcLog(funcName string, fileName string, queryName string) {
	fmt.Println("--------------------------------------------------------------", "\n", "Function   : ", funcName, "\n", "File Name  : ", fileName, "\n", "Query Name : ", queryName, "\n--------------------------------------------------------------")
}
