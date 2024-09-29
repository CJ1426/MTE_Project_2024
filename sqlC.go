package main

import (
	"database/sql"
	f "fmt"
	_ "crypto/sha256"
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

func getAllNote() map[uint]string {
	AllNote := make(map[uint]string);
	db := ConnectToDB("./db/note.db");
	//f.Println(reflect.TypeOf(db));
	//f.Println(db);
	rows, err := db.Query("select * from notes order by id desc;");
	CheckError(err);
	//f.Println(reflect.TypeOf(rows));
	for rows.Next() {
		var id uint;
		var note string;
		err := rows.Scan(&id, &note);
		CheckError(err);
		//f.Println(id);
		//f.Println(note);
		AllNote[id] = note;
	}
	rows.Close();
	db.Close();
	return AllNote;
	//rows.Close();
}

func AddNote(note string) {
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
