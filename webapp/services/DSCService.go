package services

import (
	"strings"

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

func (s *DSCService) GetAllSystem(loggedinid, search string, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
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
			LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = ts.id
			LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id
		WHERE
			upper(tlrp.object_type) = upper('system') `

	if loggedinid != "" {
		a := toolkit.ToInt(loggedinid, "")
		toolkit.Println(a)
		q += `AND tp.bank_id = '` + toolkit.ToString(a) + `' `
	}

	if search != "" {
		q += `
			AND (
				upper(ts.system_name) LIKE upper('%` + search + `%')
				OR upper(ts.itam_id) LIKE upper('%` + search + `%')
				OR upper(tp.first_name) LIKE upper('%` + search + `%')
				OR upper(tp.bank_id) LIKE upper('%` + search + `%')
			) `
	}

	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewSystemModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  pageNumber,
		RowsPerPage: rowsPerPage,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetTableName(systemID int, search string, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
			ts.id as tsid,
			tmt.id,
			tmt.name as table_name,
			tmc.id as colid,
			tmc.name as column_name,
			tmc.alias_name,
			tmc.cde `

	q += s.getSystemRightTableFROMandWHERE(systemID, search, nil)

	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewSystemModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  pageNumber,
		RowsPerPage: rowsPerPage,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetInterfacesRightTable(systemID int, search string, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
			tmt.id,
			ts.id as tsid,
			tmc.id as colid,
			tmc.alias_name as listof_cde,
			tmc.imm_prec_system_id,
			ips.system_name as imm_prec_system_name,
			tmc.Imm_Prec_System_SLA,
			tmc.Imm_Prec_System_OLA,
			tmc.imm_succ_system_id,
			iss.system_name as imm_succ_system_name,
			tmc.Imm_Succ_System_SLA,
			tmc.Imm_Succ_System_OLA,
			tdp.name as list_downstream_process,
			tdp.owner_id as downstream_owner `

	q += s.getInterfacesRightTableFROMandWHERE(systemID, search, nil)

	err := h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewSystemModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  pageNumber,
		RowsPerPage: rowsPerPage,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetDetails(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `
		SELECT DISTINCT
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
			ips.system_name as imm_prec_system_name,
			tmc.imm_succ_system_id,
			iss.system_name as imm_succ_system_name,
			tmc.threshold,
			tlcp.People_ID as domain_owner `

	if strings.Contains(payload.GetString("Which"), "interfaces") == true {
		toolkit.Println("interfaces")
		q += s.getInterfacesRightTableFROMandWHERE(payload.GetInt("Left"), "", payload)
	} else {
		toolkit.Println("system")
		q += s.getSystemRightTableFROMandWHERE(payload.GetInt("Left"), "", payload)
	}

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

func (s *DSCService) GetddSource(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := `SELECT DISTINCT
			tmt.name as table_name,
			tmc.name as column_name,
			tmc.alias_name `

	if strings.Contains(payload.GetString("Which"), "interfaces") == true {
		toolkit.Println("interfaces")
		q += s.getInterfacesRightTableFROMandWHERE(payload.GetInt("Left"), "", nil)
	} else {
		toolkit.Println("system")
		q += s.getSystemRightTableFROMandWHERE(payload.GetInt("Left"), "", nil)
	}

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

func (s *DSCService) getSystemRightTableFROMandWHERE(systemID int, search string, payload toolkit.M) string {
	q := `FROM tbl_system ts
			LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = ts.id
			LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id
			inner join tbl_md_resource tmr ON ts.id = tmr.system_id
			inner join tbl_md_table tmt ON tmr.id = tmt.resource_id
			inner join tbl_md_column tmc ON tmt.id = tmc.table_id
			LEFT JOIN tbl_business_term tbt ON tmc.business_term_id = tbt.id
			LEFT JOIN tbl_subcategory tsc ON tbt.parent_id = tsc.id
			LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
			LEFT JOIN tbl_link_category_people tlcp ON tlcp.category_id = tc.id
			LEFT JOIN tbl_policy tpol ON tbt.policy_id = tpol.id
			inner join tbl_link_column_interface ci on tmc.id = ci.column_id
			left outer join tbl_system ips on ci.imm_prec_system_id = ips.id
			left outer join tbl_system iss on ci.imm_succ_system_id = iss.id
			inner join
			(
				SELECT
				DISTINCT ts.id as sys_id, tmr.id as res_id, tmt.id as tab_id 
				FROM tbl_system ts
					inner join tbl_md_resource tmr ON ts.id = tmr.system_id
					inner join tbl_md_table tmt ON tmr.id = tmt.resource_id
					inner join tbl_md_column tmc ON tmt.id = tmc.table_id
				WHERE ts.id = ` + toolkit.ToString(systemID) + ` `

	if search != "" {
		q += `
				AND (
					upper(tmt.name) LIKE upper('%` + search + `%')
					OR upper(tmc.name) LIKE upper('%` + search + `%')
					OR upper(tmc.alias_name) LIKE upper('%` + search + `%')
					OR upper(tmc.cde) LIKE upper('%` + search + `%')
				) `
	}

	q += `		AND CDE = 1
			) cde ON ts.id = cde.sys_id and tmr.id = cde.res_id and tmt.id = cde.tab_id `

	if payload != nil {
		q += `WHERE ROWNUM = 1 `
		if payload.GetString("TableName") == "" {
			if payload.GetString("Right") != "" {
				q += `AND tmt.id = ` + payload.GetString("Right") + ` `
			}

			if payload.GetString("Column") != "" {
				q += `AND tmc.id = ` + payload.GetString("Column") + ` `
			}
		} else {
			q += `AND tmt.name = '` + payload.GetString("TableName") + `' `
		}

		if payload.GetString("ColumnName") != "" {
			q += ` AND tmc.name = '` + payload.GetString("ColumnName") + `' `
		}

		if payload.GetString("ScreenLabel") != "" {
			q += ` AND tmc.alias_name = '` + payload.GetString("ScreenLabel") + `' `
		}
	}
	q += `	ORDER BY tmt.name, tmc.name `

	return q
}

func (s *DSCService) getInterfacesRightTableFROMandWHERE(systemID int, search string, payload toolkit.M) string {
	q := `FROM tbl_system ts
			LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = ts.id
			LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id
			inner join tbl_md_resource res ON ts.id = res.system_id
			inner join tbl_md_table tmt ON res.id = tmt.resource_id
			inner join tbl_md_column tmc ON tmt.id = tmc.table_id
			inner join tbl_link_column_interface ci on tmc.id = ci.column_id
			left outer join tbl_system ips on ci.imm_prec_system_id = ips.id
			left outer join tbl_system iss on ci.imm_succ_system_id = iss.id
			LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.imm_prec_system_id = ips.id
			LEFT JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id
			LEFT JOIN tbl_business_term tbt ON tmc.business_term_id = tbt.id
			LEFT JOIN tbl_subcategory tsc ON tbt.parent_id = tsc.id
			LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
			LEFT JOIN tbl_link_category_people tlcp ON tlcp.category_id = tc.id
			LEFT JOIN tbl_policy tpol ON tbt.policy_id = tpol.id
		WHERE ts.id = ` + toolkit.ToString(systemID) + ` `

	if payload.GetString("TableName") == "" {
		if payload.GetString("Right") != "" {
			q += `AND tmt.id = ` + payload.GetString("Right") + ` `
		}

		if payload.GetString("Column") != "" {
			q += `AND tmc.id = ` + payload.GetString("Column") + ` `
		}
	} else {
		q += `AND tmt.name = '` + payload.GetString("TableName") + `' `
	}

	if payload.GetString("ColumnName") != "" {
		q += ` AND tmc.name = '` + payload.GetString("ColumnName") + `' `
	}

	if payload.GetString("ScreenLabel") != "" {
		q += ` AND tmc.alias_name = '` + payload.GetString("ScreenLabel") + `' `
	}

	if search != "" {
		q += `
		AND (
			upper(tmc.alias_name) LIKE upper('%` + search + `%')
			OR upper(ips.system_name) LIKE upper('%` + search + `%')
			OR upper(tmc.Imm_Prec_System_SLA) LIKE upper('%` + search + `%')
			OR upper(tmc.Imm_Prec_System_OLA) LIKE upper('%` + search + `%')
			OR upper(iss.system_name) LIKE upper('%` + search + `%')
			OR upper(tmc.Imm_Succ_System_SLA) LIKE upper('%` + search + `%')
			OR upper(tmc.Imm_Succ_System_OLA) LIKE upper('%` + search + `%')
			OR upper(tdp.name) LIKE upper('%` + search + `%')
			OR upper(tdp.owner_id) LIKE upper('%` + search + `%')
		) `
	}

	q += `AND (ips.system_name is not null or iss.system_name is not null)
		AND tmc.cde = 1 `

	return q
}
