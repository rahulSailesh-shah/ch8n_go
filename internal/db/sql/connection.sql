-- name: CreateConnection :one
INSERT INTO connection (workflow_id, source_node_id, target_node_id, from_output, to_input)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetConnectionByID :one
SELECT * FROM connection WHERE id = $1 AND workflow_id = $2;

-- name: GetConnectionsByWorkflowID :many
SELECT * FROM connection WHERE workflow_id = $1;

-- name: ListConnections :many
SELECT * FROM connection ORDER BY updated_at DESC;

-- name: UpdateConnection :one
UPDATE connection
SET source_node_id = $3, target_node_id = $4, updated_at = NOW()
WHERE id = $1 AND workflow_id = $2
RETURNING *;

-- name: DeleteConnection :exec
DELETE FROM connection WHERE id = $1 AND workflow_id = $2;

-- name: DeleteConnectionsByWorkflowID :exec
DELETE FROM connection WHERE workflow_id = $1;
