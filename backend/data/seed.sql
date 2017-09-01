/* 

Data for Development


run: psql -d telehealth-app -f seed.sql

*/

insert into user_account(email, password) values('jndeutsch@gmail.com', '823d289b1c06a79bf86eccf2e6154a524cde958e');
insert into user_account(email, password) values('mary@gmail.com', '823d289b1c06a79bf86eccf2e6154a524cde958e');
insert into user_account(email, password) values('david@gmail.com', '823d289b1c06a79bf86eccf2e6154a524cde958e');
insert into user_account(email, password) values('tina@gmail.com', '823d289b1c06a79bf86eccf2e6154a524cde958e');
insert into user_account(email, password) values('baltizar@gmail.com', '823d289b1c06a79bf86eccf2e6154a524cde958e');



insert into patient(id, first_name, last_name, state, country) values(1,'Jason', 'Deutsch', 'Austin', 'TX');
insert into patient(id, first_name, last_name, state, country) values(2,'Mary', 'Smith', 'Austin', 'TX');
insert into patient(id, first_name, last_name, state, country) values(3,'David', 'Fez', 'Austin', 'TX');



