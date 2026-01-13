-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id            INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    name          VARCHAR(255) NOT NULL,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE lists
(
    id          INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE users_lists
(
    id INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    user_id INT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    list_id INT REFERENCES lists (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE items
(
    id          INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done        BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE list_items
(
    id      INT GENERATED ALWAYS AS IDENTITY UNIQUE,
    item_id INT REFERENCES items (id) ON DELETE CASCADE NOT NULL,
    list_id INT REFERENCES lists (id) ON DELETE CASCADE NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE list_items;

DROP TABLE items;

DROP TABLE users_lists;

DROP TABLE lists;

DROP TABLE users;
-- +goose StatementEnd
