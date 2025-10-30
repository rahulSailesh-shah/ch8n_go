-- +goose Up
-- +goose StatementBegin
-- Remove the default UUID generation so we can provide our own IDs
ALTER TABLE node ALTER COLUMN id DROP DEFAULT;
ALTER TABLE connection ALTER COLUMN id DROP DEFAULT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Restore the default UUID generation
ALTER TABLE node ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE connection ALTER COLUMN id SET DEFAULT uuid_generate_v4();
-- +goose StatementEnd
