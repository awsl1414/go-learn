/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-14 10:39:22
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-14 16:24:56
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
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接server端，此处禁用安全传输，没有加密和验证 grpc.Dial() 已弃用
	conn, err := grpc.NewClient("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "嗨嗨嗨"})

	fmt.Println(resp.GetResposeMsg())
}
