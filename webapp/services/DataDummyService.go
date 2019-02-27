package services

import (
	"log"
	"math/rand"
	"time"

	"github.com/eaciit/toolkit"
	"github.com/icrowley/fake"

	h "eaciit/datacatalogue/webapp/helpers"
	m "eaciit/datacatalogue/webapp/models"
)

func (s *DSCService) CreateUserDummyData() error {
	toolkit.Println("CreateUserAdminData")
	data := m.NewSysUserModel()
	data.ID = 1
	data.Username = 123
	data.Password = "Password.1"
	data.Email = "eaciit@eaciit.com"
	data.Name = "eaciit"
	data.Status = 1
	data.Role = "Admin,DSC,DDO,DPO,RFO"
	data.CreatedAt = time.Now().String()
	data.UpdatedAt = time.Now().String()

	data.Password = NewUserService().HashPassword(data.Password)

	err := h.NewDBcmd().Insert(h.InsertParam{
		TableName: m.NewSysUserModel().TableName(),
		Data:      data,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
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

func (s *DSCService) CreateMDResourceDummyData() error {
	toolkit.Println("CreateMDResourceDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewMDResource().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.MDResource, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewMDResource()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Type = fake.Words()
		mdt.Description = fake.Words()
		mdt.System_ID = fake.Day()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewMDResource().TableName(),
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
		mdt.Schema_Name = fake.Words()
		mdt.Name = fake.Words()
		mdt.UUID = fake.Words()
		mdt.Type = fake.Words()
		mdt.Description = fake.Words()
		mdt.Business_Term_ID = fake.Day()
		mdt.Status = rand.Intn(9)
		mdt.Record_Category = fake.Words()

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

func (s *DSCService) CreatePeopleDummyData() error {
	toolkit.Println("CreatePeopleDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewPeopleModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.People, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewPeopleModel()
		mdt.ID = i
		mdt.First_Name = fake.Words()
		mdt.Last_Name = fake.Words()
		mdt.Bank_ID = fake.WordsN(7)[0:7]
		mdt.Email_ID = fake.Words()
		mdt.Function = fake.Words()
		mdt.Org_Unit = fake.Words()
		mdt.Status = fake.WordsN(25)[0:25]

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewPeopleModel().TableName(),
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
		mdt.Derived = rand.Intn(9)
		mdt.Derivation_Logic = fake.Words()
		mdt.Status = rand.Intn(9)
		mdt.Alias_Name = fake.Words()
		mdt.CDE = rand.Intn(9)
		mdt.Sourced_from_Upstream = rand.Intn(9)
		mdt.System_Checks = fake.Words()
		mdt.Imm_Prec_System_ID = fake.Day()
		mdt.Imm_Prec_System_SLA = rand.Intn(9)
		mdt.Imm_Prec_System_OLA = rand.Intn(9)
		mdt.Imm_Succ_System_ID = fake.Day()
		mdt.Imm_Succ_System_SLA = rand.Intn(9)
		mdt.Imm_Succ_System_OLA = rand.Intn(9)
		mdt.Data_SLA_Signed = rand.Intn(9)
		mdt.Golden_Source = rand.Intn(9)
		mdt.DQ_Standards = fake.Words()
		mdt.Threshold = fake.Day()
		mdt.DPO_DQ_Standards = fake.Words()
		mdt.DPO_Threshold = fake.Day()
		mdt.DDO_DQ_Standards = fake.Words()
		mdt.DDO_Threshold = fake.Day()
		mdt.PII_Flag = rand.Intn(9)
		mdt.Record_Category = fake.Words()

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
		mdt.CDE = rand.Intn(9)
		mdt.CDE_Rationale = fake.Words()
		mdt.Mandatory = rand.Intn(9)
		mdt.Policy_ID = fake.Day()
		mdt.Policy_Guidance = fake.Words()
		mdt.DQ_Standards = fake.Words()
		mdt.Threshold = fake.Day()
		mdt.Golden_Source_System_ID = fake.Day()
		mdt.Golden_Source_ITAM_ID = fake.Day()
		mdt.Golden_Source_TableName_ID = fake.Day()
		mdt.Golden_Source_Column_ID = fake.Day()
		mdt.Target_Golden_Source_ID = fake.Day()
		mdt.DDO_DQ_Standards = fake.Words()
		mdt.DDO_Threshold = fake.Day()

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

func (s *DSCService) CreateDSProcessesDummyData() error {
	toolkit.Println("CreateDSProcessesDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewDSProcessesModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.DSProcesses, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewDSProcessesModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Owner_ID = fake.Day()
		mdt.Owner_Name = fake.Words()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewDSProcessesModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateDSProcessesDetailDummyData() error {
	toolkit.Println("CreateDSProcessesDetailDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewDSProcessesDetailModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.DSProcessDetail, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewDSProcessesDetailModel()
		mdt.ID = i
		mdt.Process_ID = fake.Day()
		mdt.Business_Term_ID = fake.Day()
		mdt.Segment_ID = fake.Day()
		mdt.Imm_Prec_System_ID = fake.Day()
		mdt.Ultimate_Source_System_ID = fake.Day()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewDSProcessesDetailModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateSegmentDummyData() error {
	toolkit.Println("CreateSegmentDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewSegmentModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.Segment, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewSegmentModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Subdomain_ID = fake.Day()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewSegmentModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateLinkSubcategoryPeopleDummyData() error {
	toolkit.Println("CreateLinkSubcategoryPeopleDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewLinkSubcategoryPeopleModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.LinkSubcategoryPeople, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewLinkSubcategoryPeopleModel()
		mdt.ID = i
		mdt.Subcategory_ID = fake.Day()
		mdt.People_ID = fake.Day()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewLinkSubcategoryPeopleModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
