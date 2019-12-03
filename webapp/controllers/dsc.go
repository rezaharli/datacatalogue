package controllers

import (
	"fmt"
	"time"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"

	h "eaciit/datacatalogue/webapp/helpers"
	s "eaciit/datacatalogue/webapp/services"
)

type DSC struct {
	*Base
}

func NewDSCController() *DSC {
	return new(DSC)
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
	globalFilter := payload.Get("GlobalFilters")
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

	tableRows, _, err := s.NewDSCService().GetEdmpDDTechnicalTable(system, globalFilter, colFilter, pagination, defaultSort)
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
	globalFilter := payload.Get("GlobalFilters")
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

	tableRows, _, err := s.NewDSCService().GetEdmpDDBusinessTable(system, globalFilter, colFilter, pagination, defaultSort)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, tableRows)
	toolkit.Println("Process Time:", time.Since(queryTime).Seconds(), "\n------------------------------------------------------------------------")
}
