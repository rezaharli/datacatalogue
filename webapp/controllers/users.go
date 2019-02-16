package controllers

import (
	"time"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
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
		h.WriteResultError(k, res, err.Error())
		return
	}

	ok, user, err := s.NewUserService().Authenticate(payload.GetInt("username"), payload.GetString("password"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	if !ok {
		h.WriteResultErrorOK(k, res, "Invalid username or password")
		return
	}

	h.WriteResultOK(k, res, user)
}

func (c *Users) GetAll(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}

	sortKey := payload.GetString("sort-key")
	sortOrder := payload.GetString("sort-order")
	skip := payload.GetInt("skip")
	take := payload.GetInt("take")

	users, _, err := s.NewUserService().GetAll(sortKey, sortOrder, skip, take, toolkit.M{})
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, users)
}

func (c *Users) Register(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	user := m.NewSysUserModel()
	user.Username = payload.GetInt("Username")
	user.Password = payload.GetString("Password")
	user.Name = payload.GetString("Name")

	ok, err := s.NewUserService().Insert(user)
	if !ok && err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}
	if ok && err != nil {
		h.WriteResultErrorOK(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, true)
}

func (c *Users) Update(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := m.NewSysUserModel()
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	payload.UpdatedAt = time.Now()
	toolkit.Println(payload)
	ok, err := s.NewUserService().Update(payload)
	if !ok && err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}
	if ok && err != nil {
		h.WriteResultErrorOK(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, true)
}

func (c *Users) Delete(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	err = s.NewUserService().DeleteByUsername(payload.GetInt("Username"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, true)
}
