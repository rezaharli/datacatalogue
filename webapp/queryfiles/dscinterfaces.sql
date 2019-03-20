-- name: left-grid
SELECT DISTINCT
        ts.id,
        ts.system_name 	as system_name,
        ts.itam_id		as itam_id,
        tp.first_name	as dataset_custodian,
        tp.bank_id		as bank_id
    FROM tbl_system ts 
        LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = ts.id and tlrp.Object_type = 'SYSTEM'
        LEFT JOIN Tbl_Role rl_sys ON tlrp.role_id = rl_sys.id and rl_sys.role_name = 'Dataset Custodian'
        LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id 
        
        inner join tbl_md_resource tmr ON ts.id = tmr.system_id
        inner join
            (
                SELECT
                DISTINCT ts.id as sys_id, tmr.id as res_id, tmt.id as tab_id 
                FROM tbl_system ts
                    inner join tbl_md_resource tmr ON ts.id = tmr.system_id
                    inner join tbl_md_table tmt ON tmr.id = tmt.resource_id
                    inner join tbl_md_column tmc ON tmt.id = tmc.table_id
                WHERE CDE = 1
            ) cde ON ts.id = cde.sys_id and tmr.id = cde.res_id

-- name: right-grid
SELECT DISTINCT
        tmt.id,
        ts.id                               as tsid,
        tmc.id                              as colid,
        tmt.name                            as table_name,
        tmc.name                            as column_name,
        tmc.alias_name                      as list_of_cde,
        ips.system_name                     as imm_prec_system_name,
        tmc.Imm_Prec_System_SLA             as Imm_Prec_System_SLA,
        tmc.Imm_Prec_System_OLA             as Imm_Prec_System_OLA,
        iss.system_name                     as imm_succ_system_name,
        tmc.Imm_Succ_System_SLA             as Imm_Succ_System_SLA,
        tmc.Imm_Succ_System_OLA             as Imm_Succ_System_OLA,
        tdp.name                            as list_downstream_process,
        ppl.first_name||' '||ppl.last_name  as downstream_process_owner
    FROM tbl_system ts
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
        LEFT JOIN Tbl_People ppl ON tlrp_sdo.people_id = ppl.id
    WHERE ts.id = '?'
        AND (ips.system_name is not null or iss.system_name is not null) AND tmc.cde = 1

-- name: details
SELECT DISTINCT
		ts.id,
		ts.system_name						as system_name,
		ts.itam_id							as itam_id,
		tp.first_name						as dataset_custodian,
		tp.bank_id							as bank_id,
		tmc.alias_name						as business_alias_name,
		tmt.name 							as table_name,
		tmc.name 							as column_name,
		tmc.description 					as business_alias_description,
		tmc.cde								as cde_yes_no,
		tmc.status							as status,
		tmc.data_type						as data_type,
		tmc.data_format						as data_format,
		tmc.data_length						as data_length,
		tmc.example							as example,
		tmc.derived							as derived_yes_no,
		tmc.Derivation_Logic				as derivation_logic,
		tmc.Sourced_from_Upstream			as sourced_from_upstream_yes_no,
		tmc.System_Checks					as system_checks,
		tc.Name 							as domain,
		tsc.name 							as subdomain,
		ppl.first_name||' '||ppl.last_name 	as domain_owner,
		tbt.bt_name							as business_term,
		tbt.description 					as business_term_description,
		tpol.info_asset_name				as information_asset_names,
		tpol.description 					as information_asset_description,
		tpol.confidentiality				as confidentiality,
		tpol.integrity						as integrity,
		tpol.availability					as availability,
		tpol.overall_cia_rating				as overall_cia_rating,
		tmt.record_category					as record_categories,
		tmc.pii_flag						as pii_flag,
		ips.system_name 					as imm_preceeding_system,
		iss.system_name 					as imm_succeeding_system,
		tmc.threshold						as threshold
    FROM tbl_system ts
        inner join tbl_md_resource res ON ts.id = res.system_id
        inner join tbl_md_table tmt ON res.id = tmt.resource_id
        inner join tbl_md_column tmc ON tmt.id = tmc.table_id
        
        left join tbl_link_column_interface ci on tmc.id = ci.column_id
        left join tbl_system ips on ci.imm_prec_system_id = ips.id
        left join tbl_system iss on ci.imm_succ_system_id = iss.id
        LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.column_id = tmc.id
        LEFT JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id 
        
        LEFT JOIN Tbl_Link_Role_People tlrp_sdo ON tlrp_sdo.Object_ID = tdp.id and tlrp_sdo.object_type = 'SUBCATEGORY'
        LEFT JOIN Tbl_Role rl ON tlrp_sdo.role_id = rl.id and rl.role_name = 'Data Domain Owner'
        LEFT JOIN Tbl_People ppl ON tlrp_sdo.people_id = ppl.id
        LEFT JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = ts.id and tlrp.Object_type = 'SYSTEM'
        LEFT JOIN Tbl_Role rl_sys ON tlrp.role_id = rl_sys.id and rl_sys.role_name = 'Dataset Custodian'
        LEFT JOIN tbl_people tp ON tlrp.people_id = tp.id 
        
        LEFT JOIN tbl_business_term tbt ON tmc.business_term_id = tbt.id
        LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
        LEFT JOIN tbl_category tc ON tsc.category_id = tc.id
        LEFT JOIN tbl_policy tpol ON tbt.policy_id = tpol.id
        
        LEFT JOIN Tbl_Subcategory tsc ON tbt.parent_id = tsc.id
    WHERE ts.id = '?'
		AND (ips.system_name is not null or iss.system_name is not null) AND tmc.cde = 1
        AND 
        (
            (
                tmt.id = '?'
                AND tmc.id = '?'
            ) OR (
                tmt.name = '?'
                AND tmc.name = '?'
                AND tmc.alias_name = '?'
            )
        )

-- name: details-dropdown
SELECT DISTINCT
		tmt.name 	    as table_name,
		tmc.name 	    as column_name,
		tmc.alias_name  as business_alias_name
    FROM tbl_system ts
        inner join tbl_md_resource res ON ts.id = res.system_id
        inner join tbl_md_table tmt ON res.id = tmt.resource_id
        inner join tbl_md_column tmc ON tmt.id = tmc.table_id
        
        left join tbl_link_column_interface ci on tmc.id = ci.column_id
        left join tbl_system ips on ci.imm_prec_system_id = ips.id
        left join tbl_system iss on ci.imm_succ_system_id = iss.id
        LEFT JOIN tbl_ds_process_detail tdpd ON tdpd.column_id = tmc.id
        LEFT JOIN tbl_ds_processes tdp ON tdpd.process_id = tdp.id 
        
        LEFT JOIN Tbl_Link_Role_People tlrp_sdo ON tlrp_sdo.Object_ID = tdp.id and tlrp_sdo.object_type = 'SUBCATEGORY'
        LEFT JOIN Tbl_Role rl ON tlrp_sdo.role_id = rl.id and rl.role_name = 'Data Domain Owner'
        LEFT JOIN Tbl_People ppl ON tlrp_sdo.people_id = ppl.id
    WHERE ts.id = '?'
		AND (ips.system_name is not null or iss.system_name is not null) AND tmc.cde = 1