package services

import (
	"crypto/md5"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/icrowley/fake"

	"github.com/eaciit/clit"
	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/dbflex"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type EdmpUserService struct {
	*Base
}

func NewEdmpUserService() *EdmpUserService {
	ret := new(EdmpUserService)
	return ret
}

func (s *EdmpUserService) HashPassword(password string) string {
	hashedPassword := md5.New()
	io.WriteString(hashedPassword, fmt.Sprintf("%s%s", password, PASSWORD_SALT))

	return fmt.Sprintf("%x", hashedPassword.Sum(nil))
}

func (s *EdmpUserService) Authenticate(username int, password string) (bool, *m.EdmpUser, error) {
	users, err := s.authenticate(username, password)
	if err != nil {
		return false, m.NewEdmpUserModel(), err
	}

	if len(users) == 0 {
		//if not found create new user
		newUser := m.NewEdmpUserModel()
		newUser.ID = username
		newUser.USERNAME = username
		newUser.PASSWORD = password
		newUser.EMAIL = toolkit.ToString(username)
		newUser.NAME = toolkit.ToString(username)
		newUser.STATUS = 1
		newUser.ROLE = "DSC,DDO,DPO,RFO"
		newUser.CREATEDAT = time.Now().String()
		newUser.UPDATEDAT = time.Now().String()

		ldapConf := clit.Config("default", "LDAP", "").(map[string]interface{})
		if ldapConf["IsEnabled"].(string) == "auto" {
			newUser.PASSWORD = toolkit.ToString(username)
		} else {
			newUser.PASSWORD = password
		}

		ok, err := NewEdmpUserService().Insert(newUser)
		if !ok && err != nil {
			return false, m.NewEdmpUserModel(), err
		}

		users, err = s.authenticate(username, password)
		if err != nil {
			return false, m.NewEdmpUserModel(), err
		}

		if len(users) == 0 {
			return false, m.NewEdmpUserModel(), nil
		}
	}

	return true, &(users[0]), nil
}

func (s *EdmpUserService) authenticate(username int, password string) ([]m.EdmpUser, error) {
	filter := []*dbflex.Filter{}
	filter = append(filter, dbflex.Eq("status", 1))
	filter = append(filter, dbflex.Eq("username", username))

	if clit.Config("default", "LDAP", "") != "" {
		ldapConf := clit.Config("default", "LDAP", "").(map[string]interface{})
		if ldapConf["IsEnabled"].(string) == "true" || ldapConf["IsEnabled"].(string) == "auto" {
			//do nothing
		} else {
			filter = append(filter, dbflex.Eq("password", s.HashPassword(password)))
		}
	} else {
		filter = append(filter, dbflex.Eq("password", s.HashPassword(password)))
	}

	users := make([]m.EdmpUser, 0)
	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: m.NewEdmpUserModel().TableName(),
		Clause:    dbflex.And(filter...),
		Result:    &users,
	})

	return users, err
}

func (s *EdmpUserService) GetAll(tabs, loggedinid, search string, searchDD, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "users.sql")
	gridArgs.QueryName = "users"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	gridArgs.Colnames = append(gridArgs.Colnames,
		"USERNAME", "EMAIL", "NAME", "ROLE", "STATUS", "CREATEDAT", "UPDATEDAT",
	)

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colFilterM.GetString(colname))
		}
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *EdmpUserService) getRegisteredPeople(username int) ([]m.People, error) {
	people := make([]m.People, 0)
	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: m.NewPeopleModel().TableName(),
		Clause:    dbflex.Eq("BANK_ID", toolkit.ToString(username)),
		Result:    &people,
	})
	if err != nil {
		return nil, err
	}

	if len(people) > 0 {
		return people, nil
	}

	return nil, err
}

func (s *EdmpUserService) insertLinkRolePeople(people m.People, role string) (err error) {
	roles := strings.Split(role, ",")
	for _, role := range roles {
		q := `SELECT max(ID)+1 as NEWID from ` + m.NewLinkRolePeopleModel().TableName()

		resultRows := make([]toolkit.M, 0)
		err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
			TableName: m.NewLinkRolePeopleModel().TableName(),
			SqlQuery:  q,
			Results:   &resultRows,
		})
		if err != nil {
			return
		}

		toolkit.Println("role--------", role)
		if role == "DSC" {
			newlrp := m.NewLinkRolePeopleModel()
			newlrp.ID = resultRows[0].GetInt("NEWID")
			newlrp.People_ID = people.ID
			newlrp.Role_ID = 1
			newlrp.Object_Type = "SYSTEM"
			newlrp.Object_ID = fake.Day() ///// WARNING !!!!!!!!!!!

			toolkit.Println(newlrp)
			err = h.NewDBcmd().Insert(h.InsertParam{
				Data:      newlrp,
				TableName: newlrp.TableName(),
			})
			if err != nil {
				return
			}
		} else if role == "DPO" {
			newlrp := m.NewLinkRolePeopleModel()
			newlrp.ID = resultRows[0].GetInt("NEWID")
			newlrp.People_ID = people.ID
			newlrp.Role_ID = 2
			newlrp.Object_Type = "PROCESSES"
			newlrp.Object_ID = fake.Day() ///// WARNING !!!!!!!!!!!

			toolkit.Println(newlrp)
			err = h.NewDBcmd().Insert(h.InsertParam{
				Data:      newlrp,
				TableName: newlrp.TableName(),
			})
			if err != nil {
				return
			}
		}
	}

	return
}

