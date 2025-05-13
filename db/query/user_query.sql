-- name: ListUsers :many
SELECT 
    * 
FROM 
    user 
LIMIT 10;

-- name: FindUserById :one
SELECT
    *
FROM 
    user
WHERE
    user_id = ?
LIMIT 1;

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

-- name: UpdateUserById :exec
UPDATE user
SET
    email = ?,
    password = ?
WHERE
    user_id = ?;

-- name: DeleteUserById :exec
DELETE FROM user
WHERE user_id = ?;

-- name: UserExistsById :one
SELECT EXISTS (
    SELECT 1 FROM user WHERE user_id = ?
);

-- name: UserExistsByEmail :one
SELECT EXISTS (
    SELECT 1 FROM user WHERE email = ?
);
