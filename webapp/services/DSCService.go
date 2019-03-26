package services

import (
	"path/filepath"
	"strings"

	"github.com/eaciit/clit"
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

func (s *DSCService) GetAllSystem(tabs, loggedinid, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) ([]toolkit.M, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	if loggedinid != "" {
		args = append(args, loggedinid)
	}

	///////// --------------------------------------------------DROPDOWN FILTER
	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	filterSystemName := ""
	if search != "" {
		filterSystemName = search
	} else {
		filterSystemName = searchDDM.GetString("SystemName")
	}

	args = append(args, filterSystemName)
	args = append(args, searchDDM.GetString("ItamID"))

	///////// --------------------------------------------------COLUMN FILTER
	colFilterM, err := toolkit.ToM(colFilter)
	if err != nil {
		args = append(args, "", "", "", "")
	} else {
		args = append(args,
			colFilterM.GetString("SYSTEM_NAME"),
			colFilterM.GetString("ITAM_ID"),
			colFilterM.GetString("DATASET_CUSTODIAN"),
			colFilterM.GetString("BANK_ID"),
		)
	}

	///////// --------------------------------------------------BUILD QUERY FROM ARGS
	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err = h.BuildQueryFromFile(filePath, "left-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	///////// --------------------------------------------------EXECUTE
	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
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

func (s *DSCService) GetTableName(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(systemID))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "right-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	filterTableName := searchDDM.GetString("TableName")
	filterColumnName := searchDDM.GetString("ColumnName")

	///////// DROPDOWN FILTER
	q = `SELECT * FROM (
		` + q + `
	) WHERE (
		upper(table_name) LIKE upper('%` + filterTableName + `%')
		AND upper(column_name) LIKE upper('%` + filterColumnName + `%')
	) `

	colFilterM, err := toolkit.ToM(colFilter)
	cf := make([]string, 0)
	if err != nil {
		cf = append(cf, "", "", "", "")
	} else {
		cf = append(cf, colFilterM.GetString("TABLE_NAME"), colFilterM.GetString("COLUMN_NAME"), colFilterM.GetString("BUSINESS_ALIAS_NAME"), colFilterM.GetString("CDE_YES_NO"))
	}

	///////// COLUMN FILTER
	if cf[0] != "" || cf[1] != "" || cf[2] != "" || cf[3] != "" {
		q += `AND (
			upper(table_name) LIKE upper('%` + cf[0] + `%')
			AND upper(column_name) LIKE upper('%` + cf[1] + `%')
			AND upper(business_alias_name) LIKE upper('%` + cf[2] + `%')
			AND upper(cde_yes_no) LIKE upper('%` + cf[3] + `%')
		) `
	}

	///////// COUNT
	q = `SELECT res.*, 
			COUNT(DISTINCT table_name) OVER () COUNT_table_name,
			COUNT(DISTINCT column_name) OVER () COUNT_column_name,
			COUNT(DISTINCT business_alias_name) OVER () COUNT_business_alias_name,
			(SELECT COUNT (cde_yes_no) FROM ( ` + q + `) res2 WHERE cde_yes_no = 1) COUNT_cde_yes_no
		FROM (
			` + q + `
		) res `

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
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

func (s *DSCService) GetInterfacesRightTable(tabs string, systemID int, search string, searchDD, colFilter interface{}, pageNumber, rowsPerPage int, filter toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(systemID))

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "right-grid", args...)
	if err != nil {
		return nil, 0, err
	}

	searchDDM, err := toolkit.ToM(searchDD)
	if err != nil {
		return nil, 0, err
	}

	filterTableName := searchDDM.GetString("TableName")
	filterColumnName := searchDDM.GetString("ColumnName")

	///////// FILTER
	q = `SELECT * FROM (
		` + q + `
	) WHERE (
		upper(table_name) LIKE upper('%` + filterTableName + `%')
		AND upper(column_name) LIKE upper('%` + filterColumnName + `%')
	) `

	colFilterM, err := toolkit.ToM(colFilter)
	cf := make([]string, 0)
	if err != nil {
		cf = append(cf, "", "", "", "", "", "", "", "", "")
	} else {
		cf = append(cf,
			colFilterM.GetString("LIST_OF_CDE"),
			colFilterM.GetString("IMM_PREC_SYSTEM_NAME"),
			colFilterM.GetString("IMM_PREC_SYSTEM_SLA"),
			colFilterM.GetString("IMM_PREC_SYSTEM_OLA"),
			colFilterM.GetString("IMM_SUCC_SYSTEM_NAME"),
			colFilterM.GetString("IMM_SUCC_SYSTEM_SLA"),
			colFilterM.GetString("IMM_SUCC_SYSTEM_OLA"),
			colFilterM.GetString("LIST_DOWNSTREAM_PROCESS"),
			colFilterM.GetString("DOWNSTREAM_PROCESS_OWNER"),
		)
	}

	if cf[0] != "" || cf[1] != "" || cf[2] != "" || cf[3] != "" || cf[4] != "" || cf[5] != "" || cf[6] != "" || cf[7] != "" || cf[8] != "" {
		q += `AND (
			upper(list_of_cde) LIKE upper('%` + cf[0] + `%')
			AND upper(imm_prec_system_name) LIKE upper('%` + cf[1] + `%')
			AND upper(Imm_Prec_System_SLA) LIKE upper('%` + cf[2] + `%')
			AND upper(Imm_Prec_System_OLA) LIKE upper('%` + cf[3] + `%')
			AND upper(imm_succ_system_name) LIKE upper('%` + cf[4] + `%')
			AND upper(Imm_Succ_System_SLA) LIKE upper('%` + cf[5] + `%')
			AND upper(Imm_Succ_System_OLA) LIKE upper('%` + cf[6] + `%')
			AND upper(list_downstream_process) LIKE upper('%` + cf[7] + `%')
			AND upper(downstream_process_owner) LIKE upper('%` + cf[8] + `%')
		) `
	}

	///////// COUNT
	q = `SELECT res.*, 
			COUNT(DISTINCT list_of_cde) OVER () COUNT_list_of_cde,
			COUNT(DISTINCT imm_prec_system_name) OVER () COUNT_imm_prec_system_name,
			COUNT(DISTINCT Imm_Prec_System_SLA) OVER () COUNT_Imm_Prec_System_SLA,
			COUNT(DISTINCT Imm_Prec_System_OLA) OVER () COUNT_Imm_Prec_System_OLA,
			COUNT(DISTINCT imm_succ_system_name) OVER () COUNT_imm_succ_system_name,
			COUNT(DISTINCT Imm_Succ_System_SLA) OVER () COUNT_Imm_Succ_System_SLA,
			COUNT(DISTINCT Imm_Succ_System_OLA) OVER () COUNT_Imm_Succ_System_OLA,
			COUNT(DISTINCT list_downstream_process) OVER () COUNT_list_downstream_process,
			COUNT(DISTINCT downstream_process_owner) OVER () COUNT_downstream_process_owner
		FROM ( ` + q + `) res `

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName:   m.NewSystemModel().TableName(),
		SqlQuery:    q,
		Results:     &resultRows,
		PageNumber:  pageNumber,
		RowsPerPage: rowsPerPage,
		GroupCol:    "list_of_cde",
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) GetDetails(payload toolkit.M) (interface{}, int, error) {
	resultRows := make([]toolkit.M, 0)
	resultTotal := 0

	tabs := ""

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

	if strings.Contains(payload.GetString("Which"), "interfaces") == true {
		tabs = "dscinterfaces"
	} else if strings.Contains(payload.GetString("Which"), "my") == true {
		tabs = "dscmy"
	} else {
		tabs = "dscall"
	}

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "details", args...)
	if err != nil {
		return nil, 0, err
	}

	otherArgs := make([]string, 0)
	if payload.GetString("TableName") != "" && payload.GetString("ColumnName") != "" {
		otherArgs = append(otherArgs, "", "")
	} else {
		otherArgs = append(otherArgs, payload.GetString("Right"), payload.GetString("Column"))
	}

	otherArgs = append(otherArgs, payload.GetString("TableName"), payload.GetString("ColumnName"), payload.GetString("ScreenLabel"))

	///////// FILTER
	q = `SELECT * FROM (
		` + q + `
	) 
	WHERE (
			tmtid = '` + otherArgs[0] + `'
			AND tmcid = '` + otherArgs[1] + `'
		) OR (
			table_name = '` + otherArgs[2] + `'
			AND column_name = '` + otherArgs[3] + `' `
	if otherArgs[4] != "" {
		q += `AND business_alias_name = '` + otherArgs[4] + `' `
	}
	q += `) `

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
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

	tabs := ""

	q := ""
	args := make([]interface{}, 0)

	args = append(args, toolkit.ToString(payload.GetInt("Left")))

	if strings.Contains(payload.GetString("Which"), "interfaces") == true {
		tabs = "dscinterfaces"
	} else if strings.Contains(payload.GetString("Which"), "my") == true {
		tabs = "dscmy"
	} else {
		tabs = "dscall"
	}

	filePath := filepath.Join(clit.ExeDir(), "queryfiles", tabs+".sql")
	q, err := h.BuildQueryFromFile(filePath, "details", args...)
	if err != nil {
		return nil, 0, err
	}

	///////// FILTER
	q = `SELECT DISTINCT table_name, column_name, business_alias_name 
		FROM (
		` + q + `
	) `

	err = h.NewDBcmd().ExecuteSQLQuery(h.SqlQueryParam{
		TableName: m.NewCategoryModel().TableName(),
		SqlQuery:  q,
		Results:   &resultRows,
	})

	if err != nil {
		return nil, 0, err
	}

	return resultRows, resultTotal, nil
}

