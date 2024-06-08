package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// if strings.HasPrefix(GetIP(r), "192.168.") {
		// 	return true
		// }
		// return false
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var MsgTypeMap = map[int]string{
	websocket.TextMessage:   "Text",
	websocket.BinaryMessage: "Binary",
	websocket.CloseMessage:  "Close",
	websocket.PingMessage:   "Ping",
	websocket.PongMessage:   "Pong",
}

type Message struct {
	MessageType int
	Payload     []byte
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func ServeWS(w http.ResponseWriter, r *http.Request) {
	log.Println("WS Connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Connected to", GetIP(r))
	defer conn.Close()

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			errStr := "Connection Error " + err.Error() + "\n"
			conn.WriteMessage(websocket.CloseMessage, []byte(errStr))
			return
		}
		log.Println("Received:", MsgTypeMap[mt], string(msg))
	}
}
