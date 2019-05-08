-- name: dpo-view
SELECT DISTINCT
        TC.NAME                                                                     AS DSP_NAME,
        TP.FIRST_NAME||' '||TP.LAST_NAME||' | '||TP.BANK_ID                         AS DSP_OWNER,
        TP.BANK_ID                                                                  AS BANK_ID,
        COUNT(DISTINCT TC.NAME) OVER ()                                             AS COUNT_DSP_NAME,
        COUNT(DISTINCT TP.FIRST_NAME||' '||TP.LAST_NAME||' | '||TP.BANK_ID) OVER () AS COUNT_DSP_OWNER,
        COUNT(DISTINCT TP.BANK_ID) OVER ()                                          AS COUNT_BANK_ID
    FROM
        TBL_DS_PROCESSES TC
        INNER JOIN TBL_DS_PROCESS_DETAIL TSC ON TSC.PROCESS_ID = TC.ID 
        INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TSC.ID AND UPPER(TLRP.OBJECT_TYPE) = 'PROCESSES'
        INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
        INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
    WHERE ('ALL' = '?' OR TP.BANK_ID = '?')
        AND (
            upper(NVL(TC.NAME, ' ')) LIKE upper('%?%')
            AND upper(NVL(TP.FIRST_NAME||' '||TP.LAST_NAME||' | '||TP.BANK_ID, ' ')) LIKE upper('%?%')
        )
    ORDER BY TC.NAME

-- name: dpo-view-homepage
SELECT DATA_ELEMENTS, UPSTREAM_DATA FROM 
    (
    SELECT 
        COUNT(DISTINCT CD.ALIAS_NAME)         AS DATA_ELEMENTS
    FROM
        TBL_DS_PROCESSES TC
        INNER JOIN TBL_DS_PROCESS_DETAIL TSC ON TSC.PROCESS_ID = TC.ID 
        INNER JOIN TBL_MD_COLUMN COL ON TSC.COLUMN_ID = COL.ID
        INNER JOIN TBL_MD_COLUMN_DETAILS CD ON CD.COLUMN_ID = COL.ID
        INNER JOIN TBL_MD_TABLE TAB ON TAB.ID = COL.TABLE_ID
        INNER JOIN TBL_MD_RESOURCE RES ON RES.ID = TAB.RESOURCE_ID
        INNER JOIN TBL_SYSTEM SYS ON SYS.ID = RES.SYSTEM_ID
    WHERE TC.NAME = '?'
    ) DATA_ELEMENTS, 
    (
    SELECT COUNT(1) AS UPSTREAM_DATA FROM TBL_DS_PROCESSES WHERE 1 = 2 -- PLACEHOLDER --
    )
    UPSTREAM_DATA

-- name: dpo-dataelements
SELECT DISTINCT
        TC.NAME                                         AS DSP_NAME,
        COUNT(DISTINCT CD.ALIAS_NAME) OVER ()         AS COUNT_DATA_ELEMENTS,
        SYS.SYSTEM_NAME                               AS SYSTEM_NAME,
        SYS.ITAM_ID                                   AS ITAM_ID,
        CD.ALIAS_NAME                                 AS ALIAS_NAME,
        CD.CDE                                        AS CDE,
        TAB.NAME                                      AS TABLE_NAME,
        COL.NAME                                      AS COLUMN_NAME,
        ULT.SYSTEM_NAME                               AS ULT_SYSTEM_NAME,
        CD.DATA_SLA_SIGNED                            AS DATA_SLA
    FROM
        TBL_DS_PROCESSES TC
        INNER JOIN TBL_DS_PROCESS_DETAIL TSC ON TSC.PROCESS_ID = TC.ID 
        INNER JOIN TBL_MD_COLUMN COL ON TSC.COLUMN_ID = COL.ID
        INNER JOIN TBL_MD_COLUMN_DETAILS CD ON CD.COLUMN_ID = COL.ID
        INNER JOIN TBL_MD_TABLE TAB ON TAB.ID = COL.TABLE_ID
        INNER JOIN TBL_MD_RESOURCE RES ON RES.ID = TAB.RESOURCE_ID
        INNER JOIN TBL_SYSTEM SYS ON SYS.ID = RES.SYSTEM_ID
        LEFT OUTER JOIN TBL_SYSTEM ULT ON ULT.ID = TSC.ULTIMATE_SOURCE_SYSTEM_ID
    WHERE TC.NAME = '?'
        AND (
            upper(NVL(SYS.SYSTEM_NAME, ' ')) LIKE upper('%?%')
            AND upper(NVL(SYS.ITAM_ID, ' ')) LIKE upper('%?%')
            AND upper(NVL(CD.ALIAS_NAME, ' ')) LIKE upper('%?%')
            AND upper(CD.CDE) LIKE upper('%?%')
            AND upper(NVL(TAB.NAME, ' ')) LIKE upper('%?%')
            AND upper(NVL(COL.NAME, ' ')) LIKE upper('%?%')
            AND upper(NVL(ULT.SYSTEM_NAME, ' ')) LIKE upper('%?%')
            AND upper(CD.DATA_SLA_SIGNED) LIKE upper('%?%')
        )
    ORDER BY SYS.SYSTEM_NAME, TAB.NAME, COL.NAME

