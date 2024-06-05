package server

import "github.com/gorilla/websocket"

type Client struct {
	Connection *websocket.Conn
	Egress     chan []byte
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Connection: conn,
		Egress:     make(chan []byte),
	}
}

type ClientList map[*Client]bool
