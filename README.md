# tcp-server

A simple CLI to handle TCP server/client connections

# Installing

Make sure you have `go` installed and that `$(go env GOPATH)/bin` is in your path, then:

```sh
go install github.com/pedromotita/tcp-server@latest
```

# Usage

## Build

To build the CLI run:

```sh  
cd tcp-server/
go build -o tcp
```

## Start up a server

To start up a TCP server must provide a host and a port to listen to:

```sh
./tcp server run --host localhost --port 9001
```

## Conenct a client to a server

To connect a client to a server you must provide a host and port to connect, and a message to send to the server:

```sh
./tcp client connect --host localhost --port 9001 -m "Hello from client"
```

## Contributing

Feel free to open PRs to fix bugs and add new features

