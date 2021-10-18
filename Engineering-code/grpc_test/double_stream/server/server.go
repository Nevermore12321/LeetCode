package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "grpc_test/double_stream/proto"
)

const (
	// Address 监听地址
	Address string = ":1234"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	//- 创建 net 监听 listener 实例
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")

	//- 启动 grpc server 实例
	grpcServer := grpc.NewServer()

	s := pb.BothStream{}
	//- 将 Service 结构体注册进 grpc server
	pb.RegisterStreamServer(grpcServer, &s)

	//- 开始启动服务端，通过 listener，grpc server 开启监听
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
