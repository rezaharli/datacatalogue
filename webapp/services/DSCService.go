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
	for _, val := range mdTables {
		mdColumns := make([]m.MDColumn, 0)
		err := h.NewDBcmd().GetBy(h.GetByParam{
			TableName: m.NewMDColumnModel().TableName(),
			Clause:    dbflex.Eq("table_id", val.ID),
			Result:    &mdColumns,
		})
		if err != nil {
			return nil, 0, err
		}

		mdTable, _ := toolkit.ToM(val)
		mdTable.Set("Columns", mdColumns)
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
		system.System_Name = fake.WordsN(1)
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
		mdt.UUID = fake.WordsN(1)
		mdt.Type = fake.WordsN(1)
		mdt.Description = fake.Words()
		mdt.Business_Term_ID = fake.Day()
		mdt.Status = true
		mdt.Imm_Prec_System_ID = fake.Day()
		mdt.Imm_Succ_System_ID = fake.Day()
		mdt.Golden_Source = true
		mdt.Data_SLA_Signed = true
		mdt.DSC_DQ_Standards = fake.WordsN(1)
		mdt.DSC_Threshold = fake.Day()
		mdt.DPO_DQ_Standards = fake.WordsN(1)
		mdt.DPO_Threshold = fake.Day()
		mdt.DDO_DQ_Standards = fake.WordsN(1)
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
		mdt.Name = fake.WordsN(1)
		mdt.UUID = fake.WordsN(1)
		mdt.Type = fake.WordsN(1)
		mdt.Description = fake.WordsN(1)
		mdt.Business_Term_ID = fake.Day()
		mdt.Data_Type = fake.WordsN(1)
		mdt.Data_Format = fake.WordsN(1)
		mdt.Data_Length = fake.Day()
		mdt.Example = fake.WordsN(1)
		mdt.Derived = true
		mdt.Derivation_Logic = fake.WordsN(1)
		mdt.Mandatory = true
		mdt.Status = true
		mdt.Alias_Name = fake.WordsN(1)
		mdt.CDE = true
		mdt.Sourced_from_Upstream = true
		mdt.System_Checks = fake.WordsN(1)
		mdt.Imm_Prec_System_ID = fake.Day()
		mdt.Imm_Succ_System_ID = fake.Day()
		mdt.Data_SLA_Signed = true
		mdt.Golden_Source = true
		mdt.DQ_Standards = fake.WordsN(1)
		mdt.Threshold = fake.Day()
		mdt.DPO_DQ_Standards = fake.WordsN(1)
		mdt.DPO_Threshold = fake.Day()
		mdt.DDO_DQ_Standards = fake.WordsN(1)
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
