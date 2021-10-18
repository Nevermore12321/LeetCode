GRPC

[toc]



## RPC

**RPC（Remote Procedure Call）远程过程调用**，简单的理解是一个节点请求另一个节点提供的服务。所谓 RPC(remote procedure call 远程过程调用) 框架实际是提供了一套机制，使得应用程序之间可以进行通信，而且也遵从server/client模型。使用的时候客户端调用 server 端提供的接口就像是调用本地的函数一样。


过程：
- 首先客户端需要告诉服务器，需要调用的函数，这里函数和进程 ID 存在一个映射，客户端远程调用时，需要查一下函数，找到对应的 ID，然后执行函数的代码。
- 客户端需要把本地参数传给远程函数，本地调用的过程中，直接压栈即可，但是在远程调用过程中不再同一个内存里，无法直接传递函数的参数，因此需要客户端把参数转换成字节流，传给服务端，然后服务端将字节流转换成自身能读取的格式，是一个序列化和反序列化的过程。
- 数据准备好了之后，如何进行传输？网络传输层需要把调用的 ID 和序列化后的参数传给服务端，然后把计算好的结果序列化传给客户端，因此 TCP 层即可完成上述过程，gRPC 中采用的是 HTTP2 协议。



## GRPC

与许多 RPC 系统一样，gRPC 基于定义服务的思想，指定可以使用其参数和返回类型远程调用的方法。默认情况下，gRPC 使用协议缓冲区作为接口定义语言（IDL）来描述服务接口和有效负载消息的结构。


### GRPC 的优势

**GRPC的优势**
- 多语言：语言中立，支持多种语言。
- 轻量级、高性能：序列化支持 PB(Protocol Buffer) 和 JSON，PB 是一种语言无关的高性能序列化框架。
- 可插拔
- IDL：基于文件定义服务，通过 proto3 工具生成指定语言的数据结构、服务端接口以及客户端 Stub。
- 设计理念
    - 移动端：基于标准的 HTTP2 设计，支持双向流、消息头压缩、单 TCP 的多路复用、服务端推送等特性，这些特性使得 gRPC 在移动端设备上更加省电和节省网络流量。
- 服务而非对象、消息而非引用：促进微服务的系统间粗粒度消息交互设计理念。
- 负载无关的：不同的服务需要使用不同的消息类型编码，例如 protocol buffers、JSON、XML 和 Thrift。
- 流：Streaming API。
- 阻塞式和非阻塞式：支持异步和同步处理在客户端和服务端间交互的消息序列。
- 元数据交换：常见的横切关注点，如认证或跟踪，依赖数据交换。
- 标准化状态码：客户端通常以有限的方式响应 API 调用返回的错误。





### 服务的定义

**gRPC 基于如下思想**：定义一个服务， 指定其可以被远程调用的方法及其参数和返回类型。gRPC 默认使用 protocol buffers 作为接口定义语言，来描述服务接口和有效载荷消息结构。


例如：
```protobuf
service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  required string greeting = 1;
}

message HelloResponse {
  required string reply = 1;
}
```

### gRPC 允许定义四类服务方法

1. 单向 RPC，即客户端发送一个请求给服务端，从服务端获取一个应答，就像一次普通的函数调用。
```protobuf
rpc SayHello(HelloRequest) returns (HelloResponse){
    ...
}
```
2. 服务端流式 RPC，即客户端发送一个请求给服务端，可获取一个数据流用来读取一系列消息。客户端从返回的数据流里一直读取直到没有更多消息为止。
    - 也就是 grpc 的 返回值中有一个数据流，客户端可以从返回值，也就是数据流中读取数据
```protobuf
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){
    ...
}
```
3. 客户端流式 RPC，即客户端用提供的一个数据流写入并发送一系列消息给服务端。一旦客户端完成消息写入，就等待服务端读取这些消息并返回应答。
    - 也就是 grpc 函数的接收参数中有一个客户端发来的数据流，服务端可以从数据流中读取数据
```protobuf
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {
    ...
}
```
4. 双向流式 RPC，即两边都可以分别通过一个读写数据流来发送一系列消息。这两个数据流操作是相互独立的，所以客户端和服务端能按其希望的任意顺序读写，例如：服务端可以在写应答前等待所有的客户端消息，或者它可以先读一个消息再写一个消息，或者是读写相结合的其他方式。每个数据流里消息的顺序会被保持。
    - 也就是 grpc 函数的接受参数有客户端提供的数据流，函数的返回值是服务端提供的数据流，两边可以同时读写数据。
