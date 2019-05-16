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
    ORDER BY TIME DESC