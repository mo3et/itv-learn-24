package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// TCP Socket: https://github.com/astaxie/build-web-application-with-golang

// TCP server
func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	// 设置连接超时时间,当超过设置时间时，连接自动关闭。
	// conn, err = net.DialTimeout("tcp", tcpAddr.String(), 5*time.Hour)
	checkError(err)
	for {
		conn, err := listener.AcceptTCP()
		// 设置keepAlive属性。操作系统层在tcp上没有数据和ACK的时候，会间隔性的发送keepalive包，
		// 操作系统可以通过该包来判断一个tcp连接是否已经断开，这个功能和我们通常在应用层加的心跳包的功能类似。
		conn.SetKeepAlive(true)
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// 读取超时时间，当超过设置时间，连接自动关闭。
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout

	request := make([]byte, 128) // set maxium request length to 128B to prevent flood attack
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

/*
在上面这个例子中，我们使用conn.Read()不断读取客户端发来的请求。由于我们需要保持与客户端的长连接，
所以不能在读取完一次请求后就关闭连接。由于conn.SetReadDeadline()设置了超时，
当一定时间内客户端无请求发送，conn便会自动关闭，下面的for循环即会因为连接已关闭而跳出。
需要注意的是，request在创建时需要指定一个最大长度以防止flood attack；每次读取到请求处理完毕后，
需要清理request，因为conn.Read()会将新读取到的内容append到原内容之后。
*/
