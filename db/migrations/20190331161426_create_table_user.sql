create table "user"
(
    user_id  serial not null
        constraint user_pk
            primary key,
    password varchar(250),
    username varchar(250),
    email    varchar(250)
);

create unique index "User_user-id_uindex"
    on "user" (user_id);

create unique index user_username_uindex
    on "user" (username);

create unique index user_email_uindex
    on "user" (email);
