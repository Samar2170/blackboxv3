package main

import (
	blackboxv3 "blackboxv3/api/apigen"
	"blackboxv3/cmd"
	"log"
	"net"

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
	log.Println("Starting server...")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	blackboxv3.RegisterBlackBoxServiceServer(cmd.Server, &cmd.GrpcServer{})
	if err := cmd.Server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Println("Server started")
	}
}

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic(err)
// 	}
// 	hostname := os.Getenv("HOSTNAME")
// 	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("OK"))
// 	})
// 	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
// 		qrCode := getQRCode(hostname)
// 		w.Write(qrCode)
// 	})
// 	http.HandleFunc("/socket", ws.ServeWS)
// 	log.Println("Server started on " + hostname + ":8080")
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		log.Fatal(http.ListenAndServe(":8080", nil))
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }

// func testClient() {
// 	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/socket", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	c := wsclient.NewClient(conn)
// 	c.WriteMessage(websocket.TextMessage, []byte("Hello World"))
// 	c.WriteMessage(websocket.TextMessage, []byte("Duniya Khatam Ho Jayegi"))
// }
