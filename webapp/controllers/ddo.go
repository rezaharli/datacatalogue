package controllers

import (
	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"

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

	pageNumber := pagination.GetInt("page")
	rowsPerPage := pagination.GetInt("rowsPerPage")

	systems, _, err := s.NewDDOService().GetLeftTable(tabs, loggedinId, search, searchDD, colFilter, pageNumber, rowsPerPage, toolkit.M{})
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
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	detailsBusinessMetadata, ddSourceBusinessMetadata, err := c.Base.GetDetails(payload, s.NewDDOService().GetDetailsBusinessMetadataFromDomain, s.NewDDOService().GetddSourceBusinessMetadataFromDomain)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	detailsDownstreamUsage, ddSourceDownstreamUsage, err := c.Base.GetDetails(payload, s.NewDDOService().GetDetailsDownstreamUsageOfBusinessTerm, s.NewDDOService().GetddSourceDownstreamUsageOfBusinessTerm)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	detailsBTResiding, ddSourceBTResiding, err := c.Base.GetDetails(payload, s.NewDDOService().GetDetailsBTResiding, s.NewDDOService().GetddSourceBTResiding)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data := toolkit.M{}
	data.Set("DetailsBusinessMetadata", detailsBusinessMetadata)
	data.Set("DDSourceBusinessMetadata", ddSourceBusinessMetadata)
	data.Set("DetailsDownstreamUsage", detailsDownstreamUsage)
	data.Set("DDSourceDownstreamUsage", ddSourceDownstreamUsage)
	data.Set("DetailsBTResiding", detailsBTResiding)
	data.Set("DDSourceBTResiding", ddSourceBTResiding)

	h.WriteResultOK(k, res, data)
}