-- name: details
SELECT DISTINCT
        tdp.id,
        tdp.Name 						as downstream_process,
        TP.FIRST_NAME||' '|| TP.LAST_NAME 					as process_owner,
        tp.bank_id						as bank_id,
        ''		            as cde_name,
        ''				as cde_rationale,

        isys.system_name				as imm_system,
        isys.itam_id					as imm_itam_id,
        itmt.name 						as imm_table_name,
        itmc.name 						as imm_column_name,

        ''				as imm_screen_label_name,
        '' 				as imm_business_description,
        ''					as imm_derived_yes_no,
        ''			as imm_derivation_logic,
        '' 				as imm_dq_requirements,
        ''					as imm_threshold,
        ''			as imm_data_sla_signed,

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
        '' 						        as domain,
        '' 						        as subdomain,
        '' 					            as domain_owner,
        tbt.bt_name 					as business_term,
        tbt.description 				as business_term_description,
        ''			as dpo_dq_standards,
        tbt.dq_standards				as dq_standards_bt_level,
        ''				as dpo_threshold,
        ''					as threshold_bt_level
	FROM
		tbl_ds_processes tdp
		INNER JOIN tbl_ds_process_detail tdpd ON tdpd.process_id = tdp.id
		INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TDP.ID AND TLRP.OBJECT_TYPE = 'PROCESSES'
		INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DOWNSTREAM PROCESS OWNER'
		INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID

		LEFT JOIN tbl_md_column itmc ON tdpd.column_id = itmc.id
        INNER JOIN TBL_MD_COLUMN_DETAILS CD ON CD.COLUMN_ID = itmc.ID
		LEFT JOIN tbl_md_table itmt ON itmc.table_id = itmt.id
		LEFT JOIN tbl_md_resource itmr ON itmr.id = itmt.resource_id
		LEFT JOIN tbl_system isys ON isys.id = itmr.system_id and tdpd.IMM_PREC_SYSTEM_ID = isys.id
    LEFT JOIN tbl_system usys ON tdpd.ultimate_source_system_id = usys.id
    
		LEFT JOIN tbl_business_term tbt ON itmt.business_term_id = tbt.id
	WHERE
		tdp.name = '?' AND CD.alias_name = '?'

