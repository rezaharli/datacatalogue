package services

import (
	"git.eaciitapp.com/sebar/dbflex"
	"github.com/eaciit/toolkit"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type DSCService struct {
}

func NewDSCService() *DSCService {
	ret := new(DSCService)
	return ret
}

func (s *DSCService) GetAllSystem(sortKey, sortOrder string, skip, take int, filter toolkit.M) ([]m.System, int, error) {
	resultRows := make([]m.System, 0)
	resultTotal := 0

	err := h.NewDBcmd().GetAll(h.GetAllParam{
		Filter:      filter,
		SortKey:     sortKey,
		SortOrder:   sortOrder,
		Skip:        skip,
		Take:        take,
		TableName:   m.NewSystemModel().TableName(),
		ResultRows:  &resultRows,
		ResultTotal: &resultTotal,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetTableName(systemID int) (interface{}, int, error) {
	mdTables := make([]m.MDTable, 0)
	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: m.NewMDTableModel().TableName(),
		Clause:    dbflex.Eq("imm_prec_system_id", systemID),
		Result:    &mdTables,
	})
	if err != nil {
		return nil, 0, err
	}

	res := []toolkit.M{}
	for _, mdTable := range mdTables {
		mdColumns := make([]m.MDColumn, 0)
		err := h.NewDBcmd().GetBy(h.GetByParam{
			TableName: m.NewMDColumnModel().TableName(),
			Clause:    dbflex.Eq("table_id", mdTable.ID),
			Result:    &mdColumns,
		})
		if err != nil {
			return nil, 0, err
		}

		resCol := []toolkit.M{}
		for _, mdCol := range mdColumns {
			businessTerm := make([]m.BusinessTerm, 0)
			err := h.NewDBcmd().GetBy(h.GetByParam{
				TableName: m.NewBusinessTermModel().TableName(),
				Clause:    dbflex.Eq("id", mdCol.Business_Term_ID),
				Result:    &businessTerm,
			})
			if err != nil {
				return nil, 0, err
			}

			resBus := []toolkit.M{}
			for _, buster := range businessTerm {
				subCats := make([]m.SubCategory, 0)
				err := h.NewDBcmd().GetBy(h.GetByParam{
					TableName: m.NewSubCategoryModel().TableName(),
					Clause:    dbflex.Eq("id", buster.Parent_ID),
					Result:    &subCats,
				})
				if err != nil {
					return nil, 0, err
				}

				resSubcat := []toolkit.M{}
				for _, subcat := range subCats {
					subCats := make([]m.Category, 0)
					err := h.NewDBcmd().GetBy(h.GetByParam{
						TableName: m.NewCategoryModel().TableName(),
						Clause:    dbflex.Eq("id", subcat.Category_ID),
						Result:    &subCats,
					})
					if err != nil {
						return nil, 0, err
					}

					sc, _ := toolkit.ToM(subcat)
					tmpSubCat := m.NewCategoryModel()
					tmpSubCat = &subCats[0]
					sc.Set("Category", tmpSubCat)
					resSubcat = append(resSubcat, sc)
				}

				policies := make([]m.Policy, 0)
				err = h.NewDBcmd().GetBy(h.GetByParam{
					TableName: m.NewPolicyModel().TableName(),
					Clause:    dbflex.Eq("id", buster.Policy_ID),
					Result:    &policies,
				})
				if err != nil {
					return nil, 0, err
				}

				bt, _ := toolkit.ToM(buster)

				tmpSubCat := &resSubcat[0]
				bt.Set("SubCategory", tmpSubCat)

				tmpPol := m.NewPolicyModel()
				tmpPol = &policies[0]
				bt.Set("Policy", tmpPol)

				resBus = append(resBus, bt)
			}

			mdColumn, _ := toolkit.ToM(mdCol)
			tmpBusinessTerm := &resBus[0]
			mdColumn.Set("BusinessTerms", tmpBusinessTerm)
			resCol = append(resCol, mdColumn)
		}

		mdTable, _ := toolkit.ToM(mdTable)
		mdTable.Set("Columns", resCol)
		res = append(res, mdTable)
	}

	return res, 1, nil
}
