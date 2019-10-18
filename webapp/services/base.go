package services

import (
	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/eaciit/clit"
	"github.com/eaciit/toolkit"
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

	Filter string
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

	args = append(args, toolkit.ToString(headerArgs.Filter))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", fileName)
	q, err := h.BuildQueryFromFile(filePath, queryName, []string{}, args...)
	if err != nil {
		return nil, err
	}

	funcLog(funcName(), fileName, queryName)

	///////// FILTER
	q = `SELECT DISTINCT ` + headerArgs.FieldName + `
		FROM (
		` + q + `
	) `

	additionalWhere := make(map[string]interface{}, 0)
	additionalWhere[headerArgs.FieldName] = headerArgs.Filter

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:       m.NewCategoryModel().TableName(),
		SqlQuery:        q,
		Results:         &resultRows,
		AdditionalWhere: additionalWhere,
		PageNumber:      1,
		RowsPerPage:     -1,
		GroupCol:        headerArgs.FieldName,
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
}

func (s *Base) ExecuteGridQueryFromFile(gridArgs GridArgs) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)

	log.Println("ExecuteGridQueryFromFile", 115)

	resultTotal := 0

	args := make([]interface{}, 0)
	args = append(args, gridArgs.MainArgs...)
	args = append(args, gridArgs.DropdownFilter...)

	log.Println("ExecuteGridQueryFromFile", 124)

	///////// --------------------------------------------------BUILD QUERY FROM ARGS
	q, err := h.BuildQueryFromFile(gridArgs.QueryFilePath, gridArgs.QueryName, gridArgs.Colnames, args...)
	log.Println("ExecuteGridQueryFromFile", 128)
	if err != nil {
		log.Println("ExecuteGridQueryFromFile", 130)
		return nil, 0, err
	}

	additionalWhere := make(map[string]interface{}, len(gridArgs.Colnames))
	log.Println("ExecuteGridQueryFromFile", 135)
	for i, colname := range gridArgs.Colnames {

		log.Println("ExecuteGridQueryFromFile", 138)
		log.Println("ExecuteGridQueryFromFile", 139, gridArgs.ColumnFilter[i])
		if gridArgs.ColumnFilter[i] != nil {
			log.Println("ExecuteGridQueryFromFile", 141)
			log.Println("ExecuteGridQueryFromFile", 142, gridArgs.ColumnFilter[i])
			log.Println("ExecuteGridQueryFromFile", 143)
			log.Println("ExecuteGridQueryFromFile", 144, reflect.TypeOf(gridArgs.ColumnFilter[i]))
			log.Println("ExecuteGridQueryFromFile", 145)
			log.Println("ExecuteGridQueryFromFile", 146, reflect.TypeOf(gridArgs.ColumnFilter[i]).Kind())
			switch reflect.TypeOf(gridArgs.ColumnFilter[i]).Kind() {
			case reflect.Slice:
				log.Println("ExecuteGridQueryFromFile", 149)
				log.Println("ExecuteGridQueryFromFile", 150, reflect.ValueOf(gridArgs.ColumnFilter[i]))
				s := reflect.ValueOf(gridArgs.ColumnFilter[i])

				log.Println("ExecuteGridQueryFromFile", 153)
				log.Println("ExecuteGridQueryFromFile", 154, s.Len())
				if s.Len() > 0 {
					log.Println("ExecuteGridQueryFromFile", 156)
					additionalWhere[colname] = gridArgs.ColumnFilter[i]
				}
			default:
				log.Println("ExecuteGridQueryFromFile", 160)
				if toolkit.ToString(gridArgs.ColumnFilter[i]) != "" {
					log.Println("ExecuteGridQueryFromFile", 162)
					additionalWhere[colname] = gridArgs.ColumnFilter[i]
				}
			}
		}
	}

	log.Println("ExecuteGridQueryFromFile", 169)
	splitted := strings.Split(gridArgs.OrderBy, ".")
	log.Println("ExecuteGridQueryFromFile", 171)
	orderBy := splitted[len(splitted)-1]
	log.Println("ExecuteGridQueryFromFile", 173)

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
