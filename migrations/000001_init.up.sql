CREATE SCHEMA todo_app;

CREATE TABLE todo_app.users (
    id              UUID                            PRIMARY KEY,
    version         BIGINT          NOT NULL    DEFAULT 1,
    name            VARCHAR(100)    NOT NULL,
    first_name      VARCHAR(100)    NOT NULL    CHECK(char_length(first_name) BETWEEN 3 AND 100),
    last_name       VARCHAR(100)    NOT NULL    CHECK(char_length(last_name) BETWEEN 3 AND 100),
    phone_number    VARCHAR(15)                 CHECK(
        phone_number ~ '^\+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    ),
    created_at      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE todo_app.tasks (
    id              UUID                        PRIMARY KEY,
    version         BIGINT          NOT NULL    DEFAULT 1,
    title           VARCHAR(100)    NOT NULL    CHECK (char_length(title) BETWEEN 1 AND 100),
    description     VARCHAR(1000)               CHECK (char_length(description) BETWEEN 1 AND 1000),
    completed       BOOLEAN         NOT NULL    DEFAULT False,
    completed_at    TIMESTAMP,
    created_at      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       NOT NULL    DEFAULT CURRENT_TIMESTAMP,

    CHECK (
        (completed = True AND completed_at IS NOT NULL and completed_at >= created_at)
        OR
        (completed = False AND completed_at IS NULL)
    ),

    author_user_id  UUID            NOT NULL    REFERENCES todo_app.users (id)
);