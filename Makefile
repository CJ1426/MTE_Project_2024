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

podman: todo
	podman build -t mte_p . && podman save -o todo.tar mte_p && gzip todo.tar && podman image rm mte_p

docker: todo
	docker build -t mte_p . && docker save -o todo.tar mte_p && gzip todo.tar && docker image rm mte_p
