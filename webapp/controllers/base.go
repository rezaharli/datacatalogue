package controllers

import (
	"github.com/eaciit/toolkit"
)

type Base struct {
}

func NewBaseController() *Base {
	return new(Base)
}

func (c *Base) GetDetails(payload toolkit.M,
	detailsFunction func(payload toolkit.M) (interface{}, int, error),
	dropdownFunction func(payload toolkit.M) (interface{}, int, error),
) (interface{}, interface{}, error) {

	detail, _, err := detailsFunction(payload)
	if err != nil {
		return nil, nil, err
	}

	ddSource, _, err := dropdownFunction(payload)
	if err != nil {
		return nil, nil, err
	}

	return detail, ddSource, nil
}
