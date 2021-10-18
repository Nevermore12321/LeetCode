package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_test/client_stream/proto"
	"log"
	"strconv"
	"time"
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
	grpcClient := pb.NewStreamClientClient(conn)


	//- 调用注册的Service中的方法（也就是函数接口）
	stream, err := grpcClient.RouteList(context.Background())
	if err != nil {
		log.Fatalf("grpcClient.RouteList err: %v", err)
	}

	//- 向 数据流中，写入数据
	for i := 0; i < 10; i++ {
		// 向数据流中写入
		err := stream.Send(&pb.StreamRequest{
			StreamData: "stream client rpc: " + strconv.Itoa(i),
		})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	//  写完数据后，接受服务端发送来的响应
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
	log.Println(response)
}
