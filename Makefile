PodBuilder ?= podman

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
	[ -e todo.tar.gz ] && rm -rf todo.tar.gz || echo "first time?"
	$(PodBuilder) build -t mte_p . && $(PodBuilder) save -o todo.tar mte_p && gzip todo.tar && $(PodBuilder) image rm mte_p
