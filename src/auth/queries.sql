-- name: get_by_uid^
select
  uid,
  email,
  password_hash,
  creation_date
from
  credentials
where
  uid = :uid
limit
  1
;


-- name: get_by_email^
select
  uid,
  email,
  password_hash,
  creation_date
from
  credentials
where
  email = :email
limit
  1
;


-- name: create<!
insert into
  credentials (uid, email, password_hash, creation_date)
values
  (
    gen_random_uuid (),
    :email,
    :password_hash,
    current_timestamp
  )
returning
  uid
;

-- name: create_uid<!
insert into
  credentials (uid, email, password_hash, creation_date)
values
  (
    :uid,
    :email,
    :password_hash,
    current_timestamp
  )
returning
  uid
;


-- name: update<!
update credentials
set
  email = :email,
  password_hash = :password_hash
where
  uid = :uid
returning
  uid
;


-- name: delete!
delete from credentials
where
  uid = :uid
;
