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
	for i := 0; i < 200; i++ {
		mdt := m.NewPriorityReportsModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Owner_ID = fake.Day()
		mdt.Lead_ID = fake.Day()
		mdt.Sub_Risk_Type_ID = fake.Day()
		mdt.Rationale = fake.Words()

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
	for i := 0; i < 200; i++ {
		mdt := m.NewCRMModel()
		mdt.ID = i
		mdt.Name = fake.Words()
		mdt.Prority_Report_ID = fake.Day()
		mdt.CRM_Rationale = fake.Words()

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

func (s *DSCService) CreateLinkCRMCDEDummyData() error {
	toolkit.Println("CreateLinkCRMCDEDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewLinkCRMCDEModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.LinkCRMCDE, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewLinkCRMCDEModel()
		mdt.ID = i
		mdt.CRM_ID = fake.Day()
		mdt.CDE_ID = fake.Day()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewLinkCRMCDEModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateLinkCategoryPeopleDummyData() error {
	toolkit.Println("CreateLinkCategoryPeopleDummyData")
	err := h.NewDBcmd().Delete(h.DeleteParam{
		TableName: m.NewLinkCategoryPeopleModel().TableName(),
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	data := make([]*m.LinkCategoryPeople, 0)
	for i := 0; i < 200; i++ {
		mdt := m.NewLinkCategoryPeopleModel()
		mdt.ID = i
		mdt.Category_ID = fake.Day()
		mdt.People_ID = fake.Day()

		data = append(data, mdt)
	}

	err = h.NewDBcmd().Insert(h.InsertParam{
		TableName:       m.NewLinkCategoryPeopleModel().TableName(),
		Data:            data,
		ContinueOnError: true,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *DSCService) CreateTables() error {
	resultRows := make([]toolkit.M, 0)

	q := `CREATE TABLE Tbl_System
	( ID int NOT NULL,
	  System_Name varchar2(200) NOT NULL,
	  ITAM_ID int NOT NULL,
	  CONSTRAINT tbl_system_pk PRIMARY KEY (ID)
	);
	
	CREATE TABLE Tbl_People
	( ID int NOT NULL,
	  First_Name varchar2(200),
	  Last_Name varchar2(200),
	  Bank_ID varchar2(7),
	  Email_ID varchar2(50),
	  Function varchar2(50) NOT NULL,
	  Org_Unit varchar2(50) NOT NULL,
	  Status varchar2(25) NOT NULL,
	  CONSTRAINT tbl_people_pk PRIMARY KEY (ID)
	);
	
	CREATE TABLE Tbl_Role
	( ID int NOT NULL,
	  Role_Name varchar2(200) NOT NULL,
	  Role_Type varchar2(50) NOT NULL,
	  Role_Description varchar2(4000) NOT NULL,
	  CONSTRAINT tbl_role_pk PRIMARY KEY (ID)
	);
	
	CREATE TABLE Tbl_Category
	( ID int NOT NULL,
	  Name varchar2(200) NOT NULL,
	  Type varchar2(50) NOT NULL,
	  CONSTRAINT tbl_category_pk PRIMARY KEY (ID)
	);
	
	CREATE TABLE Tbl_Policy
	( ID int NOT NULL,
	  Info_Asset_Name varchar2(200) NOT NULL,
	  Description varchar2(4000) NOT NULL,
	  Confidentiality int,
	  Integrity int,
	  Availability int,
	  Overall_CIA_Rating int,
	  CONSTRAINT tbl_policy_pk PRIMARY KEY (ID)
	);
	
	CREATE TABLE Tbl_MD_Resource
	( ID int NOT NULL,
	  Name varchar2(200) NOT NULL,
	  Type varchar2(50) NOT NULL,
	  Description varchar2(200),
	  System_ID int,
	  CONSTRAINT tbl_resource_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_resource_system
		FOREIGN KEY (System_ID)
		REFERENCES Tbl_System(ID)
	);
	
	CREATE TABLE Tbl_Link_Role_People
	( ID int NOT NULL,
	  Role_ID int NOT NULL,
	  People_ID int NOT NULL,
	  Object_Type varchar(100),
	  Object_ID int,
	  CONSTRAINT tbl_rolepeople_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_rolepeople_role
		FOREIGN KEY (Role_ID)
		REFERENCES Tbl_Role(ID) on delete cascade,
	  CONSTRAINT fk_rolepeople_people
		FOREIGN KEY (People_ID)
		REFERENCES Tbl_People(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_Subcategory
	( ID int NOT NULL,
	  Name varchar(200) NOT NULL,
	  Type VARCHAR(50) NOT NULL,
	  Category_ID int NOT NULL,
	  CONSTRAINT tbl_subcategory_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_subcategory_category
		FOREIGN KEY (Category_ID)
		REFERENCES Tbl_Category(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_Link_Category_People
	( ID int NOT NULL,
	  Category_ID int NOT NULL,
	  People_ID int NOT NULL,
	  CONSTRAINT tbl_categorypeople_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_categorypeople_category
		FOREIGN KEY (Category_ID)
		REFERENCES Tbl_Category(ID) on delete cascade,
	  CONSTRAINT fk_categorypeople_people
		FOREIGN KEY (People_ID)
		REFERENCES Tbl_People(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_Link_Subcategory_People
	( ID int NOT NULL,
	  Subcategory_ID int NOT NULL,
	  People_ID int NOT NULL,
	  CONSTRAINT tbl_subcategorypeople_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_subcategorypeople_subcategory
		FOREIGN KEY (Subcategory_ID)
		REFERENCES Tbl_Subcategory(ID) on delete cascade,
	  CONSTRAINT fk_subcategorypeople_people
		FOREIGN KEY (People_ID)
		REFERENCES Tbl_People(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_Business_Term
	( ID int NOT NULL,
	  BT_Name VARCHAR(200) NOT NULL,
	  Parent_ID int NOT NULL,
	  Description VARCHAR(4000) NOT NULL,
	  CDE NUMBER(1,0),
	  CDE_Rationale VARCHAR(2000),
	  Mandatory NUMBER(1,0),
	  Policy_ID int,
	  Policy_Guidance VARCHAR(200),
	  DQ_Standards VARCHAR(200),
	  Threshold int,
	  Golden_Source_System_ID int,
	  Golden_Source_ITAM_ID int,
	  Golden_Source_TableName_ID int,
	  Golden_Source_Column_ID int,
	  Target_Golden_Source_ID int,
	  DDO_DQ_Standards VARCHAR(200),
	  DDO_Threshold int,
	  CONSTRAINT tbl_businessterm_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_businessterm_parent
		FOREIGN KEY (Parent_ID)
		REFERENCES Tbl_Subcategory(ID) on delete cascade,
	  CONSTRAINT fk_businessterm_policy
		FOREIGN KEY (Policy_ID)
		REFERENCES Tbl_Policy(ID) on delete cascade,
	  CONSTRAINT fk_businessterm_goldensourcesystem
		FOREIGN KEY (Golden_Source_System_ID)
		REFERENCES Tbl_System(ID) on delete cascade,
	  CONSTRAINT fk_businessterm_goldensource
		FOREIGN KEY (Target_Golden_Source_ID)
		REFERENCES Tbl_System(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_MD_Table
	( ID int NOT NULL,
	  Resource_ID int NOT NULL,
	  Schema_Name VARCHAR(200) NOT NULL,
	  Name VARCHAR(200) NOT NULL,
	  UUID VARCHAR(200),
	  Type VARCHAR(50),
	  Description VARCHAR(4000),
	  Business_Term_ID int,
	  Status NUMBER(1,0),
	  Record_Category VARCHAR(200),
	  CONSTRAINT tbl_table_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_table_resource
		FOREIGN KEY (Resource_ID)
		REFERENCES Tbl_MD_Resource(ID) on delete cascade,
	  CONSTRAINT fk_table_businessterm
		FOREIGN KEY (Business_Term_ID)
		REFERENCES Tbl_Business_Term(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_MD_Column
	( ID int NOT NULL,
	  Table_ID int NOT NULL,
	  Name VARCHAR(200) NOT NULL,
	  UUID VARCHAR(200) NOT NULL,
	  Type VARCHAR(50) NOT NULL,
	  Description VARCHAR(4000) NOT NULL,
	  Business_Term_ID int NOT NULL,
	  Data_Type VARCHAR(50) NOT NULL,
	  Data_Format VARCHAR(50) NOT NULL,
	  Data_Length int NOT NULL,
	  Example VARCHAR(50),
	  Derived NUMBER(1,0),
	  Derivation_Logic VARCHAR(4000),
	  Status NUMBER(1,0),
	  Alias_Name VARCHAR(200),
	  CDE NUMBER(1,0),
	  Sourced_from_Upstream NUMBER(1,0),
	  System_Checks VARCHAR(200),
	  Imm_Prec_System_ID int,
	  Imm_Prec_System_SLA NUMBER(1,0),
	  Imm_Prec_System_OLA NUMBER(1,0),
	  Imm_Succ_System_ID int,
	  Imm_Succ_System_SLA NUMBER(1,0),
	  Imm_Succ_System_OLA NUMBER(1,0),
	  Data_SLA_Signed NUMBER(1,0),
	  Golden_Source NUMBER(1,0),
	  DQ_Standards VARCHAR(200),
	  Threshold int,
	  DPO_DQ_Standards VARCHAR(200),
	  DPO_Threshold int,
	  DDO_DQ_Standards VARCHAR(200),
	  DDO_Threshold int,
	  PII_Flag NUMBER(1,0),
	  Record_Category VARCHAR(200),
	  CONSTRAINT tbl_column_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_column_table
		FOREIGN KEY (Table_ID)
		REFERENCES Tbl_MD_Table(ID) on delete cascade,
	  CONSTRAINT fk_column_businessterm
		FOREIGN KEY (Business_Term_ID)
		REFERENCES Tbl_Business_Term(ID) on delete cascade,
	  CONSTRAINT fk_column_precsystem
		FOREIGN KEY (Imm_Prec_System_ID)
		REFERENCES Tbl_System(ID) on delete cascade,
	  CONSTRAINT fk_column_succsystem
		FOREIGN KEY (Imm_Succ_System_ID)
		REFERENCES Tbl_System(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_DS_Processes
	( ID int NOT NULL,
	  Name VARCHAR(200) NOT NULL,
	  Owner_ID int,
	  Owner_Name VARCHAR(200) NOT NULL,
	  CONSTRAINT tbl_processes_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_processes_owner
		FOREIGN KEY (Owner_ID)
		REFERENCES Tbl_People(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_Segment
	( ID int NOT NULL,
	  Name VARCHAR(200) NOT NULL,
	  Subdomain_ID int NOT NULL,
	  CONSTRAINT tbl_segment_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_segment_owner
		FOREIGN KEY (Subdomain_ID)
		REFERENCES Tbl_Category(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_DS_Process_Detail
	( ID int NOT NULL,
	  Process_ID int NOT NULL,
	  Business_Term_ID int NOT NULL,
	  Segment_ID int NOT NULL,
	  Imm_Prec_System_ID int,
	  Ultimate_Source_System_ID int,
	  CONSTRAINT tbl_processdetail_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_processdetail_process
		FOREIGN KEY (Process_ID)
		REFERENCES Tbl_DS_Processes(ID) on delete cascade,
	  CONSTRAINT fk_processdetail_businessterm
		FOREIGN KEY (Business_Term_ID)
		REFERENCES Tbl_Business_Term(ID) on delete cascade,
	  CONSTRAINT fk_processdetail_segment
		FOREIGN KEY (Segment_ID)
		REFERENCES Tbl_Segment(ID) on delete cascade,
	  CONSTRAINT fk_processdetail_precsystem
		FOREIGN KEY (Imm_Prec_System_ID)
		REFERENCES Tbl_System(ID) on delete cascade,
	  CONSTRAINT fk_processdetail_sourcesystem
		FOREIGN KEY (Ultimate_Source_System_ID)
		REFERENCES Tbl_System(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_Priority_Reports
	( ID int NOT NULL,
	  Name VARCHAR(200) NOT NULL,
	  Owner_ID int NOT NULL,
	  Lead_ID int NOT NULL,
	  Sub_Risk_Type_ID int NOT NULL,
	  Rationale VARCHAR(2000) NOT NULL,
	  CONSTRAINT tbl_priorityreports_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_priorityreports_owner
		FOREIGN KEY (Owner_ID)
		REFERENCES Tbl_People(ID) on delete cascade,
	  CONSTRAINT fk_priorityreports_lead
		FOREIGN KEY (Lead_ID)
		REFERENCES Tbl_People(ID) on delete cascade,
	  CONSTRAINT fk_priorityreports_subcategory
		FOREIGN KEY (Sub_Risk_Type_ID)
		REFERENCES Tbl_Subcategory(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_CRM
	( ID int NOT NULL,
	  Name VARCHAR(300) NOT NULL,
	  Prority_Report_ID int NOT NULL,
	  CRM_Rationale VARCHAR(2000) NOT NULL,
	  CONSTRAINT tbl_crm_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_crm_priorityreports
		FOREIGN KEY (Prority_Report_ID)
		REFERENCES Tbl_Priority_Reports(ID) on delete cascade
	);
	
	CREATE TABLE Tbl_Link_CRM_CDE
	( ID int NOT NULL,
	  CRM_ID int NOT NULL,
	  CDE_ID int NOT NULL,
	  CONSTRAINT tbl_crmcde_pk PRIMARY KEY (ID),
	  CONSTRAINT fk_crmcde_crm
		FOREIGN KEY (CRM_ID)
		REFERENCES Tbl_CRM(ID) on delete cascade,
	  CONSTRAINT fk_crmcde_cde
		FOREIGN KEY (CDE_ID)
		REFERENCES Tbl_Business_Term(ID) on delete cascade
	);
	
	CREATE TABLE tbl_users
	( ID NUMBER GENERATED by default on null as IDENTITY,
	  username int NOT NULL,
	  password varchar(100) NOT NULL,
	  email varchar(100) NOT NULL,
	  name varchar(100) NOT NULL,
	  status int NOT NULL,
	  role varchar(100) NOT NULL,
	  createdat varchar(100) NOT NULL,
	  updatedat varchar(100) NOT NULL,
	  CONSTRAINT tbl_sysuser_pk PRIMARY KEY (ID)
	);`
	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewCategoryModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	if err != nil {
		return err
	}

	return nil
}
