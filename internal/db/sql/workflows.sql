-- name: CreateWorkflow :one
INSERT INTO workflow (name, description, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetWorkflowByID :one
SELECT * FROM workflow WHERE id = $1 AND user_id = $2;

-- name: GetWorkflowsByUserID :many
SELECT id, name, description, user_id, created_at, updated_at, COUNT(*) OVER() as total_count
FROM workflow
WHERE user_id = $1
    AND (CASE WHEN $2::text != '' THEN name ILIKE '%' || $2 || '%' ELSE TRUE END)
ORDER BY updated_at DESC
LIMIT $3 OFFSET $4;

-- name: ListWorkflows :many
SELECT * FROM workflow ORDER BY updated_at DESC;



-- name: UpdateWorkflow :one
UPDATE workflow
SET name = $2, description = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteWorkflow :exec
DELETE FROM workflow WHERE id = $1;
