/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-07 16:56:16
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-07 20:22:19
 * @FilePath: /go-learn/http/request/main.go
 * @Description:
 *
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func printBody(r *http.Response) {
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func clientDo(request *http.Request) *http.Response {
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	return resp
}

func requestByParams(baseURL string) *http.Response {
	request, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("name", "xiaoming")
	params.Add("age", "18")

	// fmt.Println(params.Encode())

	request.URL.RawQuery = params.Encode()

	return clientDo(request)
}

func requestByHead(baseURL string) *http.Response {
	request, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("user-agent", "chrome")

	return clientDo(request)
}

func main() {
	// 如何设置请求的查询参数，http://httpbin.org/get?name=test&age=18
	// 如何定制请求头，如修改 user-agent
	var baseURL string = "http://httpbin.org/get"

	respose1 := requestByParams(baseURL)
	respose2 := requestByHead(baseURL)

	printBody(respose1)
	printBody(respose2)
}
