package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"time"

	pb "grpc_test/double_stream/proto"
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
	grpcClient := pb.NewStreamClient(conn)


	//- 调用服务端的Conversations方法，获取流
	stream, err := grpcClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("grpcClient.RouteList err: %v", err)
	}

	//- 向 数据流中，写入数据
	for i := 0; i < 10; i++ {
		// 向数据流中写入
		err := stream.Send(&pb.StreamRequest{
			Question: "stream client rpc: " + strconv.Itoa(i),
		})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}

		recv, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}

		// 打印返回值
		log.Println(recv.Answer)

		time.Sleep(1 * time.Second)
	}

	//  写完数据后，接受服务端发送来的响应
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
}
