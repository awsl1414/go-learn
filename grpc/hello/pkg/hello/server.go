/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-29 22:39:16
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-31 21:19:26
 * @FilePath: /go-learn/grpc/hello/pkg/hello/server.go
 * @Description:
 *
 */
package hello

import context "context"

type Server struct {
	UnimplementedSayHelloServer
}

// 重写 SayHello 函数
func (s Server) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{
		ResponseMsg: "hello " + req.RequestName,
	}, nil
}
