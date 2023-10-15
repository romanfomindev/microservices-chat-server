-- +goose Up
-- +goose StatementBegin
create table chats
(
    id         serial primary key,
    name       text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chats;
-- +goose StatementEnd
