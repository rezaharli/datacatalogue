package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-ldap/ldap"

	"github.com/eaciit/clit"
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

func (c *Users) CheckSession(k *knot.WebContext) {
	res := toolkit.NewResult()

	isSession := k.GetSession("IsSession", false)

	h.WriteResultOK(k, res, isSession)
}

func (c *Users) Authenticate(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	go func() {
		if clit.Config("default", "LDAP", "") != "" {
			ldapConf := clit.Config("default", "LDAP", "").(map[string]interface{})
			if ldapConf["IsEnabled"].(string) == "true" {
				isSuccess, _, err := h.TryToLoginUsingLDAP(payload.GetString("username"), payload.GetString("password"))

				toolkit.Println("Success ?? >> ", isSuccess)
				toolkit.Println("Err ?? >> ", err)

				if !(isSuccess && err == nil) {
					h.WriteResultErrorOK(k, res, "LDAP login fail.")
					return
				}

				toolkit.Println("------------------------------------------------------------------------------------------")

				// var ldapServer = "ldap.itd.umich.edu"
				// var ldapPort = uint16(389)
				// var ldapTLSPort = uint16(636)
				// var baseDN = "dc=umich,dc=edu"
				// var filter = []string{
				// 	"(cn=cis-fac)",
				// 	"(&(owner=*)(cn=cis-fac))",
				// 	"(&(objectclass=rfc822mailgroup)(cn=*Computer*))",
				// 	"(&(objectclass=rfc822mailgroup)(cn=*Mathematics*))"}
				var attributes = []string{}

				fmt.Printf("TestSearch: starting...\n")
				ldapConf := clit.Config("default", "LDAP", "").(map[string]interface{})
				l, err := ldap.Dial("tcp", strings.TrimSpace(ldapConf["Host"].(string)))
				if err != nil {
					toolkit.Println(err.Error())
					return
				}
				defer l.Close()

				searchRequest := ldap.NewSearchRequest(
					strings.TrimSpace(ldapConf["BaseDN"].(string)),
					ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
					"("+strings.TrimSpace(ldapConf["UserAuthAttr"].(string))+"="+payload.GetString("username")+")",
					attributes,
					nil)

				sr, err := l.Search(searchRequest)
				if err != nil {
					toolkit.Println(err.Error())
					return
				}

				fmt.Printf("TestSearch: %s -> num of entries = %d\n", searchRequest.Filter, len(sr.Entries))
				fmt.Println(sr.Entries)

				for _, v := range sr.Entries {
					for _, str := range attributes {

						toolkit.Println("v ---", v)
						toolkit.Println("str ---", str)

						val := ""
						if len(v.GetAttributeValues(str)) > 1 {
							val = strings.Join(v.GetAttributeValues(str), "|")
						} else {
							val = v.GetAttributeValue(str)
						}

						toolkit.Println("val ---", val)
					}
				}
			}
		}
	}()

	ok, user, err := s.NewUserService().Authenticate(payload.GetInt("username"), payload.GetString("password"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	if !ok {
		h.WriteResultErrorOK(k, res, "Invalid username or password")
		return
	}

	k.SetSession("IsSession", true)
	h.WriteResultOK(k, res, user)
}

func (c *Users) GetAll(k *knot.WebContext) {
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

	systems, _, err := s.NewUserService().GetAll(tabs, loggedinId, search, searchDD, colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}

func (c *Users) Register(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := m.NewSysUserModel()
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	payload.ID = payload.Username
	ok, err := s.NewUserService().Insert(payload)
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

	payload.UpdatedAt = time.Now().String()
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

func (c *Users) SaveUsage(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	err = s.NewUserService().SaveUsage(payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, true)
}

func (c *Users) GetUsageTable(k *knot.WebContext) {
	res := toolkit.NewResult()

	payload := toolkit.M{}
	err := k.GetPayload(&payload)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	colFilter := payload.Get("Filters")
	pagination, err := toolkit.ToM(payload.Get("Pagination"))
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	systems, _, err := s.NewUserService().GetUsageTable(colFilter, pagination)
	if err != nil {
		h.WriteResultError(k, res, err.Error())
		return
	}

	h.WriteResultOK(k, res, systems)
}
