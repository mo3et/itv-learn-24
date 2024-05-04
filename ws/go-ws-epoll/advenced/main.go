package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"reflect"
	"sync"
	"syscall"

	"github.com/gorilla/websocket"
	"golang.org/x/sys/unix"
)

// ONLY UNIX

func main() {
	// create epoll
	epollFd, err := unix.EpollCreate(1)
	if err != nil {
		log.Fatalf("unable to create epoll:%v\n", err)
	}

	connections := make(map[int32]*websocket.Conn, 1000000)
	var mu sync.Mutex
	ch := make(chan int32, 50)

	// create goroutine, processing received msg
	go func() {
		for {
			select {
			case fd := <-ch:
				conn := connections[fd]
				if conn == nil {
					continue
				}
				// send msg
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Println("unable to read message:", err.Error())
					_ = conn.Close()
					// 删除 epoll 事件
					if err := unix.EpollCtrl(epollFd, syscall.EPOLL_CTL_DEL, int(fd), nil); err != nil {
						log.Println("unable to remove event")
					}
				}
				// reply msg
				if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("received:%s", string(message)))); err != nil {
					log.Println("unable to send message:", err.Error())
				}
			}
		}
	}()

	// create goroutine listen epoll event
	go func() {
		for {
			// declare 50 events, get 50 events more in once
			events := make([]unix.EpollEvent, 50)
			// peer 100ms execute once
			n, err := unix.EpollWait(epollFd, events, 100)
			if err != nil {
				log.Println("epoll wait:", err.Error())
				continue
			}

			// 获得就绪的WebSocket连接的fd
			for i := 0; i < n; i++ {
				if events[i].Fd == 0 {
					continue
				}
				ch <- events[i].Fd // 通过channel传递到另一个 gouroutine处理
			}
		}
	}()

	// bind http Service
	http.ListenAndServe(":8085", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// upgrade to websocket connect
		conn, err := NewWebsocketConnection(w, r)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// 获得文件标识符(fd)
		fd := GetSocketFD(conn.UnderlyingConn())

		// 注册事件
		if err := unix.EpollCtl(epollFd, unix.EPOLL_CTL_ADD, int(fd), &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, fd: fd}); err != nil {
			log.Println("unable to add event:%v", err.Error())
			return
		}
		// 保存到map里
		mu.Lock()
		connections[fd] = conn
		mu.Unlock()
	}))
}

var u = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }} // use default options

// WebsocketConn return web socket connection
func NewWebsocketConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return u.Upgrade(w, r, nil)
}

// GetSocketFD get socket connection fd
func GetSocketFD(conn net.Conn) int32 {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	return int32(pfdVal.FieldByName("Sysfd").Int())
}
