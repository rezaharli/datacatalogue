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

	detail, ddSource, err := c.Base.GetDetails(payload, s.NewDDOService().GetDetails, s.NewDDOService().GetddSource)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data := toolkit.M{}
	data.Set("Detail", detail)
	data.Set("DDSource", ddSource)

	h.WriteResultOK(k, res, data)
}

func (c *DDO) GetDetailsBusinessMetadataFromDomain(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data, _, err := s.NewDDOService().GetDetailsBusinessMetadataFromDomain(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, data)
}

func (c *DDO) GetddSourceBusinessMetadataFromDomain(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data, _, err := s.NewDDOService().GetddSourceBusinessMetadataFromDomain(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, data)
}

func (c *DDO) GetDetailsDownstreamUsageOfBusinessTerm(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data, _, err := s.NewDDOService().GetDetailsDownstreamUsageOfBusinessTerm(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, data)
}

func (c *DDO) GetddSourceDownstreamUsageOfBusinessTerm(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data, _, err := s.NewDDOService().GetddSourceDownstreamUsageOfBusinessTerm(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, data)
}

func (c *DDO) GetDetailsBTResiding(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data, _, err := s.NewDDOService().GetDetailsBTResiding(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, data)
}

func (c *DDO) GetddSourceBTResiding(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data, _, err := s.NewDDOService().GetddSourceBTResiding(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, data)
}
