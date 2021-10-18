package proto

import "golang.org/x/net/context"

// 定义服务
type MyService struct {}

//  实现 service 中的 函数接口
func (t *MyService) SayHello(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	//  初始化 response
	res := TestResponse{
		Status: 200,
		Value: "Hello " + req.GetName() + ", id " + string(req.GetId()),
	}
	//  返回
	return &res, nil
}