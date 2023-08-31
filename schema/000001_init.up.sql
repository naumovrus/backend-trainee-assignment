CREATE TABLE users
(
    id            serial       not null unique,
    name      varchar(255) not null unique
);

CREATE TABLE segments
(
    id          serial       not null unique,
    name       varchar(255) not null
);

CREATE TABLE users_segments
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    segment_id int references segments (id) on delete cascade not null
);

