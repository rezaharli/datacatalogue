package services

import (
	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"

	"github.com/eaciit/toolkit"
)

type Base struct {
}

func NewBaseService() *Base {
	return new(Base)
}

type GridArgs struct {
	QueryFilePath string
	QueryName     string

	MainArgs       []interface{}
	ColumnFilter   []interface{}
	DropdownFilter []interface{}
	GroupCol       string

	PageNumber  int
	RowsPerPage int
}

func (s *Base) ExecuteGridQueryFromFile(gridArgs GridArgs) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	args := make([]interface{}, 0)
	args = append(args, gridArgs.MainArgs...)
	args = append(args, gridArgs.ColumnFilter...)
	args = append(args, gridArgs.DropdownFilter...)

	///////// --------------------------------------------------BUILD QUERY FROM ARGS
	q, err := h.BuildQueryFromFile(gridArgs.QueryFilePath, gridArgs.QueryName, args...)
	if err != nil {
		return nil, 0, err
	}

	///////// --------------------------------------------------EXECUTE
	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewSystemModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		ResultTotal: resultTotal,
		PageNumber:  gridArgs.PageNumber,
		RowsPerPage: gridArgs.RowsPerPage,
		GroupCol:    gridArgs.GroupCol,
	})
	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}
