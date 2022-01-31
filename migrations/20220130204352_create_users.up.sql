
CREATE TABLE users (
    id bigserial NOT NULL primary key,
    email varchar NOT NULL unique,
    encrypted_password varchar NOT NULL
);