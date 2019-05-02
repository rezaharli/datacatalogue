-- name: dsc-view
SELECT DISTINCT
		TS.ID,
		TS.SYSTEM_NAME 				AS SYSTEM_NAME,
		TS.ITAM_ID					AS ITAM_ID,
		TP.FIRST_NAME||' '||TP.LAST_NAME 			AS DATASET_CUSTODIAN,
		TP.BANK_ID					AS BANK_ID,
		COUNT(DISTINCT TS.SYSTEM_NAME) OVER ()		AS COUNT_SYSTEM_NAME,
		COUNT(DISTINCT TS.ITAM_ID) OVER ()			AS COUNT_ITAM_ID,
		COUNT(DISTINCT TP.FIRST_NAME||' '||TP.LAST_NAME) OVER ()	AS COUNT_DATASET_CUSTODIAN,
		COUNT(DISTINCT TP.BANK_ID) OVER ()			AS COUNT_BANK_ID
	FROM 
		TBL_SYSTEM TS 
		INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TS.ID AND TLRP.OBJECT_TYPE = 'SYSTEM'
		INNER JOIN TBL_ROLE RL_SYS ON TLRP.ROLE_ID = RL_SYS.ID AND UPPER(RL_SYS.ROLE_NAME) = 'DATASET CUSTODIAN'
		INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID 
	WHERE ('ALL' = '?' OR TP.BANK_ID = '?')
		AND (
			upper(TS.SYSTEM_NAME) LIKE upper('%?%')
			AND upper(TS.ITAM_ID) LIKE upper('%?%')
			AND upper(TP.FIRST_NAME||' '||TP.LAST_NAME) LIKE upper('%?%')
			AND upper(TP.BANK_ID) LIKE upper('%?%')
		)
	ORDER BY TS.SYSTEM_NAME

