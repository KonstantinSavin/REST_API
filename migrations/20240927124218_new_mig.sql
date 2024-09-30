-- Active: 1727443958061@@localhost@5432@music_lib_db
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS songs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    song_name VARCHAR NOT NULL,
    group_name VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS songs;
-- +goose StatementEnd
-- goose -dir migrations postgres "postgres://effmob:effmob@localhost/music_lib_db?sslmode=disable" up