package client

import (
	"fmt"
	"net"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) Connect(host, port string) (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Write(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))

	defer conn.Close()

	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)

	_, err = conn.Read(buffer)
	if err != nil {
		return err
	}

	fmt.Print(string(buffer))

	return err
}
