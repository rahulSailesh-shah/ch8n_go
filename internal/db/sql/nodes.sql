-- name: CreateNode :one
INSERT INTO node (workflow_id, name, type, position, data)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetNodeByID :one
SELECT * FROM node WHERE id = $1 AND workflow_id = $2;

-- name: GetNodesByWorkflowID :many
SELECT * FROM node WHERE workflow_id = $1;

-- name: ListNodes :many
SELECT * FROM node ORDER BY updated_at DESC;

-- name: UpdateNode :one
UPDATE node
SET name = $3, type = $4, position = $5, data = $6, updated_at = NOW()
WHERE id = $1 AND workflow_id = $2
RETURNING *;

-- name: DeleteNode :exec
DELETE FROM node WHERE id = $1 AND workflow_id = $2;
