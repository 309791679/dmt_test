package handler

import (

	"OrderTakeaway/proto"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"

	context2 "context"
)


//第六步 实现接口具体功能
type UsercenterServer struct{}

func (s UsercenterServer) CreateUser(ctx context2.Context, request *grpc_.UniversalUserRequest) (*grpc_.UniversalResponse, error) {
	panic("implement me")
}

func (s UsercenterServer) ChangeUser(ctx context2.Context, request *grpc_.UniversalUserRequest) (*grpc_.UniversalResponse, error) {
	panic("implement me")
}

func (s UsercenterServer) QueryUser(ctx context2.Context, request *grpc_.UniversalUserRequest) (*grpc_.UniversalResponse, error) {

	fmt.Println("不是go-zero,进入延时时间",time.Now().String())
	time.Sleep(time.Duration(20*time.Second))
	fmt.Println("执行成功",time.Now().String())
	return &grpc_.UniversalResponse{},status.Error(codes.OK, "不支持的模式")
}

