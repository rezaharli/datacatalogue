package controllers

import (
	"strings"
	"time"

	"github.com/novalagung/gubrak"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"

	"eaciit/datacatalogue/webapp/helpers"
	h "eaciit/datacatalogue/webapp/helpers"
	s "eaciit/datacatalogue/webapp/services"
)

type DSC struct {
	*Base
}

func NewDSCController() *DSC {
	return new(DSC)
}

func (c *DSC) GetAllSystems(k *knot.WebContext) {
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

	systems, _, err := s.NewDSCService().GetAllSystem(tabs, loggedinId, search, searchDD, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DSC) GetHomepageCounts(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDSCService().GetHomepageCounts(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DSC) GetCDETable(k *knot.WebContext) {
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

	systems, _, err := s.NewDSCService().GetCDETable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DSC) GetCDPTable(k *knot.WebContext) {
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

	systems, _, err := s.NewDSCService().GetCDPTable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DSC) GetCDPCDETable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	system := payload.GetString("System")
	dspName := payload.GetString("DspName")
	colFilter := payload.Get("Filters")
	pagination, err := toolkit.ToM(payload.Get("Pagination"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDSCService().GetCDPCDETable(system, dspName, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DSC) GetInterfacesTable(k *knot.WebContext) {
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

	systems, _, err := s.NewDSCService().GetInterfacesTable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DSC) GetInterfacesCDETable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	system := payload.GetString("System")
	dspName := payload.GetString("DspName")
	colFilter := payload.Get("Filters")
	pagination, err := toolkit.ToM(payload.Get("Pagination"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDSCService().GetInterfacesCDETable(system, dspName, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DSC) GetDDTable(k *knot.WebContext) {
	queryTime := time.Now()
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

	tableRows, _, err := s.NewDSCService().GetDDTable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	ret := toolkit.M{}
	ret.Set("Flat", tableRows)
	ret.Set("Grouped", tableRows)

	h.WriteResultOK(k, res, ret)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
}

func (c *DSC) GetTableName(k *knot.WebContext) {
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

	systems, _, err := s.NewDSCService().GetTableName(tabs, systemID, search, searchDD, colFilter, pageNumber, rowsPerPage, toolkit.M{})
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DSC) GetInterfacesRightTable(k *knot.WebContext) {
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

	systems, _, err := s.NewDSCService().GetInterfacesRightTable(tabs, systemID, search, searchDD, colFilter, pageNumber, rowsPerPage, toolkit.M{})
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DSC) GetDetails(k *knot.WebContext) {
	processTime := time.Now()
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	detail, _, err := s.NewDSCService().GetDetails(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	groupedDetail, err := gubrak.GroupBy(detail, func(each toolkit.M) string {
		return each.GetString("ID")
	})
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	mappedDetails, err := gubrak.Map(helpers.ObjectKeys(groupedDetail), func(key string, i int) toolkit.M {
		tmp := groupedDetail.(map[string]([]toolkit.M))

		ret := toolkit.M{}
		ret["ID"] = key
		ret["Values"] = tmp[key]

		return ret
	})
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	selectedDetail := toolkit.M{}
	detailSources := mappedDetails.([]toolkit.M)
	if len(detailSources) > 0 {
		detailSource := detailSources[0]
		detailValues := detailSource.Get("Values").([]toolkit.M)

		doInterrupt := func(v string, conds, expectedreses []string) string {
			if v == conds[1] {
				return expectedreses[1]
			} else if v == conds[0] {
				return expectedreses[0]
			}

			return v
		}

		for _, key := range helpers.ObjectKeys(detailValues[0]) {
			mappedValues, err := gubrak.Map(detailValues, func(v toolkit.M, i int) string {
				stringVal := strings.TrimSpace(toolkit.ToString(v[key]))

				switch key {
				case "CDE_YES_NO":
					stringVal = doInterrupt(stringVal, []string{"0", "1"}, []string{"No", "Yes"})
					break
				case "STATUS":
					stringVal = doInterrupt(stringVal, []string{"0", "1"}, []string{"INACTIVE", "ACTIVE"})
					break
				case "DERIVED_YES_NO":
					stringVal = doInterrupt(stringVal, []string{"0", "1"}, []string{"No", "Yes"})
					break
				case "SOURCED_FROM_UPSTREAM_YES_NO":
					stringVal = doInterrupt(stringVal, []string{"0", "1"}, []string{"No", "Yes"})
					break
				case "IMM_PREC_DERIVED":
					stringVal = doInterrupt(stringVal, []string{"0", "1"}, []string{"No", "Yes"})
					break
				case "IMM_SUCC_DERIVED":
					stringVal = doInterrupt(stringVal, []string{"0", "1"}, []string{"No", "Yes"})
					break
				}

				if stringVal == "" {
					stringVal = "NA"
				}

				return stringVal
			})
			if err != nil {
				h.WriteResultError(k, res, err.Error())
				return
			}

			uniqueValues, err := gubrak.Uniq(mappedValues)
			if err != nil {
				h.WriteResultError(k, res, err.Error())
				return
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

	ddSource, _, err := s.NewDSCService().GetddSource(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	mappedddSource, err := gubrak.Map(ddSource, func(v toolkit.M, i int) toolkit.M {
		for _, key := range helpers.ObjectKeys(v) {
			v[key] = strings.TrimSpace(toolkit.ToString(v[key]))
			if v[key] == "" {
				v[key] = "NA"
			}
		}

		return v
	})

	data := toolkit.M{}
	data.Set("SelectedDetail", selectedDetail)
	data.Set("DDSource", mappedddSource)

	ddVal := toolkit.M{}
	ddVal.Set("ddTableSelected", strings.Split(selectedDetail["TABLE_NAME"].(string), ", ")[0])
	ddVal.Set("ddColumnSelected", strings.Split(selectedDetail["COLUMN_NAME"].(string), ", ")[0])
	ddVal.Set("ddScreenLabelSelected", strings.Split(selectedDetail["BUSINESS_ALIAS_NAME"].(string), ", ")[0])
	ddVal.Set("ddBusinessTermSelected", strings.Split(selectedDetail["BUSINESS_TERM"].(string), ", ")[0])
	ddVal.Set("ddPrecSelected", strings.Split(selectedDetail["IMM_PRECEEDING_SYSTEM"].(string), ", ")[0])
	ddVal.Set("ddPrecIncomingSelected", strings.Split(selectedDetail["IMM_PREC_INCOMING"].(string), ", ")[0])
	ddVal.Set("ddSuccSelected", strings.Split(selectedDetail["IMM_SUCCEEDING_SYSTEM"].(string), ", ")[0])
	ddVal.Set("ddSuccIncomingSelected", strings.Split(selectedDetail["IMM_SUCC_INCOMING"].(string), ", ")[0])
	data.Set("DDVal", ddVal)

	h.WriteResultOK(k, res, data)
	toolkit.Println("processTime:", time.Since(processTime).Seconds(), "\n------------------------------------------------------------------------------------")
}
