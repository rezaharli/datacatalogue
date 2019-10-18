package helpers

import (
	"log"
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
	log.Println("ExecuteSQLQuery", 235)

	sqlQuery := param.SqlQuery
	if !(param.PageNumber == 0 && param.RowsPerPage == 0) {
		log.Println("ExecuteSQLQuery", 140)
		sqlQuery = `SELECT a.*, `

		if param.GroupCol != "" {
			log.Println("ExecuteSQLQuery", 244)
			if param.GroupCol == "-" {
				log.Println("ExecuteSQLQuery", 246)
				sqlQuery += `rownum r__, `
				sqlQuery += `COUNT(*) OVER () RESULT_COUNT `
			} else {
				log.Println("ExecuteSQLQuery", 250)
				sqlQuery += `DENSE_RANK() OVER (ORDER BY ` + param.GroupCol + ` ASC ) AS r__, `
				sqlQuery += `COUNT(DISTINCT  ` + param.GroupCol + `) OVER () RESULT_COUNT `
			}
		} else {
			log.Println("ExecuteSQLQuery", 255)
			sqlQuery += `DENSE_RANK() OVER (ORDER BY id ASC ) AS r__, `
			sqlQuery += `COUNT(DISTINCT  id) OVER () RESULT_COUNT `
		}
		log.Println("ExecuteSQLQuery", 259)

		if param.OrderBy != "" {
			log.Println("ExecuteSQLQuery", 262)
			param.OrderBy = `ORDER BY regexp_replace(regexp_replace(UPPER(` + param.OrderBy + `), '(\d+)', lpad('0', 20, '0')||'\1'), '0*?(\d{21}(\D|$))', '\1') `
		}
		log.Println("ExecuteSQLQuery", 265)

		if param.IsDescending {
			log.Println("ExecuteSQLQuery", 268)
			param.OrderBy += `DESC `
		}
		log.Println("ExecuteSQLQuery", 271)

		sqlQuery += `
			FROM
			(
				` + param.SqlQuery + `
			) a `

		log.Println("ExecuteSQLQuery", 279)

		if len(param.AdditionalWhere) > 0 {
			log.Println("ExecuteSQLQuery", 282)
			sqlQuery += `WHERE( `

			i := 0
			for key, val := range param.AdditionalWhere {
				log.Println("ExecuteSQLQuery", 287)
				if i != 0 {
					log.Println("ExecuteSQLQuery", 289)
					sqlQuery += `AND `
				}
				i++
				log.Println("ExecuteSQLQuery", 293)

				if filterType, ok := param.ColumnFilterType[key]; ok && filterType != nil {
					log.Println("ExecuteSQLQuery", 296)
					if filterType.(string) == "eq" {
						log.Println("ExecuteSQLQuery", 298)
						appendAdditionalWhere := func(value interface{}) {
							log.Println("ExecuteSQLQuery", 300)
							intVal, err := strconv.Atoi(toolkit.ToString(value))
							log.Println("ExecuteSQLQuery", 302)
							if err != nil {
								log.Println("ExecuteSQLQuery", 304)
								if value == "NA" {
									log.Println("ExecuteSQLQuery", 306)
									sqlQuery += `
								upper(` + key + `) IS NULL `
								} else {
									log.Println("ExecuteSQLQuery", 310)
									replacedVal := strings.ReplaceAll(toolkit.ToString(value), "'", "''")
									sqlQuery += `
									upper(NVL(` + key + `, ' ')) = upper('` + replacedVal + `') `
								}
								log.Println("ExecuteSQLQuery", 315)
							} else {
								log.Println("ExecuteSQLQuery", 317)
								sqlQuery += `
								upper(` + key + `) = upper('` + toolkit.ToString(intVal) + `') `
							}
							log.Println("ExecuteSQLQuery", 321)
						}

						log.Println("ExecuteSQLQuery", 324)
						log.Println("ExecuteSQLQuery", 325, val)
						log.Println("ExecuteSQLQuery", 326)
						log.Println("ExecuteSQLQuery", 327, reflect.TypeOf(val))
						log.Println("ExecuteSQLQuery", 328)
						log.Println("ExecuteSQLQuery", 329, reflect.TypeOf(val).Kind())
						switch reflect.TypeOf(val).Kind() {
						case reflect.Slice:
							log.Println("ExecuteSQLQuery", 332)

							log.Println("ExecuteSQLQuery", 334, reflect.ValueOf(val))
							s := reflect.ValueOf(val)

							log.Println("ExecuteSQLQuery", 337)
							log.Println("ExecuteSQLQuery", 337, s.Len())
							for i := 0; i < s.Len(); i++ {
								log.Println("ExecuteSQLQuery", 340)
								log.Println("ExecuteSQLQuery", 341, s.Index(i))
								log.Println("ExecuteSQLQuery", 342, s.Index(i).Interface())
								appendAdditionalWhere(s.Index(i).Interface())

								log.Println("ExecuteSQLQuery", 345)
								if i != s.Len()-1 {
									log.Println("ExecuteSQLQuery", 347)
									sqlQuery += `
									OR `
								}
							}
						default:
							log.Println("ExecuteSQLQuery", 353)
							log.Println("ExecuteSQLQuery", 354, val)
							appendAdditionalWhere(val)
							log.Println("ExecuteSQLQuery", 356)
						}

						log.Println("ExecuteSQLQuery", 359)
						continue
					}
				}

				log.Println("ExecuteSQLQuery", 364)
				log.Println("ExecuteSQLQuery", 365, val)
				log.Println("ExecuteSQLQuery", 366, toolkit.ToString(val))
				intVal, err := strconv.Atoi(toolkit.ToString(val))
				if err != nil {
					log.Println("ExecuteSQLQuery", 369)
					if val == "NA" {
						log.Println("ExecuteSQLQuery", 371)
						sqlQuery += `
						upper(` + key + `) IS NULL `
					} else {
						log.Println("ExecuteSQLQuery", 375)
						replacedVal := strings.ReplaceAll(toolkit.ToString(val), "'", "''")
						log.Println("ExecuteSQLQuery", 377)
						sqlQuery += `
							upper(NVL(` + key + `, ' ')) LIKE upper('%` + replacedVal + `%') `
						log.Println("ExecuteSQLQuery", 379)
					}
					log.Println("ExecuteSQLQuery", 382)
				} else {
					log.Println("ExecuteSQLQuery", 384)
					sqlQuery += `
						upper(` + key + `) LIKE upper('%` + toolkit.ToString(intVal) + `%') `
				}
			}

			log.Println("ExecuteSQLQuery", 390)
			sqlQuery += `
				) `
		}

		log.Println("ExecuteSQLQuery", 395)
		sqlQuery += param.OrderBy

		log.Println("ExecuteSQLQuery", 398)
		if param.RowsPerPage > 0 {
			log.Println("ExecuteSQLQuery", 400)

			sqlQuery = `SELECT * FROM
				(
					` + sqlQuery + ` 
				) WHERE r__ 
				BETWEEN ` + toolkit.ToString(((param.PageNumber-1)*param.RowsPerPage)+1) + ` 
				AND ` + toolkit.ToString(param.PageNumber*param.RowsPerPage) + ` `
		}
		log.Println("ExecuteSQLQuery", 409)
	}
	log.Println("ExecuteSQLQuery", 411)

	conn := Database()
	log.Println("ExecuteSQLQuery", 414)
	cursor := conn.Cursor(dbflex.From(param.TableName).SQL(sqlQuery), nil)
	log.Println("ExecuteSQLQuery", 416)
	defer cursor.Close()
	log.Println("ExecuteSQLQuery", 418)

	err := cursor.Fetchs(param.Results, 0)
	log.Println("ExecuteSQLQuery", 421)
	log.Println("ExecuteSQLQuery", 422, err)

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
