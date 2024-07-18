/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-14 10:39:22
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-18 10:03:02
 * @FilePath: /go-learn/grpc/hello-client/main.go
 * @Description:
 *
 */
package main

import (
	"context"
	"fmt"
	pb "hello-client/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
}

type ClientTokenAuth struct{}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "awsl1414",
		"appKey": "123456",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return true
}

func main() {
	// tls 认证
	creds, _ := credentials.NewClientTLSFromFile("../key/test.pem", "*.sailor.work") // serverNameOverride 应从浏览器获取

	//连接到 server 端，此处启用安全传输
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	// 连接 server 端，此处禁用安全传输，没有加密和验证 grpc.Dial() 已弃用
	// conn, err := grpc.NewClient("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "嗨嗨嗨"})

	fmt.Println(resp.GetResposeMsg())
}
