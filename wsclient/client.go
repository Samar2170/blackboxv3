package wsclient

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		connection: conn,
	}
}
func (c *Client) ReadMessages() {
	for {
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break // Break the loop to close conn & Cleanup
		}
		log.Println("MessageType: ", messageType)
		log.Println("Payload: ", string(payload))
	}
}

func (c *Client) WriteMessage(messageType int, payload []byte) error {
	return c.connection.WriteMessage(messageType, payload)
}
