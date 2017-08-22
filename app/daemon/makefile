run:
	go run apple.go

start:
	nohup ./apple 1>apple.log 2>&1 &

stop:
	ps aux | grep "\./apple" | grep -v grep | awk '{print $2}' | xargs kill

status:
	ps aux | grep "\./apple"

log:
	tail -f apple.log

build:
	go build -o apple apple.go
	chmod a+x apple

install:
	go build -o apple apple.go
	chmod a+x apple
	mv apple $(GOPATH)/bin/apple

upload:
	GOOS=linux GOARCH=amd64 go build -o apple apple.go
	chmod a+x apple
	oss put apple && rm apple

download:
	oss get apple
	chmod a+x apple

linux:
	GOOS=linux GOARCH=amd64 go build -o apple apple.go
	chmod a+x apple

mac:
	GOOS=darwin GOARCH=amd64 go build -o apple apple.go
	chmod a+x apple

createdb:
	psql -c "CREATE USER meta CREATEDB;"
	psql -c "ALTER USER meta with PASSWORD 'meta';"
	psql postgres meta -c "CREATE DATABASE meta;"

setup:
	psql meta meta < apple.ddl

clean:
	rm -rf apple apple.log

.PHONY: run clean

