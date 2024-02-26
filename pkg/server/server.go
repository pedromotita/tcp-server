package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	host string
	port string
}

func New(host, port string) *Server {
	return &Server{host, port}
}

func (s *Server) Run() {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)

	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buffer))

	conn.Write([]byte("Hello from server\n"))
}