func (s *DSCService) getSystemRightTableFROMandWHERE(systemID int, searchDDM, payload toolkit.M) string {
	q := `FROM tbl_system ts
			LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = ts.id and tlrp.Object_type = 'SYSTEM'
			LEFT JOIN Tbl_Role rl_sys ON tlrp.role_id = rl_sys.id and rl_sys.role_name = 'Dataset Custodian'
			LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id

			inner join tbl_md_resource tmr ON ts.id = tmr.system_id
			inner join tbl_md_table tmt ON tmr.id = tmt.resource_id
			inner join tbl_md_column tmc ON tmt.id = tmc.table_id
			
			LEFT JOIN tbl_business_term tbt ON tmc.business_term_id = tbt.id
			
			LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
			LEFT JOIN Tbl_Link_Role_People tlrp_sdo ON tlrp_sdo.Object_ID = tsc.id AND tlrp_sdo.object_type = 'SUBCATEGORY'
			LEFT JOIN Tbl_Role rl ON tlrp_sdo.role_id = rl.id and rl.role_name = 'Data Domain Owner'
			LEFT JOIN Tbl_People ppl ON tlrp_sdo.people_id = ppl.id
			
			LEFT JOIN tbl_policy tpol ON tbt.policy_id = tpol.id
			
			left join tbl_link_column_interface ci on tmc.id = ci.column_id
			left join tbl_system ips on ci.imm_prec_system_id = ips.id
			left join tbl_system iss on ci.imm_succ_system_id = iss.id
			LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
			LEFT JOIN tbl_link_category_people tlcp ON tlcp.category_id = tc.id
			inner join
			(
				SELECT
				DISTINCT ts.id as sys_id, tmr.id as res_id, tmt.id as tab_id
				FROM tbl_system ts
					inner join tbl_md_resource tmr ON ts.id = tmr.system_id
					inner join tbl_md_table tmt ON tmr.id = tmt.resource_id
					inner join tbl_md_column tmc ON tmt.id = tmc.table_id
				WHERE 
					ts.id = ` + toolkit.ToString(systemID) + ` `

	if searchDDM != nil {
		if searchDDM.GetString("TableName") != "" {
			q += `AND upper(tmt.name) LIKE upper('%` + searchDDM.GetString("TableName") + `%') `
		}

		if searchDDM.GetString("ColumnName") != "" {
			q += `AND upper(tmc.name) LIKE upper('%` + searchDDM.GetString("ColumnName") + `%') `
		}
	}

	q += `		AND CDE = 1
			) cde ON ts.id = cde.sys_id and tmr.id = cde.res_id and tmt.id = cde.tab_id `

	if payload != nil {
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

func (s *DSCService) getInterfacesRightTableFROMandWHERE(systemID int, searchDDM, payload toolkit.M) string {
	q := `FROM tbl_system ts
			inner join tbl_md_resource res ON ts.id = res.system_id
			inner join tbl_md_table tmt ON res.id = tmt.resource_id
			inner join tbl_md_column tmc ON tmt.id = tmc.table_id
			
			left join tbl_link_column_interface ci on tmc.id = ci.column_id
			left join tbl_system ips on ci.imm_prec_system_id = ips.id
			left join tbl_system iss on ci.imm_succ_system_id = iss.id
			LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.column_id = tmc.id
			LEFT JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id 
			
			LEFT JOIN Tbl_Link_Role_People tlrp_sdo ON tlrp_sdo.Object_ID = tdp.id and tlrp_sdo.object_type = 'PROCESSES'
			LEFT JOIN Tbl_Role rl ON tlrp_sdo.role_id = rl.id and rl.role_name = 'Downstream Process Owner'
			LEFT JOIN Tbl_People ppl ON tlrp_sdo.people_id = ppl.id `

	if payload != nil {
		q += `
			LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = ts.id and tlrp.Object_type = 'SYSTEM'
			LEFT JOIN Tbl_Role rl_sys ON tlrp.role_id = rl_sys.id and rl_sys.role_name = 'Dataset Custodian'
			LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id 
			
			LEFT JOIN tbl_business_term tbt ON tmc.business_term_id = tbt.id
			LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
			LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
			LEFT JOIN tbl_policy tpol ON tbt.policy_id = tpol.id
			
			LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id `
	}

	q += `WHERE ts.id = ` + toolkit.ToString(systemID) + `
			AND (ips.system_name is not null or iss.system_name is not null) AND tmc.cde = 1 `

	if searchDDM != nil {
		if searchDDM.GetString("TableName") != "" {
			q += `AND upper(tmt.name) LIKE upper('%` + searchDDM.GetString("SystemName") + `%') `
		}

		if searchDDM.GetString("ColumnName") != "" {
			q += `AND upper(tmc.name) LIKE upper('%` + searchDDM.GetString("ColumnName") + `%') `
		}
	}

	if payload != nil {
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

	return q
}
