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
		tbt.id,
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
		tdp.Name as downstream_process,
		tdp.owner_id as process_owner,
		tp.bank_id,
		tdpd.business_term_id,
		tdpd.business_term_id as cde_name,
		ts.name as Segment,
		tdpd.imm_prec_system_id,
		tdpd.ultimate_source_system_id,
		tbt.description as business_description,
		tbt.cde_rationale,
		tsy.system_name,
		tsy.itam_id,
		tmt.name as table_name,
		tmc.name as column_name,
		tmc.derived,
		tmc.derivation_logic,
		tmc.dq_standards as dq_requirements,
		tmc.threshold,
		tmc.data_sla_signed,
		tmc.golden_source,
		tbt.golden_source_system_id,
		tc.name as domain,
		tsc.name as subdomain,
		tlcp.people_id as domain_owner,
		tbt.bt_name as business_term,
		tmc.dpo_dq_standards,
		tbt.dq_standards,
		tmc.dpo_threshold
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
		tdp.id = '?'