package services

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/icrowley/fake"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/dbflex"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type UserService struct {
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

func (s *UserService) GetAll(sortKey, sortOrder string, skip, take int, filter toolkit.M) ([]m.SysUser, int, error) {
	resultRows := make([]m.SysUser, 0)
	resultTotal := 0

	err := h.NewDBcmd().GetAll(h.GetAllParam{
		Filter:      filter,
		SortKey:     sortKey,
		SortOrder:   sortOrder,
		Skip:        skip,
		Take:        take,
		TableName:   m.NewSysUserModel().TableName(),
		ResultRows:  &resultRows,
		ResultTotal: &resultTotal,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *UserService) getRegisteredPeople(username int) ([]m.People, error) {
	people := make([]m.People, 0)
	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: m.NewPeopleModel().TableName(),
		Clause:    dbflex.Eq("ID", username),
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

	registeredPeople, err := s.getRegisteredPeople(data.Username)
	if err != nil {
		return false, err
	}

	if len(registeredPeople) > 0 {
		data.Password = s.HashPassword(data.Password)

		dataM, err := toolkit.ToM(data)
		if err != nil {
			return false, err
		}

		dataM.Unset("ID")

		err = h.NewDBcmd().Insert(h.InsertParam{
			TableName: m.NewSysUserModel().TableName(),
			Data:      dataM,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				return true, fmt.Errorf("Different data with same ID already exists")
			} else {
				return false, err
			}
		}

		err = s.insertLinkRolePeople(registeredPeople[0], data.Role)
		if err != nil {
			return false, err
		}

		return true, nil
	}

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
