package services

import (
	"git.eaciitapp.com/sebar/dbflex"
	"github.com/eaciit/toolkit"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

type DPOService struct {
}

func NewDPOService() *DPOService {
	ret := new(DPOService)
	return ret
}

func (s *DPOService) GetLeftTable(sortKey, sortOrder string, skip, take int, filter toolkit.M) ([]m.DSProcesses, int, error) {
	resultRows := make([]m.DSProcesses, 0)
	resultTotal := 0

	err := h.NewDBcmd().GetAll(h.GetAllParam{
		Filter:      filter,
		SortKey:     sortKey,
		SortOrder:   sortOrder,
		Skip:        skip,
		Take:        take,
		TableName:   m.NewDSProcessesModel().TableName(),
		ResultRows:  &resultRows,
		ResultTotal: &resultTotal,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DPOService) GetRightTable(processID int) (interface{}, int, error) {
	dsProcessDets := make([]m.DSProcessDetail, 0)
	err := h.NewDBcmd().GetBy(h.GetByParam{
		TableName: m.NewDSProcessesDetailModel().TableName(),
		Clause:    dbflex.Eq("process_id", processID),
		Result:    &dsProcessDets,
	})
	if err != nil {
		return nil, 0, err
	}

	res := []toolkit.M{}
	for _, dsProcessDet := range dsProcessDets {
		businessTerms := make([]m.BusinessTerm, 0)
		err := h.NewDBcmd().GetBy(h.GetByParam{
			TableName: m.NewBusinessTermModel().TableName(),
			Clause:    dbflex.Eq("id", dsProcessDet.Business_Term_ID),
			Result:    &businessTerms,
		})
		if err != nil {
			return nil, 0, err
		}

		resSubCat := []toolkit.M{}
		for _, buster := range businessTerms {
			subCats := make([]m.SubCategory, 0)
			err := h.NewDBcmd().GetBy(h.GetByParam{
				TableName: m.NewSubCategoryModel().TableName(),
				Clause:    dbflex.Eq("id", buster.Parent_ID),
				Result:    &subCats,
			})
			if err != nil {
				return nil, 0, err
			}

			resCats := []toolkit.M{}
			for _, subCat := range subCats {
				cats := make([]m.Category, 0)
				err := h.NewDBcmd().GetBy(h.GetByParam{
					TableName: m.NewCategoryModel().TableName(),
					Clause:    dbflex.Eq("id", subCat.Category_ID),
					Result:    &cats,
				})
				if err != nil {
					return nil, 0, err
				}

				subcat, _ := toolkit.ToM(subCat)

				tmpcat := m.NewCategoryModel()
				tmpcat = &cats[0]
				subcat.Set("Category", tmpcat)
				resCats = append(resCats, subcat)
			}

			busTerm, _ := toolkit.ToM(buster)

			toolkit.Println(len(resCats), resCats)
			tmpsubcat := &resCats[0]
			busTerm.Set("SubCategory", tmpsubcat)
			resSubCat = append(resSubCat, busTerm)
		}

		systems := make([]m.System, 0)
		err = h.NewDBcmd().GetBy(h.GetByParam{
			TableName: m.NewSystemModel().TableName(),
			Clause:    dbflex.Eq("id", dsProcessDet.Imm_Prec_System_ID),
			Result:    &systems,
		})
		if err != nil {
			return nil, 0, err
		}

		resSys := []toolkit.M{}
		for _, system := range systems {
			cols := make([]m.MDColumn, 0)
			err := h.NewDBcmd().GetBy(h.GetByParam{
				TableName: m.NewMDColumnModel().TableName(),
				Clause:    dbflex.Eq("imm_prec_system_id", system.ID),
				Result:    &cols,
			})
			if err != nil {
				return nil, 0, err
			}

			resCols := []toolkit.M{}
			for _, col := range cols {
				cats := make([]m.MDTable, 0)
				err := h.NewDBcmd().GetBy(h.GetByParam{
					TableName: m.NewMDTableModel().TableName(),
					Clause:    dbflex.Eq("id", col.Table_ID),
					Result:    &cats,
				})
				if err != nil {
					return nil, 0, err
				}

				c, _ := toolkit.ToM(col)

				tmpcat := m.NewMDTableModel()
				tmpcat = &cats[0]
				c.Set("MDTable", tmpcat)
				resCols = append(resCols, c)
			}

			sys, _ := toolkit.ToM(system)

			tmpcol := &resCols[0]
			sys.Set("Columns", tmpcol)
			resSys = append(resSys, sys)
		}

		dsProcessDet, _ := toolkit.ToM(dsProcessDet)

		tmpbuster := &resSubCat[0]
		dsProcessDet.Set("BusinessTerm", tmpbuster)

		tmpSys := &resSys[0]
		dsProcessDet.Set("System", tmpSys)

		res = append(res, dsProcessDet)
	}

	return res, 1, nil
}
