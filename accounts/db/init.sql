--  This scripts run on the database's initialization.

-- create a function which updates the updated_at timestamp
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- create a table
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

    created_at  TIMESTAMP       NOT NULL    DEFAULT NOW(),
    updated_at  TIMESTAMP       NOT NULL    DEFAULT NOW(),
    deleted_at  TIMESTAMP
);

-- create a trigger
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON accounts
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
