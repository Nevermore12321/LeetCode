package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_test/single_cs/proto"
	"log"
)

const (
	// Address 连接地址
	Address string = ":1234"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: ", err)
	}

	// 退出后，关闭连接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("connection close err: ", err)
		}
	}(conn)

	// 建立gRPC连接
	client := pb.NewTestClient(conn)

	// 创建发送请求的结构体
	request := pb.TestRequest{
		Name: "gsh",
		Id:   1,
	}

	// 在服务端，我们已经将服务注册到了 grpc 中，因此可以直接调用我们的服务(SayHello)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC}
	response, err := client.SayHello(context.Background(), &request)
	if err != nil {
		log.Fatalf("Call FirstRpc err: %v", err)
	}


	// 打印返回值
	log.Println(response)
}