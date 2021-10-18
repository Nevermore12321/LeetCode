package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"

	pb "grpc_test/tiimeout_sample/proto"
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
	clientDeadLine := time.Now().Add(time.Duration(3 * time.Second))
	ctx, cancelFunc := context.WithDeadline(context.Background(), clientDeadLine)
	defer cancelFunc()

	response, err := client.SayHello(ctx, &request)
	if err != nil {
		//获取错误状态
		fromError, ok := status.FromError(err)
		if ok {
			//判断是否为调用超时
			if fromError.Code() == codes.DeadlineExceeded {
				log.Fatalln("FirstRpc timeout!")
			}
		}
		log.Fatalf("Call FirstRpc err: %v", err)
	}

	// 打印返回值
	log.Println(response)
}