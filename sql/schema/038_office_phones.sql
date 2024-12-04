-- +goose Up
CREATE TABLE office_phones (
    office_id INTEGER,
    phones_id INTEGER,
    CONSTRAINT fk_offices
        FOREIGN KEY (office_id)
        REFERENCES offices(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_phones
        FOREIGN KEY (phones_id)
        REFERENCES phones(id)
        ON DELETE CASCADE,
    PRIMARY KEY (office_id, phones_id)
);

-- +goose Down
DROP TABLE office_phones;