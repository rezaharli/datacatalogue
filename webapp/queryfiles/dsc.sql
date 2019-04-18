-- name: dsc-view
SELECT * 
	FROM (
		SELECT * 
			FROM (
				SELECT res.*, 
						COUNT(DISTINCT system_name) OVER ()			as COUNT_SYSTEM_NAME,
						COUNT(DISTINCT itam_id) OVER ()				as COUNT_ITAM_ID,
						COUNT(DISTINCT dataset_custodian) OVER ()	as COUNT_DATASET_CUSTODIAN,
						COUNT(DISTINCT bank_id) OVER ()				as COUNT_BANK_ID
					FROM (
                        SELECT DISTINCT
                            TS.ID
                            ,TS.SYSTEM_NAME 						AS SYSTEM_NAME
							,TS.ITAM_ID								AS ITAM_ID
							,TP.FIRST_NAME||' '||TP.LAST_NAME 		AS DATASET_CUSTODIAN
							,TP.BANK_ID								AS BANK_ID
                        FROM 
                            TBL_SYSTEM TS 
							INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TS.ID AND TLRP.OBJECT_TYPE = 'SYSTEM'
							INNER JOIN TBL_ROLE RL_SYS ON TLRP.ROLE_ID = RL_SYS.ID AND UPPER(RL_SYS.ROLE_NAME) = 'DATASET CUSTODIAN'
							INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID 
						WHERE ('ALL' = '?' OR TP.BANK_ID = '?')
                        ORDER BY TS.SYSTEM_NAME
					) res
			) WHERE ( -- Main and dropdown search
				upper(system_name) LIKE upper('%?%')
				AND upper(itam_id) LIKE upper('%?%')
			)
	) WHERE ( -- Column filter
		upper(system_name) LIKE upper('%?%')
		AND upper(itam_id) LIKE upper('%?%')
		AND upper(dataset_custodian) LIKE upper('%?%')
		AND upper(bank_id) LIKE upper('%?%')
	)

-- name: dsc-view-homepage
SELECT DISTINCT CDE_COUNT, DSP_COUNT, INTERFACE_COUNT
FROM 
(
	SELECT 
		COUNT(DISTINCT TMC.ALIAS_NAME) 			AS CDE_COUNT
	FROM 
		TBL_SYSTEM TS
		INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
		INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
		INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
	WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?') AND CDE = 1
) CDE , 
(
	SELECT 
		COUNT(DISTINCT DSP.NAME) 				AS DSP_COUNT
	FROM 
		TBL_SYSTEM TS
		INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
		INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
		INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
		INNER JOIN TBL_DS_PROCESS_DETAIL DSPD ON DSPD.COLUMN_ID = TMC.ID
		INNER JOIN TBL_DS_PROCESSES DSP ON DSPD.PROCESS_ID = DSP.ID
	WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?')
) DSP , 
(
	SELECT 
		COUNT(DISTINCT SS.SYSTEM_NAME) 			AS INTERFACE_COUNT
	FROM 
		TBL_SYSTEM TS
		INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
		INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
		INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
		INNER JOIN TBL_LINK_COLUMN_INTERFACE CI ON CI.COLUMN_ID = TMC.ID
    INNER JOIN TBL_SYSTEM SS ON CI.IMM_SUCC_SYSTEM_ID = SS.ID
	WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?')
) INTERFACE

