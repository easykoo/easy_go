delete from user;
insert into user (id, username, password, full_name, gender, qq, tel, postcode,address, email, role_id, dept_id, active, locked, create_user,create_date, update_user, update_date, version)
values(1,'admin','b0baee9d279d34fa1dfd71aadb908c3f','Admin',1,111111,'11122233344','123456','长江西路130号','admin@easykoo.com',1,1,1,0,'SYSTEM',now(),'SYSTEM',now(),1);
insert into user (id, username, password, full_name, gender, qq, tel, postcode,address, email, role_id, dept_id, active, locked, create_user,create_date, update_user, update_date, version)
values(2,'user','b0baee9d279d34fa1dfd71aadb908c3f','User',1,111111,'11122233344','123456','长江西路130号','user@easykoo.com',4,1,1,0,'SYSTEM',now(),'SYSTEM',now(),1);

delete from category;
insert into category (id,  description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('101', 'Bamboo Crafts', '竹工艺品', null, 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('102', 'Christmas Tree', '圣诞树', null, 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('103', 'Crystal Crafts', '水晶工艺品', null, 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('104', 'Glass Crafts', '玻璃工艺品', null, 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('105', 'Porcelain Crafts', '陶瓷工艺品', null, 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('106', 'Power Bank Gifts', '移动电源', null, 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('107', 'Promotional Porcelain', '促销瓷器', null, 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('108', 'Resin Crafts', '树脂工艺品', null, 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('102001', 'Fiber optic Christmas tree', '光纤圣诞树', '102', 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('102002', 'Ordinary Christmas tree', '普通圣诞树', '102', 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('104001', 'Blown Glass Ornaments', '吹制玻璃饰品', '104', 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('104002', 'Glass Candle Holder', '玻璃罐', '104', 'SYSTEM', now(), 'SYSTEM', now());
insert into category (id, description, cn_description, parent_category, create_user, create_date, update_user, update_date)
values ('104003', 'Glass Crafts', '玻璃工艺品', '104', 'SYSTEM', now(), 'SYSTEM', now());

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