-- name: details-immediate-preceding-system
SELECT DISTINCT 
        DSP.ID,
        DSP.NAME                    AS PROCESS_NAME,
        CD.ALIAS_NAME               AS DATA_ELEMENT,
        SYS.SYSTEM_NAME             AS SYSTEM_NAME, 
        SYS.ITAM_ID                 AS ITAM_ID,
        TAB.NAME                    AS TABLE_NAME, 
        COL.NAME                    AS COLUMN_NAME, 
        CD.DESCRIPTION              AS DESCRIPTION, 
        CD.DERIVED                  AS DERIVED,
        CD.DERIVATION_LOGIC         AS DERIVATION_LOGIC,
        CD.DQ_STANDARDS             AS DQ_STANDARDS,
        CD.THRESHOLD                AS THRESHOLD,
        CD.DATA_SLA_SIGNED          AS DATA_SLA
    FROM TBL_DS_PROCESSES DSP
        INNER JOIN TBL_DS_PROCESS_DETAIL DSPD ON DSP.ID = DSPD.PROCESS_ID 
        INNER JOIN TBL_MD_COLUMN COL ON DSPD.COLUMN_ID = COL.ID 
        INNER JOIN TBL_MD_COLUMN_DETAILS CD ON CD.COLUMN_ID = COL.ID 
        INNER JOIN TBL_MD_TABLE TAB ON COL.TABLE_ID = TAB.ID
        INNER JOIN TBL_MD_RESOURCE RES ON TAB.RESOURCE_ID = RES.ID
        INNER JOIN TBL_SYSTEM SYS ON RES.SYSTEM_ID = SYS.ID
    WHERE 
        DSP.NAME = '?' -- DOWNSTREAM PROCESS NAME
        AND CD.ALIAS_NAME = '?' --CDE NAME

-- name: details-ultimate-source-system
SELECT DISTINCT 
        DSP.ID,
        USYS.SYSTEM_NAME             AS SYSTEM_NAME, 
        USYS.ITAM_ID                 AS ITAM_ID,
        UTAB.NAME                    AS TABLE_NAME, 
        UCOL.NAME                    AS COLUMN_NAME, 
        UCD.ALIAS_NAME                AS DATA_ELEMENT,
        UCD.DESCRIPTION              AS DESCRIPTION, 
        UCD.DERIVED                  AS DERIVED,
        UCD.DERIVATION_LOGIC         AS DERIVATION_LOGIC,
        UCD.DQ_STANDARDS             AS DQ_STANDARDS,
        UCD.THRESHOLD                AS THRESHOLD,
        CASE WHEN UCOL.ID = GSC.ID THEN 'YES' ELSE 'NO' END AS GOLDEN_SOURCE,
        GS.SYSTEM_NAME               AS GS_SYSTEM_NAME, 
        GS.ITAM_ID                   AS GF_ITAM_ID,
        GST.NAME                     AS GS_TABLE_NAME, 
        GSC.NAME                     AS GS_COLUMN_NAME, 
        GSCD.ALIAS_NAME                AS GS_DATA_ELEMENT,
        GSCD.DESCRIPTION             AS GS_DESCRIPTION, 
        GSCD.DERIVED                 AS GS_DERIVED,
        GSCD.DERIVATION_LOGIC        AS GS_DERIVATION_LOGIC
    FROM TBL_DS_PROCESSES DSP
        INNER JOIN TBL_DS_PROCESS_DETAIL DSPD ON DSP.ID = DSPD.PROCESS_ID 
        INNER JOIN TBL_MD_COLUMN COL ON DSPD.COLUMN_ID = COL.ID 
        INNER JOIN TBL_MD_COLUMN_DETAILS CD ON CD.COLUMN_ID = COL.ID 
        INNER JOIN TBL_SYSTEM USYS ON DSPD.ULTIMATE_SOURCE_SYSTEM_ID = USYS.ID
        INNER JOIN TBL_MD_RESOURCE URES ON USYS.ID = URES.SYSTEM_ID
        INNER JOIN TBL_MD_TABLE UTAB ON UTAB.RESOURCE_ID = URES.ID
        INNER JOIN TBL_MD_COLUMN UCOL ON UCOL.TABLE_ID = UTAB.ID
        INNER JOIN TBL_MD_COLUMN_DETAILS UCD ON UCD.COLUMN_ID = UCOL.ID AND TRIM(UPPER(DSPD.USS_CDE_NAME)) = TRIM(UPPER(UCD.ALIAS_NAME))
        LEFT OUTER JOIN TBL_LINK_COLUMN_GOLDEN_SOURCE CGS ON COL.ID = CGS.COLUMN_ID
        LEFT OUTER JOIN TBL_MD_COLUMN GSC ON GSC.ID = CGS.GOLDEN_SOURCE_COLUMN_ID
        LEFT OUTER JOIN TBL_MD_COLUMN_DETAILS GSCD ON GSCD.COLUMN_ID = CGS.ID
        LEFT OUTER JOIN TBL_MD_TABLE GST ON GST.ID = GSC.TABLE_ID
        LEFT OUTER JOIN TBL_MD_RESOURCE GSR ON GSR.ID = GST.RESOURCE_ID
        LEFT OUTER JOIN TBL_SYSTEM GS ON GS.ID = GSR.SYSTEM_ID
    WHERE 
        DSP.NAME = '?' -- DOWNSTREAM PROCESS NAME
        AND CD.ALIAS_NAME = '?' --CDE NAME

