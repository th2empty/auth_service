CREATE TABLE users
(
    id serial not null unique,
    username VARCHAR(255) not null,
    password_hash VARCHAR(255) not null
);