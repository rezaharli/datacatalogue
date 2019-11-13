package controllers

import (
	"fmt"
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

func (c *DSC) GetIARCTable(k *knot.WebContext) {
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

	tableRows, _, err := s.NewDSCService().GetIARCTable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	ret := toolkit.M{}
	ret.Set("Flat", tableRows)
	ret.Set("Grouped", tableRows)

	h.WriteResultOK(k, res, tableRows)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
}

func (c *DSC) GetIARCPersonalDataTable(k *knot.WebContext) {
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

	tableRows, _, err := s.NewDSCService().GetIARCPersonalDataTable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, tableRows)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
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

	h.WriteResultOK(k, res, tableRows)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
}

func (c *DSC) GetEdmpDDdropdowns(k *knot.WebContext) {
	queryTime := time.Now()
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	mappedddSource, err := c.GetDropdownSource(payload, s.NewDSCService().GetEdmpDDDropdowns)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	ret := toolkit.M{}
	ret.Set("MappedDDSource", mappedddSource)

	h.WriteResultOK(k, res, ret)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
}

func (c *DSC) GetEdmpIarcdropdowns(k *knot.WebContext) {
	queryTime := time.Now()
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	mappedddSource, err := c.GetDropdownSource(payload, s.NewDSCService().GetEdmpIarcDropdowns)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	ret := toolkit.M{}
	ret.Set("MappedDDSource", mappedddSource)

	h.WriteResultOK(k, res, ret)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
}

func (c *DSC) GetEdmpDDTechnicalTable(k *knot.WebContext) {
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

	defaultSortInt := payload.Get("DefaultSort").([]interface{})
	defaultSort := make([]string, len(defaultSortInt))
	for i, v := range defaultSortInt {
		defaultSort[i] = fmt.Sprint(v)
	}

	tableRows, _, err := s.NewDSCService().GetEdmpDDTechnicalTable(system, colFilter, pagination, defaultSort)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, tableRows)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
}

func (c *DSC) GetEdmpDDBusinessTable(k *knot.WebContext) {
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

	defaultSortInt := payload.Get("DefaultSort").([]interface{})
	defaultSort := make([]string, len(defaultSortInt))
	for i, v := range defaultSortInt {
		defaultSort[i] = fmt.Sprint(v)
	}

	tableRows, _, err := s.NewDSCService().GetEdmpDDBusinessTable(system, colFilter, pagination, defaultSort)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, tableRows)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
}

func (c *DSC) GetEdmpDDConsumptionTable(k *knot.WebContext) {
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

	defaultSortInt := payload.Get("DefaultSort").([]interface{})
	defaultSort := make([]string, len(defaultSortInt))
	for i, v := range defaultSortInt {
		defaultSort[i] = fmt.Sprint(v)
	}

	tableRows, _, err := s.NewDSCService().GetEdmpDDConsumptionTable(system, colFilter, pagination, defaultSort)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, tableRows)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
}

func (c *DSC) GetEdmpIarcPersonalTable(k *knot.WebContext) {
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

	defaultSortInt := payload.Get("DefaultSort").([]interface{})
	defaultSort := make([]string, len(defaultSortInt))
	for i, v := range defaultSortInt {
		defaultSort[i] = fmt.Sprint(v)
	}

	tableRows, _, err := s.NewDSCService().GetEdmpIarcPersonalTable(system, colFilter, pagination, defaultSort)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, tableRows)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
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

	selectedDetailLeftPanel, ddValLeftPanel, mappedddSourceLeftPanel, err := c.GetDetailsLeftPanel(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	selectedDetail, ddVal, mappedddSource, err := c.GetDetailsRightPanel(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data := toolkit.M{}
	data.Set("SelectedDetailLeftPanel", selectedDetailLeftPanel)
	data.Set("DDSourceLeftPanel", mappedddSourceLeftPanel)
	data.Set("DDValLeftPanel", ddValLeftPanel)

	data.Set("SelectedDetail", selectedDetail)
	data.Set("DDSource", mappedddSource)
	data.Set("DDVal", ddVal)

	h.WriteResultOK(k, res, data)
	toolkit.Println("processTime:", time.Since(processTime).Seconds(), "\n------------------------------------------------------------------------------------")
}

func (c *DSC) GetDetailsLeftPanel(payload toolkit.M) (interface{}, interface{}, interface{}, error) {
	mappedDetails, mappedddSource, err := c.GetDetailAndDropdown(payload, s.NewDSCService().GetDetailsLeftPanel, nil)
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

func (c *DSC) GetDetailsRightPanel(payload toolkit.M) (interface{}, interface{}, interface{}, error) {
	doInterrupt := func(v string, conds, expectedreses []string) string {
		if v == conds[1] {
			return expectedreses[1]
		} else if v == conds[0] {
			return expectedreses[0]
		}

		return v
	}

	mappedDetails, mappedddSource, err := c.GetDetailAndDropdown(payload, s.NewDSCService().GetDetailsRightPanel, s.NewDSCService().GetddSourceRightPanel)
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
			case "TABLE_NAME":
				ddVal.Set("ddTableSelected", uniqueValues.([]string)[0])
				break
			case "COLUMN_NAME":
				ddVal.Set("ddColumnSelected", uniqueValues.([]string)[0])
				break
			case "BUSINESS_ALIAS_NAME":
				ddVal.Set("ddScreenLabelSelected", uniqueValues.([]string)[0])
				break
			case "BUSINESS_TERM":
				ddVal.Set("ddBusinessTermSelected", uniqueValues.([]string)[0])
				break
			case "IMM_PRECEEDING_SYSTEM":
				ddVal.Set("ddPrecSelected", uniqueValues.([]string)[0])
				break
			case "IMM_PREC_INCOMING":
				ddVal.Set("ddPrecIncomingSelected", uniqueValues.([]string)[0])
				break
			case "IMM_SUCCEEDING_SYSTEM":
				ddVal.Set("ddSuccSelected", uniqueValues.([]string)[0])
				break
			case "IMM_SUCC_INCOMING":
				ddVal.Set("ddSuccIncomingSelected", uniqueValues.([]string)[0])
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
