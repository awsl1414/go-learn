/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-29 22:39:16
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-31 21:35:07
 * @FilePath: /go-learn/grpc/hello/pkg/hello/server.go
 * @Description:
 *
 */
package hello

import (
	context "context"
	"errors"

	"google.golang.org/grpc/metadata"
)

type Server struct {
	UnimplementedSayHelloServer
}

// 重写 SayHello 函数（业务）
func (s Server) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	// token认证

	// 获取元数据
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, errors.New("未传输token")
	}
	var appId, appKey string

	// 元数据是放在http头里传输的，http头是小写
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	if appId != "awsl1414" || appKey != "123456" {
		return nil, errors.New("token不正确")
	}

	return &HelloResponse{
		ResponseMsg: "hello " + req.RequestName,
	}, nil
}
