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

type UserService struct {
	*Base
}

func NewUserService() *UserService {
	ret := new(UserService)
	return ret
}

var PASSWORD_SALT = "bb733e7939a434225752c6b70dee3f18"

func (s *UserService) HashPassword(password string) string {
	hashedPassword := md5.New()
	io.WriteString(hashedPassword, fmt.Sprintf("%s%s", password, PASSWORD_SALT))

	return fmt.Sprintf("%x", hashedPassword.Sum(nil))
}

func (s *UserService) Authenticate(username int, password string) (bool, *m.SysUser, error) {
	users := make([]m.SysUser, 0)

	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: m.NewSysUserModel().TableName(),
		Clause: dbflex.And(
			dbflex.Eq("username", username),
			dbflex.Eq("password", s.HashPassword(password)),
			dbflex.Eq("status", 1),
		),
		Result: &users,
	})
	if err != nil {
		return false, m.NewSysUserModel(), err
	}

	if len(users) == 0 {
		return false, m.NewSysUserModel(), nil
	}

	return true, &(users[0]), nil
}

func (s *UserService) GetAll(tabs, loggedinid, search string, searchDD, colFilter interface{}, pagination toolkit.M) ([]toolkit.M, int, error) {
	gridArgs := GridArgs{}
	gridArgs.QueryFilePath = filepath.Join(clit.ExeDir(), "queryfiles", "users.sql")
	gridArgs.QueryName = "users"
	gridArgs.PageNumber = pagination.GetInt("page")
	gridArgs.RowsPerPage = pagination.GetInt("rowsPerPage")

	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter, "", "", "", "", "", "", "")
	} else {
		gridArgs.ColumnFilter = append(gridArgs.ColumnFilter,
			colFilterM.GetString("PRINCIPAL_RISK_TYPES"),
			colFilterM.GetString("RISK_SUB_TYPE"),
			colFilterM.GetString("RISK_FRAMEWORK_OWNER"),
			colFilterM.GetString("RISK_REPORTING_LEAD"),
			colFilterM.GetString("PR_COUNT"),
			colFilterM.GetString("CRM_COUNT"),
			colFilterM.GetString("CDE_COUNT"),
		)
	}

	gridArgs.OrderBy = pagination.GetString("sortBy")
	descending := pagination.Get("descending")
	if descending != nil {
		gridArgs.IsDescending = descending.(bool)
	}

	return s.Base.ExecuteGridQueryFromFile(gridArgs)
}

func (s *UserService) getRegisteredPeople(username int) ([]m.People, error) {
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

func (s *UserService) insertLinkRolePeople(people m.People, role string) (err error) {
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

func (s *UserService) Insert(data *m.SysUser) (bool, error) {
	data.CreatedAt = time.Now().String()
	data.UpdatedAt = time.Now().String()

	users := make([]m.SysUser, 0)
	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: m.NewSysUserModel().TableName(),
		Clause:    dbflex.Eq("username", data.Username),
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
	data.Password = s.HashPassword(data.Password)
	dataM, err := toolkit.ToM(data)
	if err != nil {
		return false, err
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName: m.NewSysUserModel().TableName(),
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

func (s *UserService) Update(data *m.SysUser) (bool, error) {
	data.UpdatedAt = time.Now().String()

	rows := make([]m.SysUser, 0)
	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: m.NewSysUserModel().TableName(),
		Clause:    dbflex.Eq("username", data.Username),
		Result:    &rows,
	})
	if err != nil {
		return false, nil
	}
	if len(rows) == 0 {
		return true, fmt.Errorf("User not found")
	}

	if data.Password == "" {
		// password not updated
		data.Password = rows[0].Password
	} else {
		// new password used
		data.Password = s.HashPassword(data.Password)
	}

	err = h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewSysUserModel().TableName(),
		Clause:    dbflex.Eq("username", data.Username),
	})

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName: m.NewSysUserModel().TableName(),
		Data:      data,
	})

	// err = h.NewDBcmd().Save(h.SaveParam{
	// 	TableName: s.TableName,
	// 	Data:      data,
	// })
	return true, err
}

func (s *UserService) DeleteByUsername(username int) error {
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
		TableName: m.NewSysUserModel().TableName(),
		Clause:    dbflex.Eq("username", username),
	})
	return err
}

func (s *UserService) SaveUsage(data toolkit.M) error {
	data.Set("Time", time.Now().String())

	data.Unset("ID")

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
