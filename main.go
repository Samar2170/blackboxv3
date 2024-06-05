package main

import (
	"blackboxv3/client"
	"blackboxv3/server"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/skip2/go-qrcode"
)

func getQRCode(ipAddr string) []byte {
	png, err := qrcode.Encode(ipAddr, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}
	return png
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	hostname := os.Getenv("HOSTNAME")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		qrCode := getQRCode(hostname)
		w.Write(qrCode)
	})
	http.HandleFunc("/", server.ServeWS)
	log.Println("Server started on " + hostname + ":8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func testClient() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	c := client.NewClient(conn)
	c.WriteMessage(websocket.TextMessage, []byte("Hello World"))
	c.WriteMessage(websocket.TextMessage, []byte("Duniya Khatam Ho Jayegi"))
}
