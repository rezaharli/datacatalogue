package helpers

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	_ "gopkg.in/goracle.v2"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/dbflex"
)

type DBcmd struct{}

func NewDBcmd() *DBcmd {
	return new(DBcmd)
}

type GetAllParam struct {
	Filter      toolkit.M
	SortKey     string
	SortOrder   string
	Skip        int
	Take        int
	TableName   string
	ResultRows  interface{}
	ResultTotal *int
}

func (DBcmd) GetAll(param GetAllParam) error {
	sortKey := param.SortKey
	if sortKey == "" {
		sortKey = "CreatedAt"
	}

	sortOrder := -1
	if param.SortKey == "asc" {
		sortOrder = 1
	}

	pipeRows := make([]toolkit.M, 0)

	if len(param.Filter) > 0 {
		pipeRows = append(pipeRows, toolkit.M{"$match": param.Filter})
	}

	pipeRows = append(pipeRows, toolkit.M{"$sort": toolkit.M{sortKey: sortOrder}})

	if param.Skip > 0 {
		pipeRows = append(pipeRows, toolkit.M{"$skip": param.Skip})
	}
	if param.Take > 0 {
		pipeRows = append(pipeRows, toolkit.M{"$limit": param.Take})
	}

	err := Database().Cursor(
		dbflex.From(param.TableName).Select(), nil,
	).Fetchs(param.ResultRows, 0)
	if err != nil {
		return err
	}

	// ------------ count

	pipeTotal := make([]toolkit.M, 0)
	pipeTotal = append(pipeTotal, toolkit.M{
		"$group": toolkit.M{
			"_id": nil,
			"total": toolkit.M{
				"$sum": 1,
			},
		},
	})

	resultTotal := make([]toolkit.M, 0)
	err = Database().Cursor(
		dbflex.From(param.TableName).Select(), nil,
	).Fetchs(&resultTotal, 0)
	if err != nil {
		return err
	}

	if len(resultTotal) > 0 && param.ResultTotal != nil {
		*param.ResultTotal = resultTotal[0].GetInt("total")
	}

	return nil
}

type GetByParam struct {
	TableName string
	Clause    *dbflex.Filter
	Result    interface{}
}

func (DBcmd) GetBy(param GetByParam) error {
	cursor := Database().Cursor(
		dbflex.
			From(param.TableName).
			Where(param.Clause).
			Select(),
		nil)
	err := cursor.Fetchs(param.Result, 0)
	if err != nil {
		return err
	}

	return nil
}

type AggrParam struct {
	TableName string
	Pipe      []toolkit.M
	Result    interface{}
}

func (DBcmd) Aggr(param AggrParam) error {
	err := Database().Cursor(
		dbflex.From(param.TableName).Command("pipe"),
		toolkit.M{"pipe": param.Pipe},
	).Fetchs(param.Result, 0)
	if err != nil {
		return err
	}

	return nil
}

type InsertParam struct {
	TableName       string
	Data            interface{}
	ContinueOnError bool
}

func (DBcmd) Insert(param InsertParam) error {
	query, err := Database().Prepare(
		dbflex.From(param.TableName).Insert(),
	)
	if err != nil {
		return err
	}

	var lastError error

	switch reflect.TypeOf(param.Data).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(param.Data)

		for i := 0; i < s.Len(); i++ {
			_, err = query.Execute(toolkit.M{}.Set("data", s.Index(i).Interface()))
			lastError = err
			if !param.ContinueOnError && err != nil {
				return err
			}
		}
	default:
		_, err = query.Execute(toolkit.M{}.Set("data", param.Data))
		lastError = err
	}

	return lastError
}

type SaveParam struct {
	TableName       string
	Data            interface{}
	ContinueOnError bool
}

func (DBcmd) Save(param SaveParam) error {
	query, err := Database().Prepare(
		dbflex.From(param.TableName).Save(),
	)
	if err != nil {
		return err
	}

	var lastError error

	switch reflect.TypeOf(param.Data).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(param.Data)

		for i := 0; i < s.Len(); i++ {
			_, err = query.Execute(toolkit.M{}.Set("data", s.Index(i).Interface()))
			lastError = err
			if !param.ContinueOnError && err != nil {
				return err
			}
		}
	default:
		_, err = query.Execute(toolkit.M{}.Set("data", param.Data))
		lastError = err
	}

	return lastError
}

type DeleteParam struct {
	TableName string
	Clause    *dbflex.Filter
}

func (DBcmd) Delete(param DeleteParam) error {
	query := dbflex.From(param.TableName)
	if param.Clause != nil {
		query = query.Where(param.Clause)
	}
	query = query.Delete()

	_, err := Database().Execute(query, nil)
	return err
}

