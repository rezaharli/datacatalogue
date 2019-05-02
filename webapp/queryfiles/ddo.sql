-- name: ddo-view
SELECT DISTINCT
        TC.NAME                                     				AS DATA_DOMAIN,
        TSC.NAME                                    				AS SUB_DOMAINS,
        TP.FIRST_NAME||' '||TP.LAST_NAME            				AS SUB_DOMAIN_OWNER,
        TP.BANK_ID                                  				AS BANK_ID,
        COUNT(DISTINCT TC.NAME) OVER ()         					AS COUNT_DATA_DOMAIN,
        COUNT(DISTINCT TSC.NAME) OVER ()         					AS COUNT_SUB_DOMAINS,      
        COUNT(DISTINCT TP.FIRST_NAME||' '||TP.LAST_NAME) OVER ()    AS COUNT_SUB_DOMAIN_OWNER,
        COUNT(DISTINCT TP.BANK_ID) OVER ()             				AS COUNT_BANK_ID
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TSC.ID AND UPPER(TLRP.OBJECT_TYPE) = 'SUBCATEGORY'
        INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DATA DOMAIN OWNER'
        INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
    WHERE ('ALL' = '?' OR TP.BANK_ID = '?')
        AND (
            upper(NVL(TC.NAME, ' ')) LIKE upper('%?%')
            AND upper(NVL(TSC.NAME, ' ')) LIKE upper('%?%')
            AND upper(NVL(TP.FIRST_NAME||' '||TP.LAST_NAME, ' ')) LIKE upper('%?%')
            AND upper(NVL(TP.BANK_ID, ' ')) LIKE upper('%?%')
        )
    ORDER BY TC.NAME, TSC.NAME

-- name: ddo-view-homepage
SELECT BT_COUNT, BT_SYSTEMS_COUNT, DP_COUNT
    FROM 
    (
    SELECT 
        COUNT(DISTINCT BT.BT_NAME) AS BT_COUNT
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        INNER JOIN TBL_BUSINESS_TERM BT ON TSC.ID = BT.PARENT_ID
    WHERE TSC.NAME = '?' -- SUBCATEGORY NAME
    ) BT, 
    (
    SELECT 
        COUNT(DISTINCT SYS.SYSTEM_NAME) AS BT_SYSTEMS_COUNT
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        INNER JOIN TBL_BUSINESS_TERM BT ON TSC.ID = BT.PARENT_ID
        INNER JOIN TBL_LINK_COLUMN_BUSINESS_TERM CBT ON BT.ID = CBT.BUSINESS_TERM_ID
        INNER JOIN TBL_MD_COLUMN COL ON CBT.COLUMN_ID = COL.ID
        INNER JOIN TBL_MD_TABLE TAB ON COL.TABLE_ID = TAB.ID
        INNER JOIN TBL_MD_RESOURCE RES ON TAB.RESOURCE_ID = RES.ID
        INNER JOIN TBL_SYSTEM SYS ON RES.SYSTEM_ID = SYS.ID
    WHERE TSC.NAME = '?' -- SUBCATEGORY NAME
    ) BT_SYSTEMS,
    (
    SELECT 
        COUNT(DISTINCT DP.NAME) AS DP_COUNT
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        INNER JOIN TBL_BUSINESS_TERM BT ON TSC.ID = BT.PARENT_ID
        INNER JOIN TBL_LINK_COLUMN_BUSINESS_TERM CBT ON BT.ID = CBT.BUSINESS_TERM_ID
        INNER JOIN TBL_MD_COLUMN COL ON CBT.COLUMN_ID = COL.ID
        INNER JOIN TBL_DS_PROCESS_DETAIL DPD ON COL.ID = DPD.COLUMN_ID
        INNER JOIN TBL_DS_PROCESSES DP ON DP.ID = DPD.PROCESS_ID
    WHERE TSC.NAME = '?' -- SUBCATEGORY NAME
    ) DP

