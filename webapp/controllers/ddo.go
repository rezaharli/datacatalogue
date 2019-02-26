package controllers

import (
	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"

	h "eaciit/datacatalogue/webapp/helpers"
	s "eaciit/datacatalogue/webapp/services"
)

type DDO struct {
}

func NewDDOController() *DDO {
	return new(DDO)
}

func (c *DDO) GetLeftTable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}

	sortKey := payload.GetString("sort-key")
	sortOrder := payload.GetString("sort-order")
	skip := payload.GetInt("skip")
	take := payload.GetInt("take")

	systems, _, err := s.NewDDOService().GetLeftTable(sortKey, sortOrder, skip, take, toolkit.M{})
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

	// sortKey := payload.GetString("sort-key")
	// sortOrder := payload.GetString("sort-order")
	// skip := payload.GetInt("skip")
	// take := payload.GetInt("take")

	systems, _, err := s.NewDDOService().GetRightTable(payload.GetInt("SystemID"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}
