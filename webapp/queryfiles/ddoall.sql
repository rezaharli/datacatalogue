-- name: left-grid
SELECT * 
    FROM (
        SELECT * 
            FROM (
                SELECT res.*, 
                        COUNT(DISTINCT sub_domains) OVER ()         as COUNT_sub_domains,
                        COUNT(DISTINCT data_domain) OVER ()         as COUNT_data_domain,
                        COUNT(DISTINCT sub_domain_owner) OVER ()    as COUNT_sub_domain_owner,
                        COUNT(DISTINCT bank_id) OVER ()             as COUNT_bank_id
                    FROM (
                        SELECT DISTINCT
                                tsc.id,
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
                    ) res
            ) WHERE ( -- Main and dropdown search
                upper(sub_domains) LIKE upper('%?%')
                AND upper(data_domain) LIKE upper('%?%')
                AND upper(sub_domain_owner) LIKE upper('%?%')
            )
    ) WHERE ( -- Column filter
        upper(sub_domains) LIKE upper('%?%')
        AND upper(data_domain) LIKE upper('%?%')
        AND upper(sub_domain_owner) LIKE upper('%?%')
        AND upper(bank_id) LIKE upper('%?%')
    )

-- name: right-grid
SELECT DISTINCT
        tbt.id,
        tsc.id as tscid,
        tbt.BT_Name	as business_term,
        tbt.Description	as bt_description,
        tbt.CDE		as cde_yes_no
    FROM
        tbl_category tc
        INNER JOIN tbl_subcategory tsc ON tsc.category_id = tc.id
        INNER JOIN tbl_business_term tbt ON tbt.parent_id = tsc.id
    WHERE tsc.id = '?'

-- name: details-business-metadata-from-domain
SELECT DISTINCT 
        tsc.id,
        tbt.id                                as tbtid,
        TSC.NAME                              AS SUB_DOMAIN,
        TC.NAME                               AS DATA_DOMAIN, 
        TP.FIRST_NAME||' '|| TP.LAST_NAME     AS SUBDOMAIN_OWNER,
        TP.BANK_ID                            AS BANK_ID,
        TBT.BT_NAME                           AS BUSINESS_TERM,
        TBT.DESCRIPTION                       AS BT_DESCRIPTION,
        TMC.ALIAS_NAME                        AS BUSINESS_ALIAS,
        TBT.CDE                               AS CDE_YES_NO,
        TMC.DQ_STANDARDS||' '||TMC.THRESHOLD  AS DQ_STANDARDS,
        TBT.POLICY_GUIDANCE                   AS POLICY_GUIDANCE,
        TMC.MANDATORY                         AS MANDATORY,
        GS_SYS.SYSTEM_NAME                    AS GS_SYSTEM,
        GS_SYS.ITAM_ID                        AS GS_ITAM_ID,
        TMT_GS.NAME                           AS GS_TABLE_NAME,
        TMC_GS.NAME                           AS GS_COLUMN_NAME,
        TGS_SYS.SYSTEM_NAME                   AS TGS_SYSTEM,
        TGS_SYS.ITAM_ID                       AS TGS_ITAM_ID,
        TPOL.INFO_ASSET_NAME                  AS INFO_ASSET_NAMES,
        TPOL.DESCRIPTION                      AS INFO_ASSET_DESC,
        TPOL.CONFIDENTIALITY                  AS CONFIDENTIALITY,
        TPOL.INTEGRITY                        AS INTEGRITY,
        TPOL.AVAILABILITY                     AS AVAILABILITY,
        TPOL.OVERALL_CIA_RATING               AS OVERALL_CIA_RATING,
        TMC.RECORD_CATEGORY                   AS RECORD_CATEGORIES,
        TMC.PII_FLAG                          AS PII_FLAG
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID
        INNER JOIN TBL_BUSINESS_TERM TBT ON TBT.PARENT_ID = TSC.ID
        INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TSC.ID AND TLRP.OBJECT_TYPE = 'SUBCATEGORY'
        INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DATA DOMAIN OWNER'
        INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID -- SUBDOMAIN OWNER
        INNER JOIN TBL_MD_COLUMN TMC ON TMC.BUSINESS_TERM_ID = TBT.ID
        INNER JOIN TBL_MD_TABLE TMT ON TMC.TABLE_ID = TMT.ID
        INNER JOIN TBL_MD_RESOURCE TMR ON TMT.RESOURCE_ID = TMR.ID
        INNER JOIN TBL_SYSTEM TS ON TMR.SYSTEM_ID = TS.ID
        LEFT JOIN TBL_POLICY TPOL ON TBT.POLICY_ID = TPOL.ID
        LEFT JOIN TBL_SYSTEM TGS_SYS ON TBT.TARGET_GOLDEN_SOURCE_ID = TGS_SYS.ID
        LEFT JOIN TBL_SYSTEM GS_SYS ON TBT.GOLDEN_SOURCE_SYSTEM_ID = GS_SYS.ID
        INNER JOIN TBL_MD_COLUMN TMC_GS ON  TMC_GS.BUSINESS_TERM_ID = TBT.ID
        INNER JOIN TBL_MD_TABLE TMT_GS ON TMC_GS.TABLE_ID = TMT_GS.ID 
        LEFT JOIN TBL_MD_RESOURCE TMR_GS ON TMR_GS.SYSTEM_ID = GS_SYS.ID AND TMT_GS.RESOURCE_ID = TMR_GS.ID
    WHERE
        tsc.id = '?'

