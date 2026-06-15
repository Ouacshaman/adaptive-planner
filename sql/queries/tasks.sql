-- name: CreateTask :one
INSERT INTO tasks (id, created_at, updated_at, task)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1
)
RETURNING *;
