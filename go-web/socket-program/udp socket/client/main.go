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
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_, err = conn.Write([]byte("anything"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}

/*
通过上面的代码我们可以看出：首先程序将用户的输入作为参数service传入net.ResolveUDPAddr获取一个udpAddr,
然后把UDPAddr传入DialUDP后创建了一个UDP连接conn，通过conn来发送请求信息，
最后通过ioutil.ReadAll从conn中读取全部的文本，也就是服务端响应反馈的信息。
*/
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
