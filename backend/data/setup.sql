create table user_account(
  id            serial primary key,
  email         text,
  password      text, 
  disabled      bool default false,
  confirmed     bool default false, 
  role          text,
  created_at    timestamp default current_timestamp
);

alter table user_account
add constraint order_unique unique (email);

create table user_session(
  id             serial primary key, 
  uuid           text not null unique,
  user_id        int references user_account(id),
  login_time     timestamp not null
);


/* User Types */

create table provider(
  id           int primary key references user_account(id), 
  first_name   text,
  last_name    text,
  vidyo_room   text,
  phone_number text,
  credential   text[]
);

create table patient(
  id          int primary key references user_account(id),
  first_name  text not null,
  last_name   text not null,
  state       text not null,
  country     text not null,
  created_at  timestamp default current_timestamp
);


/* Patient/Provider Pairings */

create table pairing(
  id         serial primary key,
  patient    int references patient(id),	
  provider   int references provider(id),
  active     bool default true
  created_at timestamp default current_timestamp
);

/* Scheduling and Sessions */

create table availability(
  id          serial primary key,
  start_time  timestamp not null,
  end_time    timestamp not null,
  created_at  timestamp default current_timestamp
);

create table appointment(
  id          serial primary key,
  patinet_id  int refrences patient(id),
  provider_id int references provider(id),
  location    text 
  appt_day    date not null,
  start_time  time not null,
  duration    int not null, 
  cancelled   bool default false,
  completed   bool default false,
  created_at  timestamp default current_timestamp
);


/* Surveys, Forms and Questionnaires */

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
    user_id     int references user_account(id),
    answer      text
);
