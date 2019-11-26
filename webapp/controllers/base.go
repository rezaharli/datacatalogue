package controllers

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strings"

	"git.eaciitapp.com/sebar/knot"
	"github.com/eaciit/clit"
	"github.com/eaciit/toolkit"

	"github.com/novalagung/gubrak"

	"eaciit/datacatalogue/webapp/helpers"
	h "eaciit/datacatalogue/webapp/helpers"
	s "eaciit/datacatalogue/webapp/services"
)

type Base struct {
}

func NewBaseController() *Base {
	return new(Base)
}

func (c *Base) GetHeaderOpts(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	headerArgs := s.HeaderArgs{
		Filename:  payload.GetString("Filename"),
		Queryname: payload.GetString("Queryname"),
		FieldName: payload.GetString("FieldName"),

		Param1:        payload.GetString("System"),
		Param2:        payload.GetString("DspName"),
		Filter:        payload.GetString("Filter"),
		GlobalFilters: payload.Get("GlobalFilters"),
		ColumnFilters: payload.Get("Filters"),
	}

	if payload.Has("LoggedInID") == true {
		headerArgs.LoggedInID = payload.GetString("LoggedInID")
	} else {
		headerArgs.LoggedInID = "-"
	}

	resultRows, err := s.NewBaseService().GetHeaderOpts(headerArgs)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	resultArray, err := gubrak.Map(resultRows, func(v toolkit.M) string {
		return v.GetString(headerArgs.FieldName)
	})

	h.WriteResultOK(k, res, resultArray)
}

func (c *Base) GetRowCount(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	headerArgs := s.HeaderArgs{
		Filename:  payload.GetString("Filename"),
		Queryname: payload.GetString("Queryname"),
		FieldName: payload.GetString("FieldName"),

		Param1:        payload.GetString("System"),
		Param2:        payload.GetString("DspName"),
		Filter:        payload.GetString("Filter"),
		GlobalFilters: payload.Get("GlobalFilters"),
		ColumnFilters: payload.Get("Filters"),
	}

	if payload.Has("LoggedInID") == true {
		headerArgs.LoggedInID = payload.GetString("LoggedInID")
	} else {
		headerArgs.LoggedInID = "-"
	}

	resultRows, err := s.NewBaseService().GetRowCount(headerArgs)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, resultRows[0].GetInt("RESULT_COUNT"))
}

func (c *Base) ExportToCsv(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	rowSelect := payload.Get("RowSelect").([]interface{})

	headerArgs := s.HeaderArgs{
		Filename:  payload.GetString("Filename"),
		Queryname: payload.GetString("Queryname"),
		Headers:   payload.Get("Headers").([]interface{}),

		Param1:        payload.GetString("System"),
		Param2:        payload.GetString("DspName"),
		Filter:        payload.GetString("Filter"),
		GlobalFilters: payload.Get("GlobalFilters"),
		ColumnFilters: payload.Get("Filters"),
	}

	if payload.Has("LoggedInID") == true {
		headerArgs.LoggedInID = payload.GetString("LoggedInID")
	} else {
		headerArgs.LoggedInID = "-"
	}

	resultRows := []toolkit.M{}
	if len(rowSelect) > 0 {
		for _, row := range rowSelect {
			data, err := toolkit.ToM(row)
			if err != nil {
				h.WriteResultError(k, res, err.Error())
				return
			}

			resultRows = append(resultRows, data)
		}
	} else {
		resultRows, err = s.NewBaseService().GetExportData(headerArgs)
		if err != nil {
			h.WriteResultError(k, res, err.Error())
			return
		}
	}

	folderPath := filepath.Join(clit.ExeDir(), "csv")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, os.ModePerm)
	}

	fileName := "result.csv"
	os.Remove(filepath.Join(folderPath, fileName))
	file, err := os.Create(filepath.Join(folderPath, fileName))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.Comma = '\t'
	defer w.Flush()

	var headers []string
	for _, value := range headerArgs.Headers {
		mapVal := value.(map[string]interface{})
		if mapVal["exportable"].(bool) == true {
			headers = append(headers, mapVal["text"].(string))
		}
	}

	if err := w.Write(headers); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	for _, resultRow := range resultRows {
		var row []string

		for _, value := range headerArgs.Headers {
			mapVal := value.(map[string]interface{})
			if mapVal["exportable"].(bool) == true {
				row = append(row, resultRow.GetString(mapVal["value"].(string)))
			}
		}

		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	h.WriteResultOK(k, res, fileName)
}

func (c *Base) GetDetails(payload toolkit.M,
	detailsFunction func(payload toolkit.M) (interface{}, int, error),
	dropdownFunction func(payload toolkit.M) (interface{}, int, error),
) (interface{}, interface{}, error) {

	detail, _, err := detailsFunction(payload)
	if err != nil {
		return nil, nil, err
	}

	ddSource, _, err := dropdownFunction(payload)
	if err != nil {
		return nil, nil, err
	}

	return detail, ddSource, nil
}

func (c *Base) GetDetailAndDropdown(payload toolkit.M, detailsFunction func(payload toolkit.M) (interface{}, int, error), dropdownFunction func(payload toolkit.M) (interface{}, int, error)) (interface{}, interface{}, error) {
	mappedDetails, err := c.GetDetail(payload, detailsFunction)
	if err != nil {
		return nil, nil, err
	}

	var mappedddSource interface{}
	if dropdownFunction != nil {
		mappedddSource, err = c.GetDropdownSource(payload, dropdownFunction)
		if err != nil {
			return nil, nil, err
		}
	}

	return mappedDetails, mappedddSource, nil
}

func (c *Base) GetDetail(payload toolkit.M, detailsFunction func(payload toolkit.M) (interface{}, int, error)) (interface{}, error) {
	detail, _, err := detailsFunction(payload)
	if err != nil {
		return nil, err
	}

	groupedDetail, err := gubrak.GroupBy(detail, func(each toolkit.M) string {
		return each.GetString("ID")
	})
	if err != nil {
		return nil, err
	}

	mappedDetails, err := gubrak.Map(helpers.ObjectKeys(groupedDetail), func(key string, i int) toolkit.M {
		tmp := groupedDetail.(map[string]([]toolkit.M))

		ret := toolkit.M{}
		ret["ID"] = key
		ret["Values"] = tmp[key]

		return ret
	})
	if err != nil {
		return nil, err
	}

	return mappedDetails, nil
}

func (c *Base) GetDropdownSource(payload toolkit.M, dropdownFunction func(payload toolkit.M) (interface{}, int, error)) (interface{}, error) {
	ddSource, _, err := dropdownFunction(payload)
	if err != nil {
		return nil, err
	}

	mappedddSource, err := gubrak.Map(ddSource, func(v toolkit.M, i int) toolkit.M {
		for _, key := range helpers.ObjectKeys(v) {
			stringVal := toolkit.ToString(v[key])
			trimmedStringVal := strings.TrimSpace(stringVal)

			if trimmedStringVal == "" {
				stringVal = "NA"
			}

			v[key] = stringVal
		}

		return v
	})

	return mappedddSource, nil
}
