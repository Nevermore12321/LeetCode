package proto

import "context"

type ServerInterceptor struct {}

func (service *ServerInterceptor) SayHello(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	// 初始化 response
	response := TestResponse{
		Status: 200,
		Value:  "Success",
	}

	return &response, nil
}