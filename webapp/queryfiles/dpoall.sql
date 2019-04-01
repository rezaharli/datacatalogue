-- name: left-grid
SELECT * 
	FROM (
		SELECT * 
			FROM (
				SELECT res.*, 
							COUNT(DISTINCT DOWNSTREAM_PROCESS) OVER () 	as COUNT_DOWNSTREAM_PROCESS,
							COUNT(DISTINCT PROCESS_OWNER) OVER () 			as COUNT_PROCESS_OWNER,
							COUNT(DISTINCT BANK_ID) OVER () 						as COUNT_BANK_ID
						FROM (
							SELECT DISTINCT
									TDP.id,
									TDP.NAME                              AS downstream_process,
									TP.FIRST_NAME||' '|| TP.LAST_NAME     AS process_owner,
									TP.BANK_ID                            AS bank_id
								FROM TBL_DS_PROCESSES TDP
									INNER JOIN TBL_DS_PROCESS_DETAIL TDPD ON TDP.ID = TDPD.PROCESS_ID
									INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TDP.ID AND TLRP.OBJECT_TYPE = 'PROCESSES'
									INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
									INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
								WHERE TDP.NAME = 'USFRR'
								ORDER BY TDP.NAME
						) res
			) WHERE ( -- Main and dropdown search
				upper(downstream_process) LIKE upper('%?%')
				AND upper(process_owner) LIKE upper('%?%')
			) 
	) WHERE ( -- Column filter
			upper(DOWNSTREAM_PROCESS) LIKE upper('%?%')
			AND upper(PROCESS_OWNER) LIKE upper('%?%')
			AND upper(BANK_ID) LIKE upper('%?%')
		)

-- name: right-grid
SELECT * 
	FROM (
		SELECT * 
			FROM (
				SELECT res.*, 
						COUNT(DISTINCT CDE_NAME) OVER () COUNT_CDE_NAME,
						COUNT(DISTINCT SEGMENT) OVER () COUNT_SEGMENT,
						COUNT(DISTINCT IMM_PREC_SYSTEM) OVER () COUNT_IMM_PREC_SYSTEM,
						COUNT(DISTINCT ULT_SOURCE_SYSTEM) OVER () COUNT_ULT_SOURCE_SYSTEM,
						COUNT(DISTINCT BUSINESS_DESCRIPTION) OVER () COUNT_BUSINESS_DESCRIPTION,
						COUNT(DISTINCT CDE_RATIONALE) OVER () COUNT_CDE_RATIONALE
					FROM (
						SELECT DISTINCT
							TDPD.ID,
							TDP.ID	AS tdpid,
							TMC.ALIAS_NAME          AS CDE_NAME,
							TDPD.SEGMENT_NAME       AS SEGMENT, 
							IPS.SYSTEM_NAME         AS IMM_PREC_SYSTEM,
							USS.SYSTEM_NAME         AS ULT_SOURCE_SYSTEM,
							TMC.DESCRIPTION          AS BUSINESS_DESCRIPTION, 
							TBT.CDE_RATIONALE        AS CDE_RATIONALE
							FROM 
								TBL_DS_PROCESSES TDP
								INNER JOIN TBL_DS_PROCESS_DETAIL TDPD ON TDP.ID = TDPD.PROCESS_ID
								INNER JOIN TBL_MD_COLUMN TMC ON TDPD.COLUMN_ID = TMC.ID
								LEFT JOIN TBL_SYSTEM IPS ON TDPD.IMM_PREC_SYSTEM_ID = IPS.ID
								LEFT JOIN TBL_SYSTEM USS ON TDPD.ULTIMATE_SOURCE_SYSTEM_ID = USS.ID
								LEFT JOIN TBL_BUSINESS_TERM TBT ON TMC.BUSINESS_TERM_ID = TBT.ID
							WHERE TDP.ID = '?'
							ORDER BY TMC.ALIAS_NAME
					) res
			) WHERE ( -- Main and dropdown search
				upper(NVL(CDE_NAME,' ')) LIKE upper('%?%')
			) 
	) WHERE ( -- Column filter
		upper(NVL(CDE_NAME,' ')) LIKE upper('%?%')
		AND upper(NVL(SEGMENT,' ')) LIKE upper('%?%')
		AND upper(NVL(IMM_PREC_SYSTEM,' ')) LIKE upper('%?%')
		AND upper(NVL(ULT_SOURCE_SYSTEM,' ')) LIKE upper('%?%')
		AND upper(NVL(BUSINESS_DESCRIPTION,' ')) LIKE upper('%?%')
		AND upper(NVL(CDE_RATIONALE,' ')) LIKE upper('%?%')
	)

