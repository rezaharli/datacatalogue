package controllers

import (
	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"

	h "eaciit/datacatalogue/webapp/helpers"
	s "eaciit/datacatalogue/webapp/services"
)

type DPO struct {
}

func NewDPOController() *DPO {
	return new(DPO)
}

func (c *DPO) GetLeftTable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}

	sortKey := payload.GetString("sort-key")
	sortOrder := payload.GetString("sort-order")
	skip := payload.GetInt("skip")
	take := payload.GetInt("take")

	systems, _, err := s.NewDPOService().GetLeftTable(sortKey, sortOrder, skip, take, toolkit.M{})
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

	// sortKey := payload.GetString("sort-key")
	// sortOrder := payload.GetString("sort-order")
	// skip := payload.GetInt("skip")
	// take := payload.GetInt("take")

	systems, _, err := s.NewDPOService().GetRightTable(payload.GetInt("ProcessID"))
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

	systems, _, err := s.NewDPOService().GetDetails(payload.GetInt("LeftParam"), payload.GetInt("RightParam"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}
