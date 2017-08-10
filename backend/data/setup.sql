create table user_account(
  id            serial primary key,
  email         text,
  password_hash text,
  password_salt text,
  is_disabled   bool,
  created_at    timestamp not null
);

create table user_session(
  id             serial primary key, 
  user_id        int references user_account(id),
  login_time     timestamp not null,
  last_seen_time timestamp not null
);

create table provider(
  id          int primary key references user_account(id), 
  is_admin    bool not null,
  vidyo_room  text,
  credential  text
);

create table patient(
  id          int primary key references user_account(id),
  first_name  text not null,
  last_name   text not null,
  state       text not null,
  country     text not null,
  created_at  timestamp not null
);

create table availability(
  id          serial primary key,
  start_time  timestamp not null,
  end_time    timestamp not null,
  created_at  timestamp not null
);

create table appointment(
  id         serial primary key,
  date       date not null,
  start_time time not null,
  end_time   time not null,
  created_at timestamp not null
);

