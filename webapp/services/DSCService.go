package services

import (
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

func (s *DSCService) GetAllSystem(sortKey, sortOrder string, skip, take int, filter toolkit.M) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
		ts.id,
		ts.system_name,
		ts.itam_id,
		tp.first_name,
		tp.bank_id
	FROM 
		Tbl_System ts
		JOIN Tbl_MD_Resource tmr ON tmr.system_id = ts.id
		JOIN Tbl_MD_Table tmt ON tmt.resource_id = tmr.id
		JOIN Tbl_MD_Column tmc ON tmc.table_id = tmt.id
		JOIN tbl_business_term tbt ON tmt.business_term_id = tbt.id
		JOIN tbl_ds_process_detail tdpd ON tdpd.business_term_id = tbt.id
		JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id
		JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
		JOIN tbl_category tc ON tsc.category_id = tc.id
		JOIN tbl_link_subcategory_people tlsp on tlsp.subcategory_id = tsc.id
		JOIN tbl_people tp on tlsp.people_id = tp.id
		JOIN tbl_policy tpol on tbt.policy_id = tpol.id`
	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewSystemModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetTableName(systemID int) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
		tmt.id,
		ts.id as tsid,
		tmt.name as table_name,
		tmc.name as column_name,
		tmc.alias_name,
		tmc.cde
	FROM 
		Tbl_System ts
		JOIN Tbl_MD_Resource tmr ON tmr.system_id = ts.id
		JOIN Tbl_MD_Table tmt ON tmt.resource_id = tmr.id
		JOIN Tbl_MD_Column tmc ON tmc.table_id = tmt.id
		JOIN tbl_business_term tbt ON tmt.business_term_id = tbt.id
		JOIN tbl_ds_process_detail tdpd ON tdpd.business_term_id = tbt.id
		JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id
		JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
		JOIN tbl_category tc ON tsc.category_id = tc.id
		JOIN tbl_link_subcategory_people tlsp on tlsp.subcategory_id = tsc.id
		JOIN tbl_people tp on tlsp.people_id = tp.id
		JOIN tbl_policy tpol on tbt.policy_id = tpol.id
	WHERE
		ts.id = ` + toolkit.ToString(systemID)
	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewSystemModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetInterfacesRightTable(systemID int) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
		tmt.id,
		ts.id as tsid,
		tmt.name as table_name,
		tmc.imm_prec_system_id,
		tmc.imm_succ_system_id,
		tdp.owner_id
	FROM 
		Tbl_System ts
		JOIN Tbl_MD_Resource tmr ON tmr.system_id = ts.id
		JOIN Tbl_MD_Table tmt ON tmt.resource_id = tmr.id
		JOIN Tbl_MD_Column tmc ON tmc.table_id = tmt.id
		JOIN tbl_business_term tbt ON tmt.business_term_id = tbt.id
		JOIN tbl_ds_process_detail tdpd ON tdpd.business_term_id = tbt.id
		JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id
		JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
		JOIN tbl_category tc ON tsc.category_id = tc.id
		JOIN tbl_link_subcategory_people tlsp on tlsp.subcategory_id = tsc.id
		JOIN tbl_people tp on tlsp.people_id = tp.id
		JOIN tbl_policy tpol on tbt.policy_id = tpol.id
	WHERE
		ts.id = ` + toolkit.ToString(systemID)
	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewSystemModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetDetails(leftParam int, rightParam int) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT 
			ts.id,
			ts.system_name,
			ts.itam_id,
			tp.first_name,
			tp.bank_id,
			tmc.alias_name,
			tmt.name as table_name,
			tmc.name as column_name,
			tbt.bt_name,
			tbt.description as business_description,
			tmc.cde,
			tmc.status,
			tmc.data_type,
			tmc.data_format,
			tmc.data_length,
			tmc.example,
			tmc.derived,
			tmc.Derivation_Logic,
			tmc.Sourced_from_Upstream,
			tmc.System_Checks,
			tc.Name as domain,
			tsc.name as subdomain,
			tpol.info_asset_name,
			tpol.description as info_asset_description,
			tpol.confidentiality,
			tpol.integrity,
			tpol.availability,
			tpol.overall_cia_rating,
			tmt.record_category,
			tmc.pii_flag,
			tmc.imm_prec_system_id,
			tmc.imm_succ_system_id,
			tmc.threshold
		FROM
			Tbl_System ts
			JOIN Tbl_MD_Resource tmr ON tmr.system_id = ts.id
			JOIN Tbl_MD_Table tmt ON tmt.resource_id = tmr.id
			JOIN Tbl_MD_Column tmc ON tmc.table_id = tmt.id
			JOIN tbl_business_term tbt ON tmt.business_term_id = tbt.id
			JOIN tbl_ds_process_detail tdpd ON tdpd.business_term_id = tbt.id
			JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id
			JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
			JOIN tbl_category tc ON tsc.category_id = tc.id
			JOIN tbl_link_subcategory_people tlsp on tlsp.subcategory_id = tsc.id
			JOIN tbl_people tp on tlsp.people_id = tp.id
			JOIN tbl_policy tpol on tbt.policy_id = tpol.id
		WHERE
			ts.id = ` + toolkit.ToString(leftParam) + ` and tmt.id = ` + toolkit.ToString(rightParam)
	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewCategoryModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

// func (s *DSCService) GetTableName(systemID int) (interface{}, int, error) {
// 	mdTables := make([]m.MDTable, 0)
// 	err := h.NewDBcmd().GetBy(h.GetByParam{
// 		TableName: m.NewMDTableModel().TableName(),
// 		Clause:    dbflex.Eq("imm_prec_system_id", systemID),
// 		Result:    &mdTables,
// 	})
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	res := []toolkit.M{}
// 	for _, mdTable := range mdTables {
// 		mdColumns := make([]m.MDColumn, 0)
// 		err := h.NewDBcmd().GetBy(h.GetByParam{
// 			TableName: m.NewMDColumnModel().TableName(),
// 			Clause:    dbflex.Eq("table_id", mdTable.ID),
// 			Result:    &mdColumns,
// 		})
// 		if err != nil {
// 			return nil, 0, err
// 		}

// 		resCol := []toolkit.M{}
// 		for _, mdCol := range mdColumns {
// 			businessTerm := make([]m.BusinessTerm, 0)
// 			err := h.NewDBcmd().GetBy(h.GetByParam{
// 				TableName: m.NewBusinessTermModel().TableName(),
// 				Clause:    dbflex.Eq("id", mdCol.Business_Term_ID),
// 				Result:    &businessTerm,
// 			})
// 			if err != nil {
// 				return nil, 0, err
// 			}

// 			resBus := []toolkit.M{}
// 			for _, buster := range businessTerm {
// 				subCats := make([]m.SubCategory, 0)
// 				err := h.NewDBcmd().GetBy(h.GetByParam{
// 					TableName: m.NewSubCategoryModel().TableName(),
// 					Clause:    dbflex.Eq("id", buster.Parent_ID),
// 					Result:    &subCats,
// 				})
// 				if err != nil {
// 					return nil, 0, err
// 				}

// 				resSubcat := []toolkit.M{}
// 				for _, subcat := range subCats {
// 					subCats := make([]m.Category, 0)
// 					err := h.NewDBcmd().GetBy(h.GetByParam{
// 						TableName: m.NewCategoryModel().TableName(),
// 						Clause:    dbflex.Eq("id", subcat.Category_ID),
// 						Result:    &subCats,
// 					})
// 					if err != nil {
// 						return nil, 0, err
// 					}

// 					sc, _ := toolkit.ToM(subcat)
// 					tmpSubCat := m.NewCategoryModel()
// 					tmpSubCat = &subCats[0]
// 					sc.Set("Category", tmpSubCat)
// 					resSubcat = append(resSubcat, sc)
// 				}

// 				policies := make([]m.Policy, 0)
// 				err = h.NewDBcmd().GetBy(h.GetByParam{
// 					TableName: m.NewPolicyModel().TableName(),
// 					Clause:    dbflex.Eq("id", buster.Policy_ID),
// 					Result:    &policies,
// 				})
// 				if err != nil {
// 					return nil, 0, err
// 				}

// 				bt, _ := toolkit.ToM(buster)

// 				tmpSubCat := &resSubcat[0]
// 				bt.Set("SubCategory", tmpSubCat)

// 				tmpPol := m.NewPolicyModel()
// 				tmpPol = &policies[0]
// 				bt.Set("Policy", tmpPol)

// 				resBus = append(resBus, bt)
// 			}

// 			mdColumn, _ := toolkit.ToM(mdCol)
// 			tmpBusinessTerm := &resBus[0]
// 			mdColumn.Set("BusinessTerms", tmpBusinessTerm)
// 			resCol = append(resCol, mdColumn)
// 		}

// 		mdTable, _ := toolkit.ToM(mdTable)
// 		mdTable.Set("Columns", resCol)
// 		res = append(res, mdTable)
// 	}

// 	return res, 1, nil
// }
