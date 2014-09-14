depends:
	go get code.google.com/p/gorest
	go get github.com/go-sql-driver/mysql
	go get github.com/go-xorm/xorm
	go get github.com/go-xorm/core
	go get github.com/go-xorm/cmd/xorm
	go get github.com/go-xorm/ql
	go get github.com/pelletier/go-toml

build:
	go build -v -o lilliputservice src/main.go
