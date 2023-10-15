-- +goose Up
-- +goose StatementBegin
create table chat_users
(
    id         serial primary key,
    username   text      not null,
    chat_id    int,
    created_at timestamp not null default now(),
    updated_at timestamp,
    FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chat_users;
-- +goose StatementEnd
