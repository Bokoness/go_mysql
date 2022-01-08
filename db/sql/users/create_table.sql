DROP TABLE IF EXISTS users;

CREATE TABLE if not exists users (
	id INT NOT NULL AUTO_INCREMENT,
	username VARCHAR(20),
	password VARCHAR(100),
	PRIMARY KEY (id)
);