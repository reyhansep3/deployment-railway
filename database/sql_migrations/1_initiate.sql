-- +migrate Up
-- +migrate StatementBegin

create table bioskop(
    id SERIAL PRIMARY KEY,
    nama VARCHAR(60) NOT NULL,
    lokasi VARCHAR(100),
    rating FLOAT NOT NULL

)

-- +migrate StatementEnd