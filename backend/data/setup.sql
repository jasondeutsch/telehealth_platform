create table user_account(
  id            serial primary key,
  email         text,
  password      text, 
  disabled      bool,
  created_at    timestamp default current_timestamp
);

create table user_session(
  id             serial primary key, 
  user_id        int references user_account(id),
  login_time     timestamp not null
);

create table role (
    id       int primary key references user_account(id),
    provider bool,
    patient  bool,
    admin    bool
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
  created_at  timestamp default current_timestamp
  );

create table availability(
  id          serial primary key,
  start_time  timestamp not null,
  end_time    timestamp not null,
  created_at  timestamp default current_timestamp
);

create table appointment(
  id         serial primary key,
  date       date not null,
  start_time time not null,
  end_time   time not null,
  created_at timestamp default current_timestamp
);

create table survey(
    id          serial primary key,
    name        text,
    description text
);

create table question(
    id            serial primary key,
    survey_id     int references survey(id),
    question_text text
);

create table answer(
    question_id int primary key references question(id),
    survey_id   int references survey(id),
    user_id        int references user_account(id),
    answer      text
);
