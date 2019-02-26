package helpers

import (
	"reflect"

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
	TableName string
	SqlQuery  string

	Results interface{}
}

func (DBcmd) ExecuteSQLQuery(param SqlQueryParam) error {
	err := Database().Cursor(dbflex.From(param.TableName).SQL(param.SqlQuery), nil).Fetchs(param.Results, 0)
	return err
}
