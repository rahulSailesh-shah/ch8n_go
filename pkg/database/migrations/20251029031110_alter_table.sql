-- +goose Up
-- +goose StatementBegin
-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Convert workflow table
ALTER TABLE workflow ALTER COLUMN id DROP DEFAULT;
ALTER TABLE workflow ALTER COLUMN id TYPE UUID USING uuid_generate_v4();
ALTER TABLE workflow ALTER COLUMN id SET DEFAULT uuid_generate_v4();

-- Convert node table
ALTER TABLE node ALTER COLUMN id DROP DEFAULT;
ALTER TABLE node ALTER COLUMN id TYPE UUID USING uuid_generate_v4();
ALTER TABLE node ALTER COLUMN id SET DEFAULT uuid_generate_v4();

ALTER TABLE node ALTER COLUMN workflow_id TYPE UUID USING uuid_generate_v4();

-- Convert connection table
ALTER TABLE connection ALTER COLUMN id DROP DEFAULT;
ALTER TABLE connection ALTER COLUMN id TYPE UUID USING uuid_generate_v4();
ALTER TABLE connection ALTER COLUMN id SET DEFAULT uuid_generate_v4();

ALTER TABLE connection ALTER COLUMN workflow_id TYPE UUID USING uuid_generate_v4();
ALTER TABLE connection ALTER COLUMN source_node_id TYPE UUID USING uuid_generate_v4();
ALTER TABLE connection ALTER COLUMN target_node_id TYPE UUID USING uuid_generate_v4();

-- Add indexes
CREATE INDEX idx_node_workflow_id ON node(workflow_id);
CREATE INDEX idx_connection_workflow_id ON connection(workflow_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_connection_workflow_id;
DROP INDEX IF EXISTS idx_node_workflow_id;

ALTER TABLE connection ALTER COLUMN target_node_id TYPE INTEGER USING 0;
ALTER TABLE connection ALTER COLUMN source_node_id TYPE INTEGER USING 0;
ALTER TABLE connection ALTER COLUMN workflow_id TYPE INTEGER USING 0;
ALTER TABLE connection ALTER COLUMN id TYPE INTEGER USING 0;

ALTER TABLE node ALTER COLUMN workflow_id TYPE INTEGER USING 0;
ALTER TABLE node ALTER COLUMN id TYPE INTEGER USING 0;

ALTER TABLE workflow ALTER COLUMN id TYPE INTEGER USING 0;
-- +goose StatementEnd
