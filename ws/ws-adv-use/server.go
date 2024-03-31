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

// TODO 需要复习 wsHandler 包括心跳发送，长连接的内容

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
	// 可以在 loop 中，从ws 连接 `conn` 读写信息。根据收到的信息实现不同逻辑
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

// Advance Features
// 实现多client广播[可以维护连接列表，并遍历这些连接以发送信息]
// 实现自定义消息格式、错误处理和安全措施。

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
