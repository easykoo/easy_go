DROP SCHEMA IF EXISTS easy_go;
CREATE SCHEMA easy_go
  DEFAULT CHARACTER SET utf8
  COLLATE utf8_unicode_ci;
USE easy_go;

DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id          INT(11)     NOT NULL AUTO_INCREMENT,
  username    VARCHAR(20) NOT NULL,
  password    VARCHAR(60) NOT NULL,
  full_name   VARCHAR(30) DEFAULT NULL,
  gender      INT(1) DEFAULT '1',
  qq          VARCHAR(16) DEFAULT NULL,
  tel         VARCHAR(20) DEFAULT NULL,
  postcode    VARCHAR(10) DEFAULT NULL,
  address     VARCHAR(80) DEFAULT NULL,
  email       VARCHAR(45) DEFAULT NULL,
  role_id     INT(3)      NOT NULL DEFAULT '3',
  dept_id     INT(3)      NOT NULL DEFAULT '1',
  active      TINYINT(1) DEFAULT '0',
  locked      TINYINT(1) DEFAULT '0',
  fail_time   INT(2) DEFAULT '0',
  create_user VARCHAR(20) DEFAULT NULL,
  create_date DATETIME DEFAULT NULL,
  update_user VARCHAR(20) DEFAULT NULL,
  update_date DATETIME DEFAULT NULL,
  version     INT(11) DEFAULT 1,
  PRIMARY KEY (id),
  UNIQUE KEY username_UNIQUE (username),
  UNIQUE KEY email_UNIQUE (email)
);

DROP TABLE IF EXISTS role;
CREATE TABLE role (
  id          INT(3)      NOT NULL AUTO_INCREMENT,
  description VARCHAR(20) NOT NULL,
  create_user VARCHAR(20) DEFAULT NULL,
  create_date DATETIME DEFAULT NULL,
  update_user VARCHAR(20) DEFAULT NULL,
  update_date DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY description_UNIQUE (description)
);

DROP TABLE IF EXISTS department;
CREATE TABLE department (
  id          INT(3)      NOT NULL AUTO_INCREMENT,
  description VARCHAR(20) NOT NULL,
  create_user VARCHAR(20) DEFAULT NULL,
  create_date DATETIME DEFAULT NULL,
  update_user VARCHAR(20) DEFAULT NULL,
  update_date DATETIME DEFAULT NULL,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS module;
CREATE TABLE module (
  id          INT(3)      NOT NULL,
  description VARCHAR(40) NOT NULL,
  create_user VARCHAR(20) DEFAULT NULL,
  create_date DATETIME DEFAULT NULL,
  update_user VARCHAR(20) DEFAULT NULL,
  update_date DATETIME DEFAULT NULL,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS privilege;
CREATE TABLE privilege (
  module_id   INT(11) NOT NULL,
  role_id     INT(11) NOT NULL,
  dept_id     INT(11) NOT NULL,
  create_user VARCHAR(20) DEFAULT NULL,
  create_date DATETIME DEFAULT NULL,
  update_user VARCHAR(20) DEFAULT NULL,
  update_date DATETIME DEFAULT NULL
);

DROP TABLE IF EXISTS feedback;
CREATE TABLE feedback (
  id          INT(11)      NOT NULL AUTO_INCREMENT,
  email       VARCHAR(45)  NOT NULL,
  name        VARCHAR(20)  NOT NULL,
  content     VARCHAR(200) NOT NULL,
  viewed      TINYINT(1) DEFAULT '0',
  create_date DATETIME DEFAULT NULL,
  view_date   DATETIME DEFAULT NULL,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS news;
CREATE TABLE news (
  id          INT(11)       NOT NULL AUTO_INCREMENT,
  title       VARCHAR(60)   NOT NULL,
  content     VARCHAR(2000) NOT NULL,
  lang        VARCHAR(5)    NOT NULL,
  priority    INT(1)        NULL DEFAULT 5,
  create_user VARCHAR(20) DEFAULT NULL,
  create_date DATETIME DEFAULT NULL,
  update_user VARCHAR(20) DEFAULT NULL,
  update_date DATETIME DEFAULT NULL,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS blog;
CREATE TABLE blog (
  id          INT(11)     NOT NULL AUTO_INCREMENT,
  title       VARCHAR(60) NOT NULL,
  content     BLOB        NOT NULL,
  state      VARCHAR(10) NOT NULL,
  priority    INT(1)      NULL DEFAULT 5,
  publish_date DATETIME DEFAULT NULL,
  create_user VARCHAR(20) DEFAULT NULL,
  create_date DATETIME DEFAULT NULL,
  update_user VARCHAR(20) DEFAULT NULL,
  update_date DATETIME DEFAULT NULL,
  version     INT(11) DEFAULT 1,
  PRIMARY KEY (id)
);