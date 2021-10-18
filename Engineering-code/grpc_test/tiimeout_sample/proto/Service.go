package proto

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"runtime"
	"strconv"
	"time"
)

// 定义服务
type MyService struct {}

//  实现 service 中的 函数接口
func (t *MyService) SayHello(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	//  函数接口 开始处，进行超时处理
	//  设置一个 resChan，用来查看是否超时，如果超时，会从 ctx.Done() chan 中读取到数据，如果在超时前，resChan 返回结果，那么正常返回
	resChan := make(chan *TestResponse, 1)

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

func handler(ctx context.Context, req *TestRequest, resChan chan<- *TestResponse) {
	//  开启监听，判断是否超时
	select {
	// 如果 从 ctx.Done() 中拿到数据，则表示超时
	case <-ctx.Done():
		//  超时处理
		log.Println(ctx.Err())
		runtime.Goexit()
	//  处理请求，并返回结果，这里模拟 需要处理 4s，即超时
	case <-time.After(2 * time.Second):
		//  初始化 response
		res := TestResponse{
			Status: 200,
			Value: "Hello " + req.GetName() + ", id " + strconv.FormatInt(req.GetId(), 10),
		}
		//  处理完成后，往 resChan 中返回结果
		resChan <- &res
	}
}