```protobuf
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){
    ...
}
```



### GRPC 的安装

#### grpc 安装与配置

1. 安装 protobuf
    - 下载地址：https://github.com/protocolbuffers/protobuf/releases
    - 将解压后在bin目录找到protoc.exe，然后把bin目录添加到环境变量
    - 运行`protoc --version`，检查是否安装成功

2. 安装 golang 工具包
    - 安装 golang 的proto工具包
    ```bash
    go get -u github.com/golang/protobuf/proto
    ```
    - 安装 goalng 的proto编译支持
    ```bash
    go get -u github.com/golang/protobuf/protoc-gen-go
    ```
    - 安装 gRPC 包
    ```bash
    go get -u google.golang.org/grpc
    ```


### GRPC 的使用

#### 1. 单向 RPC 使用

1. 新建test.proto文件
```
syntax = "proto3"; //  使用 protobuf3

//  package 名称
package grpc;

//
option go_package = "./";


//  定义发送请求信息
message testRequest {
  //  定义发送的参数
  //  格式： 参数名 参数类型 = 标识号（不重复）
  string name = 1;
  int64 id = 2;
}

// 定义响应信息
message testResponse {
  //  定义响应的参数
  //  格式： 参数名 参数类型 = 标识号（不重复）
  int64 status = 1;
  string value = 2;
}

// 定义服务
//  可以定义多个服务，也可以在同一个服务中，定义多个函数接口
service test {
  rpc FirstRpc(testRequest) returns (testResponse) {};
}


```
2. 编译proto文件, 生成 go 文件
    - 命令：`protoc --proto_path=IMPORT_PATH --<lang>_out=DST_DIR path/to/file.proto`
    - `--proto_path=IMPORT_PATH`：可以在 .proto 文件中 import 其他的 .proto 文件，proto_path 即用来指定其他 .proto 文件的查找目录。如果没有引入其他的 .proto 文件，该参数可以省略。
    - `--<lang>_out=DST_DIR`：指定生成代码的目标文件夹，例如 –go_out=. 即生成 GO 代码在当前文件夹，另外支持 cpp/java/python/ruby/objc/csharp/php 等语言
```
C:\install-tools\goland\golang-project\go-learn\grpc-test>protoc --go_out=plugins=grpc:./ ./test.proto

C:\install-tools\goland\golang-project\go-learn\grpc-test>dir
2021/04/13  15:44    <DIR>          .
2021/04/13  15:44    <DIR>          ..
2021/04/13  15:44             9,454 test.pb.go
2021/04/13  15:28               678 test.proto
               2 个文件         10,132 字节
```
3. 创建Server端
    - 定义 Service 的结构体
    - 实现 Service 结构体的方法，也就是service内的函数接口
    - 创建 net 监听 listener 实例
    - 启动 grpc server 实例
    - 将 Service 结构体注册进 grpc server
    - 开始启动服务端，通过 listener，grpc server 开启监听
```go
package main

import (
	"context"
	pb "go-learn/grpc-test"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 定义服务
type test struct {}

//  实现 service 中的 函数接口
func (t *test) FirstRpc(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	//  初始化 response
	res := pb.TestResponse{
		Status: 200,
		Value: "Hello " + req.GetName() + ", id " + string(req.GetId()),
	}
	//  返回
	return &res, nil
}

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

	//  在启动的 grpc server 中，注册我们定义的服务
	pb.RegisterTestServer(grpcServer, &test{})

	// 用服务器 Serve() 方法以及我们的端口信息 实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}

}

```
4. 创建Client端
    - 连接服务器（ip+port）
    - 建立gRPC连接
    -  创建 请求的结构体
    - 调用注册的Service中的方法（也就是函数接口）
    - 获取 response 并处理
    - 关闭 连接