type SqlQueryParam struct {
	TableName        string
	SqlQuery         string
	PageNumber       int
	RowsPerPage      int
	OrderBy          string
	IsDescending     bool
	GroupCol         string
	AdditionalWhere  map[string]interface{}
	ColumnFilterType map[string]interface{}

	Results     interface{}
	ResultTotal int
}

func (DBcmd) ExecuteSQLQuery(param SqlQueryParam) error {
	queryTime := time.Now()

	sqlQuery := param.SqlQuery
	if !(param.PageNumber == 0 && param.RowsPerPage == 0) {
		sqlQuery = ""

		//split the query
		splittedFROM := strings.Split(strings.ReplaceAll(param.SqlQuery, "\n", " "), " FROM ")
		splittedWHERE := strings.Split(splittedFROM[1], " WHERE ")

		var splittedORDERBY []string
		if len(splittedWHERE) > 1 {
			splittedORDERBY = strings.Split(splittedWHERE[1], "ORDER BY")
		}

		selectQuery := splittedFROM[0]
		fromQuery := splittedWHERE[0]

		whereQuery := ""
		orderbyQuery := ""
		if len(splittedORDERBY) > 0 {
			whereQuery = splittedORDERBY[0]
			if len(splittedORDERBY) > 1 {
				orderbyQuery = splittedORDERBY[1]
			}
		}

		if len(param.AdditionalWhere) > 0 {
			additionalWhereQuery := `( `

			i := 0
			for key, val := range param.AdditionalWhere {
				if i != 0 {
					additionalWhereQuery += `AND `
				}
				i++

				if filterType, ok := param.ColumnFilterType[key]; ok && filterType != nil {
					if filterType.(string) == "eq" {
						appendAdditionalWhere := func(value interface{}) {
							intVal, err := strconv.Atoi(toolkit.ToString(value))
							if err != nil {
								if value == "NA" {
									additionalWhereQuery += `upper(` + key + `) IS NULL `
								} else {
									replacedVal := strings.ReplaceAll(toolkit.ToString(value), "'", "''")
									additionalWhereQuery += `upper(NVL(` + key + `, ' ')) = upper('` + replacedVal + `') `
								}
							} else {
								additionalWhereQuery += `upper(` + key + `) = upper('` + toolkit.ToString(intVal) + `') `
							}
						}

						switch reflect.TypeOf(val).Kind() {
						case reflect.Slice:
							s := reflect.ValueOf(val)

							for i := 0; i < s.Len(); i++ {
								if i == 0 {
									additionalWhereQuery += `( `
								}

								appendAdditionalWhere(s.Index(i).Interface())

								if i != s.Len()-1 {
									additionalWhereQuery += `OR `
								} else {
									additionalWhereQuery += `) `
								}
							}
						default:
							appendAdditionalWhere(val)
						}

						continue
					}
				}

				appendAdditionalWhere := func(value interface{}) {
					intVal, err := strconv.Atoi(toolkit.ToString(value))
					if err != nil {
						if value == "NA" {
							additionalWhereQuery += `upper(` + key + `) IS NULL `
						} else {
							replacedVal := strings.ReplaceAll(toolkit.ToString(value), "'", "''")
							additionalWhereQuery += `upper(NVL(` + key + `, ' ')) LIKE upper('%` + replacedVal + `%') `
						}
					} else {
						additionalWhereQuery += `upper(` + key + `) LIKE upper('%` + toolkit.ToString(intVal) + `%') `
					}
				}

				switch reflect.TypeOf(val).Kind() {
				case reflect.Slice:
					s := reflect.ValueOf(val)

					for i := 0; i < s.Len(); i++ {
						if i == 0 {
							additionalWhereQuery += `( `
						}

						appendAdditionalWhere(s.Index(i).Interface())

						if i != s.Len()-1 {
							additionalWhereQuery += `OR `
						} else {
							additionalWhereQuery += `) `
						}
					}
				default:
					appendAdditionalWhere(val)
				}
			}

			additionalWhereQuery += `) `

			if strings.TrimSpace(whereQuery) != "" {
				whereQuery = whereQuery + "\nAND " + additionalWhereQuery
			} else {
				whereQuery = additionalWhereQuery
			}
		}

		if param.RowsPerPage > 0 {
			selectQuery = selectQuery + ",rownum row_num"
		}

		if param.GroupCol != "" {
			if param.GroupCol == "-" {
				// 		sqlQuery += `rownum r__, `
				selectQuery += `,COUNT(*) OVER () RESULT_COUNT `
			} else {
				// 		sqlQuery += `DENSE_RANK() OVER (ORDER BY ` + param.GroupCol + ` ASC ) AS r__, `
				// sqlQuery += `COUNT(DISTINCT  ` + param.GroupCol + `) OVER () RESULT_COUNT `
				selectQuery += `,COUNT(DISTINCT rownum) OVER () RESULT_COUNT `
			}
		} else {
			// 	sqlQuery += `DENSE_RANK() OVER (ORDER BY id ASC ) AS r__, `
			selectQuery += `,COUNT(DISTINCT rownum) OVER () RESULT_COUNT `
		}

		// combine it back
		selectQuery = strings.ReplaceAll(selectQuery, ",", ",\n")
		param.SqlQuery = strings.TrimSpace(selectQuery) + "\nFROM\n" + strings.TrimSpace(fromQuery)
		if strings.TrimSpace(whereQuery) != "" {
			param.SqlQuery = strings.TrimSpace(param.SqlQuery) + "\nWHERE\n" + strings.TrimSpace(whereQuery)
		}
		if strings.TrimSpace(orderbyQuery) != "" {
			param.SqlQuery = strings.TrimSpace(param.SqlQuery) + "\nORDER BY\n" + strings.TrimSpace(orderbyQuery)
		}

		sqlQuery += "SELECT a.* FROM (\n"
		sqlQuery += param.SqlQuery + "\n"
		sqlQuery += ") a "

		if param.RowsPerPage > 0 {
			sqlQuery = sqlQuery + `WHERE row_num BETWEEN ` + toolkit.ToString(((param.PageNumber-1)*param.RowsPerPage)+1) + ` AND ` + toolkit.ToString(param.PageNumber*param.RowsPerPage) + " "
		}

		if param.OrderBy != "" {
			param.OrderBy = `ORDER BY regexp_replace(regexp_replace(UPPER(` + param.OrderBy + `), '(\d+)', lpad('0', 20, '0')||'\1'), '0*?(\d{21}(\D|$))', '\1') `
		}

		if param.IsDescending {
			param.OrderBy += `DESC `
		}

		sqlQuery += param.OrderBy
	}

	conn := Database()
	cursor := conn.Cursor(dbflex.From(param.TableName).SQL(sqlQuery), nil)
	defer cursor.Close()

	err := cursor.Fetchs(param.Results, 0)
	if err != nil {
		return err
	}

	toolkit.Println(sqlQuery, "\nqueryTime:", time.Since(queryTime).Seconds())
	a := param.Results.(*[]toolkit.M)
	toolkit.Println("fetched results:", len(*a), "\n--------------------------------------------------------------")
	return err
}

