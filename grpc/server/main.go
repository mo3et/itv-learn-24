package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.Um
}

func main() {
	// listen 8972 port
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer() // 创建gRPC服务器

	// 启动服务
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
