package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"

	pb "grpc_test/tls_and_token/proto"
)

const (
	// Address 监听地址
	Address string = ":1234"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.listen err: %v", err)
	}

	log.Println(Address + "net.listening ...")

	//  新建 grpc 服务器实例
	creds, err := credentials.NewServerTLSFromFile("../tls_token/server.pem", "../tls_token/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	//  添加拦截器
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		//拦截普通方法请求，验证Token
		err := pb.CheckToken(ctx)

		if err != nil {
			return nil, err
		}
		// 继续处理请求
		return handler(ctx, req)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))

	//  初始化服务
	s := pb.MyService{}

	//  在启动的 grpc server 中，注册我们定义的服务
	pb.RegisterTestServer(grpcServer, &s)

	// 用服务器 Serve() 方法以及我们的端口信息 实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
