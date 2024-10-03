Needed : 
    - github.com/mattn/go-sqlite3
    - github.com/a-h/templ
    - github.com/a-h/templ/cmd/templ@latest

setup
    go get
    PATH=$(go env GOPATH)/bin:$PATH

load docker image:
    docker load -i todo.tar.gz