-- name: details
SELECT DISTINCT
tdp.id,
tdp.Name 						as downstream_process,
TP.FIRST_NAME||' '|| TP.LAST_NAME 					as process_owner,
tp.bank_id						as bank_id,
itmc.alias_name		            as cde_name,
tbt.cde_rationale				as cde_rationale,

isys.system_name				as imm_system,
isys.itam_id					as imm_itam_id,
itmt.name 						as imm_table_name,
itmc.name 						as imm_column_name,

tbt.bt_name						as imm_screen_label_name,
itmc.description 				as imm_business_description,
itmc.derived					as imm_derived_yes_no,
itmc.derivation_logic			as imm_derivation_logic,
itmc.dq_standards 				as imm_dq_requirements,
itmc.threshold					as imm_threshold,
itmc.data_sla_signed			as imm_data_sla_signed,

usys.system_name				as ult_system,
usys.itam_id         			as ULT_ITAM_ID,
''								as ult_table_name,
''								as ult_column_name,
''								as ult_screen_label_name,
'' 								as ult_business_description,
'' 								as ult_derived_yes_no,
'' 								as ult_derivation_logic,
'' 								as ult_dq_requirements,
'' 								as ult_threshold,
'' 								as golden_source,
'' 								as gs_system_name,
'' 								as gs_itam_id,
'' 								as gs_table_name,
'' 								as gs_column_name,
''								as gs_screen_name,
''								as gs_business_desc,
''								as gs_derived_yes_no,
''								as gs_derivation_logic,
tc.name 						as domain,
tsc.name 						as subdomain,
tlcp.people_id 					as domain_owner,
tbt.bt_name 					as business_term,
tbt.description 				as business_term_description,
itmc.dpo_dq_standards			as dpo_dq_standards,
tbt.dq_standards				as dq_standards_bt_level,
itmc.dpo_threshold				as dpo_threshold,
itmc.threshold					as threshold_bt_level
	FROM
		tbl_ds_processes tdp
		INNER JOIN tbl_ds_process_detail tdpd ON tdpd.process_id = tdp.id
		INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TDP.ID AND TLRP.OBJECT_TYPE = 'PROCESSES'
		INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
		INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID

		LEFT JOIN tbl_md_column itmc ON tdpd.column_id = itmc.id
		LEFT JOIN tbl_md_table itmt ON itmc.table_id = itmt.id
		LEFT JOIN tbl_md_resource itmr ON itmr.id = itmt.resource_id
		LEFT JOIN tbl_system isys ON isys.id = itmr.system_id and tdpd.IMM_PREC_SYSTEM_ID = isys.id
    LEFT JOIN tbl_system usys ON tdpd.ultimate_source_system_id = usys.id
    
		LEFT JOIN tbl_business_term tbt ON itmt.business_term_id = tbt.id
		LEFT JOIN tbl_link_category_people tlcp ON tlcp.people_id = tp.id
		LEFT JOIN tbl_category tc ON tlcp.category_id = tc.id
		LEFT JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
	WHERE
		tdp.id = '?' AND itmc.alias_name = '?'
