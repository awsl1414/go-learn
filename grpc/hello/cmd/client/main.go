/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-29 22:50:31
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-31 21:22:32
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

func main() {

	var opts []grpc.DialOption
	// 空加密和验证
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 连接 server 端，此处禁用安全传输
	conn, err := grpc.NewClient(":8888", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)

	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "awsl1414"})

	fmt.Println(resp.GetResponseMsg())

}
