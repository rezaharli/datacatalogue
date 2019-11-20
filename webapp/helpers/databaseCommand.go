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
	TableName    string
	SqlQuery     string
	SelectFields []string

	PageNumber  int
	RowsPerPage int

	DefaultSort  []string
	OrderBy      string
	IsDescending bool

	GroupCol string

	GlobalFilterWhere map[string]interface{}
	GlobalFilterType  map[string]interface{}
	ColumnFilterWhere map[string]interface{}
	ColumnFilterType  map[string]interface{}

	Results     interface{}
	ResultTotal int
}

func (s *DBcmd) BuildQuery(param *SqlQueryParam) error {
	var err error

	//split the query
	tempSplittedFrom := strings.Split(strings.ReplaceAll(param.SqlQuery, "\n", " "), " FROM ")
	var splittedFROM []string
	splittedFROM = append(splittedFROM, tempSplittedFrom[0])
	if len(tempSplittedFrom) > 2 {
		var temp []string
		for i := 1; i < len(tempSplittedFrom); i++ {
			temp = append(temp, tempSplittedFrom[i])
		}

		splittedFROM = append(splittedFROM, strings.Join(temp, " FROM "))
	} else {
		splittedFROM = append(splittedFROM, tempSplittedFrom[1])
	}

	tempSplittedWHERE := strings.Split(splittedFROM[1], " WHERE ")
	var splittedWHERE []string
	if len(tempSplittedWHERE) > 2 {
		var temp []string
		for i := 0; i < len(tempSplittedWHERE)-1; i++ {
			temp = append(temp, tempSplittedWHERE[i])
		}
		splittedWHERE = append(splittedWHERE, strings.Join(temp, " WHERE "))
		splittedWHERE = append(splittedWHERE, tempSplittedWHERE[len(tempSplittedWHERE)-1])
	} else if len(tempSplittedWHERE) >= 2 {
		splittedWHERE = append(splittedWHERE, tempSplittedWHERE[0])
		splittedWHERE = append(splittedWHERE, tempSplittedWHERE[1])
	} else {
		splittedWHERE = append(splittedWHERE, tempSplittedWHERE[0])
	}

	var splittedGROUPBY []string
	var splittedORDERBY []string
	if len(splittedWHERE) > 1 {
		splittedGROUPBY = strings.Split(splittedWHERE[1], "GROUP BY")
		if len(splittedGROUPBY) > 1 {
			splittedORDERBY = strings.Split(splittedGROUPBY[1], "ORDER BY")
		} else {
			splittedORDERBY = strings.Split(splittedWHERE[1], "ORDER BY")
		}
	}

	selectQuery := splittedFROM[0]
	fromQuery := splittedWHERE[0]

	whereQuery := ""
	if len(splittedGROUPBY) >= 2 {
		whereQuery = splittedGROUPBY[0]
	} else {
		if len(splittedGROUPBY) >= 1 {
			whereQuery = splittedORDERBY[0]
		}
	}

	groupbyQuery := ""
	if len(splittedGROUPBY) >= 2 {
		if len(splittedORDERBY) >= 1 {
			groupbyQuery = splittedORDERBY[0]
		} else {
			groupbyQuery = splittedGROUPBY[1]
		}
	}

	orderbyQuery := ""
	if len(splittedORDERBY) >= 2 {
		orderbyQuery = splittedORDERBY[1]
	}

	isContainsDistinct := strings.Contains(strings.ToUpper(selectQuery), strings.ToUpper("DISTINCT"))
	tmpSelect := TabToSpace(strings.ReplaceAll(strings.ReplaceAll(selectQuery, "SELECT ", " "), "DISTINCT ", " "))
	lines := []string{}
	fixedlines := map[string]string{}

	for {
		line := ""
		i := strings.Index(tmpSelect, " AS ")
		if i > -1 {
			beforeAS := tmpSelect[:i]
			afterAS := tmpSelect[i+1:]

			j := strings.Index(afterAS, ",")
			if j > -1 {
				beforeComma := afterAS[:j]
				afterComma := afterAS[j+1:]

				line = beforeAS + beforeComma
				tmpSelect = afterComma
			} else {
				break
			}
		} else {
			break
		}

		lines = append(lines, line)
	}

	mapAS := map[string]string{}
	for _, line := range lines {
		splittedAS := strings.Split(TabToSpace(line), " AS ")

		if strings.TrimSpace(splittedAS[1]) == "ID" {
			mapAS["ID"] = "ID"
		} else {
			mapAS[strings.TrimSpace(splittedAS[1])] = strings.TrimSpace(splittedAS[0])
		}

		//recreate select query if SelectFields defined
		if len(param.SelectFields) > 0 {
			orderFields := strings.Split(orderbyQuery, ",")
			for i := range orderFields {
				orderFields[i] = strings.ReplaceAll(orderFields[i], " ASC", "")
				orderFields[i] = strings.ReplaceAll(orderFields[i], " DESC", "")
				orderFields[i] = strings.TrimSpace(orderFields[i])
			}

			for _, field := range param.SelectFields {
				fieldName := strings.TrimSpace(splittedAS[0])
				fieldAlias := strings.TrimSpace(splittedAS[1])

				isFieldInsideOrder := false
				for _, field := range orderFields {
					if strings.ToUpper(field) == strings.ToUpper(fieldName) || strings.ToUpper(field) == strings.ToUpper(fieldAlias) {
						isFieldInsideOrder = true
					}
				}

				if strings.ToUpper(fieldAlias) == strings.ToUpper(field) || isFieldInsideOrder == true {
					fixedlines[fieldAlias] = line
				}
			}
		}
	}

	//recreate select query if SelectFields defined
	if len(param.SelectFields) > 0 {
		selectQuery = "SELECT "
		if isContainsDistinct {
			selectQuery += "DISTINCT\n"
		}

		i := 0
		for _, line := range fixedlines {
			selectQuery += line
			if i < len(fixedlines)-1 {
				selectQuery += ", "
			}
			i++
		}
	}

	appendWhereQueryWithFilters := func(paramFilterWhere, paramFilterType map[string]interface{}) {
		additionalWhereQuery := `( `

		i := 0
		for key, val := range paramFilterWhere {
			if i != 0 {
				additionalWhereQuery += `AND `
			}
			i++

			if filterType, ok := paramFilterType[key]; ok && filterType != nil {
				if filterType.(string) == "eq" {
					appendAdditionalWhere := func(value interface{}) {
						intVal, err := strconv.Atoi(toolkit.ToString(value))
						if err != nil {
							if value == "NA" {
								if mapAS[key] != "" {
									additionalWhereQuery += `upper(` + mapAS[key] + `) IS NULL `
								} else {
									additionalWhereQuery += `upper(` + key + `) IS NULL `
								}
							} else {
								replacedVal := strings.ReplaceAll(toolkit.ToString(value), "'", "''")

								if mapAS[key] != "" {
									additionalWhereQuery += `upper(NVL(` + mapAS[key] + `, ' ')) = upper('` + replacedVal + `') `
								} else {
									additionalWhereQuery += `upper(NVL(` + key + `, ' ')) = upper('` + replacedVal + `') `
								}
							}
						} else {
							if mapAS[key] != "" {
								additionalWhereQuery += `upper(` + mapAS[key] + `) = upper('` + toolkit.ToString(intVal) + `') `
							} else {
								additionalWhereQuery += `upper(` + key + `) = upper('` + toolkit.ToString(intVal) + `') `
							}
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
						if mapAS[key] != "" {
							additionalWhereQuery += `upper(` + mapAS[key] + `) IS NULL `
						} else {
							additionalWhereQuery += `upper(` + key + `) IS NULL `
						}
					} else {
						replacedVal := strings.ReplaceAll(toolkit.ToString(value), "'", "''")

						if mapAS[key] != "" {
							additionalWhereQuery += `upper(NVL(` + mapAS[key] + `, ' ')) LIKE upper('%` + replacedVal + `%') `
						} else {
							additionalWhereQuery += `upper(NVL(` + key + `, ' ')) LIKE upper('%` + replacedVal + `%') `
						}
					}
				} else {
					if mapAS[key] != "" {
						additionalWhereQuery += `upper(` + mapAS[key] + `) LIKE upper('%` + toolkit.ToString(intVal) + `%') `
					} else {
						additionalWhereQuery += `upper(` + key + `) LIKE upper('%` + toolkit.ToString(intVal) + `%') `
					}
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

	if len(param.GlobalFilterWhere) > 0 {
		appendWhereQueryWithFilters(param.GlobalFilterWhere, param.GlobalFilterType)
	}

	if len(param.ColumnFilterWhere) > 0 {
		appendWhereQueryWithFilters(param.ColumnFilterWhere, param.ColumnFilterType)
	}

	orders := []string{}
	if len(param.DefaultSort) > 0 {
		for _, val := range param.DefaultSort {
			order := val + " ASC"
			orders = append(orders, order)
		}
	}

	if param.OrderBy != "" {
		order := param.OrderBy + ` `

		if param.IsDescending {
			order += `DESC `
		}

		orders = append(orders, order)
	}

	if strings.TrimSpace(orderbyQuery) != "" {
		orderbyQuery = strings.Join(orders, ", ") + ", " + orderbyQuery
	} else {
		orderbyQuery = strings.Join(orders, ", ")
	}

	// combine it back
	selectQuery = strings.ReplaceAll(selectQuery, ",", ",\n")
	param.SqlQuery = strings.TrimSpace(selectQuery) + "\nFROM\n" + strings.TrimSpace(fromQuery)
	if strings.TrimSpace(whereQuery) != "" {
		param.SqlQuery = strings.TrimSpace(param.SqlQuery) + "\nWHERE\n" + strings.TrimSpace(whereQuery)
	}
	if strings.TrimSpace(groupbyQuery) != "" {
		param.SqlQuery = strings.TrimSpace(param.SqlQuery) + "\nGROUP BY\n" + strings.TrimSpace(groupbyQuery)
	}
	if strings.TrimSpace(orderbyQuery) != "" {
		param.SqlQuery = strings.TrimSpace(param.SqlQuery) + "\nORDER BY\n" + strings.TrimSpace(orderbyQuery)
	}

	return err
}

func (s *DBcmd) ExecuteSQLQuery(param SqlQueryParam) error {
	queryTime := time.Now()
	var err error

	sqlQuery := param.SqlQuery
	if !(param.PageNumber == 0 && param.RowsPerPage == 0) {
		err = s.BuildQuery(&param)

		sqlQuery = "SELECT * FROM (\n"
		sqlQuery += "	SELECT t.*, rownum as rn FROM (\n"
		sqlQuery += "		" + param.SqlQuery + "\n"
		sqlQuery += "	) t "

		if param.RowsPerPage > 0 {
			sqlQuery = sqlQuery + " WHERE rownum <= " + toolkit.ToString(param.PageNumber*param.RowsPerPage) + " "
		}

		sqlQuery += ") "

		if param.RowsPerPage > 0 {
			sqlQuery = sqlQuery + `WHERE rn >= ` + toolkit.ToString(((param.PageNumber-1)*param.RowsPerPage)+1) + " "
		}
	}

	conn := Database()
	cursor := conn.Cursor(dbflex.From(param.TableName).SQL(sqlQuery), nil)
	defer cursor.Close()

	err = cursor.Fetchs(param.Results, 0)
	if err != nil {
		return err
	}

	toolkit.Println(sqlQuery, "\nqueryTime:", time.Since(queryTime).Seconds())
	a := param.Results.(*[]toolkit.M)
	toolkit.Println("fetched results:", len(*a), "\n--------------------------------------------------------------")
	return err
}

func (s *DBcmd) ExecuteSQLQueryHeaderOpts(param SqlQueryParam) error {
	queryTime := time.Now()
	var err error

	sqlQuery := param.SqlQuery
	if !(param.PageNumber == 0 && param.RowsPerPage == 0) {
		err = s.BuildQuery(&param)

		sqlQuery = param.SqlQuery
	}

	conn := Database()
	cursor := conn.Cursor(dbflex.From(param.TableName).SQL(sqlQuery), nil)
	defer cursor.Close()

	err = cursor.Fetchs(param.Results, 0)
	if err != nil {
		return err
	}

	toolkit.Println(sqlQuery, "\nqueryTime:", time.Since(queryTime).Seconds())
	a := param.Results.(*[]toolkit.M)
	toolkit.Println("fetched results:", len(*a), "\n--------------------------------------------------------------")
	return err
}

func (DBcmd) ExecuteSQLQueryRowCount(param SqlQueryParam) error {
	sqlQuery := param.SqlQuery
	if !(param.PageNumber == 0 && param.RowsPerPage == 0) {
		sqlQuery = ""

		//split the query
		tempSplittedFrom := strings.Split(strings.ReplaceAll(param.SqlQuery, "\n", " "), " FROM ")
		var splittedFROM []string
		splittedFROM = append(splittedFROM, tempSplittedFrom[0])
		if len(tempSplittedFrom) > 2 {
			var temp []string
			for i := 1; i < len(tempSplittedFrom); i++ {
				temp = append(temp, tempSplittedFrom[i])
			}

			splittedFROM = append(splittedFROM, strings.Join(temp, " FROM "))
		} else {
			splittedFROM = append(splittedFROM, tempSplittedFrom[1])
		}

		tempSplittedWHERE := strings.Split(splittedFROM[1], " WHERE ")
		var splittedWHERE []string
		if len(tempSplittedWHERE) > 2 {
			var temp []string
			for i := 0; i < len(tempSplittedWHERE)-1; i++ {
				temp = append(temp, tempSplittedWHERE[i])
			}
			splittedWHERE = append(splittedWHERE, strings.Join(temp, " WHERE "))
			splittedWHERE = append(splittedWHERE, tempSplittedWHERE[len(tempSplittedWHERE)-1])
		} else if len(tempSplittedWHERE) >= 2 {
			splittedWHERE = append(splittedWHERE, tempSplittedWHERE[0])
			splittedWHERE = append(splittedWHERE, tempSplittedWHERE[1])
		} else {
			splittedWHERE = append(splittedWHERE, tempSplittedWHERE[0])
		}

		var splittedGROUPBY []string
		var splittedORDERBY []string
		if len(splittedWHERE) > 1 {
			splittedGROUPBY = strings.Split(splittedWHERE[1], "GROUP BY")
			if len(splittedGROUPBY) > 1 {
				splittedORDERBY = strings.Split(splittedGROUPBY[1], "ORDER BY")
			} else {
				splittedORDERBY = strings.Split(splittedWHERE[1], "ORDER BY")
			}
		}

		selectQuery := splittedFROM[0]
		fromQuery := splittedWHERE[0]

		whereQuery := ""
		if len(splittedGROUPBY) >= 2 {
			whereQuery = splittedGROUPBY[0]
		} else {
			if len(splittedGROUPBY) >= 1 {
				whereQuery = splittedORDERBY[0]
			}
		}

		groupbyQuery := ""
		if len(splittedGROUPBY) >= 2 {
			if len(splittedORDERBY) >= 1 {
				groupbyQuery = splittedORDERBY[0]
			} else {
				groupbyQuery = splittedGROUPBY[1]
			}
		}

		orderbyQuery := ""
		if len(splittedORDERBY) >= 2 {
			orderbyQuery = splittedORDERBY[1]
		}

		tmpSelect := TabToSpace(strings.ReplaceAll(strings.ReplaceAll(selectQuery, "SELECT ", " "), "DISTINCT ", " "))
		lines := []string{}

		for {
			line := ""
			i := strings.Index(tmpSelect, " AS ")
			if i > -1 {
				beforeAS := tmpSelect[:i]
				afterAS := tmpSelect[i+1:]

				j := strings.Index(afterAS, ",")
				if j > -1 {
					beforeComma := afterAS[:j]
					afterComma := afterAS[j+1:]

					line = beforeAS + beforeComma
					tmpSelect = afterComma
				} else {
					break
				}
			} else {
				break
			}

			lines = append(lines, line)
		}

		mapAS := map[string]string{}
		for _, line := range lines {
			splittedAS := strings.Split(TabToSpace(line), " AS ")

			if strings.TrimSpace(splittedAS[1]) == "ID" {
				mapAS["ID"] = "ID"
			} else {
				mapAS[strings.TrimSpace(splittedAS[1])] = strings.TrimSpace(splittedAS[0])
			}
		}

		appendWhereQueryWithFilters := func(paramFilterWhere, paramFilterType map[string]interface{}) {
			additionalWhereQuery := `( `

			i := 0
			for key, val := range paramFilterWhere {
				if i != 0 {
					additionalWhereQuery += `AND `
				}
				i++

				if filterType, ok := paramFilterType[key]; ok && filterType != nil {
					if filterType.(string) == "eq" {
						appendAdditionalWhere := func(value interface{}) {
							intVal, err := strconv.Atoi(toolkit.ToString(value))
							if err != nil {
								if value == "NA" {
									if mapAS[key] != "" {
										additionalWhereQuery += `upper(` + mapAS[key] + `) IS NULL `
									} else {
										additionalWhereQuery += `upper(` + key + `) IS NULL `
									}
								} else {
									replacedVal := strings.ReplaceAll(toolkit.ToString(value), "'", "''")

									if mapAS[key] != "" {
										additionalWhereQuery += `upper(NVL(` + mapAS[key] + `, ' ')) = upper('` + replacedVal + `') `
									} else {
										additionalWhereQuery += `upper(NVL(` + key + `, ' ')) = upper('` + replacedVal + `') `
									}
								}
							} else {
								if mapAS[key] != "" {
									additionalWhereQuery += `upper(` + mapAS[key] + `) = upper('` + toolkit.ToString(intVal) + `') `
								} else {
									additionalWhereQuery += `upper(` + key + `) = upper('` + toolkit.ToString(intVal) + `') `
								}
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
							if mapAS[key] != "" {
								additionalWhereQuery += `upper(` + mapAS[key] + `) IS NULL `
							} else {
								additionalWhereQuery += `upper(` + key + `) IS NULL `
							}
						} else {
							replacedVal := strings.ReplaceAll(toolkit.ToString(value), "'", "''")

							if mapAS[key] != "" {
								additionalWhereQuery += `upper(NVL(` + mapAS[key] + `, ' ')) LIKE upper('%` + replacedVal + `%') `
							} else {
								additionalWhereQuery += `upper(NVL(` + key + `, ' ')) LIKE upper('%` + replacedVal + `%') `
							}
						}
					} else {
						if mapAS[key] != "" {
							additionalWhereQuery += `upper(` + mapAS[key] + `) LIKE upper('%` + toolkit.ToString(intVal) + `%') `
						} else {
							additionalWhereQuery += `upper(` + key + `) LIKE upper('%` + toolkit.ToString(intVal) + `%') `
						}
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

		if len(param.GlobalFilterWhere) > 0 {
			appendWhereQueryWithFilters(param.GlobalFilterWhere, param.GlobalFilterType)
		}

		if len(param.ColumnFilterWhere) > 0 {
			appendWhereQueryWithFilters(param.ColumnFilterWhere, param.ColumnFilterType)
		}

		//replace the query with count query
		selectQuery = `SELECT COUNT(DISTINCT rownum) RESULT_COUNT `

		// combine it back
		selectQuery = strings.ReplaceAll(selectQuery, ",", ",\n")
		param.SqlQuery = strings.TrimSpace(selectQuery) + "\nFROM\n" + strings.TrimSpace(fromQuery)
		if strings.TrimSpace(whereQuery) != "" {
			param.SqlQuery = strings.TrimSpace(param.SqlQuery) + "\nWHERE\n" + strings.TrimSpace(whereQuery)
		}
		if strings.TrimSpace(groupbyQuery) != "" {
			param.SqlQuery = strings.TrimSpace(param.SqlQuery) + "\nGROUP BY\n" + strings.TrimSpace(groupbyQuery)
		}
		if strings.TrimSpace(orderbyQuery) != "" {
			param.SqlQuery = strings.TrimSpace(param.SqlQuery) + "\nORDER BY\n" + strings.TrimSpace(orderbyQuery)
		}

		sqlQuery += param.SqlQuery + "\n"
	}

	conn := Database()
	cursor := conn.Cursor(dbflex.From(param.TableName).SQL(sqlQuery), nil)
	defer cursor.Close()

	err := cursor.Fetchs(param.Results, 0)
	if err != nil {
		return err
	}

	return err
}
