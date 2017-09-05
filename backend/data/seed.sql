/* 

Data for Development


run: psql -d telehealth-app -f seed.sql

*/

insert into user_account(email, password, role) values('a@email.com', '823d289b1c06a79bf86eccf2e6154a524cde958e', '{"patient"}');
insert into user_account(email, password, role) values('b@email.com', '823d289b1c06a79bf86eccf2e6154a524cde958e', '{"patient"}');
insert into user_account(email, password, role) values('c@email.com', '823d289b1c06a79bf86eccf2e6154a524cde958e', '{"patient"}');
insert into user_account(email, password, role) values('d@email.com', '823d289b1c06a79bf86eccf2e6154a524cde958e', '{"patient"}');
insert into user_account(email, password, role) values('e@email.com', '823d289b1c06a79bf86eccf2e6154a524cde958e', '{"patient"}');
insert into user_account(email, password) values('f@email.com', '823d289b1c06a79bf86eccf2e6154a524cde958e');
insert into user_account(email, password) values('g@email.com', '823d289b1c06a79bf86eccf2e6154a524cde958e');
insert into user_account(email, password) values('h@email.com', '823d289b1c06a79bf86eccf2e6154a524cde958e');
insert into user_account(email, password) values('i@email.com', '823d289b1c06a79bf86eccf2e6154a524cde958e');
insert into user_account(email, password) values('j@email.com', '823d289b1c06a79bf86eccf2e6154a524cde958e');
insert into user_account(email, password, role) values('k@email.com', '5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8','{"admin"}');  
/* ^^^ password is password */



insert into patient(id, first_name, last_name, state, country) values(1,'Jason', 'Deutsch', 'Austin', 'TX');
insert into patient(id, first_name, last_name, state, country) values(2,'Mary', 'Smith', 'Austin', 'TX');
insert into patient(id, first_name, last_name, state, country) values(3,'David', 'Fez', 'Austin', 'TX');
insert into patient(id, first_name, last_name, state, country) values(4,'Jason', 'Deutsch', 'Austin', 'TX');
insert into patient(id, first_name, last_name, state, country) values(5,'Mary', 'Smith', 'Austin', 'TX');