```go
package main

//func main() {
//	utils.SortDuration(ch02_Sort.SelectionSort, "SelectionSort")
//	utils.SortDuration(ch02_Sort.SelectionSortAdvanced, "SelectSortAdvanced")
//}

import (
	pb "algorithm4/grpc-test"
	"context"
	grpc "google.golang.org/grpc"
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
	defer conn.Close()

	// 建立gRPC连接
	grpcClient := pb.NewTestClient(conn)

	// 创建发送请求的结构体
	request := pb.TestRequest{
		Name: "gsh",
		Id: 1,
	}

	// 在服务端，我们已经将服务注册到了 grpc 中，因此可以直接调用我们的服务(FirstRpc方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	response, err := grpcClient.FirstRpc(context.Background(), &request)
	if err != nil {
		log.Fatalf("Call FirstRpc err: %v", err)
	}

	// 打印返回值
	log.Println(response)

}
```


#### 2. 服务端流式 RPC 使用


**服务端流式RPC**：客户端发送请求到服务器，拿到一个流去读取返回的消息序列。 客户端读取返回的流，直到里面没有任何消息。

情景模拟：实时获取股票走势。  
1. 客户端要获取某原油股的实时走势，客户端发送一个请求
2. 服务端实时返回该股票的走势


使用步骤：
1. 创建proto文件, 函数接口返回stream 数据流，客户端可以读取
```
syntax = "proto3"; //  使用 protobuf3

//  package 名称
package my_lib;

// import path
option go_package = "./";

message StreamRequest{
  // 定义发送的参数，采用驼峰命名方式，小写加下划线，如：student_name
  // 参数类型 参数名 = 标识号(不可重复)
  string data = 1;
}


message StreamResponse{
  // 流式响应数据
  string stream_value = 1;
}


// 定义服务（可定义多个服务,每个服务可定义多个接口）
service StreamServer{
  // 服务端流式rpc，在响应数据前添加stream
  rpc ListValue(StreamRequest)returns(stream StreamResponse){};
}
```

2. 生成 go 文件
```
protoc --go_out=plugins=grpc:./ ./server_stream.proto
```
3. 编写 server 端代码
    - 可以通过流，写入数据
```
package main

import (
	grpc "google.golang.org/grpc"
	mylib "grpc-client/my_lib"
	"log"
	"net"
	"strconv"
	"time"
)

type StreamServer struct {}

// ListValue 实现ListValue方法
func (ss *StreamServer)ListValue(req *mylib.StreamRequest, stream mylib.StreamServer_ListValueServer) error {
	for i := 0; i < 10; i++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := stream.Send(&mylib.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(i),
		})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

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

	//- 将 Service 结构体注册进 grpc server
	mylib.RegisterStreamServerServer(grpcServer, &StreamServer{})


	//- 开始启动服务端，通过 listener，grpc server 开启监听
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

```
4. 编写 client 端代码
    - 可以读取流中的数据
```
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-client/my_lib"
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

```


#### 3. 客户端流式 RPC 使用

**客户端流式RPC**：与服务端流式RPC相反，客户端不断的向服务端发送数据流，而在发送结束后，由服务端返回一个响应。



情景模拟：客户端大量数据上传到服务端。


使用步骤：
1. 创建proto文件, 函数接口接收参数 stream 数据流，客户端可以写入，服务端接受
```
syntax = "proto3"; //  使用 protobuf3

//  package 名称
package my_lib;

// import path
option go_package = "./";

//  流式请求
message StreamRequest{
  string stream_data = 1;
}

//  普通响应
message SimpleResponse{
  //  响应状态码
  int64 status = 1;
  //  响应值
  string value = 2;
}


// 定义服务（可定义多个服务,每个服务可定义多个接口）
service StreamClient{
  // 客户端流式rpc，在请求的参数前添加stream, 也就是客户端的stream，客户端写，服务端读
  rpc RouteList (stream StreamRequest) returns (SimpleResponse){};
}
```
2. 生成 go 文件
```
protoc --go_out=plugins=grpc:./ ./client_stream.proto
```
3. 编写 server 端代码
    - 可以通过流，接收数据
    - 接收完数据后，发送响应给客户端
