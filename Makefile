todo: main.go ./db/note.db sqlC.go *_templ.go
	go build -o todo .

templ: src/*.templ
	templ generate; mv src/*_templ.go .

./db/note.db: note.sql
	sqlite3 db/note.db < note.sql

run: todo
	./todo

clean:
	rm -rf ./todo

cleanAll:
	rm -rf ./todo ./db/note.db ./*_templ.go
