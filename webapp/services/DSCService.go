package services

import (
	"log"

	"git.eaciitapp.com/sebar/dbflex"
	"github.com/eaciit/toolkit"
	"github.com/icrowley/fake"

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

func (s *DSCService) CreateSystemDummyData() error {
	toolkit.Println("CreateSystemDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewSystemModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.System, 0)
	for i := 0; i < 200; i++ {
		system := m.NewSystemModel()
		system.ID = i
		system.System_Name = fake.Words()
		system.ITAM_ID = fake.Day()
		system.MD_Resource_ID = fake.Day()

		data = append(data, system)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewSystemModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateMDTableDummyData() error {
	toolkit.Println("CreateMDTableDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewMDTableModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.MDTable, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewMDTableModel()
		mdt.ID = i
		mdt.Resource_ID = fake.Day()
		mdt.Name = fake.Company()
		mdt.UUID = fake.Words()
		mdt.Type = fake.Words()
		mdt.Description = fake.Words()
		mdt.Business_Term_ID = fake.Day()
		mdt.Status = true
		mdt.Imm_Prec_System_ID = fake.Day()
		mdt.Imm_Succ_System_ID = fake.Day()
		mdt.Golden_Source = true
		mdt.Data_SLA_Signed = true
		mdt.DSC_DQ_Standards = fake.Words()
		mdt.DSC_Threshold = fake.Day()
		mdt.DPO_DQ_Standards = fake.Words()
		mdt.DPO_Threshold = fake.Day()
		mdt.DDO_DQ_Standards = fake.Words()
		mdt.DDO_Threshold = fake.Day()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewMDTableModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateMDColumnDummyData() error {
	toolkit.Println("CreateMDColumnDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewMDColumnModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.MDColumn, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewMDColumnModel()
		mdt.ID = i
		mdt.Table_ID = fake.Day()
		mdt.Name = fake.Words()
		mdt.UUID = fake.Words()
		mdt.Type = fake.Words()
		mdt.Description = fake.Words()
		mdt.Business_Term_ID = fake.Day()
		mdt.Data_Type = fake.Words()
		mdt.Data_Format = fake.Words()
		mdt.Data_Length = fake.Day()
		mdt.Example = fake.Words()
		mdt.Derived = true
		mdt.Derivation_Logic = fake.Words()
		mdt.Mandatory = true
		mdt.Status = true
		mdt.Alias_Name = fake.Words()
		mdt.CDE = true
		mdt.Sourced_from_Upstream = true
		mdt.System_Checks = fake.Words()
		mdt.Imm_Prec_System_ID = fake.Day()
		mdt.Imm_Succ_System_ID = fake.Day()
		mdt.Data_SLA_Signed = true
		mdt.Golden_Source = true
		mdt.DQ_Standards = fake.Words()
		mdt.Threshold = fake.Day()
		mdt.DPO_DQ_Standards = fake.Words()
		mdt.DPO_Threshold = fake.Day()
		mdt.DDO_DQ_Standards = fake.Words()
		mdt.DDO_Threshold = fake.Day()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewMDColumnModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateBusinessTermDummyData() error {
	toolkit.Println("CreateBusinessTermDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewBusinessTermModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.BusinessTerm, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewBusinessTermModel()
		mdt.ID = i
		mdt.BT_Name = fake.Words()
		mdt.Parent_ID = fake.Day()
		mdt.Description = fake.Words()
		mdt.CDE = true
		mdt.CDE_Rationale = fake.Words()
		mdt.Status = true
		mdt.Policy_ID = fake.Day()
		mdt.DQ_Standards = fake.Words()
		mdt.Threshold = fake.Day()
		mdt.Golden_Source_ID = fake.Day()
		mdt.Target_Golden_Source_ID = fake.Day()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewBusinessTermModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateSubCategoryDummyData() error {
	toolkit.Println("CreateSubCategoryDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewSubCategoryModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.SubCategory, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewSubCategoryModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Type = fake.Words()
		mdt.Category_ID = fake.Day()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewSubCategoryModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateCategoryDummyData() error {
	toolkit.Println("CreateCategoryDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewCategoryModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.Category, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewCategoryModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Type = fake.Words()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewCategoryModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreatePolicyDummyData() error {
	toolkit.Println("CreatePolicyDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewPolicyModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.Policy, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewPolicyModel()
		mdt.ID = i
		mdt.Info_Asset_Name = fake.Words()
		mdt.Description = fake.Words()
		mdt.Confidentiality = fake.Day()
		mdt.Integrity = fake.Day()
		mdt.Availability = fake.Day()
		mdt.Overall_CIA_Rating = fake.Day()
		mdt.Record_Category = fake.Words()
		mdt.PII_Flag = true
		mdt.Policy_Guidance = fake.Words()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewPolicyModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
