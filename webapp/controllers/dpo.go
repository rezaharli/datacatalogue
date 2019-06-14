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

type DPO struct {
	*Base
}

func NewDPOController() *DPO {
	return new(DPO)
}

func (c *DPO) GetLeftTable(k *knot.WebContext) {
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

	systems, _, err := s.NewDPOService().GetLeftTable(tabs, loggedinId, search, searchDD, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DPO) GetHomepageCounts(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewDPOService().GetHomepageCounts(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DPO) GetDataelementsTable(k *knot.WebContext) {
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

	systems, _, err := s.NewDPOService().GetDataelementsTable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DPO) GetDatalineageTable(k *knot.WebContext) {
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

	systems, _, err := s.NewDPOService().GetDatalineageTable(system, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DPO) GetRightTable(k *knot.WebContext) {
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

	systems, _, err := s.NewDPOService().GetRightTable(tabs, systemID, search, searchDD, colFilter, pageNumber, rowsPerPage, toolkit.M{})
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *DPO) GetDetails(k *knot.WebContext) {
	processTime := time.Now()
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	selectedDetailImmediatePrecedingSystem, ddValImmediatePrecedingSystem, mappedddSourceImmediatePrecedingSystem, err := c.GetDetailsImmediatePrecedingSystem(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	selectedDetailUltimateSourceSystem, ddValUltimateSourceSystem, mappedddSourceUltimateSourceSystem, err := c.GetDetailsUltimateSourceSystem(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	selectedDetailDomainView, ddValDomainView, mappedddSourceDomainView, err := c.GetDetailsDomainView(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	selectedDetailDataStandards, ddValDataStandards, mappedddSourceDataStandards, err := c.GetDetailsDataStandards(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data := toolkit.M{}
	data.Set("SelectedDetailImmediatePrecedingSystem", selectedDetailImmediatePrecedingSystem)
	data.Set("DDSourceImmediatePrecedingSystem", mappedddSourceImmediatePrecedingSystem)
	data.Set("DDValImmediatePrecedingSystem", ddValImmediatePrecedingSystem)

	data.Set("SelectedDetailUltimateSourceSystem", selectedDetailUltimateSourceSystem)
	data.Set("DDSourceUltimateSourceSystem", mappedddSourceUltimateSourceSystem)
	data.Set("DDValUltimateSourceSystem", ddValUltimateSourceSystem)

	data.Set("SelectedDetailDomainView", selectedDetailDomainView)
	data.Set("DDSourceDomainView", mappedddSourceDomainView)
	data.Set("DDValDomainView", ddValDomainView)

	data.Set("SelectedDetailDataStandards", selectedDetailDataStandards)
	data.Set("DDSourceDataStandards", mappedddSourceDataStandards)
	data.Set("DDValDataStandards", ddValDataStandards)

	h.WriteResultOK(k, res, data)
	toolkit.Println("processTime:", time.Since(processTime).Seconds(), "\n------------------------------------------------------------------------------------")
}

func (c *DPO) GetDetailsImmediatePrecedingSystem(payload toolkit.M) (interface{}, interface{}, interface{}, error) {
	doInterrupt := func(v string, conds, expectedreses []string) string {
		if v == conds[1] {
			return expectedreses[1]
		} else if v == conds[0] {
			return expectedreses[0]
		}

		return v
	}

	mappedDetails, mappedddSource, err := c.GetDetailAndDropdown(payload, s.NewDPOService().GetDetailsImmediatePrecedingSystem, s.NewDPOService().GetddSourceImmediatePrecedingSystem)
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
					case "DERIVED":
						stringVal = doInterrupt(stringVal, []string{"0", "1"}, []string{"No", "Yes"})
						break
					case "THRESHOLD":
						if stringVal == "0" {
							stringVal = "NA"
						}
						break
					case "DATA_SLA":
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
				ddVal.Set("ddImmSystemNameSelected", uniqueValues.([]string)[0])
				break
			case "ITAM_ID":
				ddVal.Set("ddImmItamIDSelected", uniqueValues.([]string)[0])
				break
			case "TABLE_NAME":
				ddVal.Set("ddImmTableNameSelected", uniqueValues.([]string)[0])
				break
			case "COLUMN_NAME":
				ddVal.Set("ddImmColumnNameSelected", uniqueValues.([]string)[0])
				break
			case "DATA_ELEMENT":
				ddVal.Set("ddImmScreenLabelSelected", uniqueValues.([]string)[0])
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
	} else {
		selectedDetail = nil

		ddVal.Set("ddImmSystemNameSelected", "NA")
		ddVal.Set("ddImmItamIDSelected", "NA")
		ddVal.Set("ddImmTableNameSelected", "NA")
		ddVal.Set("ddImmColumnNameSelected", "NA")
		ddVal.Set("ddImmScreenLabelSelected", "NA")
	}

	return selectedDetail, ddVal, mappedddSource, err
}

func (c *DPO) GetDetailsUltimateSourceSystem(payload toolkit.M) (interface{}, interface{}, interface{}, error) {
	doInterrupt := func(v string, conds, expectedreses []string) string {
		if v == conds[1] {
			return expectedreses[1]
		} else if v == conds[0] {
			return expectedreses[0]
		}

		return v
	}

	mappedDetails, mappedddSource, err := c.GetDetailAndDropdown(payload, s.NewDPOService().GetDetailsUltimateSourceSystem, s.NewDPOService().GetddSourceUltimateSourceSystem)
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
					case "DERIVED":
						stringVal = doInterrupt(stringVal, []string{"0", "1"}, []string{"No", "Yes"})
						break
					case "THRESHOLD":
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
				ddVal.Set("ddUltSystemNameSelected", uniqueValues.([]string)[0])
				break
			case "ITAM_ID":
				ddVal.Set("ddUltItamIDSelected", uniqueValues.([]string)[0])
				break
			case "TABLE_NAME":
				ddVal.Set("ddUltTableNameSelected", uniqueValues.([]string)[0])
				break
			case "COLUMN_NAME":
				ddVal.Set("ddUltColumnNameSelected", uniqueValues.([]string)[0])
				break
			case "DATA_ELEMENT":
				ddVal.Set("ddUltScreenLabelSelected", uniqueValues.([]string)[0])
				break
			case "GOLDEN_SOURCE":
				ddVal.Set("ddUltGoldenSourceSelected", uniqueValues.([]string)[0])
				break
			case "GS_SYSTEM_NAME":
				ddVal.Set("ddUltGsSystemNameSelected", uniqueValues.([]string)[0])
				break
			case "GF_ITAM_ID":
				ddVal.Set("ddUltGsItamIdSelected", uniqueValues.([]string)[0])
				break
			case "GS_TABLE_NAME":
				ddVal.Set("ddUltGsTableNameSelected", uniqueValues.([]string)[0])
				break
			case "GS_COLUMN_NAME":
				ddVal.Set("ddUltGsColumnNameSelected", uniqueValues.([]string)[0])
				break
			case "GS_DATA_ELEMENT":
				ddVal.Set("ddUltGsDataElementSelected", uniqueValues.([]string)[0])
				break
			case "GS_DESCRIPTION":
				ddVal.Set("ddUltGsDescriptionSelected", uniqueValues.([]string)[0])
				break
			case "GS_DERIVED":
				ddVal.Set("ddUltGsDerivedSelected", uniqueValues.([]string)[0])
				break
			case "GS_DERIVATION_LOGIC":
				ddVal.Set("ddUltGsDerivationLogicSelected", uniqueValues.([]string)[0])
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
	} else {
		selectedDetail = nil

		ddVal.Set("ddUltSystemNameSelected", "NA")
		ddVal.Set("ddUltItamIDSelected", "NA")
		ddVal.Set("ddUltTableNameSelected", "NA")
		ddVal.Set("ddUltColumnNameSelected", "NA")
		ddVal.Set("ddUltScreenLabelSelected", "NA")
		ddVal.Set("ddUltGoldenSourceSelected", "NA")
		ddVal.Set("ddUltGsSystemNameSelected", "NA")
		ddVal.Set("ddUltGsItamIdSelected", "NA")
		ddVal.Set("ddUltGsTableNameSelected", "NA")
		ddVal.Set("ddUltGsColumnNameSelected", "NA")
		ddVal.Set("ddUltGsDataElementSelected", "NA")
		ddVal.Set("ddUltGsDescriptionSelected", "NA")
		ddVal.Set("ddUltGsDerivedSelected", "NA")
		ddVal.Set("ddUltGsDerivationLogicSelected", "NA")
	}

	return selectedDetail, ddVal, mappedddSource, err
}

func (c *DPO) GetDetailsDomainView(payload toolkit.M) (interface{}, interface{}, interface{}, error) {
	mappedDetails, mappedddSource, err := c.GetDetailAndDropdown(payload, s.NewDPOService().GetDetailsDomainView, s.NewDPOService().GetddDetailsDomainView)
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

			switch key {
			case "DOMAIN_NAME":
				ddVal.Set("ddDomainNameSelected", uniqueValues.([]string)[0])
				break
			case "SUBDOMAIN_NAME":
				ddVal.Set("ddSubdomainNameSelected", uniqueValues.([]string)[0])
				break
			case "SUBDOMAIN_OWNER":
				ddVal.Set("ddSubdomainOwnerSelected", uniqueValues.([]string)[0])
				break
			case "BUSINESS_TERM":
				ddVal.Set("ddBusinessTermSelected", uniqueValues.([]string)[0])
				break
			case "BUSINESS_TERM_DESCRIPTION":
				ddVal.Set("ddBusinessTermDescriptionSelected", uniqueValues.([]string)[0])
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
	} else {
		selectedDetail = nil

		ddVal.Set("ddDomainNameSelected", "NA")
		ddVal.Set("ddSubdomainNameSelected", "NA")
		ddVal.Set("ddSubdomainOwnerSelected", "NA")
		ddVal.Set("ddBusinessTermSelected", "NA")
		ddVal.Set("ddBusinessTermDescriptionSelected", "NA")
	}

	return selectedDetail, ddVal, mappedddSource, err
}

func (c *DPO) GetDetailsDataStandards(payload toolkit.M) (interface{}, interface{}, interface{}, error) {
	mappedDetails, mappedddSource, err := c.GetDetailAndDropdown(payload, s.NewDPOService().GetDetailsDataStandards, nil)
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
	} else {
		selectedDetail = nil
	}

	return selectedDetail, ddVal, mappedddSource, err
}
