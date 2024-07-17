/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-14 10:39:28
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-17 23:05:28
 * @FilePath: /go-learn/grpc/hello-server/main.go
 * @Description:
 *
 */
package main

import (
	"context"
	"fmt"
	pb "hello-server/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {

	creds, _ := credentials.NewServerTLSFromFile("grpc/key/test.pem", "grpc/key/test.key")

	// 开启端口
	listen, _ := net.Listen("tcp", "127.0.0.1:9090")

	// 创建grpc服务
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// 在grpc服务端注册自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
