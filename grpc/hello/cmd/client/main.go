/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-29 22:50:31
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-08 16:04:16
 * @FilePath: /go-learn/grpc/hello/cmd/client/main.go
 * @Description:
 *
 */
package main

import (
	"context"
	"fmt"
	pb "hello/pkg/hello"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 定义一个接口 PerRPCCredentials，用于每次 RPC 调用时传递认证信息
type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
}

// 定义一个结构体 ClientTokenAuth，实现 PerRPCCredentials 接口
type ClientTokenAuth struct{}

// 实现 GetRequestMetadata 方法，返回每次 RPC 调用时需要的认证信息
func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "awsl1414",
		"appKey": "123456",
	}, nil
}

// 实现 RequireTransportSecurity 方法，返回是否需要传输安全
func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 定义 gRPC 连接选项
	var opts []grpc.DialOption
	// 使用空加密和验证（不安全的传输）
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 添加自定义的每次 RPC 调用认证信息
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	// 连接 server 端，此处禁用安全传输
	// conn, err := grpc.Dial(":8888", opts...)
	conn, err := grpc.NewClient(":8888", opts...)
	if err != nil {
		// 如果连接失败，记录错误并退出
		log.Fatalf("did not connect: %v", err)
	}

	// 在函数结束时关闭连接
	defer conn.Close()

	// 使用连接创建一个新的 SayHello 客户端
	client := pb.NewSayHelloClient(conn)

	// 调用 SayHello 方法并传递请求
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "awsl1414"})
	if err != nil {
		// 如果调用失败，记录错误
		log.Fatalf("could not greet: %v", err)
	}

	// 输出响应的消息
	fmt.Println(resp.GetResponseMsg())
}
