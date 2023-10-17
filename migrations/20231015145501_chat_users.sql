-- +goose Up
-- +goose StatementBegin
create table chat_users
(
    id         serial primary key,
    username   text      not null,
    chat_id    int,
    created_at timestamp not null default now(),
    updated_at timestamp,
    foreign key (chat_id) references chats (id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chat_users;
-- +goose StatementEnd
