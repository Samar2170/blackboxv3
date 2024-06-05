package client

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

// Client is a websocket client, basically a frontend visitor
type Client struct {
	// the websocket connection
	connection *websocket.Conn

	// manager is the manager used to manage the client
}

// NewClient is used to initialize a new Client with all required values initialized
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		connection: conn,
	}
}
func (c *Client) ReadMessages() {
	// Loop Forever
	for {
		// ReadMessage is used to read the next message in queue
		// in the connection
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			// If Connection is closed, we will Recieve an error here
			// We only want to log Strange errors, but not simple Disconnection
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break // Break the loop to close conn & Cleanup
		}
		log.Println("MessageType: ", messageType)
		log.Println("Payload: ", string(payload))
	}
}

// WriteMessage is used to write a message to the client
func (c *Client) WriteMessage(messageType int, payload []byte) error {
	return c.connection.WriteMessage(messageType, payload)
}
