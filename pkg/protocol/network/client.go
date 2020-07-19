package network

import (
	"log"
	"net"
)

type Client struct {
	Conn net.Conn
}

func NewClient(dns string) *Client {
	conn, err := net.Dial("tcp", dns)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{Conn: conn}
}
