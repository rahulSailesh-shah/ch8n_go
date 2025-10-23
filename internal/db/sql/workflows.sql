-- name: CreateWorkflow :one
INSERT INTO workflow (name, description, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetWorkflowByID :one
SELECT * FROM workflow WHERE id = $1;

-- name: GetWorkflowsByUserID :many
SELECT * FROM workflow where user_id = $1;

-- name: ListWorkflows :many
SELECT * FROM workflow ORDER BY id;

-- name: UpdateWorkflow :one
UPDATE workflow
SET name = $2, description = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteWorkflow :exec
DELETE FROM workflow WHERE id = $1;