```
package main

import (
	grpc "google.golang.org/grpc"
	mylib "grpc-client/my_lib"
	"io"
	"log"
	"net"
)

type StreamClient struct {}

// server 端 实现 service 中的方法 RouteList
func (sc StreamClient) RouteList(stream mylib.StreamClient_RouteListServer) error  {
	//  服务端，拿到客户端提供的 数据流， 可以读取其中的数据
	for {
		// 从数据流中读取数据
		res, err := stream.Recv()
		if err == io.EOF {
			//  如果 读取完成，则发送 响应 后 ，关闭 数据流
			return stream.SendAndClose(&mylib.SimpleResponse{
				Status: 200,
				Value: "OK",
			})
		} else if err != nil {
			return err
		}

		//  打印从数据流中读取的数据
		log.Println(res.StreamData)
	}
}

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

	//- 将 Service 结构体注册进 grpc server
	mylib.RegisterStreamClientServer(grpcServer, &StreamClient{})


	//- 开始启动服务端，通过 listener，grpc server 开启监听
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

```
4. 编写 client 端代码
    - 可以读取流中的数据
    - 发送完数据后，等待接收服务端发来的响应
```
package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-client/my_lib"
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

```

#### 4. 双向流式 RPC 使用

**双向流式RPC**：客户端和服务端双方使用读写流去发送一个消息序列，两个流独立操作，双方可以同时发送和同时接收。

情景模拟：双方对话（可以一问一答、一问多答、多问一答，形式灵活）。

使用步骤：
1. 创建proto文件, 函数接口接收参数 stream 数据流，客户端可以写入，服务端接受
```
syntax = "proto3"; //  使用 protobuf3

//  package 名称
package my_lib;

// import path
option go_package = "./";


// 定义流式请求信息
message StreamRequest{
  //流请求参数
  string question = 1;
}

// 定义流式响应信息
message StreamResponse{
  //流响应数据
  string answer = 1;
}

service Stream{
  // 双向流式rpc，同时在请求参数前和响应参数前加上stream
  rpc Conversations(stream StreamRequest) returns(stream StreamResponse){};
}
```
2. 生成 go 文件
```
protoc --go_out=plugins=grpc:./ ./both_stream.proto
```
3. 编写 server 端代码
    - 可以通过server端提供的数据流，写入数据
    - 也可以通过client端提供的数据流，读取数据
```
package main

import (
	grpc "google.golang.org/grpc"
	mylib "grpc-client/my_lib"
	"io"
	"log"
	"net"
	"strconv"
)

type BothStream struct{}

func (bs *BothStream) Conversations(stream mylib.BothStream_ConversationsServer) error {
	//  用于计数请求
	n := 1

	//  开始读写数据
	for {
		//  读取数据，读取问题
		recv, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		//  写入数据，回答问题
		err = stream.Send(&mylib.StreamResponse{
			Answer: "from client stream , server anwser : the " + strconv.Itoa(n) + " question is " + recv.Question,
		})
		if err != nil {
			return err
		}

		n ++
		log.Printf("from stream client question is: %s", recv.Question)
	}

}

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

	//- 将 Service 结构体注册进 grpc server
	mylib.RegisterBothStreamServer(grpcServer, &BothStream{})

	//- 开始启动服务端，通过 listener，grpc server 开启监听
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

```
4. 编写 client 端代码
    - 可以通过server端提供的数据流，写入数据
    - 也可以通过client端提供的数据流，读取数据
```
package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-client/my_lib"
	"io"
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
	grpcClient := pb.NewBothStreamClient(conn)


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

```



### GRPC 的高级使用


#### 1. 超时设置

##### GRPC 超时的原理

超时设置，指的是某个请求到服务器后，所用时间的最大值，超过这个时间视为超时。

GRPC 如果没有设置超时时间，那么默认的超时时间很长，所有在运行的请求都占用大量资源且可能运行很长的时间，导致服务资源损耗过高，使得后来的请求响应过慢，甚至会引起整个进程崩溃。



**设置超时的方法**

- 客户端在发送请求时，设置超时时间
- 客户端在调用 service 中的方法时，会传入 context.Context 上下文
- 在 context 中设置超时时间
- 将设置好的超时时间的 context 传入被调用的 service 中的方法


##### context 包简介

**context 包简介：**

Go 语言中的每一个请求的都是通过一个单独的 Goroutine 进行处理的，HTTP/RPC 请求的处理器往往都会启动新的 Goroutine 访问数据库和 RPC 服务，我们可能会创建多个 Goroutine 来处理一次请求，而 **Context 的主要作用就是在不同的 Goroutine 之间同步请求特定的数据、取消信号以及处理请求的截止日期**。

