package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	Out  chan<- string
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

// var timeout = 10 * time.Second
var timeout = 10 * time.Minute

// 广播
func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-message:
			for cli := range clients {
				cli.Out <- msg
			}
		case cli := <-entering:
			clients[cli] = true

			cli.Out <- "Present:"
			for c := range clients {
				cli.Out <- c.Name
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

func handleConn(conn net.Conn) {
	out := make(chan string, 10)
	go clientWriter(conn, out)

	in := make(chan string)
	go clientReader(conn, in)

	var (
		who       string
		nameTimer = time.NewTimer(timeout)
	)

	out <- "Enter your name:"

	select {
	case name := <-in:
		who = name
	case <-nameTimer.C:
		conn.Close()
		return
	}

	cli := client{out, who}

	out <- "You are " + who

	message <- who + " has arrived"

	entering <- cli

	idle := time.NewTimer(timeout)

	// Label 用于循环语句或代码块
	// 以便跳出循环，提高灵活性和可读性
Loop:
	for {
		select {
		case msg := <-in:
			message <- who + ": " + msg
			idle.Reset(timeout)
		case <-idle.C:
			conn.Close()
			// 跳出标签为 Loop 的循环
			break Loop

		}
	}
	leaving <- cli

	message <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
