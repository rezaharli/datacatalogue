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