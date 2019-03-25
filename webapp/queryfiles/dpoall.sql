-- name: left-grid
SELECT DISTINCT
		tdp.id,
		tdp.Name 		as downstream_process,
		tdp.owner_id 	as process_owner,
		tp.bank_id		as bank_id
	FROM
		tbl_ds_processes tdp
		LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = tdp.id
		LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id

-- name: right-grid
SELECT DISTINCT
		tdpd.id,
		tdp.id 							as tdpid,
		tdpd.business_term_id 			as cde_name,
		ts.name 						as segment,
		tdpd.imm_prec_system_id			as imm_prec_system,
		tdpd.ultimate_source_system_id	as ult_source_system,
		tbt.description 				as business_description,
		tbt.cde_rationale				as cde_rationale
	FROM 
		tbl_ds_processes tdp
		LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = tdp.id
		LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id

		LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.process_id = tdp.id
		LEFT JOIN tbl_segment ts ON tdpd.segment_id = ts.id
		LEFT JOIN tbl_business_term tbt ON tdpd.business_term_id = tbt.id
		LEFT JOIN tbl_system tsy ON tdpd.imm_prec_system_id = tsy.id
		LEFT JOIN tbl_md_table tmt ON tmt.business_term_id = tbt.id
		LEFT JOIN tbl_md_column tmc ON tmc.table_id = tmt.id
		LEFT JOIN tbl_category tc ON ts.subdomain_id = tc.id
		LEFT JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
		LEFT JOIN tbl_link_category_people tlcp ON tlcp.category_id = tc.id
	WHERE tdp.id = '?'

-- name: details
SELECT DISTINCT
		tdp.id,
		tdp.Name 						as downstream_process,
		tdp.owner_id 					as process_owner,
		tp.bank_id						as bank_id,
		tdpd.business_term_id 			as cde_name,
		tbt.cde_rationale				as cde_rationale,
		tdpd.imm_prec_system_id			as imm_system,
		tsy.itam_id						as imm_itam_id,
		tmt.name 						as imm_table_name,
		tmc.name 						as imm_column_name,
		tbt.bt_name						as imm_screen_label_name,
		tbt.description 				as imm_business_description,
		tmc.derived						as imm_derived_yes_no,
		tmc.derivation_logic			as imm_derivation_logic,
		tmc.dq_standards 				as imm_dq_requirements,
		tmc.threshold					as imm_threshold,
		tmc.data_sla_signed				as imm_data_sla_signed,
		tdpd.imm_prec_system_id			as ult_system,
		tsy.itam_id						as ult_itam_id,
		tmt.name 						as ult_table_name,
		tmc.name 						as ult_column_name,
		''								as ult_screen_label_name,
		tbt.description 				as ult_business_description,
		tmc.derived						as ult_derived_yes_no,
		tmc.derivation_logic			as ult_derivation_logic,
		tmc.dq_standards 				as ult_dq_requirements,
		tmc.threshold					as ult_threshold,
		tmc.golden_source				as golden_source,
		tbt.golden_source_system_id		as gs_system_name,
		tbt.Golden_Source_ITAM_ID		as gs_itam_id,
		tbt.Golden_Source_TableName_ID	as gs_table_name,
		tbt.Golden_Source_Column_ID		as gs_column_name,
		''								as gs_screen_name,
		''								as gs_business_desc,
		''								as gs_derived_yes_no,
		''								as gs_derivation_logic,
		tc.name 						as domain,
		tsc.name 						as subdomain,
		tlcp.people_id 					as domain_owner,
		tbt.bt_name 					as business_term,
		tbt.description 				as business_term_description,
		tmc.dpo_dq_standards			as dpo_dq_standards,
		tbt.dq_standards				as dq_standards_bt_level,
		tmc.dpo_threshold				as dpo_threshold,
		tmc.threshold					as threshold_bt_level
	FROM
		tbl_ds_processes tdp
		INNER JOIN tbl_ds_process_detail tdpd ON tdpd.process_id = tdp.id
		LEFT JOIN tbl_people tp ON tdp.owner_id = tp.id
		LEFT JOIN tbl_segment ts ON tdpd.segment_id = ts.id
		LEFT JOIN tbl_business_term tbt ON tdpd.business_term_id = tbt.id
		LEFT JOIN tbl_system tsy ON tdpd.imm_prec_system_id = tsy.id
		LEFT JOIN tbl_md_table tmt ON tmt.business_term_id = tbt.id
		LEFT JOIN tbl_md_column tmc ON tmc.table_id = tmt.id
		LEFT JOIN tbl_category tc ON ts.subdomain_id = tc.id
		LEFT JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
		LEFT JOIN tbl_link_category_people tlcp ON tlcp.category_id = tc.id
	WHERE
		tdp.id = '?' AND tdpd.id = '?'