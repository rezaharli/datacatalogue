-- name: left-grid
SELECT DISTINCT
        tc.id,
        tc.name     as data_domain,
        tsc.name    as sub_domains
    FROM
        tbl_category tc
        INNER JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
    WHERE
        upper(tc.name) LIKE upper('%?%')
        OR upper(tsc.name) LIKE upper('%?%')

-- name: right-grid
SELECT DISTINCT
        tbt.id,
        tsc.category_id as tscid,
        tbt.BT_Name		as business_term,
        tbt.Description	as business_term_description,
        tbt.CDE			as cde_yes_no
    FROM
        tbl_category tc
        INNER JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
        INNER JOIN tbl_business_term tbt ON tbt.parent_id = tsc.id
        LEFT JOIN tbl_link_subcategory_people tlscp ON tlscp.subcategory_id = tsc.id
        LEFT JOIN Tbl_People	tp on tlscp.people_id = tp.id
        LEFT join tbl_md_column tmc on tmc.business_term_id = tbt.id
        LEFT join tbl_md_table tmt on tmc.table_id = tmt.id
        LEFT join tbl_md_resource tmr on tmt.resource_id = tmr.id
        LEFT join tbl_system ts on tmr.system_id = ts.id
        LEFT join tbl_policy tpo on tbt.policy_id = tpo.id
        LEFT join tbl_ds_process_detail tdpd on tdpd.column_id = tmc.id
        LEFT join tbl_ds_processes tdp on tdpd.process_id = tdp.id
    WHERE tsc.category_id = '?'
        AND (
            upper(tbt.BT_Name) LIKE upper('%?%')
            OR upper(tbt.Description) LIKE upper('%?%')
        )

-- name: details
SELECT DISTINCT
        tc.id,
        tc.name as data_domain,
        tsc.name as sub_domain,
        tlscp.people_id as sub_domain_owner,
        tp.bank_id,
        tbt.bt_name,
        tbt.description as business_term_description,
        tmc.alias_name,
        tbt.cde,
        tbt.dq_standards,
        tbt.policy_guidance,
        tbt.mandatory,
        tbt.golden_source_system_id,
        ts.itam_id,
        tmt.name as golden_source_table_name,
        tmc.name as golden_source_column_name,
        tpo.info_asset_name,
        tpo.description as info_asset_desc,
        tpo.confidentiality,
        tpo.integrity,
        tpo.availability,
        tpo.overall_cia_rating,
        tmc.record_category,
        tmc.pii_flag,
        tdp.name as downstream_process_name,
        tdp.owner_name as downstream_process_owner,
        ts.system_name,
        tmc.name as column_name,
        tmc.ddo_threshold
    FROM
        tbl_category tc
        INNER JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
        INNER JOIN tbl_business_term tbt ON tbt.parent_id = tsc.id
        LEFT JOIN tbl_link_subcategory_people tlscp ON tlscp.subcategory_id = tsc.id
        LEFT JOIN Tbl_People	tp on tlscp.people_id = tp.id
        LEFT join tbl_md_column tmc on tmc.business_term_id = tbt.id
        LEFT join tbl_md_table tmt on tmc.table_id = tmt.id
        LEFT join tbl_md_resource tmr on tmt.resource_id = tmr.id
        LEFT join tbl_system ts on tmr.system_id = ts.id
        LEFT join tbl_policy tpo on tbt.policy_id = tpo.id
        LEFT join tbl_ds_process_detail tdpd on tdpd.column_id = tmc.id
        LEFT join tbl_ds_processes tdp on tdpd.process_id = tdp.id
    WHERE
        tc.id = '?' 
        AND tbt.id = '?'