package main

import (
	"database/sql"
	f "fmt"
	 "crypto/sha256"
	_ "github.com/mattn/go-sqlite3"
)

func CheckError(err error) {
	if err != nil {
		f.Println("----------------------------------------------------");
		f.Println(err);
		f.Println("----------------------------------------------------");
	}
}

func ConnectToDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath);
	CheckError(err);
	createTable := "create table user(uid integer PRIMARY KEY,uname varchar(20),passd varchar(60));"
	_, err = db.Exec(createTable);
	CheckError(err);
	createTable = "create table notes(id integer PRIMARY KEY, userid integer, note varchar(100), FOREIGN KEY(userid) REFERENCES user(uid));"
	_, err = db.Exec(createTable);
	CheckError(err);
	return db;
}

func getAllNote(userid string) map[uint]string {
	AllNote := make(map[uint]string);
	db := ConnectToDB("./db/note.db");
	rows, err := db.Query("select * from notes where userid = ? order by id desc;", userid);
	CheckError(err);
	for rows.Next() {
		var id uint;
		var note string;
		err := rows.Scan(&id, &note);
		CheckError(err);
		AllNote[id] = note;
	}
	rows.Close();
	db.Close();
	return AllNote;
}

func AddNote(note string, userid string) {
	db := ConnectToDB("./db/note.db");
	statement, err := db.Prepare("insert into notes(note) values (?);");
	CheckError(err);
	_ , err = statement.Exec(note);
	CheckError(err);
	db.Close();
}

func DeleteNote(id string) {
	db := ConnectToDB("./db/note.db");
	statement, err := db.Prepare("delete from notes where id = ?;");
	CheckError(err);
	_ , err = statement.Exec(id);
	CheckError(err);
	db.Close();
}
// -1 = nouser
func CheckAccount(uname string, passd string) int {
	command := "select uid from user where uname = ?";
	db := ConnectToDB("./db/note.db");
	if passd != "" {
		passd = f.Sprintf("%x", sha256.Sum256([]byte (passd)));
		command += " And passd = \"" + passd + "\"";
	}
	id := -1;
	rows, err := db.Query(command, uname);
	CheckError(err);
	if rows.Next() {
		err = rows.Scan(&id);
		CheckError(err);
	}
	rows.Close();
	db.Close();
	return id;
}
//-1 = error
func CreateUser(uname string, passd string) int {
	if (CheckAccount(uname, "") == -1) {
		//convert pass to hash
		passd = f.Sprintf("%x", sha256.Sum256([]byte (passd)));
		//connect to db
		db := ConnectToDB("./db/note.db");
		//insert data to table
		statement, err := db.Prepare("insert into user(uname, passd) values(?, ?);");
		res, err := statement.Exec(uname, passd);
		CheckError(err);
		//get last id
		id, err := res.LastInsertId();
		CheckError(err);
		return int (id);
	}
	return -1;
}
