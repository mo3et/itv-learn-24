package main

import (
	"fmt"
	"net"
	"os"
)

// tcp client
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// result, err := ioutil.ReadAll(conn)
	result := make([]byte, 256)
	_, err = conn.Read(result)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

/*
通过上面的代码我们可以看出：首先程序将用户的输入作为参数service传入net.ResolveTCPAddr获取一个tcpAddr,
然后把tcpAddr传入DialTCP后创建了一个TCP连接conn，通过conn来发送请求信息，
最后通过ioutil.ReadAll从conn中读取全部的文本，也就是服务端响应反馈的信息。
*/
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
