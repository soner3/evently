-- name: FindUserById :one
SELECT
    *
FROM 
    user
WHERE
    user_id = ?
LIMIT 1;

-- name: FindUserByIdWithReferences :many
SELECT
    u.*,
    e.*
FROM
    user u
LEFT JOIN
    event e
ON u.user_id = e.user_id
WHERE
    1 = 1
    AND u.user_id = ?;

-- name: FindUserByEmail :one
SELECT
    *
FROM 
    user
WHERE
    email = ?
LIMIT 1;

-- name: CreateUser :exec
INSERT INTO user (
    user_id,
    email,
    password
) VALUES (
    ?, ?, ?
);

-- name: DeleteUserById :exec
DELETE FROM user
WHERE user_id = ?;

-- name: UserExistsByEmail :one
SELECT EXISTS (
    SELECT 1 FROM user WHERE email = ?
);


