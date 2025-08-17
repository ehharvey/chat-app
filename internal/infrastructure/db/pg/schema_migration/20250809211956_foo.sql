-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS foo (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE foo;
-- +goose StatementEnd
