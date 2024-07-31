/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-29 22:43:29
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-31 21:26:15
 * @FilePath: /go-learn/grpc/hello/cmd/server/main.go
 * @Description:
 *
 */
package main

import (
	"fmt"
	pb "hello/pkg/hello"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 开启端口
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen create err: ", err)
	}
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 在grpc服务端注册自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &pb.Server{})
	// 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("failed to server: ", err)
	}

}