-- name: dsc-view-homepage
SELECT DISTINCT 
		CDE_COUNT, DSP_COUNT, INTERFACE_COUNT
	FROM (
		SELECT 
				COUNT(DISTINCT TMCD.ALIAS_NAME) 		AS CDE_COUNT
			FROM 
				TBL_SYSTEM TS
				INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
				INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
				INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
				INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID
			WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?') AND CDE = 1
	) CDE , (
		SELECT 
				COUNT(DISTINCT DSP.NAME) 		AS DSP_COUNT
			FROM 
				TBL_SYSTEM TS
				INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
				INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
				INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
				INNER JOIN TBL_DS_PROCESS_DETAIL DSPD ON DSPD.COLUMN_ID = TMC.ID
				INNER JOIN TBL_DS_PROCESSES DSP ON DSPD.PROCESS_ID = DSP.ID
			WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?')
	) DSP, (
		SELECT 
				COUNT(DISTINCT SS.SYSTEM_NAME) 		AS INTERFACE_COUNT
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
SELECT DISTINCT 
		TMC.ID,
		TMT.ID					AS TMTID,
		TS.ID					AS TSID,
		TMC.ID					AS COLID,
		TS.SYSTEM_NAME				AS SYSTEM_NAME,
		TMCD.ALIAS_NAME                        			AS CDE,
		TMCD.DESCRIPTION                       			AS DESCRIPTION,
		TMT.DISPLAY_NAME                              			AS TABLE_NAME,
		TMC.DISPLAY_NAME                              			AS COLUMN_NAME,
		TDP.NAME                              				AS DSP_NAME,
		TP.FIRST_NAME||' '|| TP.LAST_NAME     			AS PROCESS_OWNER,
		TP.BANK_ID                            				AS BANK_ID,
		COUNT(DISTINCT TMCD.ALIAS_NAME) OVER () 		AS CDE_COUNT,
		COUNT(DISTINCT TMCD.ALIAS_NAME) OVER ()		AS COUNT_CDE,
		COUNT(DISTINCT TMCD.DESCRIPTION) OVER ()		AS COUNT_DESCRIPTION,
		COUNT(DISTINCT TMT.DISPLAY_NAME) OVER ()		AS COUNT_TABLE_NAME,
		COUNT(DISTINCT TMC.DISPLAY_NAME) OVER ()		AS COUNT_COLUMN_NAME,
		COUNT(DISTINCT TDP.NAME) OVER ()			AS COUNT_DSP_NAME,
		COUNT(DISTINCT TP.FIRST_NAME||' '|| TP.LAST_NAME) OVER ()	AS COUNT_PROCESS_OWNER
	FROM 
		TBL_SYSTEM TS
		INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
		INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
		INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
		INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID
		LEFT OUTER JOIN TBL_DS_PROCESS_DETAIL TDPD ON TDPD.COLUMN_ID = TMC.ID
		LEFT OUTER JOIN TBL_DS_PROCESSES TDP ON  TDP.ID = TDPD.PROCESS_ID
		LEFT OUTER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TDP.ID AND TLRP.OBJECT_TYPE = 'PROCESSES'
		LEFT OUTER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
		LEFT OUTER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
	WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?') AND CDE = 1
		AND (
			upper(NVL(TMCD.ALIAS_NAME,' ')) LIKE upper('%?%')
			AND upper(NVL(TMCD.DESCRIPTION,' ')) LIKE upper('%?%')
			AND upper(NVL(TMT.DISPLAY_NAME,' ')) LIKE upper('%?%')
			AND upper(NVL(TMC.DISPLAY_NAME,' ')) LIKE upper('%?%')
			AND upper(NVL(TDP.NAME,' ')) LIKE upper('%?%')
			AND upper(NVL(TP.FIRST_NAME||' '|| TP.LAST_NAME,' ')) LIKE upper('%?%')
		)

-- name: dsc-view-cdp
SELECT  DISTINCT
		TDP.ID,
		TS.SYSTEM_NAME					AS SYSTEM_NAME,
		TDP.NAME                              					AS DSP_NAME,
		TP.FIRST_NAME||' '|| TP.LAST_NAME     				AS PROCESS_OWNER,
		TP.BANK_ID                            					AS BANK_ID,
		COUNT(DISTINCT TMCD.ALIAS_NAME) OVER (PARTITION BY TDP.NAME)	AS CDE_COUNT,
		COUNT(TMCD.ALIAS_NAME) OVER () 				AS TOTAL,
		COUNT(DISTINCT TDP.NAME) OVER ()	  			AS COUNT_DSP_NAME
	FROM 
		TBL_SYSTEM TS
		INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
		INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
		INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
		INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID
		INNER JOIN TBL_DS_PROCESS_DETAIL TDPD ON TDPD.COLUMN_ID = TMC.ID
		INNER JOIN TBL_DS_PROCESSES TDP ON  TDP.ID = TDPD.PROCESS_ID
		INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TDP.ID AND TLRP.OBJECT_TYPE = 'PROCESSES'
		INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
		INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
	WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?') AND CDE = 1
		AND (
			upper(NVL(TDP.NAME, ' ')) LIKE upper('%?%')
			AND upper(NVL(TP.FIRST_NAME||' '|| TP.LAST_NAME, ' ')) LIKE upper('%?%')
			--AND CDE_COUNT LIKE upper('%?%')
		)
	GROUP BY TS.SYSTEM_NAME, TDP.ID, TDP.NAME, TP.FIRST_NAME||' '|| TP.LAST_NAME, TP.BANK_ID, TMCD.ALIAS_NAME

-- name: dsc-view-cdp-cde
SELECT  DISTINCT 
		TMC.ID,
		TMT.ID					AS TMTID,
		TS.ID					AS TSID,
		TMC.ID					AS COLID,
		TS.SYSTEM_NAME				AS SYSTEM_NAME,
		TDP.NAME                              				AS DSP_NAME,
		TP.FIRST_NAME||' '|| TP.LAST_NAME     			AS PROCESS_OWNER,
		TP.BANK_ID                            				AS BANK_ID,
		COUNT(DISTINCT TMCD.ALIAS_NAME) OVER ()     		AS CDE_COUNT,
		TMCD.ALIAS_NAME                        			AS CDE,
		TMCD.DESCRIPTION                       			AS DESCRIPTION,
		TMT.DISPLAY_NAME                              			AS TABLE_NAME,
		TMC.DISPLAY_NAME                              			AS COLUMN_NAME,
		COUNT(DISTINCT NVL(TMCD.ALIAS_NAME, ' ')) OVER ()		AS COUNT_CDE,
		COUNT(DISTINCT TMCD.DESCRIPTION) OVER ()		AS COUNT_DESCRIPTION,
		COUNT(DISTINCT TMT.DISPLAY_NAME) OVER ()		AS COUNT_TABLE_NAME,
		COUNT(DISTINCT TMC.DISPLAY_NAME) OVER ()		AS COUNT_COLUMN_NAME
	FROM 
		TBL_SYSTEM TS
		INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
		INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
		INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
		INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID
		INNER JOIN TBL_DS_PROCESS_DETAIL TDPD ON TDPD.COLUMN_ID = TMC.ID
		INNER JOIN TBL_DS_PROCESSES TDP ON  TDP.ID = TDPD.PROCESS_ID
		INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TDP.ID AND TLRP.OBJECT_TYPE = 'PROCESSES'
		INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
		INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
	WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?') AND CDE = 1 AND UPPER(TDP.NAME) = UPPER('?')
		AND (
			upper(NVL(TMCD.ALIAS_NAME,' ')) LIKE upper('%?%')
			AND upper(NVL(TMCD.DESCRIPTION,' ')) LIKE upper('%?%')
			AND upper(NVL(TMT.DISPLAY_NAME,' ')) LIKE upper('%?%')
			AND upper(NVL(TMC.DISPLAY_NAME,' ')) LIKE upper('%?%')
		)

-- name: dsc-view-interfaces
SELECT ID, SYSTEM_NAME, IMM_INTERFACE, CDE_COUNT, COUNT_IMM_INTERFACE, PROCESS_OWNER
FROM (
SELECT  DISTINCT
		SS.ID,
		UPPER(TS.SYSTEM_NAME)                     		AS SYSTEM_NAME,
		COUNT(DISTINCT SS.SYSTEM_NAME) OVER()     	AS INTERFACE_COUNT,
		SS.SYSTEM_NAME                            		AS IMM_INTERFACE,
		COUNT(DISTINCT TMCD.ALIAS_NAME) OVER(PARTITION BY SS.SYSTEM_NAME)      	   	AS CDE_COUNT,
		COUNT(DISTINCT SS.SYSTEM_NAME) OVER ()	AS COUNT_IMM_INTERFACE,
    TP.FIRST_NAME||' '|| TP.LAST_NAME     			AS PROCESS_OWNER
	FROM 
		TBL_SYSTEM TS
		INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
		INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
		INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
		INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID
		INNER JOIN TBL_LINK_COLUMN_INTERFACE CI ON CI.COLUMN_ID = TMC.ID
		INNER JOIN TBL_SYSTEM SS ON CI.IMM_SUCC_SYSTEM_ID = SS.ID
		LEFT JOIN TBL_DS_PROCESS_DETAIL DPD ON DPD.COLUMN_ID = TMC.ID
		LEFT JOIN TBL_DS_PROCESSES TDP ON DPD.PROCESS_ID = TDP.ID
		LEFT JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TDP.ID AND TLRP.OBJECT_TYPE = 'PROCESSES'
		LEFT JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
		LEFT JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID    
	WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?')
		AND (
			upper(NVL(SS.SYSTEM_NAME,' ')) LIKE upper('%?%')
			AND upper(NVL(TP.FIRST_NAME||' '|| TP.LAST_NAME,' ')) LIKE upper('%?%')
		)
    ) RES WHERE CDE_COUNT <> 0
    ORDER BY IMM_INTERFACE, PROCESS_OWNER

-- name: dsc-view-interfaces-cde
SELECT  DISTINCT 
		SS.ID,
		TMT.ID					AS TMTID,
		TS.ID					AS TSID,
		TMC.ID					AS COLID,
		UPPER(TS.SYSTEM_NAME)				AS SYSTEM_NAME,
		SS.SYSTEM_NAME				AS IMM_INTERFACE,
		COUNT(DISTINCT TMCD.ALIAS_NAME) OVER ()		AS CDE_COUNT,
		TMCD.ALIAS_NAME                        			AS CDE,
		TMCD.DESCRIPTION                       			AS DESCRIPTION,
		TMT.DISPLAY_NAME                              			AS TABLE_NAME,
		TMC.DISPLAY_NAME                              			AS COLUMN_NAME,
		COUNT(DISTINCT TMCD.ALIAS_NAME) OVER ()		AS COUNT_CDE,
		COUNT(DISTINCT TMCD.DESCRIPTION) OVER ()		AS COUNT_DESCRIPTION,
		COUNT(DISTINCT TMT.DISPLAY_NAME) OVER ()		AS COUNT_TABLE_NAME,
		COUNT(DISTINCT TMC.DISPLAY_NAME) OVER ()		AS COUNT_COLUMN_NAME
	FROM 
		TBL_SYSTEM TS
		INNER JOIN TBL_MD_RESOURCE TMR ON TS.ID = TMR.SYSTEM_ID
		INNER JOIN TBL_MD_TABLE TMT ON TMR.ID = TMT.RESOURCE_ID
		INNER JOIN TBL_MD_COLUMN TMC ON TMT.ID = TMC.TABLE_ID
		INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID
		INNER JOIN TBL_LINK_COLUMN_INTERFACE CI ON CI.COLUMN_ID = TMC.ID
		INNER JOIN TBL_SYSTEM SS ON CI.IMM_SUCC_SYSTEM_ID = SS.ID
	WHERE UPPER(TS.SYSTEM_NAME) = UPPER('?') AND UPPER(SS.SYSTEM_NAME) = UPPER('?')
		AND (
			upper(NVL(TMCD.ALIAS_NAME,' ')) LIKE upper('%?%')
			AND upper(NVL(TMCD.DESCRIPTION,' ')) LIKE upper('%?%')
			AND upper(NVL(TMT.DISPLAY_NAME,' ')) LIKE upper('%?%')
			AND upper(NVL(TMC.DISPLAY_NAME,' ')) LIKE upper('%?%')
		)

-- name: details
SELECT DISTINCT
		ts.id,
		tmt.id					as tmtid,
		tmc.id					as tmcid,
		ts.system_name				as system_name,
		ts.itam_id					as itam_id,
		tp.first_name||' '||tp.last_name			as dataset_custodian,
		tp.bank_id					as bank_id,
		tmcd.alias_name				as business_alias_name,
		tmt.DISPLAY_NAME 				as table_name,
		tmc.DISPLAY_NAME				as column_name,
		tmcd.description 				as business_alias_description,
		tmcd.cde					as cde_yes_no,
		tmcd.status					as status,
		tmc.data_type					as data_type,
		tmc.data_format				as data_format,
		tmc.data_length				as data_length,
		tmcd.example					as example,
		tmcd.derived					as derived_yes_no,
		tmcd.Derivation_Logic				as derivation_logic,
		tmcd.Sourced_from_Upstream				as sourced_from_upstream_yes_no,
		tmcd.System_Checks				as system_checks,
		tc.Name 					as domain,
		tsc.name 					as subdomain,
		ppl.first_name||' '||ppl.last_name 			as domain_owner,
		tbt.bt_name					as business_term,
		tbt.description 					as business_term_description,
		tpol.info_asset_name				as information_asset_names,
		tpol.description 				as information_asset_description,
		tpol.confidentiality				as confidentiality,
		tpol.integrity					as integrity,
		tpol.availability				as availability,
		tpol.overall_cia_rating				as overall_cia_rating,
		tmt.record_category				as record_categories,
		tmcd.pii_flag					as pii_flag,
		ips.system_name 				as imm_preceeding_system,
		ci.INCOMING_CDE_NAME				as imm_prec_incoming,
		ci.INCOMING_DERIVED				as imm_prec_derived,
		ci.INCOMING_DERIVATION_LOGIC			as imm_prec_derivation_logic,
		iss.system_name 				as imm_succeeding_system,
		ci.OUTGOING_CDE_NAME				as imm_succ_incoming,
		ci.OUTGOING_DERIVED				as imm_succ_derived,
		ci.OUTGOING_DERIVATION_LOGIC			as imm_succ_derivation_logic,
		tmcd.DQ_STANDARDS||' '||tmcd.threshold		as threshold
	FROM tbl_system ts
		LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = ts.id and tlrp.Object_type = 'SYSTEM'
		LEFT JOIN Tbl_Role rl_sys ON tlrp.role_id = rl_sys.id and rl_sys.role_name = 'Dataset Custodian'
		LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id
		inner join tbl_md_resource tmr ON ts.id = tmr.system_id
		inner join tbl_md_table tmt ON tmr.id = tmt.resource_id
		inner join tbl_md_column tmc ON tmt.id = tmc.table_id
    		INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID
		LEFT JOIN tbl_link_column_business_term LCBT ON TMC.ID = LCBT.COLUMN_ID 
		LEFT JOIN tbl_business_term tbt ON LCBT.business_term_id = tbt.id
		LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
		LEFT JOIN Tbl_Link_Role_People tlrp_sdo ON tlrp_sdo.Object_ID = tsc.id AND tlrp_sdo.object_type = 'SUBCATEGORY'
		LEFT JOIN Tbl_Role rl ON tlrp_sdo.role_id = rl.id and rl.role_name = 'Data Domain Owner'
		LEFT JOIN Tbl_People ppl ON tlrp_sdo.people_id = ppl.id
		LEFT JOIN tbl_policy tpol ON tbt.policy_id = tpol.id
		left join tbl_link_column_interface ci on tmc.id = ci.column_id
		left join tbl_system ips on ci.imm_prec_system_id = ips.id
		left join tbl_system iss on ci.imm_succ_system_id = iss.id
		LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
		inner join
		(
			SELECT
			DISTINCT ts.id as sys_id, tmr.id as res_id, tmt.id as tab_id
			FROM tbl_system ts
				inner join tbl_md_resource tmr ON ts.id = tmr.system_id
				inner join tbl_md_table tmt ON tmr.id = tmt.resource_id
				inner join tbl_md_column tmc ON tmt.id = tmc.table_id
        				INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID
			WHERE 
				ts.id = '?'
				AND CDE = 1
		) cde ON ts.id = cde.sys_id and tmr.id = cde.res_id and tmt.id = cde.tab_id
	ORDER BY tmt.DISPLAY_NAME, tmc.DISPLAY_NAME


-- name: dsc-view-dd
SELECT DISTINCT
		ts.id,
		tmt.id				as tmtid,
		tmc.id				as tmcid,
		ts.system_name			as system_name,
		ts.itam_id				as itam_id,
		tp.first_name||' '||tp.last_name		as dataset_custodian,
		tp.bank_id				as bank_id,
		tmcd.alias_name			as business_alias_name,
		tmt.DISPLAY_NAME			as table_name,
		tmc.DISPLAY_NAME			as column_name,
		tmcd.description			as business_alias_description,
		tmcd.cde				as cde_yes_no,
		tmc.data_type				as data_type,
		tmc.data_length			as data_length,
		tmcd.example				as example,
		tmcd.derived				as derived_yes_no,
		tmcd.Derivation_Logic			as derivation_logic,
		case when tmcd.Sourced_from_Upstream = 1 then 'Yes' else 'No' End	as sourced_from_upstream_yes_no,
		tmcd.System_Checks			as system_checks,
		tc.Name				as domain,
		tsc.name				as subdomain,
		ppl.first_name||' '||ppl.last_name		as domain_owner,
		tbt.bt_name				as business_term,
		tbt.description 				as business_term_description,
		tpol.info_asset_name			as information_asset_names,
		tpol.description 			as information_asset_description,
		tpol.confidentiality			as confidentiality,
		tpol.integrity				as integrity,
		tpol.availability			as availability,
		tpol.overall_cia_rating			as overall_cia_rating,
		tmt.record_category			as record_categories,
		tmcd.pii_flag				as pii_flag
	FROM tbl_system ts
		LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = ts.id and tlrp.Object_type = 'SYSTEM'
		LEFT JOIN Tbl_Role rl_sys ON tlrp.role_id = rl_sys.id and rl_sys.role_name = 'Dataset Custodian'
		LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id
		inner join tbl_md_resource tmr ON ts.id = tmr.system_id
		inner join tbl_md_table tmt ON tmr.id = tmt.resource_id
		inner join tbl_md_column tmc ON tmt.id = tmc.table_id
		INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID		
		LEFT JOIN tbl_link_column_business_term LCBT ON TMC.ID = LCBT.COLUMN_ID 
		LEFT JOIN tbl_business_term tbt ON LCBT.business_term_id = tbt.id
		LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
		LEFT JOIN Tbl_Link_Role_People tlrp_sdo ON tlrp_sdo.Object_ID = tsc.id AND tlrp_sdo.object_type = 'SUBCATEGORY'
		LEFT JOIN Tbl_Role rl ON tlrp_sdo.role_id = rl.id and rl.role_name = 'Data Domain Owner'
		LEFT JOIN Tbl_People ppl ON tlrp_sdo.people_id = ppl.id
		LEFT JOIN tbl_policy tpol ON tbt.policy_id = tpol.id
		left join tbl_link_column_interface ci on tmc.id = ci.column_id
		left join tbl_system ips on ci.imm_prec_system_id = ips.id
		left join tbl_system iss on ci.imm_succ_system_id = iss.id
		LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
		inner join
		(
			SELECT
			DISTINCT ts.id as sys_id, tmr.id as res_id, tmt.id as tab_id
			FROM tbl_system ts
				inner join tbl_md_resource tmr ON ts.id = tmr.system_id
				inner join tbl_md_table tmt ON tmr.id = tmt.resource_id
				inner join tbl_md_column tmc ON tmt.id = tmc.table_id
        INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID
		WHERE 
				UPPER(ts.system_name) = UPPER('?')
				AND CDE = 1
		) cde ON ts.id = cde.sys_id and tmr.id = cde.res_id and tmt.id = cde.tab_id
	WHERE
		upper(NVL(ts.system_name, ' ')) LIKE upper('%?%')
		AND upper(NVL(ts.itam_id, ' ')) LIKE upper('%?%')
		AND upper(NVL(tmt.name, ' ')) LIKE upper('%?%')
		AND  upper(NVL(tmc.name, ' ')) LIKE upper('%?%')
		AND  upper(NVL(tmcd.alias_name, ' ')) LIKE upper('%?%')
		AND  upper(NVL(tmcd.description, ' ')) LIKE upper('%?%')
		AND  upper(tmcd.cde) LIKE upper('%?%')
		AND  upper(NVL(tmc.data_type, ' ')) LIKE upper('%?%')
		AND  upper(tmc.data_length) LIKE upper('%?%')
		AND  upper(NVL(tmcd.example, ' ')) LIKE upper('%?%')
		AND  upper(tmcd.derived) LIKE upper('%?%')
		AND  upper(NVL(tmcd.Derivation_Logic, ' ')) LIKE upper('%?%')
		AND  upper(tmcd.Sourced_from_Upstream) LIKE upper('%?%')
		AND  upper(NVL(tmcd.System_Checks, ' ')) LIKE upper('%?%')
		AND  upper(NVL(tc.Name, ' ')) LIKE upper('%?%')
		AND  upper(NVL(tsc.name, ' ')) LIKE upper('%?%')
		AND  upper(NVL(ppl.first_name||' '||ppl.last_name, ' ')) LIKE upper('%?%')
		AND  upper(NVL(tbt.bt_name, ' ')) LIKE upper('%?%')
		AND  upper(NVL(tbt.description, ' ')) LIKE upper('%?%')
		AND  upper(NVL(tpol.info_asset_name, ' ')) LIKE upper('%?%')
		AND  upper(NVL(tpol.description, ' ')) LIKE upper('%?%')
		AND  NVL(to_char(tpol.confidentiality), ' ') LIKE upper('%?%')
		AND  NVL(to_char(tpol.integrity), ' ') LIKE upper('%?%')
		AND  NVL(to_char(tpol.availability), ' ') LIKE upper('%?%')
		AND  NVL(to_char(tpol.overall_cia_rating), ' ') LIKE upper('%?%')
		AND  upper(NVL(tmt.record_category, ' ')) LIKE upper('%?%')
		AND  upper(tmcd.pii_flag) LIKE upper('%?%')
	ORDER BY tmt.DISPLAY_NAME, tmc.DISPLAY_NAME