-- name: details-domain-view
SELECT DISTINCT 
        TC.ID,
        TC.NAME                               AS DOMAIN_NAME, 
        TSC.NAME                              AS SUBDOMAIN_NAME, 
        TP.FIRST_NAME||' '|| TP.LAST_NAME     AS SUBDOMAIN_OWNER,
        TP.BANK_ID                            AS SUBDOMAIN_OWNER_BANKID,
        TBT.BT_NAME                           AS BUSINESS_TERM,
        TBT.DESCRIPTION                       AS BUSINESS_TERM_DESCRIPTION
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID
        INNER JOIN TBL_BUSINESS_TERM TBT ON TBT.PARENT_ID = TSC.ID
        INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TSC.ID AND TLRP.OBJECT_TYPE = 'SUBCATEGORY'
        INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DATA DOMAIN OWNER'
        INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID -- SUBDOMAIN OWNER
        INNER JOIN TBL_LINK_COLUMN_BUSINESS_TERM CBT ON TBT.ID = CBT.BUSINESS_TERM_ID
        INNER JOIN TBL_MD_COLUMN TMC ON TMC.ID = CBT.COLUMN_ID
        INNER JOIN TBL_MD_COLUMN_DETAILS TMCD ON TMCD.COLUMN_ID = TMC.ID
    WHERE UPPER(TMCD.ALIAS_NAME) = UPPER('?')  --CDE NAME

-- name: details-data-standards
SELECT DISTINCT
        TDP.ID,
        ITMCD.DQ_STANDARDS			     AS DPO_DQ_STANDARDS,
        TBT.DQ_STANDARDS				 AS DQ_STANDARDS_BT_LEVEL,
        ITMCD.THRESHOLD				     AS DPO_THRESHOLD,
        TBT.THRESHOLD					 AS THRESHOLD_BT_LEVEL
    FROM
        TBL_DS_PROCESSES TDP
        INNER JOIN TBL_DS_PROCESS_DETAIL TDPD ON TDPD.PROCESS_ID = TDP.ID
        LEFT JOIN TBL_MD_COLUMN ITMC ON TDPD.COLUMN_ID = ITMC.ID
        LEFT JOIN TBL_MD_COLUMN_DETAILS ITMCD ON ITMCD.COLUMN_ID = ITMC.ID
        LEFT JOIN TBL_MD_TABLE ITMT ON ITMC.TABLE_ID = ITMT.ID
        LEFT JOIN TBL_MD_RESOURCE ITMR ON ITMR.ID = ITMT.RESOURCE_ID
        LEFT JOIN TBL_SYSTEM ISYS ON ISYS.ID = ITMR.SYSTEM_ID AND TDPD.IMM_PREC_SYSTEM_ID = ISYS.ID
        LEFT JOIN TBL_LINK_COLUMN_BUSINESS_TERM CBT ON CBT.COLUMN_ID = ITMC.ID
        LEFT JOIN TBL_BUSINESS_TERM TBT ON CBT.BUSINESS_TERM_ID = TBT.ID
    WHERE
        TDP.NAME = '?' -- DOWNSTREAM PROCESS NAME
        AND UPPER(ITMCD.ALIAS_NAME) = UPPER('?') --CDE NAME