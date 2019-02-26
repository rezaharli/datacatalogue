package services

import (
	"git.eaciitapp.com/sebar/dbflex"
	"github.com/eaciit/toolkit"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type DDOService struct {
}

func NewDDOService() *DDOService {
	ret := new(DDOService)
	return ret
}

func (s *DDOService) GetLeftTable(sortKey, sortOrder string, skip, take int, filter toolkit.M) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewCategoryModel().TableName(),
		SqlQuery:  "SELECT tc.id, tc.name as domain, tsc.name as sub_domain FROM tbl_category tc JOIN tbl_subcategory tsc ON tsc.category_id = tc.id",
		Results:   &resultRows,
	})

	// err := h.NewDBcmd().GetAll(h.GetAllParam{
	// 	Filter:      filter,
	// 	SortKey:     sortKey,
	// 	SortOrder:   sortOrder,
	// 	Skip:        skip,
	// 	Take:        take,
	// 	TableName:   m.NewCategoryModel().TableName(),
	// 	ResultRows:  &resultRows,
	// 	ResultTotal: &resultTotal,
	// })

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DDOService) GetRightTable(systemID int) (interface{}, int, error) {
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
