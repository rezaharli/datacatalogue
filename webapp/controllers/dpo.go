package controllers

import (
	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"

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
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	// detail, ddSource, err := c.Base.GetDetails(payload, s.NewDPOService().GetDetails, s.NewDPOService().GetddSource)
	// if err != nil {
	// 	h.WriteResultError(k, res, err.Error())
	// 	return
	// }

	detailsImmediatePrecedingSystem, ddSourceImmediatePrecedingSystem, err := c.Base.GetDetails(payload, s.NewDPOService().GetDetailsImmediatePrecedingSystem, s.NewDPOService().GetddSourceImmediatePrecedingSystem)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	detailsUltimateSourceSystem, ddSourceUltimateSourceSystem, err := c.Base.GetDetails(payload, s.NewDPOService().GetDetailsUltimateSourceSystem, s.NewDPOService().GetddSourceUltimateSourceSystem)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	detailsDomainView, ddSourceDomainView, err := c.Base.GetDetails(payload, s.NewDPOService().GetDetailsDomainView, s.NewDPOService().GetddSourceDomainView)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	detailsDataStandards, ddSourceDataStandards, err := c.Base.GetDetails(payload, s.NewDPOService().GetDetailsDataStandards, s.NewDPOService().GetddSourceDataStandards)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	data := toolkit.M{}
	data.Set("DetailsImmediatePrecedingSystem", detailsImmediatePrecedingSystem)
	data.Set("DDSourceImmediatePrecedingSystem", ddSourceImmediatePrecedingSystem)
	data.Set("DetailsUltimateSourceSystem", detailsUltimateSourceSystem)
	data.Set("DDSourceUltimateSourceSystem", ddSourceUltimateSourceSystem)
	data.Set("DetailsDomainView", detailsDomainView)
	data.Set("DDSourceDomainView", ddSourceDomainView)
	data.Set("DetailsDataStandards", detailsDataStandards)
	data.Set("DDSourceDataStandards", ddSourceDataStandards)

	h.WriteResultOK(k, res, data)
}
