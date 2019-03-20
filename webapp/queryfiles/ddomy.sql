-- name: left-grid
SELECT DISTINCT
        tc.id,
        tsc.name                            as sub_domains,
        tc.name                             as data_domain,
        tp.first_name||' '||tp.last_name    as sub_domain_owner,
        tp.bank_id                          as bank_id
    FROM
        tbl_category tc
        inner JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
        inner JOIN Tbl_Link_Role_People tlrp ON tlrp.Object_ID = tsc.id and tlrp.Object_type = 'SUBCATEGORY' and tsc.type = 'Sub Data Domain'
        inner JOIN Tbl_Role rl ON tlrp.role_id = rl.id and rl.role_name = 'Data Domain Owner'
        inner JOIN tbl_people tp ON tlrp.people_id = tp.id
    WHERE tp.bank_id = '?'
        AND upper(tc.name) LIKE upper('%?%')
        AND upper(tsc.name) LIKE upper('%?%')
        AND upper(tp.first_name||' '||tp.last_name) LIKE upper('%?%')

-- name: right-grid
SELECT DISTINCT
        tbt.id,
        tsc.category_id as tscid,
        tbt.BT_Name	as business_term,
        tbt.Description	as bt_description,
        tbt.CDE		as cde_yes_no
    FROM
        tbl_category tc
        INNER JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
        INNER JOIN tbl_business_term tbt ON tbt.parent_id = tsc.id
        LEFT JOIN tbl_link_subcategory_people tlscp ON tlscp.subcategory_id = tsc.id
        LEFT JOIN Tbl_People tp on tlscp.people_id = tp.id
        LEFT join tbl_md_column tmc on tmc.business_term_id = tbt.id
        LEFT join tbl_md_table tmt on tmc.table_id = tmt.id
        LEFT join tbl_md_resource tmr on tmt.resource_id = tmr.id
        LEFT join tbl_system ts on tmr.system_id = ts.id
        LEFT join tbl_policy tpo on tbt.policy_id = tpo.id
        LEFT join tbl_ds_process_detail tdpd on tdpd.column_id = tmc.id
        LEFT join tbl_ds_processes tdp on tdpd.process_id = tdp.id
    WHERE tsc.category_id = '?'
        AND upper(tbt.BT_Name) LIKE upper('%?%')

-- name: details
SELECT DISTINCT
        tc.id,
        tc.name                         as data_domain,
        tsc.name                        as sub_domain,
        tlscp.people_id                 as sub_domain_owner,
        tp.bank_id                      as bank_id,
        tbt.bt_name                     as business_term,
        tbt.description                 as bt_description,
        tmc.alias_name                  as business_alias,
        tbt.cde                         as cde_yes_no,
        tbt.dq_standards                as dq_standards,
        tbt.policy_guidance             as policy_guidance,
        tbt.mandatory                   as mandatory,
        tbt.golden_source_system_id     as golden_system,
        ts.itam_id                      as golden_itam_id,
        tmt.name                        as golden_table_name,
        tmc.name                        as golden_column_name,
        tbt.target_golden_source_id     as target_golden,
        ts.itam_id                      as target_golden_itam_id,
        tpo.info_asset_name             as info_asset_names,
        tpo.description                 as info_asset_desc,
        tpo.confidentiality             as confidentiality,
        tpo.integrity                   as integrity,
        tpo.availability                as availability,
        tpo.overall_cia_rating          as overall_cia_rating,
        tmc.record_category             as record_categories,
        tmc.pii_flag                    as pii_flag,
        tdp.name                        as downstream_process_name,
        tdp.owner_name                  as downstream_process_owner,
        ''                              as taken_from_golden,
        ts.system_name                  as system_name,
        ts.itam_id                      as itam_id,
        tmt.name                        as table_name,
        tmc.name                        as column_name,
        tmc.ddo_threshold               as dq_standards_by_ddo,
        ts.system_name                  as system_name_dd,
        ts.itam_id                      as itam_id_dd,
        tmt.name                        as table_name_dd,
        tmc.name                        as column_name_dd
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