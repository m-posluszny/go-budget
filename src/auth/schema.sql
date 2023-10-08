create table if not exists
  credentials (
    uid uuid primary key,
    email text unique not null,
    password_hash text not null,
    creation_date timestamp
  )
;
