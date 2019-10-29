package controllers

import (
	"strings"

	"git.eaciitapp.com/sebar/knot"
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

		Param1:       payload.GetString("System"),
		Param2:       payload.GetString("DspName"),
		Filter:       payload.GetString("Filter"),
		ScopeFilters: payload.Get("Filters"),
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