-- name: ddo-businessterm
SELECT 
        TSC.ID                                      AS TSCID,
        TC.NAME                                     AS DATA_DOMAIN,
        TSC.NAME                                    AS SUB_DOMAINS,
        TP.FIRST_NAME||' '||TP.LAST_NAME            AS SUB_DOMAIN_OWNER,
        TP.BANK_ID                                  AS BANK_ID,
        COUNT(DISTINCT BT.BT_NAME) OVER ()          AS BT_COUNT,
        BT.BT_NAME                                  AS BT_NAME,
        BT.DESCRIPTION                              AS BT_DESCRIPTION
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TSC.ID AND UPPER(TLRP.OBJECT_TYPE) = 'SUBCATEGORY' AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DATA DOMAIN OWNER'
        INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
        INNER JOIN TBL_BUSINESS_TERM BT ON TSC.ID = BT.PARENT_ID
    WHERE TSC.NAME = '?' -- SUBCATEGORY NAME
        AND (
            upper(NVL(BT.BT_NAME, ' ')) LIKE upper('%?%')
            AND upper(NVL(BT.DESCRIPTION,' ')) LIKE upper('%?%')
        )

-- name: ddo-systems
SELECT DISTINCT
        TC.NAME                                                         AS DATA_DOMAIN,
        TSC.NAME                                                        AS SUB_DOMAINS,
        TP.FIRST_NAME||' '||TP.LAST_NAME                                AS SUB_DOMAIN_OWNER,
        COUNT(DISTINCT SYS.SYSTEM_NAME) OVER ()                         AS BT_SYSTEMS_COUNT,
        SYS.SYSTEM_NAME                                                 AS SYSTEM_NAME, 
        COUNT(DISTINCT BT.BT_NAME) OVER (PARTITION BY SYS.SYSTEM_NAME)  AS BT_COUNT
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        INNER JOIN TBL_BUSINESS_TERM BT ON TSC.ID = BT.PARENT_ID
        LEFT OUTER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TSC.ID AND UPPER(TLRP.OBJECT_TYPE) = 'SUBCATEGORY' AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        LEFT OUTER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DATA DOMAIN OWNER'
        LEFT OUTER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
        LEFT OUTER JOIN TBL_LINK_COLUMN_BUSINESS_TERM CBT ON BT.ID = CBT.BUSINESS_TERM_ID
        LEFT OUTER JOIN TBL_MD_COLUMN COL ON CBT.COLUMN_ID = COL.ID
        LEFT OUTER JOIN TBL_MD_TABLE TAB ON COL.TABLE_ID = TAB.ID
        LEFT OUTER JOIN TBL_MD_RESOURCE RES ON TAB.RESOURCE_ID = RES.ID
        LEFT OUTER JOIN TBL_SYSTEM SYS ON RES.SYSTEM_ID = SYS.ID
    WHERE TSC.NAME = '?' -- SUBCATEGORY NAME
        AND (
            upper(NVL(SYS.SYSTEM_NAME, ' ')) LIKE upper('%?%')
        )

-- name: ddo-systems-businessterm
SELECT DISTINCT
        TSC.ID                                      AS TSCID,
        TC.NAME                                     AS DATA_DOMAIN,
        TSC.NAME                                    AS SUB_DOMAINS,
        TP.FIRST_NAME||' '||TP.LAST_NAME            AS SUB_DOMAIN_OWNER,
        COUNT(DISTINCT BT.BT_NAME) OVER ()          AS BT_COUNT,
        BT.BT_NAME                                  AS BT_NAME,
        TAB.ID                                      AS TMTID,
        TAB.NAME                                    AS TABLE_NAME,
        COL.ID                                      AS COLID,
        COL.NAME                                    AS COLUMN_NAME
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        INNER JOIN TBL_BUSINESS_TERM BT ON TSC.ID = BT.PARENT_ID
        LEFT OUTER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TSC.ID AND UPPER(TLRP.OBJECT_TYPE) = 'SUBCATEGORY' AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        LEFT OUTER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DATA DOMAIN OWNER'
        LEFT OUTER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
        LEFT OUTER JOIN TBL_LINK_COLUMN_BUSINESS_TERM CBT ON BT.ID = CBT.BUSINESS_TERM_ID
        LEFT OUTER JOIN TBL_MD_COLUMN COL ON CBT.COLUMN_ID = COL.ID
        LEFT OUTER JOIN TBL_MD_TABLE TAB ON COL.TABLE_ID = TAB.ID
        LEFT OUTER JOIN TBL_MD_RESOURCE RES ON TAB.RESOURCE_ID = RES.ID
        LEFT OUTER JOIN TBL_SYSTEM SYS ON RES.SYSTEM_ID = SYS.ID
    WHERE TSC.NAME = '?' -- SUBCATEGORY NAME
        AND SYS.SYSTEM_NAME = '?' -- SYSTEM NAME
        AND (
            upper(NVL(BT.BT_NAME, ' ')) LIKE upper('%?%')
            AND upper(NVL(TAB.NAME, ' ')) LIKE upper('%?%')
            AND upper(NVL(COL.NAME, ' ')) LIKE upper('%?%')
        )
    ORDER BY BT.BT_NAME, TAB.NAME, COL.NAME

