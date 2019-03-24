-- name: left-grid
SELECT DISTINCT
		tpr.id,
		tpr.name		as PRIORITY_REPORT,
		tpr.owner_id	as RR_LEAD,
		tpr.owner_id	as BANK_ID
	FROM
		Tbl_Priority_Reports tpr
		JOIN tbl_subcategory tsc ON tpr.sub_risk_type_id = tsc.id
		JOIN tbl_category tc ON tsc.category_id = tc.id
		LEFT JOIN tbl_crm tcrm ON tcrm.prority_report_id = tpr.id
		LEFT JOIN tbl_link_crm_cde tlcc ON tlcc.crm_id = tcrm.id
		LEFT JOIN tbl_business_term tbt ON tlcc.cde_id = tbt.id

-- name: right-grid
SELECT DISTINCT
		tc.id,
		tpr.id 				as tprid,
		tc.name 			as PRINCIPAL_RISK,
		tsc.name 			as RISK_SUB,
		tpr.rationale 		as PR_RATIONALE,
		tcrm.name 			as CRM_NAME,
		tcrm.crm_rationale 	as CRM_RATIONALE,
		tlcc.cde_id 		as ASSOC_CDES,
		tbt.cde_rationale	as CDE_RATIONALE
	FROM
		Tbl_Priority_Reports tpr
		JOIN tbl_subcategory tsc ON tpr.sub_risk_type_id = tsc.id
		JOIN tbl_category tc ON tsc.category_id = tc.id
		LEFT JOIN tbl_crm tcrm ON tcrm.prority_report_id = tpr.id
		LEFT JOIN tbl_link_crm_cde tlcc ON tlcc.crm_id = tcrm.id
		LEFT JOIN tbl_business_term tbt ON tlcc.cde_id = tbt.id
	WHERE
		tpr.id = '?'