-- name: dsc-view-cde
SELECT * 
	FROM (
		SELECT res.*, 
				COUNT(DISTINCT CDE) OVER ()				as COUNT_CDE,
				COUNT(DISTINCT DESCRIPTION) OVER ()		as COUNT_DESCRIPTION,
				COUNT(DISTINCT TABLE_NAME) OVER ()		as COUNT_TABLE_NAME,
				COUNT(DISTINCT COLUMN_NAME) OVER ()		as COUNT_COLUMN_NAME,
				COUNT(DISTINCT DSP_NAME) OVER ()		as COUNT_DSP_NAME,
				COUNT(DISTINCT PROCESS_OWNER) OVER ()	as COUNT_PROCESS_OWNER
			FROM (
				SELECT DISTINCT 
					TMT.ID,
					TMT.ID									AS TMTID,
					TS.ID									AS TSID,
					TMC.ID									AS COLID,
					TS.SYSTEM_NAME							AS SYSTEM_NAME,
					TMC.ALIAS_NAME                        	AS CDE,
					TMC.DESCRIPTION                       	AS DESCRIPTION,
					TMT.NAME                              	AS TABLE_NAME,
					TMC.NAME                              	AS COLUMN_NAME,
					TDP.NAME                              	AS DSP_NAME,
					TP.FIRST_NAME||' '|| TP.LAST_NAME     	AS PROCESS_OWNER,
					TP.BANK_ID                            	AS BANK_ID,
					COUNT(DISTINCT TMC.ALIAS_NAME) OVER () 	AS CDE_COUNT
				FROM 
					TBL_SYSTEM TS
					INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
					INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
					INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
					LEFT OUTER JOIN TBL_DS_PROCESS_DETAIL TDPD ON TDPD.COLUMN_ID = TMC.ID
					LEFT OUTER JOIN TBL_DS_PROCESSES TDP ON  TDP.ID = TDPD.PROCESS_ID
					LEFT OUTER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TDP.ID AND TLRP.OBJECT_TYPE = 'PROCESSES'
					LEFT OUTER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
					LEFT OUTER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
				WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?') AND CDE = 1
			) res
	) WHERE ( -- Column filter
		upper(NVL(CDE,' ')) LIKE upper('%?%')
		AND upper(NVL(DESCRIPTION,' ')) LIKE upper('%?%')
		AND upper(NVL(TABLE_NAME,' ')) LIKE upper('%?%')
		AND upper(NVL(COLUMN_NAME,' ')) LIKE upper('%?%')
		AND upper(NVL(DSP_NAME,' ')) LIKE upper('%?%')
		AND upper(NVL(PROCESS_OWNER,' ')) LIKE upper('%?%')
	)

-- name: dsc-view-cdp
SELECT * 
	FROM (
		SELECT res.*, 
				COUNT(DISTINCT DSP_NAME) OVER ()		as COUNT_DSP_NAME,
				COUNT(DISTINCT PROCESS_OWNER) OVER ()	as COUNT_PROCESS_OWNER,
				COUNT(DISTINCT CDE_COUNT) OVER ()		as COUNT_CDE_COUNT
			FROM (
				SELECT  
					TDP.ID,
					TS.SYSTEM_NAME						  AS SYSTEM_NAME,
					TDP.NAME                              AS DSP_NAME,
					TP.FIRST_NAME||' '|| TP.LAST_NAME     AS PROCESS_OWNER,
					TP.BANK_ID                            AS BANK_ID,
					COUNT(DISTINCT TMC.ALIAS_NAME)        AS CDE_COUNT,
					COUNT(DISTINCT TDP.NAME) OVER ()	  AS DSP_COUNT
				FROM 
					TBL_SYSTEM TS
					INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
					INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
					INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
					INNER JOIN TBL_DS_PROCESS_DETAIL TDPD ON TDPD.COLUMN_ID = TMC.ID
					INNER JOIN TBL_DS_PROCESSES TDP ON  TDP.ID = TDPD.PROCESS_ID
					INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TDP.ID AND TLRP.OBJECT_TYPE = 'PROCESSES'
					INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
					INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
				WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?') AND CDE = 1
				GROUP BY TS.SYSTEM_NAME, TDP.ID, TDP.NAME, TP.FIRST_NAME||' '|| TP.LAST_NAME, TP.BANK_ID
			) res
	) WHERE ( -- Column filter
		upper(NVL(DSP_NAME,' ')) LIKE upper('%?%')
		AND upper(NVL(PROCESS_OWNER,' ')) LIKE upper('%?%')
		AND CDE_COUNT LIKE upper('%?%')
	)

