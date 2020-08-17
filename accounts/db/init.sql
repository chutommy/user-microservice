--  This is script run on the first database initialization.

CREATE TABLE IF NOT EXISTS accounts (
    id          UUID            PRIMARY KEY,
    username    VARCHAR(32)     UNIQUE,
    email       VARCHAR(255)    UNIQUE NOT NULL,
    phone       VARCHAR(33)     UNIQUE,
    hpassword   VARCHAR(128)    NOT NULL,

    first_name  VARCHAR(64),
    last_name   VARCHAR(64),
    birth_day   DATE,

    perm_address VARCHAR(255),
    mail_address VARCHAR(255),

    created_at  TIMESTAMP       NOT NULL,
    updated_at  TIMESTAMP       NOT NULL,
    deleted_at  TIMESTAMP
);
