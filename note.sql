create table user(
	uid integer PRIMARY KEY,
	uname varchar(20),
	passd varchar(60)
);
create table notes(
	id integer PRIMARY KEY,
	userid integer,
	note varchar(100),
	FOREIGN KEY(userid) REFERENCES user(uid)
);
