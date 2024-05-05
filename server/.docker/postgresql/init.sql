CREATE TABLE records (
    id bigint primary key,
    short_url_key varchar(255),
    long_url varchar(255),
    accesses int,
    last_access timestamp,
    created_at timestamp
);