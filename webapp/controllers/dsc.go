package controllers

import (
	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"

	h "eaciit/datacatalogue/webapp/helpers"
	s "eaciit/datacatalogue/webapp/services"
)

type DSC struct {
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

	pageNumber := pagination.GetInt("page")
	rowsPerPage := pagination.GetInt("rowsPerPage")

	systems, _, err := s.NewDSCService().GetAllSystem(tabs, loggedinId, search, searchDD, colFilter, pageNumber, rowsPerPage, toolkit.M{})
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
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

	ddSource, _, err := s.NewDSCService().GetddSource(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data := toolkit.M{}
	data.Set("Detail", detail)
	data.Set("DDSource", ddSource)

	h.WriteResultOK(k, res, data)
}
