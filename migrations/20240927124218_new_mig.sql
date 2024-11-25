-- Active: 1727443958061@@localhost@5432@music_lib_db
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS group_names ( 
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY, 
    group_name VARCHAR NOT NULL UNIQUE 
);

CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    song_name VARCHAR NOT NULL,
    group_id INT NOT NULL,
    release_date VARCHAR,
    text VARCHAR,
    link VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS songs;
-- +goose StatementEnd
-- goose -dir migrations postgres "postgres://effmob:effmob@localhost/ml_db?sslmode=disable" up