package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	_ = http.ListenAndServe(":8085", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upgrade to websocket connect
		conn, err := NewWebsocketConnection(w, r)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		go func() { // create goroutine to processing
			defer conn.Close() // Close conn
			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					log.Println("unable to read message:", err.Error())
					return
				}
				// reply to client, or other business logic
				reply := fmt.Sprintf("received:%s", string(msg))
				if err := conn.WriteMessage(websocket.TextMessage, []byte(reply)); err != nil {
					log.Printf("unable to send message")
				}
			}
		}()
	}))
}

var u = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }} // use default options

// WebsocketConn return web socket connection
func NewWebsocketConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return u.Upgrade(w, r, nil)
}