-- name: details-downstream-usage-of-business-term
SELECT DISTINCT
        tsc.id,
        tbt.id                                  as tbtid,
        TBT.BT_NAME                             AS BUSINESS_TERM,
        TDP.NAME                                AS DOWNSTREAM_PROCESS_NAME, 
        PPL.FIRST_NAME||' '|| PPL.LAST_NAME     AS DOWNSTREAM_PROCESS_OWNER,
        'NO'                                    AS DATA_FROM_GS,
        ' '                                     AS SYSTEM_NAME,
        ' '                                     AS ITAM_ID,
        ' '                                     AS TABLE_NAME,
        ' '                                     AS COLUMN_NAME, 
        ' '                                     AS DQ_STANDARDS
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID
        INNER JOIN TBL_BUSINESS_TERM TBT ON TBT.PARENT_ID = TSC.ID

        INNER JOIN TBL_MD_COLUMN TMC ON TMC.BUSINESS_TERM_ID = TBT.ID
        INNER JOIN TBL_MD_TABLE TMT ON TMC.TABLE_ID = TMT.ID
        INNER JOIN TBL_MD_RESOURCE TMR ON TMT.RESOURCE_ID = TMR.ID
        INNER JOIN TBL_SYSTEM TS ON TMR.SYSTEM_ID = TS.ID
        
        LEFT JOIN TBL_LINK_COLUMN_INTERFACE CI ON TMC.ID = CI.COLUMN_ID
        LEFT JOIN TBL_DS_PROCESS_DETAIL TDPD ON TDPD.COLUMN_ID = TMC.ID
        LEFT JOIN TBL_DS_PROCESSES TDP ON TDPD.PROCESS_ID = TDP.ID
        
        LEFT JOIN TBL_LINK_ROLE_PEOPLE TLRP_SDO ON TLRP_SDO.OBJECT_ID = TDP.ID AND TLRP_SDO.OBJECT_TYPE = 'PROCESSES'
        LEFT JOIN TBL_ROLE RL ON TLRP_SDO.ROLE_ID = RL.ID AND RL.ROLE_NAME = 'DOWNSTREAM PROCESS OWNER'
        LEFT JOIN TBL_PEOPLE PPL ON TLRP_SDO.PEOPLE_ID = PPL.ID
        
        LEFT JOIN TBL_POLICY TPOL ON TBT.POLICY_ID = TPOL.ID
        LEFT JOIN TBL_SYSTEM GS_SYS ON TBT.GOLDEN_SOURCE_SYSTEM_ID = GS_SYS.ID
        LEFT JOIN TBL_SYSTEM TGS_SYS ON TBT.TARGET_GOLDEN_SOURCE_ID = TGS_SYS.ID
    WHERE 
        tsc.id = '?'

-- name: details-business-term-residing
SELECT DISTINCT 
        tsc.id,
        tbt.id                            as tbtid,
        TBT.BT_NAME                       AS BUSINESS_TERM,
        TS.SYSTEM_NAME                    AS SYSTEM_NAME,
        TS.ITAM_ID                        AS ITAM_ID,
        TMT.NAME                          AS TABLE_NAME,
        TMC.NAME                          AS COLUMN_NAME     
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID
        INNER JOIN TBL_BUSINESS_TERM TBT ON TBT.PARENT_ID = TSC.ID

        INNER JOIN TBL_MD_COLUMN TMC ON TMC.BUSINESS_TERM_ID = TBT.ID
        INNER JOIN TBL_MD_TABLE TMT ON TMC.TABLE_ID = TMT.ID
        INNER JOIN TBL_MD_RESOURCE TMR ON TMT.RESOURCE_ID = TMR.ID
        INNER JOIN TBL_SYSTEM TS ON TMR.SYSTEM_ID = TS.ID
    WHERE 
        tsc.id = '?'
    ORDER BY TS.SYSTEM_NAME, TMT.NAME, TMC.NAME
