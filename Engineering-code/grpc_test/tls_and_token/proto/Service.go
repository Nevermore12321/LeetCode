package proto

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type MyService struct {}

//  实现 service 中的 函数接口
func (service *MyService) SayHello(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	//  验证 token
	err := CheckToken(ctx)
	if err != nil {
		return nil, err
	}

	// 初始化 response
	response := TestResponse{
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
		return status.Errorf(codes.Unauthenticated, "Token无效: app_id=%s, app_secret=%s", TokenID, TokenName)
	}

	return nil
}