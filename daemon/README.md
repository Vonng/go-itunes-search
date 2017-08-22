# apple daemon

apple daemon that take ID as input, produce apple Application as output.


## Install

```bash
go get github.com/Vonng/go-itunes-search
cd ${GOPATH}/src/github.com/Vonng/go-itunes-search/daemon

# Setup database environment. assume you have an available local pg
# It will create a user `meta` with owns a database named `meta`
make createdb

# It will create table `apple` and `apple_queue` in database `meta`
make setup

# Build binary
make build

# Install: mv binary to your $GOPATH
make install
```

now everything is prepared for running the daemon

### Usage

some frequently used bash command can be accessed from makefile

```bash
# Start the daemon. don't forget build before start
make start

# show daemon status
make status

# Stop the daemon
make stop

# See log
make log

# Using `go run apple.go`
make
```


### Assign Task

INSERT into `apple_queue`. `apple` will take task from queue table and put result into table `apple`.
 
 task format is `TypeLetter + ID`, where `TypeLetter` could be:
 
 * `!`: stand for apple ID, i.e iTunesID, TrackID
 * `@`: stand for bundleID.
 * `#`: stand for keyword.  program will search and fetch new found app.
 * no leading letter will use bundleID by default. (for stupid client...)

e.g : `!460819018` add a iTunesID Task to queue with ID value `460819018` and `#蛤蛤` add a keyword-search task to queue with keyword `蛤蛤`

And daemon binary can handle iTunesID, BundleID, Keywords directly by:

```bash
apple a 414478124
apple b com.tencent.xin
apple k yourKeyword
```
