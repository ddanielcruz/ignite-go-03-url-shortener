CREATE TABLE authors (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    bio TEXT
);

---- create above / drop below ----

DROP TABLE authors;
