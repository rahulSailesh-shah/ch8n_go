-- +goose Up
-- +goose StatementBegin
ALTER TABLE node ALTER COLUMN id DROP DEFAULT;
-- Keep connection with auto-generation
ALTER TABLE connection ALTER COLUMN id SET DEFAULT uuid_generate_v4();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
