-- name: users
SELECT 
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