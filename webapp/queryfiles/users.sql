-- name: users
SELECT 
        ID,
        USERNAME,
        EMAIL,
        NAME,
        ROLE,
        STATUS,
        CREATEDAT,
        UPDATEDAT 
    FROM TBL_USERS
    WHERE (
        UPPER(USERNAME) LIKE UPPER('%?%')
        AND UPPER(EMAIL) LIKE UPPER('%?%')
        AND UPPER(NAME) LIKE UPPER('%?%')
        AND UPPER(ROLE) LIKE UPPER('%?%')
        AND UPPER(STATUS) LIKE UPPER('%?%')
        AND UPPER(CREATEDAT) LIKE UPPER('%?%')
        AND UPPER(UPDATEDAT) LIKE UPPER('%?%')
    )

-- name: users-usage
SELECT 
        ID,
        USERNAME,
        FULLNAME,
        ROLE,
        MODULE,
        DESCRIPTION,
        TIME,
        RESOURCEURL 
    FROM TBL_USAGE
    WHERE (
        UPPER(USERNAME) LIKE UPPER('%?%')
        AND UPPER(FULLNAME) LIKE UPPER('%?%')
        AND UPPER(ROLE) LIKE UPPER('%?%')
        AND UPPER(MODULE) LIKE UPPER('%?%')
        AND UPPER(DESCRIPTION) LIKE UPPER('%?%')
        AND UPPER(TIME) LIKE UPPER('%?%')
        AND UPPER(RESOURCEURL) LIKE UPPER('%?%')
    )
    ORDER BY TIME DESC