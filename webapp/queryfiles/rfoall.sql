-- name: left-grid
SELECT * 
	FROM (
		SELECT * 
			FROM (
				SELECT res.*, 
						COUNT(DISTINCT PRIORITY_REPORT) OVER () COUNT_PRIORITY_REPORT,
						COUNT(DISTINCT RR_LEAD) OVER () COUNT_RR_LEAD,
						COUNT(DISTINCT BANK_ID) OVER () COUNT_BANK_ID
					FROM (
						SELECT DISTINCT
								PR.id,
								PR.NAME                               AS PRIORITY_REPORT, 
								PL.FIRST_NAME||' '|| PL.LAST_NAME     AS RR_LEAD,
								PL.BANK_ID                            AS BANK_ID
							FROM TBL_PRIORITY_REPORTS PR 
							INNER JOIN TBL_PEOPLE PL ON PR.LEAD_ID = PL.ID
					) res
			) WHERE ( -- Main and dropdown search
				upper(PRIORITY_REPORT) LIKE upper('%?%')
				AND upper(RR_LEAD) LIKE upper('%?%')
			)
	) WHERE ( -- Column filter
		upper(PRIORITY_REPORT) LIKE upper('%?%')
		AND upper(RR_LEAD) LIKE upper('%?%')
		AND upper(BANK_ID) LIKE upper('%?%')
	)

-- name: right-grid
SELECT * 
	FROM (
		SELECT * 
			FROM (
				SELECT res.*, 
						COUNT(DISTINCT PRINCIPAL_RISK) OVER () COUNT_PRINCIPAL_RISK,
						COUNT(DISTINCT RISK_SUB) OVER () COUNT_RISK_SUB,
						COUNT(DISTINCT PR_NAME) OVER () COUNT_PR_NAME,
						COUNT(DISTINCT PR_RATIONALE) OVER () COUNT_PR_RATIONALE,
						COUNT(DISTINCT CRM_NAME) OVER () COUNT_CRM_NAME,
						COUNT(DISTINCT CRM_RATIONALE) OVER () COUNT_CRM_RATIONALE,
						COUNT(DISTINCT ASSOC_CDES) OVER () COUNT_ASSOC_CDES,
						COUNT(DISTINCT CDE_RATIONALE) OVER () COUNT_CDE_RATIONALE
					FROM (
						SELECT DISTINCT
								CAT.id,
								PR.id 				as tprid,
								CAT.NAME 			AS PRINCIPAL_RISK,
								SC.NAME 			AS RISK_SUB,
								PR.NAME 			AS PR_NAME,
								PR.RATIONALE 		AS PR_RATIONALE,
								CRM.NAME 			AS CRM_NAME,
								CRM.CRM_RATIONALE	as CRM_RATIONALE,
								BT.BT_NAME 			AS ASSOC_CDES,
								BT.CDE_RATIONALE	AS CDE_RATIONALE
							FROM TBL_PRIORITY_REPORTS PR 
								INNER JOIN TBL_PEOPLE PL ON PR.LEAD_ID = PL.ID
								INNER JOIN TBL_SUBCATEGORY SC ON PR.SUB_RISK_TYPE_ID = SC.ID
								INNER JOIN TBL_CATEGORY CAT ON SC.CATEGORY_ID = CAT.ID
								INNER JOIN TBL_CRM CRM ON PR.ID = CRM.PRORITY_REPORT_ID
								INNER JOIN TBL_LINK_CRM_CDE LNK ON LNK.CRM_ID = CRM.ID
								INNER JOIN TBL_BUSINESS_TERM BT ON LNK.CDE_ID = BT.ID
							WHERE 
								pr.id = '?'
							ORDER BY CRM.NAME, BT.BT_NAME
					) res
			) WHERE ( -- Main and dropdown search
				upper(PRINCIPAL_RISK) LIKE upper('%?%')
				AND upper(RISK_SUB) LIKE upper('%?%')
			)
	) WHERE ( -- Column filter
		upper(PRINCIPAL_RISK) LIKE upper('%?%')
		AND upper(RISK_SUB) LIKE upper('%?%')
		AND upper(PR_RATIONALE) LIKE upper('%?%')
		AND upper(CRM_NAME) LIKE upper('%?%')
		AND upper(CRM_RATIONALE) LIKE upper('%?%')
		AND upper(ASSOC_CDES) LIKE upper('%?%')
		AND upper(CDE_RATIONALE) LIKE upper('%?%')
	)