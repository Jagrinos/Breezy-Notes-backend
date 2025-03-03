CREATE TABLE users
(
    id    VARCHAR(50) PRIMARY KEY,
    login VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(50)        NOT NULL,
    about TEXT
);