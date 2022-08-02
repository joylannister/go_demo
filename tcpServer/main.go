package main

import (
	"context"
	"demo/mysql"
	"demo/pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) GetUserInfo(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return mysql.GetUserInfo(in)
}
func main() {
	mysql.Init()
	defer mysql.Close()
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                  // 创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