原理：
- 每一个 Context 都会从最顶层的 Goroutine 一层一层传递到最下层，这也是 Golang 中上下文最常见的使用方式
- 上下层 Goroutine 的意思是，一个 Goroutine 中有创建了一个新的 Goroutine。
- 如果没有 Context，当上层执行的操作出现错误时，下层其实不会收到错误而是会继续执行下去。

**使用方法是：**

- 首先，服务器程序为每个接受的请求创建一个 Context 实例（称为根 context，通过 context.Background() 方法创建）；
- 之后的 goroutine 接受根 context 的一个派生 Context 对象。比如通过调用根 context 的 WithCancel 方法，创建子 context；
goroutine 通过 context.Done() 方法监听取消信号。func Done() <- chan struct{} 是一个通信操作，会阻塞 goroutine，直到收到取消信号解除阻塞。（可以借助 select 语句，如果收到取消信号，就退出 goroutine；否则，默认子句是继续执行 goroutine）；
当一个 Context 被取消（比如执行了 cancelFunc()），那么该 context 派生出来的 context 也会被取消。




##### GRPC 设置超时

(以简单的 RPC 为例，即没有数据流的情况)
1. 客户端请求设置超时时间
    - 把客户端传入的 context 上下文变量，设置超时时间
    ```
	clientDeadLine := time.Now().Add(time.Duration(3 * time.Second))
	ctx, cancelFunc := context.WithDeadline(context.Background(), clientDeadLine)
	defer cancelFunc()
	response, err := grpcClient.FirstRpc(ctx, &request)
    ```
    - 在响应处理时，添加超时检测
    ```
	response, err := grpcClient.FirstRpc(ctx, &request)
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
    ```

2. 服务端判断请求是否超时
    - 当请求超时后，服务端应该停止正在进行的操作，避免资源浪费。
    - 一般地，在写Service的函数接口时前进行超时检测，发现超时就停止工作。
```
//  实现 service 中的 函数接口
func (t *test) FirstRpc(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	//  函数接口 开始处，进行超时处理
	//  设置一个 resChan，用来查看是否超时，如果超时，会从 ctx.Done() chan 中读取到数据，如果在超市前，resChan 返回结果，那么正常返回
	resChan := make(chan *pb.TestResponse, 1)

	//  handler 函数用于判断是否超时，并且如果没有超时则返回正常响应
	go handler(ctx, req, resChan)

	//  这里开始监听，是否在超时时间内处理完成
	select {
	//  如果从 resChan 中返回结果，则正常处理
	case res := <- resChan:
		return res, nil
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "Client cancelled, abandoning.")
	}
}

func handler(ctx context.Context, req *pb.TestRequest, resChan chan<- *pb.TestResponse) {
	//  开启监听，判断是否超时
	select {
	// 如果 从 ctx.Done() 中拿到数据，则表示超时
	case <-ctx.Done():
		//  超时处理
		log.Println(ctx.Err())
		runtime.Goexit()
	//  处理请求，并返回结果，这里模拟 需要处理 4s，即超时
	case <-time.After(4 * time.Second):
		//  初始化 response
		res := pb.TestResponse{
			Status: 200,
			Value: "Hello " + req.GetName() + ", id " + string(req.GetId()),
		}
		//  处理完成后，往 resChan 中返回结果
		resChan <- &res
	}
}
```



#### 2. TLS 认证 和 Token 认证



##### a. TLS 认证

TLS（Transport Layer Security，安全传输层)，TLS是建立在传输层TCP协议之上的协议，服务于应用层，它的前身是SSL（Secure Socket Layer，安全套接字层），它实现了将应用层的报文进行加密后再交由TCP进行传输的功能。



使用步骤：
1. 生成私钥（这里就不在生成根证书了）
```
openssl genrsa -out server.key 2048
```
2. 生成公钥, 主要要添加`-addext "subjectAltName = DNS:YOUR_DOMAIN_NAME"`
```
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650 -addext "subjectAltName = DNS:gsh.com"
```
3. 服务端在创建时，需要添加证书
    - 创建 TLS Server 实例
    - 再创建grpcServer的时候，将TLS实例传进去
    - `credentials.NewServerTLSFromFile`：从输入证书文件和密钥文件为服务端构造TLS凭证
    - `grpc.Creds`：返回一个ServerOption，用于设置服务器连接的凭证。
