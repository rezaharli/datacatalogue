package services

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/dbflex"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type UserService struct {
	TableName string
}

func NewUserService() *UserService {
	ret := new(UserService)
	ret.TableName = m.NewSysUserModel().TableName()
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
		TableName:   s.TableName,
		ResultRows:  &resultRows,
		ResultTotal: &resultTotal,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *UserService) Insert(data *m.SysUser) error {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	users := make([]m.SysUser, 0)
	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: s.TableName,
		Clause:    dbflex.Eq("Username", data.Username),
		Result:    &users,
	})

	if err != nil {
		return err
	}
	if len(users) > 0 {
		return fmt.Errorf("The same username already exists")
	}

	data.Password = s.HashPassword(data.Password)
	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName: s.TableName,
		Data:      data,
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return fmt.Errorf("Different data with same ID already exists")
		}
	}

	return err
}
