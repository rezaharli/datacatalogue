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

	sortKey := payload.GetString("sort-key")
	sortOrder := payload.GetString("sort-order")
	skip := payload.GetInt("skip")
	take := payload.GetInt("take")

	systems, _, err := s.NewDSCService().GetAllSystem(sortKey, sortOrder, skip, take, toolkit.M{})
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

	// sortKey := payload.GetString("sort-key")
	// sortOrder := payload.GetString("sort-order")
	// skip := payload.GetInt("skip")
	// take := payload.GetInt("take")

	systems, _, err := s.NewDSCService().GetTableName(payload.GetInt("SystemID"))
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

	// sortKey := payload.GetString("sort-key")
	// sortOrder := payload.GetString("sort-order")
	// skip := payload.GetInt("skip")
	// take := payload.GetInt("take")

	systems, _, err := s.NewDSCService().GetInterfacesRightTable(payload.GetInt("SystemID"))
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

	systems, _, err := s.NewDSCService().GetDetails(payload.GetInt("LeftParam"), payload.GetInt("RightParam"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}
