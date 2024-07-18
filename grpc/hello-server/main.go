/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-14 10:39:28
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-18 10:08:28
 * @FilePath: /go-learn/grpc/hello-server/main.go
 * @Description:
 *
 */
package main

import (
	"context"
	"errors"
	"fmt"
	pb "hello-server/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

// 业务
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输token")
	}
	var appId string
	var appKey string

	// 元数据是放在http头传输的，http头的键是小写的
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}

	if appId != "awsl1414" || appKey != "123456" {
		return nil, errors.New("token不正确")
	}

	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	// tls 认证
	creds, _ := credentials.NewServerTLSFromFile("../key/test.pem", "../key/test.key")

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
