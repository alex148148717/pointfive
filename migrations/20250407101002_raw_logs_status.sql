-- +goose Up
-- +goose StatementBegin
CREATE TYPE import_job_status AS ENUM ('import', 'upload', 'insert', 'ready');


CREATE TABLE import_job_files
(
    id     SERIAL PRIMARY KEY,
    path TEXT NOT NULL,
    time TIMESTAMP NOT NULL DEFAULT now(),
    status import_job_status NOT NULL DEFAULT 'import'
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS players;
-- +goose StatementEnd
