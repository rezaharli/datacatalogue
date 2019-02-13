package controllers

import (
	m "eaciit/datacatalogue/webapp/models"
	"net/http"

	"git.eaciitapp.com/sebar/knot"
)

type Users struct {
}

func NewUsersController() *Users {
	return new(Users)
}

func (c *Users) Authenticate(k *knot.WebContext) {
	user := m.User{
		Userid: "sample user id",
	}

	k.WriteJSON(user, http.StatusOK)
}
