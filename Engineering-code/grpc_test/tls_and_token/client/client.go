package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"

	pb "grpc_test/tls_and_token/proto"
)

// Token 定义 TOKEN 结构体
type Token struct {
	TokenID   string
	TokenName string
}

// GetRequestMetadata 实现 Token 的两个方法
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token_id": t.TokenID,
		"token_name": t.TokenName,
	}, nil
}

func (t *Token) RequireTransportSecurity() bool {
	return true
}

const (
	// Address 连接地址
	Address string = ":1234"
)

func main() {
	// 连接服务器
	// tls 认证
	cred, err := credentials.NewClientTLSFromFile("../tls_token/server.pem", "gsh.com")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	// jwt token
	token := Token{
		TokenID: "12345",
		TokenName: "gsh_token",
	}

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(cred), grpc.WithPerRPCCredentials(&token))
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
