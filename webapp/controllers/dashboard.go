package controllers

import (
	h "eaciit/datacatalogue/webapp/helpers"

	"git.eaciitapp.com/sebar/knot"
	"github.com/eaciit/toolkit"
)

type Dashboard struct {
	*Base
}

func NewDashboardController() *Dashboard {
	return new(Dashboard)
}

func (c *Dashboard) Encrypt(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	plainText := payload.GetString("PlainText")
	encrypted, err := h.Encrypt(plainText)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, encrypted)
}

func (c *Dashboard) Decrypt(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	encrypted := payload.GetString("Encrypted")
	plainText, err := h.Decrypt(encrypted)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, plainText)
}