func (s *EdmpUserService) Insert(data *m.EdmpUser) (bool, error) {
	data.CREATEDAT = time.Now().String()
	data.UPDATEDAT = time.Now().String()

	users := make([]m.EdmpUser, 0)
	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: m.NewEdmpUserModel().TableName(),
		Clause:    dbflex.Eq("username", data.USERNAME),
		Result:    &users,
	})

	if err != nil {
		return false, err
	}
	if len(users) > 0 {
		return true, fmt.Errorf("The same username already exists")
	}

	// registeredPeople, err := s.getRegisteredPeople(data.Username)
	// if err != nil {
	// 	return false, err
	// }

	// if len(registeredPeople) > 0 {
	data.PASSWORD = s.HashPassword(data.PASSWORD)
	dataM, err := toolkit.ToM(data)
	if err != nil {
		return false, err
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName: m.NewEdmpUserModel().TableName(),
		Data:      dataM,
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return true, fmt.Errorf("Different data with same ID already exists")
		}

		return false, err
	}

	// err = s.insertLinkRolePeople(registeredPeople[0], data.Role)
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
	// }

	return true, fmt.Errorf("User id is not registered in tbl_people")
}

func (s *EdmpUserService) Update(data *m.EdmpUser) (bool, error) {
	data.UPDATEDAT = time.Now().String()

	rows := make([]m.EdmpUser, 0)
	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: m.NewEdmpUserModel().TableName(),
		Clause:    dbflex.Eq("username", data.USERNAME),
		Result:    &rows,
	})
	if err != nil {
		return false, nil
	}
	if len(rows) == 0 {
		return true, fmt.Errorf("User not found")
	}

	if data.PASSWORD == "" {
		// password not updated
		data.PASSWORD = rows[0].PASSWORD
	} else {
		// new password used
		data.PASSWORD = s.HashPassword(data.PASSWORD)
	}

	err = h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewEdmpUserModel().TableName(),
		Clause:    dbflex.Eq("username", data.USERNAME),
	})

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName: m.NewEdmpUserModel().TableName(),
		Data:      data,
	})

	// err = h.NewDBcmd().Save(h.SaveParam{
	// 	TableName: s.TableName,
	// 	Data:      data,
	// })
	return true, err
}

func (s *EdmpUserService) DeleteByUsername(username int) error {
	// usernames := make([]interface{}, 0)
	// if strings.Contains(username, ",") {
	// 	for _, username := range strings.Split(username, ",") {
	// 		username, _ := url.PathUnescape(username)
	// 		usernames = append(usernames, username)
	// 	}
	// } else {
	// 	usernames = append(usernames, username)
	// }

	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewEdmpUserModel().TableName(),
		Clause:    dbflex.Eq("username", username),
	})
	return err
}

func (s *EdmpUserService) SaveUsage(data toolkit.M) error {
	data.Set("Time", time.Now())
	data.Set("ID", toolkit.ToString(time.Now().UnixNano()))

	err := h.NewDBcmd().Insert(h.InsertParam{
		TableName: m.NewUserUsageModel().TableName(),
		Data:      data,
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return err
		}
	}

	return err
}

func (s *EdmpUserService) GetUsageTable(colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "users.sql")
	gridArgs.QueryName = "users-usage"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	///////// --------------------------------------------------COLUMN FILTER
	gridArgs.Colnames = append(gridArgs.Colnames,
		"USERNAME", "FULLNAME", "ROLE", "MODULE", "DESCRIPTION", "TIME", "RESOURCEURL",
	)

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		for _, colname := range gridArgs.Colnames {
			colname = ""
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colname)
		}
	} else {
		for _, colname := range gridArgs.Colnames {
			gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, colFilterM.GetString(colname))
		}
	}

	// gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	gridArgs.GroupCol = "-"
	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}
