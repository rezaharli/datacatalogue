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
	for i := 0; i < 10000; i++ {
		system := m.NewSystemModel()
		system.ID = i
		system.System_Name = fake.Words()
		system.Display_Name = fake.Words()
		system.ITAM_ID = toolkit.ToInt(fake.DigitsN(4), "")
		system.Created_DateTime = time.Now()
		system.Modified_DateTime = time.Now()
		system.Status = rand.Intn(2)

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
	for i := 0; i < 10000; i++ {
		mdt := m.NewMDResource()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Type = fake.Words()
		mdt.Description = fake.Words()
		mdt.System_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

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
	for i := 0; i < 10000; i++ {
		mdt := m.NewPeopleModel()
		mdt.ID = i
		mdt.First_Name = fake.FirstName()
		mdt.Last_Name = fake.LastName()
		mdt.Bank_ID = "100" + toolkit.ToString(i)
		mdt.Email_ID = fake.EmailAddress()
		mdt.Function = fake.JobTitle()
		mdt.Org_Unit = fake.JobTitle()
		mdt.Status = fake.WordsN(25)[0:25]
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()

		if i == 0 {
			mdt.Bank_ID = "123"
		}

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

func (s *DSCService) CreateRoleDummyData() error {
	toolkit.Println("CreateRoleDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewRoleModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.Role, 0)
	for i := 0; i < 10000; i++ {
		mdt := m.NewRoleModel()

		mdt.ID = i
		mdt.Role_Name = fake.Words()
		mdt.Role_Type = fake.Words()
		mdt.Role_Description = fake.Words()
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

		if i%1 == 0 {
			mdt.Role_Name = "Dataset Custodian"
		}
		if i%2 == 0 {
			mdt.Role_Name = "Downstream Process Owner"
		}
		if i%3 == 0 {
			mdt.Role_Name = "Data Domain Owner"
		}

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewRoleModel().TableName(),
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
	for i := 0; i < 10000; i++ {
		mdt := m.NewCategoryModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Type = fake.Words()
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

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
	for i := 0; i < 10000; i++ {
		mdt := m.NewSubCategoryModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Type = fake.Words()
		mdt.Category_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

		if i%1 == 0 {
			mdt.Type = "Sub Data Domain"
		}

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
	for i := 0; i < 10000; i++ {
		mdt := m.NewPolicyModel()
		mdt.ID = i
		mdt.Info_Asset_Name = fake.Words()
		mdt.Description = fake.Words()
		mdt.Confidentiality = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Integrity = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Availability = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Overall_CIA_Rating = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

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
	for i := 0; i < 10000; i++ {
		mdt := m.NewBusinessTermModel()
		mdt.ID = i
		mdt.BT_Name = fake.Words()
		mdt.Parent_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Description = fake.WordsN(100)
		mdt.Mandatory = rand.Intn(2)
		mdt.Policy_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Policy_Guidance = fake.Words()
		mdt.DQ_Standards = fake.Words()
		mdt.Threshold = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Remarks = fake.WordsN(100)
		mdt.Golden_Source_System_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Golden_Source_ITAM_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Golden_Source_TableName_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Golden_Source_Column_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Target_Golden_Source_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.DDO_DQ_Standards = fake.Words()
		mdt.DDO_Threshold = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

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
	for i := 0; i < 10000; i++ {
		mdt := m.NewMDTableModel()
		mdt.ID = i
		mdt.Resource_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Schema_Name = fake.Words()
		mdt.Name = fake.Words()
		mdt.Display_Name = fake.Words()
		mdt.UUID = fake.Words()
		mdt.Type = fake.Words()
		mdt.Description = fake.Words()
		mdt.Business_Term_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Status = rand.Intn(2)
		mdt.Record_Category = fake.Words()
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()

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
	for i := 0; i < 10000; i++ {
		mdt := m.NewMDColumnModel()
		mdt.ID = i
		mdt.Table_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Name = fake.Words()
		mdt.Display_Name = fake.Words()
		mdt.UUID = fake.Words()
		mdt.Type = fake.Words()
		mdt.Data_Type = fake.Words()
		mdt.Data_Format = fake.Words()
		mdt.Data_Length = fake.Words()
		mdt.Primary_Key = rand.Intn(2)
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()

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

func (s *DSCService) CreateMDColumnDetailsDummyData() error {
	toolkit.Println("CreateMDColumnDetailsDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewMDColumnDetailsModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.MDColumnDetails, 0)
	for i := 0; i < 10000; i++ {
		mdt := m.NewMDColumnDetailsModel()
		mdt.ID = i
		mdt.Column_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Alias_Name = fake.Words()
		mdt.Description = fake.WordsN(100)
		mdt.CDE = rand.Intn(2)
		mdt.CDE_Rationale = fake.WordsN(100)
		mdt.Mandatory = rand.Intn(2)
		mdt.Sourced_from_Upstream = rand.Intn(2)
		mdt.Derived = rand.Intn(2)
		mdt.Derivation_Logic = fake.WordsN(100)
		mdt.PII_Flag = rand.Intn(2)
		mdt.Status = rand.Intn(2)
		mdt.Data_SLA_Signed = rand.Intn(2)
		mdt.Example = fake.WordsN(100)
		mdt.System_Checks = fake.Words()
		mdt.DQ_Standards = fake.WordsN(100)
		mdt.Threshold = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Record_Category = fake.Words()
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewMDColumnDetailsModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateLinkColumnBusinessTermDummyData() error {
	toolkit.Println("CreateLinkColumnBusinessTermDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewLinkColumnBusinessTermModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.LinkColumnBusinessTerm, 0)
	for i := 0; i < 10000; i++ {
		mdt := m.NewLinkColumnBusinessTermModel()
		mdt.ID = i
		mdt.Column_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Business_Term_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewLinkColumnBusinessTermModel().TableName(),
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
	for i := 0; i < 10000; i++ {
		mdt := m.NewDSProcessesModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Description = fake.WordsN(100)
		mdt.Process_ID = fake.Words()
		mdt.Owner_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Owner_Name = fake.Words()
		mdt.System_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

		if i%1 == 0 {
			mdt.Name = "USFRR"
		}
		if i%2 == 0 {
			mdt.Name = "LC-FCC | Global Name Screening"
		}
		if i%3 == 0 {
			mdt.Name = "Credit Risk | Risk Reporting | Credit Risk Reporting & MI"
		}

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
	for i := 0; i < 10000; i++ {
		mdt := m.NewSegmentModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Subdomain_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

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
	for i := 0; i < 10000; i++ {
		mdt := m.NewDSProcessesDetailModel()
		mdt.ID = i
		mdt.Process_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Column_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Segment_Name = fake.Words()
		mdt.Imm_Prec_System_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Ultimate_Source_System_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.USS_CDE_NAME = fake.Words()
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()

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

func (s *DSCService) CreateLinkDSProcessSegmentDummyData() error {
	toolkit.Println("CreateLinkDSProcessSegmentDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewLinkDSProcessSegmentModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.LinkDSProcessSegment, 0)
	for i := 0; i < 10000; i++ {
		mdt := m.NewLinkDSProcessSegmentModel()
		mdt.ID = i
		mdt.Process_ID = fake.Words()
		mdt.Segment_Name = fake.Words()
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewLinkDSProcessSegmentModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateLinkColumnInterfaceDummyData() error {
	toolkit.Println("CreateLinkColumnInterfaceDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewLinkColumnInterfaceModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.LinkColumnInterface, 0)
	for i := 0; i < 10000; i++ {
		mdt := m.NewLinkColumnInterfaceModel()
		mdt.ID = i
		mdt.Column_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.System_SLA = rand.Intn(2)
		mdt.System_OLA = rand.Intn(2)
		mdt.Imm_Prec_System_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Incoming_CDE_Name = fake.Words()
		mdt.Incoming_Derived = rand.Intn(2)
		mdt.Incoming_Derivation_Logic = fake.Words()
		mdt.Imm_Prec_Column_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Imm_Succ_System_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Outgoing_CDE_Name = fake.Words()
		mdt.Outgoing_Derived = rand.Intn(2)
		mdt.Outgoing_Derivation_Logic = fake.Words()
		mdt.Imm_Succ_Column_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewLinkColumnInterfaceModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateLinkColumnGoldenSourceDummyData() error {
	toolkit.Println("CreateLinkColumnGoldenSourceDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewLinkColumnGoldenSourceModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.LinkColumnGoldenSource, 0)
	for i := 0; i < 10000; i++ {
		mdt := m.NewLinkColumnGoldenSourceModel()
		mdt.ID = i
		mdt.Column_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Golden_Source_Column_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewLinkColumnGoldenSourceModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreatePriorityReportsDummyData() error {
	toolkit.Println("CreatePriorityReportsDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewPriorityReportsModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.PriorityReports, 0)
	for i := 0; i < 10000; i++ {
		mdt := m.NewPriorityReportsModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Owner_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Lead_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Sub_Risk_Type_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Rationale = fake.Words()
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewPriorityReportsModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateCRMDummyData() error {
	toolkit.Println("CreateCRMDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewCRMModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.CRM, 0)
	for i := 0; i < 10000; i++ {
		mdt := m.NewCRMModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Prority_Report_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.CRM_Rationale = fake.Words()
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewCRMModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateCDEDummyData() error {
	toolkit.Println("CreateCDEDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewCDEModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.CDE, 0)
	for i := 0; i < 10000; i++ {
		mdt := m.NewCDEModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.CRM_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.CDE_Rationale = fake.Words()
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()
		mdt.Status = rand.Intn(2)

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewCDEModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateLinkRolePeopleDummyData() error {
	toolkit.Println("CreateLinkRolePeopleDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewLinkRolePeopleModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.LinkRolePeople, 0)
	for i := 0; i < 10000; i++ {
		mdt := m.NewLinkRolePeopleModel()
		mdt.ID = i
		mdt.Role_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.People_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Object_Type = fake.Words()
		mdt.Object_ID = toolkit.ToInt(fake.DigitsN(4), "")
		mdt.Created_DateTime = time.Now()
		mdt.Modified_DateTime = time.Now()

		if i%1 == 0 {
			mdt.Object_Type = "SYSTEM"
		}
		if i%2 == 0 {
			mdt.Object_Type = "PROCESSES"
		}
		if i%3 == 0 {
			mdt.Object_Type = "SUBCATEGORY"
		}

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewLinkRolePeopleModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

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

	ok, err := NewUserService().Insert(data)
	if !ok && err != nil {
		log.Println(err.Error())
		return err
	}
	if ok && err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
