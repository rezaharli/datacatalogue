package controllers

import (
	"net/http"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"

	s "eaciit/datacatalogue/webapp/services"
)

type Users struct {
}

func NewUsersController() *Users {
	return new(Users)
}

func (c *Users) Authenticate(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		res.SetErrorTxt(err.Error())
		k.WriteJSON(res, http.StatusUnauthorized)
		return
	}

	ok, user, err := s.NewUserService().Authenticate(payload.GetInt("username"), payload.GetString("password"))
	if err != nil {
		res.SetErrorTxt(err.Error())
		k.WriteJSON(res, http.StatusUnauthorized)
		return
	}

	if !ok {
		res.SetErrorTxt("Invalid username or password")
		k.WriteJSON(res, http.StatusOK)
		return
	}

	res.SetData(user)
	k.WriteJSON(res, http.StatusOK)
}