-- name: dsc-view-cdp-cde
SELECT * 
	FROM (
		SELECT res.*, 
				COUNT(DISTINCT CDE) OVER ()				as COUNT_CDE,
				COUNT(DISTINCT DESCRIPTION) OVER ()		as COUNT_DESCRIPTION,
				COUNT(DISTINCT TABLE_NAME) OVER ()		as COUNT_TABLE_NAME,
				COUNT(DISTINCT COLUMN_NAME) OVER ()		as COUNT_COLUMN_NAME
			FROM (
				SELECT  DISTINCT 
					TMC.ID,
					TMT.ID										AS TMTID,
					TS.ID										AS TSID,
					TMC.ID										AS COLID,
					TS.SYSTEM_NAME								AS SYSTEM_NAME,
					TDP.NAME                              		AS DSP_NAME,
					TP.FIRST_NAME||' '|| TP.LAST_NAME     		AS PROCESS_OWNER,
					TP.BANK_ID                            		AS BANK_ID,
					COUNT(DISTINCT TMC.ALIAS_NAME) OVER ()      AS CDE_COUNT,
					TMC.ALIAS_NAME                        		AS CDE,
					TMC.DESCRIPTION                       		AS DESCRIPTION,
					TMT.NAME                              		AS TABLE_NAME,
					TMC.NAME                              		AS COLUMN_NAME	
				FROM 
					TBL_SYSTEM TS
					INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
					INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
					INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
					INNER JOIN TBL_DS_PROCESS_DETAIL TDPD ON TDPD.COLUMN_ID = TMC.ID
					INNER JOIN TBL_DS_PROCESSES TDP ON  TDP.ID = TDPD.PROCESS_ID
					INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TDP.ID AND TLRP.OBJECT_TYPE = 'PROCESSES'
					INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
					INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
				WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?') AND CDE = 1 AND UPPER(TDP.NAME) = UPPER('?')
			) res
	) WHERE ( -- Column filter
		upper(NVL(CDE,' ')) LIKE upper('%?%')
		AND upper(NVL(DESCRIPTION,' ')) LIKE upper('%?%')
		AND upper(NVL(TABLE_NAME,' ')) LIKE upper('%?%')
		AND upper(NVL(COLUMN_NAME,' ')) LIKE upper('%?%')
	)

-- name: dsc-view-interfaces
SELECT * 
	FROM (
		SELECT res.*, 
				COUNT(DISTINCT IMM_INTERFACE) OVER ()		as COUNT_IMM_INTERFACE,
				COUNT(DISTINCT CDE_COUNT) OVER ()		as COUNT_CDE_COUNT
			FROM (
				SELECT  
					SS.ID,
					UPPER(TS.SYSTEM_NAME)                     AS SYSTEM_NAME,
					COUNT(DISTINCT SS.SYSTEM_NAME) OVER()     AS INTERFACE_COUNT,
					SS.SYSTEM_NAME                            AS IMM_INTERFACE,
					COUNT(DISTINCT TMC.ALIAS_NAME)      	   AS CDE_COUNT
				FROM 
					TBL_SYSTEM TS
					INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
					INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
					INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
					INNER JOIN TBL_LINK_COLUMN_INTERFACE CI ON CI.COLUMN_ID = TMC.ID
					INNER JOIN TBL_SYSTEM SS ON CI.IMM_SUCC_SYSTEM_ID = SS.ID
				WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?')
				GROUP BY UPPER(TS.SYSTEM_NAME), SS.ID, SS.SYSTEM_NAME
			) res
	) WHERE ( -- Column filter
		upper(NVL(IMM_INTERFACE,' ')) LIKE upper('%?%')
		AND CDE_COUNT LIKE upper('%?%')
	)

-- name: dsc-view-interfaces-cde
SELECT * 
	FROM (
		SELECT res.*, 
				COUNT(DISTINCT CDE) OVER ()				as COUNT_CDE,
				COUNT(DISTINCT DESCRIPTION) OVER ()		as COUNT_DESCRIPTION,
				COUNT(DISTINCT TABLE_NAME) OVER ()		as COUNT_TABLE_NAME,
				COUNT(DISTINCT COLUMN_NAME) OVER ()		as COUNT_COLUMN_NAME
			FROM (
				SELECT  DISTINCT 
					SS.ID,
					TMT.ID										AS TMTID,
					TS.ID										AS TSID,
					TMC.ID										AS COLID,
					UPPER(TS.SYSTEM_NAME)                      AS SYSTEM_NAME,
					SS.SYSTEM_NAME                             AS IMM_INTERFACE,
					COUNT(DISTINCT TMC.ALIAS_NAME) OVER ()      AS CDE_COUNT,
					TMC.ALIAS_NAME                        		AS CDE,
					TMC.DESCRIPTION                       		AS DESCRIPTION,
					TMT.NAME                              		AS TABLE_NAME,
					TMC.NAME                              		AS COLUMN_NAME
				FROM 
					TBL_SYSTEM TS
					INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
					INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
					INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
					INNER JOIN TBL_LINK_COLUMN_INTERFACE CI ON CI.COLUMN_ID = TMC.ID
					INNER JOIN TBL_SYSTEM SS ON CI.IMM_SUCC_SYSTEM_ID = SS.ID
				WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?') AND UPPER(SS.SYSTEM_NAME) = UPPER('?')
			) res
	) WHERE ( -- Column filter
		upper(NVL(CDE,' ')) LIKE upper('%?%')
		AND upper(NVL(DESCRIPTION,' ')) LIKE upper('%?%')
		AND upper(NVL(TABLE_NAME,' ')) LIKE upper('%?%')
		AND upper(NVL(COLUMN_NAME,' ')) LIKE upper('%?%')
	)

