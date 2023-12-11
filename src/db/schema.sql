create table if not exists
  credentials (
    uid uuid primary key,
    username text unique not null,
    password_hash bytea not null
  )
;

create table if not exists
  accounts (
    uid uuid primary key,
    user_uid uuid not null references credentials(uid),
    name text not null,
    offbudget boolean not null default false,
    creation_date date not null 

  )
;

create table if not exists
  groups (
    uid uuid primary key,
    user_uid uuid not null references credentials(uid),
    name text not null
  )
;

create table if not exists
  categories (
    uid uuid primary key,
    user_uid uuid not null references credentials(uid),
    name text not null,
    group_uid uuid references categories(uid)
  )
;

create table if not exists
  transactions (
    uid uuid primary key,
    account_uid uuid not null references accounts(uid),
    value numeric(10,2) not null,
    date date not null,
    memo text not null,
    payee text not null,
    category_uid uuid references categories(uid)
  )
;

create materialized view if not exists
  transactions_view as 
    select
      t.uid as uid,
      t.account_uid as account_uid,
      t.value as value,
      t.date as date,
      t.memo as memo,
      t.payee as payee,
      t.category_uid as category_uid,
      a.name as account_name,
      a.offbudget as offbudget,
      c.name as category_name
    from
      transactions t
      left join accounts a on t.account_uid = a.uid
      left join categories c on t.category_uid = c.uid
    ;

create materialized view if not exists
  accounts_balance as
    select
      a.uid as uid,
      a.user_uid as user_uid,
      a.name as name,
      a.offbudget as offbudget,
      sum(t.value) as balance
    from
      accounts a
      left join transactions t on t.account_uid = a.uid
    group by
      a.uid
    ;
  
create materialized view if not exists
  payees as
    select
      distinct payee
    from
      transactions
    left join
      accounts on transactions.account_uid = accounts.uid
    order by
      payee
    ;