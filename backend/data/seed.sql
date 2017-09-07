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

insert into provider(id, first_name, last_name, phone_number, vidyo_room, credential) values(6,'Amy', 'Pieczarka', '5555555555','', '{"RD"}');
insert into provider(id, first_name, last_name, phone_number, vidyo_room, credential) values(7,'Jane','Doe','555555555','','{"MD"}');
insert into provider(id, first_name, last_name, phone_number, vidyo_room, credential) values(8,'John','Smith','5555555555','','{"DO"}');
insert into provider(id, first_name, last_name, phone_number, vidyo_room, credential) values(9,'Thomas','Gekkly','5555555555','', '{"PhD"}');

insert into pairing(patient, provider) values(1,6);
insert into pairing(patient, provider) values(1,7);
insert into pairing(patient, provider) values(1,8);
insert into pairing(patient, provider) values(2,6);
insert into pairing(patient, provider) values(3,7);
