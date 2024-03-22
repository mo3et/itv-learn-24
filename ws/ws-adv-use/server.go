package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// 对服务端的配置
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许跨域
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// 启动一个 goroutine 定期发送心跳包
	go func() {
		for {
			err := conn.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				log.Println(err)
				break
			}
			time.Sleep(time.Second * 30) // 每 30 秒发送一次心跳包
		}
	}()

	// more handle
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Received: %s", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
