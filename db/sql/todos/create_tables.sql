DROP TABLE IF EXISTS todos;

CREATE TABLE if not exists todos (
	id SERIAL PRIMARY KEY,
	title VARCHAR(20),
	content VARCHAR(100),
	uid INT,
	FOREIGN KEY (uid) references users (id)
);