-- name: details
SELECT 
        TSC.ID,
        TC.NAME                                                     AS DATA_DOMAIN,
        TSC.NAME                                                    AS SUB_DOMAINS,
        TP.FIRST_NAME||' '||TP.LAST_NAME                            AS SUB_DOMAIN_OWNER,
        TP.BANK_ID                                                  AS BANK_ID,
        BT.BT_NAME                                                  AS BT_NAME,
        BT.DESCRIPTION                                              AS BT_DESCRIPTION,
        BT.POLICY_GUIDANCE                                          AS POLICY_GUIDANCE,
        BT.DQ_STANDARDS||' | '||BT.THRESHOLD                        AS BUSINESS_RULES,
        BT.MANDATORY                                                AS MANDATORY,
        GS.ITAM_ID                                                  AS GS_ITAM_ID,
        GS.SYSTEM_NAME                                              AS GS_SYSTEM_NAME,
        TGS.ITAM_ID                                                 AS TGS_ITAM_ID,
        TGS.SYSTEM_NAME                                             AS TGS_SYSTEM_NAME,
        ' '                                                         AS REMARKS,
        S.ITAM_ID                                                   AS ITAM_ID,
        S.SYSTEM_NAME                                               AS SYSTEM_NAME,
        T.NAME                                                      AS TABLE_NAME,
        C.NAME                                                      AS COLUMN_NAME,
        CD.ALIAS_NAME                                               AS ALIAS_NAME,
        CD.CDE                                                      AS CDE,
        CASE WHEN S.ITAM_ID = GS.ITAM_ID THEN 'YES' ELSE 'NO' END   AS GOLDEN_SOURCE,
        CD.PII_FLAG                                                 AS PII_FLAG,
        ' '                                                         AS CIA_RATING,
        CD.RECORD_CATEGORY                                          AS RECORD_CATEGORY
    FROM
        TBL_CATEGORY TC
        INNER JOIN TBL_SUBCATEGORY TSC ON TSC.CATEGORY_ID = TC.ID AND UPPER(TSC.TYPE) = 'SUB DATA DOMAIN'
        INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TSC.ID AND UPPER(TLRP.OBJECT_TYPE) = 'SUBCATEGORY' 
        INNER JOIN TBL_ROLE RL ON TLRP.ROLE_ID = RL.ID AND UPPER(RL.ROLE_NAME) = 'DATA DOMAIN OWNER'
        INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID
        INNER JOIN TBL_BUSINESS_TERM BT ON TSC.ID = BT.PARENT_ID
        LEFT OUTER JOIN TBL_LINK_COLUMN_BUSINESS_TERM CBT ON BT.ID = CBT.BUSINESS_TERM_ID
        LEFT OUTER JOIN TBL_MD_COLUMN C ON C.ID = CBT.COLUMN_ID
        LEFT OUTER JOIN TBL_MD_COLUMN_DETAILS CD ON C.ID = CD.COLUMN_ID
        LEFT OUTER JOIN TBL_MD_TABLE T ON T.ID = C.TABLE_ID
        LEFT OUTER JOIN TBL_MD_RESOURCE R ON R.ID = T.RESOURCE_ID
        LEFT OUTER JOIN TBL_SYSTEM S ON S.ID = R.SYSTEM_ID
        LEFT OUTER JOIN TBL_LINK_COLUMN_GOLDEN_SOURCE CGS ON CBT.COLUMN_ID = CGS.COLUMN_ID
        LEFT OUTER JOIN TBL_MD_COLUMN GSC ON GSC.ID = CGS.GOLDEN_SOURCE_COLUMN_ID
        LEFT OUTER JOIN TBL_MD_TABLE GST ON GST.ID = GSC.TABLE_ID
        LEFT OUTER JOIN TBL_MD_RESOURCE GSR ON GSR.ID = GST.RESOURCE_ID
        LEFT OUTER JOIN TBL_SYSTEM GS ON GS.ID = GSR.SYSTEM_ID
        LEFT OUTER JOIN TBL_SYSTEM TGS ON TGS.ID = BT.TARGET_GOLDEN_SOURCE_ID
    WHERE TSC.NAME = '?' -- SUBCATEGORY NAME
        AND BT.BT_NAME = '?' -- BUSINESS TERM NAME