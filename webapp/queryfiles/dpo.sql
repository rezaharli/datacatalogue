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