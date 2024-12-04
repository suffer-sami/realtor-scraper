-- +goose Up
CREATE TABLE phones (
    id INTEGER PRIMARY KEY,
    ext TEXT,
    number TEXT,
    type TEXT,
    is_valid BOOLEAN,
    CONSTRAINT unique_phones_ext_and_number_and_type
        UNIQUE (ext, number, type)
);

-- +goose Down
DROP TABLE phones;