```
//  创建 TLS 实例， 并传入 公钥和私钥
creds, err := credentials.NewServerTLSFromFile("../tls-token/server.pem", "../tls-token/server.key")
if err != nil {
	log.Fatalf("Failed to generate credentials %v", err)
}

//  新建 grpc 服务器实例, 并将 tls 实例 传入
grpcServer := grpc.NewServer(grpc.Creds(creds))

log.Println(Address + "net.listening with tls ...")
```
4. 客户端配置TLS连接
    - 创建 TLS Client 实例
    - 在连接到服务器时，将 TLS 实例传进去
    - `credentials.NewClientTLSFromFile`：从输入的证书文件中为客户端构造TLS凭证。
    - grpc.WithTransportCredentials`：配置连接级别的安全凭证（例如，TLS/SSL），返回一个DialOption，用于连接服务器。
```
//  从输入的证书文件中为客户端构造TLS凭证, 这里的 domain name 写在创建公钥时的 Common Name
cred, err := credentials.NewClientTLSFromFile("../tls-token/server.pem", "gsh.com")
if err != nil {
	log.Fatalf("Failed to create TLS credentials %v", err)
}

// 连接服务器
conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(cred))
if err != nil {
	log.Fatalf("grpc.Dial err: %v", err)
}
```



##### b. Token 认证

Token 认证，类似于 HTTP 请求的 JWT TOKEN 认证，也就是在
客户端发请求时，添加Token到上下文 context.Context 中，服务器接收到请求，先从上下文中获取 Token 验证，验证通过才进行下一步处理

**使用前需要了解：**

- gRPC 中默认定义了 `PerRPCCredentials` 接口
```
type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
}
```
- 在源代码中，`PerRPCCredentials` 的描述为：`PerRPCCredentials` 定义了通用的接口，目的是用来将 Credentials 信息添加到每一个 RPC 中。
- `PerRPCCredentials` 接口中定义了两个方法，分别是：
    - `GetRequestMetadata`: 获取当前请求认证所需的元数据。uri 是请求入口点的 URI， 也就是最终返回一个 map 格式的 Token 数据
    - `RequireTransportSecurity`：是否需要基于 TLS 认证进行安全传输， 返回 True，则需要进行 TLS 认证
- 也就是说，如果需要token验证，要自己定一个 token 的结构体，然后实现结构体的这两个方法，也就实现了这个接口




使用步骤：
1. 客户端请求添加Token到上下文中
    - 定义 token 结构体
    - 实现 token 结构体的 `GetRequestMetadata` 和 `RequireTransportSecurity` 方法
    - 将 token 实例，添加到Dial方法中
```
//  定义 TOKEN 结构体
type Token struct {
	TokenID string
	TokenName string
}

//  实现 Token 的两个方法
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"token_id": t.TokenID,
		"token_name": t.TokenName,
	}, nil
}

func (t *Token) RequireTransportSecurity() bool {
	return true
}


func main() {

	//  从输入的证书文件中为客户端构造TLS凭证, 这里的 domain name 写在创建公钥时的 Common Name
	creds, err := credentials.NewClientTLSFromFile("../tls-token/server.pem", "gsh.com")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	// 创建 TOKEN 实例
	token := Token{
		TokenID:   "12345",
		TokenName: "gsh_token",
	}

	// 连接服务器, 添加 tls 证书 和 Token
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

    ....
}
```
2. 服务端拿到 Token，进行验证
    - 在 Service 的方法中，要添加验证 token 的逻辑代码
```
//  实现 service 中的 函数接口
func (service *test) FirstRpc(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	//  验证 token
	err := CheckToken(ctx)
	if err != nil {
		return nil, err
	}
	
	// 初始化 response
	response := pb.TestResponse{
		Status: 200,
		Value:  "Success",
	}

	return &response, nil
}

