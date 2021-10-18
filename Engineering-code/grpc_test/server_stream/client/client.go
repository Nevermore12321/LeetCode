package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_test/server_stream/proto"
	"io"
	"log"
)

const Address string = ":1234"

func main() {
	//- 连接服务器（ip+port）
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	//- 关闭 连接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("grpc.Dial err: %v", err)
		}
	}(conn)


	//- 建立gRPC连接
	grpcClient := pb.NewStreamServerClient(conn)

	//- 创建 请求的结构体
	req := pb.StreamRequest{
		Data: "Stream Server grpc ",
	}

	//- 调用注册的Service中的方法（也就是函数接口）
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("grpcClient.ListValue err: %v", err)
	}

	//- 获取 response 并处理
	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()

		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("stream.Recv err: %v", err)
		}

		// 打印从数据流中接收到的值
		fmt.Println(res.StreamValue)
	}
}
