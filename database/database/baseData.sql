delete from user;
insert into user (id, username, password, full_name, gender, qq, tel, postcode,address, email, role_id, dept_id, active, locked, create_user,create_date, update_user, update_date, version)
values(1,'admin','b0baee9d279d34fa1dfd71aadb908c3f','Admin',1,111111,'11122233344','123456','长江西路130号','admin@easykoo.com',1,1,1,0,'SYSTEM',now(),'SYSTEM',now(),1);
insert into user (id, username, password, full_name, gender, qq, tel, postcode,address, email, role_id, dept_id, active, locked, create_user,create_date, update_user, update_date, version)
values(2,'manager','b0baee9d279d34fa1dfd71aadb908c3f','Manager',1,111111,'11122233344','123456','长江西路130号','manager@easykoo.com',2,1,1,0,'SYSTEM',now(),'SYSTEM',now(),1);
insert into user (id, username, password, full_name, gender, qq, tel, postcode,address, email, role_id, dept_id, active, locked, create_user,create_date, update_user, update_date, version)
values(3,'employee','b0baee9d279d34fa1dfd71aadb908c3f','Employee',1,111111,'11122233344','123456','长江西路130号','employee@easykoo.com',3,1,1,0,'SYSTEM',now(),'SYSTEM',now(),1);
insert into user (id, username, password, full_name, gender, qq, tel, postcode,address, email, role_id, dept_id, active, locked, create_user,create_date, update_user, update_date, version)
values(4,'user','b0baee9d279d34fa1dfd71aadb908c3f','User',1,111111,'11122233344','123456','长江西路130号','user@easykoo.com',4,1,1,0,'SYSTEM',now(),'SYSTEM',now(),1);

delete from role;
delete from department;
insert into role (id, description, create_user, create_date, update_user, update_date)values(1,'Admin','SYSTEM',now(),'SYSTEM',now());
insert into role (id, description, create_user, create_date, update_user, update_date)values(2,'Manager','SYSTEM',now(),'SYSTEM',now());
insert into role (id, description, create_user, create_date, update_user, update_date) values(3,'Employee','SYSTEM',now(),'SYSTEM',now());
insert into role (id, description, create_user, create_date, update_user, update_date) values(4,'User','SYSTEM',now(),'SYSTEM',now());
insert into department (id, description, create_user, create_date, update_user, update_date)values(1,'Default','SYSTEM',now(),'SYSTEM',now());

delete from module;
delete from privilege;
insert into module(id, description, create_user, create_date, update_user, update_date) values (1,'Admin','SYSTEM',now(),'SYSTEM',now());
insert into module(id, description, create_user, create_date, update_user, update_date) values (2,'Account','SYSTEM',now(),'SYSTEM',now());
insert into module(id, description, create_user, create_date, update_user, update_date) values (3,'Feedback','SYSTEM',now(),'SYSTEM',now());
insert into module(id, description, create_user, create_date, update_user, update_date) values (4,'News','SYSTEM',now(),'SYSTEM',now());
insert into module(id, description, create_user, create_date, update_user, update_date) values (5,'Product','SYSTEM',now(),'SYSTEM',now());

insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (1,1,1,'SYSTEM',now(),'SYSTEM',now());
insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (2,1,1,'SYSTEM',now(),'SYSTEM',now());
insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (2,2,1,'SYSTEM',now(),'SYSTEM',now());
insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (3,1,1,'SYSTEM',now(),'SYSTEM',now());
insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (3,2,1,'SYSTEM',now(),'SYSTEM',now());
insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (3,3,1,'SYSTEM',now(),'SYSTEM',now());
insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (4,1,1,'SYSTEM',now(),'SYSTEM',now());
insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (4,2,1,'SYSTEM',now(),'SYSTEM',now());
insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (4,3,1,'SYSTEM',now(),'SYSTEM',now());
insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (5,1,1,'SYSTEM',now(),'SYSTEM',now());
insert into privilege ( module_id, role_id,dept_id, create_user, create_date, update_user, update_date) values (5,2,1,'SYSTEM',now(),'SYSTEM',now());


