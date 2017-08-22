build:
	go build -o itunes itunes.go

linux:
	GOOS=linux GOARCH=amd64 go build -o itunes itunes.go

mac:
	GOOS=darwin GOARCH=amd64 go build -o itunes itunes.go

install:
	go build -o itunes itunes.go
	mv itunes $(GOPATH)/bin/itunes

clean:
	rm -rf itunes

