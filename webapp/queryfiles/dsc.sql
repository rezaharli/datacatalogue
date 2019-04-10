-- name: dsc-view
SELECT * 
	FROM (
		SELECT * 
			FROM (
				SELECT res.*, 
						COUNT(DISTINCT system_name) OVER ()			as COUNT_SYSTEM_NAME,
						COUNT(DISTINCT itam_id) OVER ()				as COUNT_ITAM_ID,
						COUNT(DISTINCT dataset_custodian) OVER ()	as COUNT_DATASET_CUSTODIAN,
						COUNT(DISTINCT bank_id) OVER ()				as COUNT_BANK_ID
					FROM (
                        SELECT DISTINCT
                            TS.ID
                            ,TS.SYSTEM_NAME 						AS SYSTEM_NAME
                            ,TS.ITAM_ID								AS ITAM_ID
                            ,TP.FIRST_NAME||' '||TP.LAST_NAME 		AS DATASET_CUSTODIAN
                            ,TP.BANK_ID								AS BANK_ID
                        FROM 
                            TBL_SYSTEM TS 
                            INNER JOIN TBL_LINK_ROLE_PEOPLE TLRP ON TLRP.OBJECT_ID = TS.ID AND TLRP.OBJECT_TYPE = 'SYSTEM'
                            INNER JOIN TBL_ROLE RL_SYS ON TLRP.ROLE_ID = RL_SYS.ID AND UPPER(RL_SYS.ROLE_NAME) = 'DATASET CUSTODIAN'
                            INNER JOIN TBL_PEOPLE TP ON TLRP.PEOPLE_ID = TP.ID 
                        WHERE ('ALL' = '?' OR TP.BANK_ID = '?')
                        ORDER BY TS.SYSTEM_NAME
					) res
			) WHERE ( -- Main and dropdown search
				upper(system_name) LIKE upper('%?%')
				AND upper(itam_id) LIKE upper('%?%')
			)
	) WHERE ( -- Column filter
		upper(system_name) LIKE upper('%?%')
		AND upper(itam_id) LIKE upper('%?%')
		AND upper(dataset_custodian) LIKE upper('%?%')
		AND upper(bank_id) LIKE upper('%?%')
	)