type PrepareQueryParam struct {
	q                            string
	Cols                         []string
	ColsFilteredBySearchDropdown map[string]string
	ColumnFilters                []string
	SpecialCountCols             []string
	SpecialCountQueries          []string
}

func (DBcmd) PrepareQueryForGrids(p PrepareQueryParam) string {
	checkNotEmpty := func(s []string) bool {
		for _, v := range s {
			if v != "" {
				return true
			}
		}
		return false
	}

	///////// DROPDOWN FILTER
	p.q = `SELECT * FROM (
		` + p.q + `
	) WHERE ( `
	i := 0
	for key, val := range p.ColsFilteredBySearchDropdown {
		if i != 0 {
			p.q += `AND `
		}
		p.q += `upper(` + key + `) LIKE upper('%` + val + `%') `
	}
	p.q += `) `

	if checkNotEmpty(p.ColumnFilters) == true {
		///////// COLUMN FILTER
		p.q += `AND ( `
		for i, col := range p.Cols {
			if i != 0 {
				p.q += `AND `
			}
			p.q += `upper(` + col + `) LIKE upper('%` + p.ColumnFilters[i] + `%') `
		}
		p.q += `) `
	}

	return p.q
}

func (DBcmd) PrepareCountQueryForGrids(p PrepareQueryParam) string {
	ret := p.q

	///////// COUNT
	ret = `SELECT res.*, `
	for i, col := range p.Cols {
		ret += `COUNT(DISTINCT ` + col + `) OVER () COUNT_` + col
		if i != len(p.Cols)-1 {
			ret += `, `
		}
	}
	ret = `COUNT(DISTINCT table_name) OVER () COUNT_table_name,
			COUNT(DISTINCT column_name) OVER () COUNT_column_name,
			COUNT(DISTINCT business_alias_name) OVER () COUNT_business_alias_name,
			(SELECT COUNT (cde_yes_no) FROM ( ` + p.q + `) res2 WHERE cde_yes_no = 1) COUNT_cde_yes_no
		FROM (
			` + p.q + `
		) res `

	return ret
}