//  验证 Token 的函数
func CheckToken(ctx context.Context) error {
	//从上下文中获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "获取Token失败")
	}

	var (
		TokenName	string
		TokenID		string
	)

	//  从 metadata 中 拿到 token 的字段
	if value, ok := md["token_id"]; ok {
		TokenID = value[0]
	}
	if value, ok := md["token_name"]; ok {
		TokenName = value[0]
	}

	//  比较 token 是否正确
	if TokenID != "12345" || TokenName != "gsh_token" {
		return status.Errorf(codes.Unauthenticated, "Token无效: app_id=%s, app_secret=%s", appID, appSecret)
	}

	return nil
}
```


##### c. 在 server 端使用拦截器，校验 Token

上面的代码有一些问题，那就是在每一个 service 中的函数方法都需要手动添加一遍 Token 校验，重复繁琐，

使用方法：
- `grpc.UnaryServerInterceptor`：为一元拦截器，只会拦截简单 RPC 方法。
- `grpc.StreamInterceptor`: 流式 RPC 方法需要使用流式拦截器进行拦截。

```
// 普通方法：一元拦截器（grpc.UnaryInterceptor）
var interceptor grpc.UnaryServerInterceptor
interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//拦截普通方法请求，验证Token
	err = CheckToken(ctx)
	if err != nil {
		return nil,err
	}
	// 继续处理请求
	return handler(ctx, req)
}

//  新建 grpc Server 实例, 并添加拦截器
grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))
```

### GRPC - HealthCheck

#### 健康检查简介

gRPC 有一个标准的健康检测协议，在 gRPC 的所有语言实现中基本都提供了生成代码和用于设置运行状态的功能。

**HealthCheck 的用途**

- 主动健康检查 healthcheck，可以在服务提供者服务不稳定时，被消费者所感知，临时从负载均衡中摘除，减少错误请求。
- 当服务提供者重新稳定后，healthcheck 成功，重新加入到消费者的负载均衡，恢复请求。
- healthcheck，同样也被用于外挂方式的容器健康检测，或者流量检测 (k8s liveness & readiness)。


#### 健康检查应用

**优雅启动**

启动流程如下图：
![grpc HealthCheck优雅启动流程](https://github.com/Nevermore12321/LeetCode/blob/blog/go%E8%BF%9B%E9%98%B6%E8%AE%AD%E7%BB%83%E8%90%A5/grpc_healthCheck_%E4%BC%98%E9%9B%85%E5%90%AF%E5%8A%A8.png?raw=true)

1. Provider 启动，k8s 中的启动脚本会定时去检查服务的健康检查接口
2. 健康检查通过之后，服务注册脚本向注册中心注册服务（rpc://ip:port）
3. 消费者定时从服务注册中心获取服务方地址信息
4. 获取成功后，会定时的向服务方发起健康检查，健康检查通过后才会向这个地址发起请求
    1. 在运行过程中如果健康检查出现问题，会从消费者本地的负载均衡中移除


**优雅终止**


优雅终止的流程图：
![grpc HealthCheck优雅终止流程](https://github.com/Nevermore12321/LeetCode/blob/blog/go%E8%BF%9B%E9%98%B6%E8%AE%AD%E7%BB%83%E8%90%A5/grpc_HealthCheck_%E4%BC%98%E9%9B%85%E7%BB%88%E6%AD%A2%E6%B5%81%E7%A8%8B.png?raw=true)


1. 触发下线操作: 首先用户在发布平台点击发版/下线按钮
2. 发布部署平台向注册中心发起服务注销请求，在注册中心下线服务的这个节点
    - 这里在发布部署平台实现有个好处，不用每个应用都去实现一遍相同的逻辑
    - 在应用受到退出信号之后由应用主动发起注销操作也是可以的
3. 注册中心下线应用之后，消费者会获取到服务注销的事件，然后将服务方的节点从本地负载均衡当中移除
    - 注意这一步操作会有一段时间，下面的第五步并不是这一步结束了才开始，而是直接开始计算时间。
4. 发布部署平台向应用发送 SIGTERM 信号，应用捕获到之后执行
    - 将健康检查接口设置为不健康，返回错误
        - 这个时候如果消费者还在调用应用程序，调用健康检查接口发现无法通过，也会将服务节点从本地负载均衡当中移除
    - 调用 grpc/http 的 shutdown 接口，并且传递超时时间，等待连接全部关闭后退出
        - 这个超时时间一般为 2 个心跳周期
5. 发布部署平台如果发现应用程序长时间没有完成退出，发送 SIGKILL 强制退出应用
    - 这个超时时间根据应用进行设置一般为 10 - 60s