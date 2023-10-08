create table if not exists
  credentials (
    uid uuid primary key,
    username text unique not null,
    password_hash bytea not null
  )
;
