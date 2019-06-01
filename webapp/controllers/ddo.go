package controllers

import (
	"strings"
	"time"

	"github.com/eaciit/toolkit"
	"github.com/novalagung/gubrak"

	"git.eaciitapp.com/sebar/knot"

	"eaciit/datacatalogue/webapp/helpers"
	h "eaciit/datacatalogue/webapp/helpers"
	s "eaciit/datacatalogue/webapp/services"
)

type DDO struct {
	*Base
}

func NewDDOController() *DDO {
	return new(DDO)
}

func (c *DDO) GetLeftTable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	tabs := payload.GetString("Tabs")
	loggedinId := payload.GetString("LoggedInID")
	search := payload.GetString("Search")
	searchDD := payload.Get("SearchDD")
	colFilter := payload.Get("Filters")
	pagination, err := toolkit.ToM(payload.Get("Pagination"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDDOService().GetLeftTable(tabs, loggedinId, search, searchDD, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DDO) GetHomepageCounts(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDDOService().GetHomepageCounts(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DDO) GetBusinesstermTable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	system := payload.GetString("System")
	colFilter := payload.Get("Filters")
	pagination, err := toolkit.ToM(payload.Get("Pagination"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDDOService().GetBusinesstermTable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DDO) GetSystemsTable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	system := payload.GetString("System")
	colFilter := payload.Get("Filters")
	pagination, err := toolkit.ToM(payload.Get("Pagination"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDDOService().GetSystemsTable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DDO) GetSystemsBusinesstermTable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	subdomain := payload.GetString("Subdomain")
	system := payload.GetString("System")
	colFilter := payload.Get("Filters")
	pagination, err := toolkit.ToM(payload.Get("Pagination"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDDOService().GetSystemsBusinesstermTable(subdomain, system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DDO) GetDownstreamTable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	system := payload.GetString("System")
	colFilter := payload.Get("Filters")
	pagination, err := toolkit.ToM(payload.Get("Pagination"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDDOService().GetDownstreamTable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DDO) GetDownstreamBusinesstermTable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	subdomain := payload.GetString("Subdomain")
	system := payload.GetString("System")
	colFilter := payload.Get("Filters")
	pagination, err := toolkit.ToM(payload.Get("Pagination"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDDOService().GetDownstreamBusinesstermTable(subdomain, system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DDO) GetRightTable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	tabs := payload.GetString("Tabs")
	search := payload.GetString("Search")
	searchDD := payload.Get("SearchDD")
	colFilter := payload.Get("Filters")
	systemID := payload.GetInt("SystemID")
	pagination, err := toolkit.ToM(payload.Get("Pagination"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	pageNumber := pagination.GetInt("page")
	rowsPerPage := pagination.GetInt("rowsPerPage")

	systems, _, err := s.NewDDOService().GetRightTable(tabs, systemID, search, searchDD, colFilter, pageNumber, rowsPerPage, toolkit.M{})
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DDO) GetDetails(k *knot.WebContext) {
	processTime := time.Now()
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	selectedDetail, ddVal, mappedddSource, err := c.getDetails(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data := toolkit.M{}
	data.Set("SelectedDetail", selectedDetail)
	data.Set("DDSource", mappedddSource)
	data.Set("DDVal", ddVal)

	h.WriteResultOK(k, res, data)
	toolkit.Println("processTime:", time.Since(processTime).Seconds(), "\n------------------------------------------------------------------------------------")
}

func (c *DDO) getDetails(payload toolkit.M) (interface{}, interface{}, interface{}, error) {
	doInterrupt := func(v string, conds, expectedreses []string) string {
		if v == conds[1] {
			return expectedreses[1]
		} else if v == conds[0] {
			return expectedreses[0]
		}

		return v
	}

	mappedDetails, mappedddSource, err := c.GetDetailAndDropdown(payload, s.NewDDOService().GetDetails, s.NewDDOService().GetddSource)
	if err != nil {
		return nil, nil, nil, err
	}

	ddVal := toolkit.M{}
	selectedDetail := toolkit.M{}
	detailSources := mappedDetails.([]toolkit.M)
	if len(detailSources) > 0 {
		detailSource := detailSources[0]
		detailValues := detailSource.Get("Values").([]toolkit.M)

		for _, key := range helpers.ObjectKeys(detailValues[0]) {
			mappedValues, err := gubrak.Map(detailValues, func(v toolkit.M, i int) string {
				stringVal := toolkit.ToString(v[key])
				trimmedStringVal := strings.TrimSpace(stringVal)

				if trimmedStringVal == "" {
					stringVal = "NA"
				} else {
					switch key {
					case "MANDATORY":
						stringVal = doInterrupt(stringVal, []string{"0", "1"}, []string{"No", "Yes"})
						break
					case "CDE":
						stringVal = doInterrupt(stringVal, []string{"0", "1"}, []string{"No", "Yes"})
						break
					case "GOLDEN_SOURCE":
						stringVal = doInterrupt(stringVal, []string{"NO", "YES"}, []string{"No", "Yes"})
						break
					case "GS_ITAM_ID":
						if stringVal == "0" {
							stringVal = "NA"
						}
						break
					}
				}

				return stringVal
			})
			if err != nil {
				return nil, nil, nil, err
			}

			uniqueValues, err := gubrak.Uniq(mappedValues)
			if err != nil {
				return nil, nil, nil, err
			}

			switch key {
			case "SYSTEM_NAME":
				ddVal.Set("ddSystemNameSelected", uniqueValues.([]string)[0])
				break
			case "TABLE_NAME":
				ddVal.Set("ddTableNameSelected", uniqueValues.([]string)[0])
				break
			case "ALIAS_NAME":
				ddVal.Set("ddBusinessAliasNameSelected", uniqueValues.([]string)[0])
				break
			}

			joinedValues := ""
			if key == "DATASET_CUSTODIAN" || key == "BANK_ID" {
				joinedValues = strings.Join(uniqueValues.([]string), "; ")
			} else {
				joinedValues = strings.Join(uniqueValues.([]string), ", ")
			}

			selectedDetail.Set(key, joinedValues)
		}
	}

	return selectedDetail, ddVal, mappedddSource, err
}