-- name: details
SELECT DISTINCT
		ts.id,
		tmt.id									as tmtid,
		tmc.id									as tmcid,
		ts.system_name							as system_name,
		ts.itam_id								as itam_id,
		tp.first_name||' '||tp.last_name		as dataset_custodian,
		tp.bank_id								as bank_id,
		tmc.alias_name							as business_alias_name,
		tmt.name 								as table_name,
		tmc.name 								as column_name,
		tmc.description 						as business_alias_description,
		tmc.cde									as cde_yes_no,
		tmc.status								as status,
		tmc.data_type							as data_type,
		tmc.data_format							as data_format,
		tmc.data_length							as data_length,
		tmc.example								as example,
		tmc.derived								as derived_yes_no,
		tmc.Derivation_Logic					as derivation_logic,
		tmc.Sourced_from_Upstream				as sourced_from_upstream_yes_no,
		tmc.System_Checks						as system_checks,
		tc.Name 								as domain,
		tsc.name 								as subdomain,
		ppl.first_name||' '||ppl.last_name 		as domain_owner,
		tbt.bt_name								as business_term,
		tbt.description 						as business_term_description,
		tpol.info_asset_name					as information_asset_names,
		tpol.description 						as information_asset_description,
		tpol.confidentiality					as confidentiality,
		tpol.integrity							as integrity,
		tpol.availability						as availability,
		tpol.overall_cia_rating					as overall_cia_rating,
		tmt.record_category						as record_categories,
		tmc.pii_flag							as pii_flag,
		ips.system_name 						as imm_preceeding_system,
		'' 										as imm_prec_incoming,
		'' 										as imm_prec_derived,
		'' 										as imm_prec_derivation_logic,
		iss.system_name 						as imm_succeeding_system,
		'' 										as imm_succ_incoming,
		'' 										as imm_succ_derived,
		'' 										as imm_succ_derivation_logic,
		tmc.DQ_STANDARDS||' '||tmc.threshold	as threshold
	FROM tbl_system ts
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
				ts.id = '?'
				AND CDE = 1
		) cde ON ts.id = cde.sys_id and tmr.id = cde.res_id and tmt.id = cde.tab_id
	ORDER BY tmt.name, tmc.name

-- name: dsc-view-dd
SELECT DISTINCT
		ts.id,
		tmt.id									as tmtid,
		tmc.id									as tmcid,
		ts.system_name							as system_name,
		ts.itam_id								as itam_id,
		tp.first_name||' '||tp.last_name		as dataset_custodian,
		tp.bank_id								as bank_id,
		tmc.alias_name							as business_alias_name,
		tmt.name 								as table_name,
		tmc.name 								as column_name,
		tmc.description 						as business_alias_description,
		tmc.cde									as cde_yes_no,
		tmc.status								as status,
		tmc.data_type							as data_type,
		tmc.data_format							as data_format,
		tmc.data_length							as data_length,
		tmc.example								as example,
		tmc.derived								as derived_yes_no,
		tmc.Derivation_Logic					as derivation_logic,
		tmc.Sourced_from_Upstream				as sourced_from_upstream_yes_no,
		tmc.System_Checks						as system_checks,
		tc.Name 								as domain,
		tsc.name 								as subdomain,
		ppl.first_name||' '||ppl.last_name 		as domain_owner,
		tbt.bt_name								as business_term,
		tbt.description 						as business_term_description,
		tpol.info_asset_name					as information_asset_names,
		tpol.description 						as information_asset_description,
		tpol.confidentiality					as confidentiality,
		tpol.integrity							as integrity,
		tpol.availability						as availability,
		tpol.overall_cia_rating					as overall_cia_rating,
		tmt.record_category						as record_categories,
		tmc.pii_flag							as pii_flag,
		ips.system_name 						as imm_preceeding_system,
		'' 										as imm_prec_incoming,
		'' 										as imm_prec_derived,
		'' 										as imm_prec_derivation_logic,
		iss.system_name 						as imm_succeeding_system,
		'' 										as imm_succ_incoming,
		'' 										as imm_succ_derived,
		'' 										as imm_succ_derivation_logic,
		tmc.DQ_STANDARDS||' '||tmc.threshold	as threshold
	FROM tbl_system ts
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
				UPPER(ts.system_name) = UPPER('?')
				AND CDE = 1
		) cde ON ts.id = cde.sys_id and tmr.id = cde.res_id and tmt.id = cde.tab_id
	ORDER BY tmt.name, tmc.name