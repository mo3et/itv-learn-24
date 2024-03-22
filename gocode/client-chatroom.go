package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Fatal("Unable to connect to server:", err)
	}
	defer conn.Close()

	// 创建用于向服务器发送消息的channel
	sendCh := make(chan string)
	// 创建用于接收从服务器收到的消息的channel
	receiveCh := make(chan string, 10)

	go sendMessage(conn, sendCh)
	go receiveMessages(conn, receiveCh)

	go func() {
		for msg := range receiveCh {
			fmt.Print(msg)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		sendCh <- input
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading stdinput:", err)
	}
}

func sendMessage(conn net.Conn, sendCh <-chan string) {
	for {
		input := <-sendCh
		_, err := fmt.Fprintf(conn, "%s\n", input)
		if err != nil {
			log.Println("Fail to read form server:", err)
			return
		}
	}
}

// 接收信息
func receiveMessages(conn net.Conn, receiveCh chan<- string) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Fail to read form server:", err)
			return
		}
		receiveCh <- message
	}
}

// ----------------------

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:8088")
// 	if err != nil {
// 		log.Fatal("Unable to connect to server:", err)
// 	}
// 	defer conn.Close()

// 	// 创建一个goroutine来持续读取服务器发送的消息并打印到控制台
// 	go func() {
// 		reader := bufio.NewReader(conn)
// 		for {
// 			message, err := reader.ReadString('\n')
// 			if err != nil {
// 				log.Println("Failed to read from server:", err)
// 				return
// 			}
// 			fmt.Print(message)
// 		}
// 	}()

// 	// 从标准输入读取用户输入，并发送到服务器
// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		input := scanner.Text()
// 		_, err := fmt.Fprintf(conn, "%s\n", input)
// 		if err != nil {
// 			log.Println("Failed to send message to server:", err)
// 			return
// 		}
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Println("Error reading standard input:", err)
// 	}
// }
