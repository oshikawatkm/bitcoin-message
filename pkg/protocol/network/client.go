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

func (c *Client) SendMessage(msg protocol.Message) (int, error) {
	message := common.NewMessage(msg.Command(), msg.Encode())
	log.Printf("send		: %s", string(message.Command[:]))
	return c.Conn.Write(message.Encode())
}

func (c *Client) ReceiveMessage(sizr uint32) ([]byte, error) {
	buf := make([]byte, size)
	_, err := c.Conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
