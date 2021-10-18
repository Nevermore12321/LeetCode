package main
import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "grpc_test/tiimeout_sample/proto"
)

const (
	// Address 监听地址
	Address string = ":1234"
	// Network 网络通信协议
	Network string = "tcp"
)

// 启动gRPC服务器
func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.listen err: %v", err)
	}

	log.Println(Address + "net.listening ...")

	//  新建 grpc 服务器实例
	grpcServer := grpc.NewServer()

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
