/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-06 16:12:02
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-07 16:51:34
 * @FilePath: /go-learn/http/base/main.go
 * @Description:
 *
 */

package main

import (
	"fmt"
	"io"
	"net/http"
)

func get() {

	// request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := http.DefaultClient.Do(request)
	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// content, err := io.ReadAll(resp.Body)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s", content)

	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("MethodGet：%s", content)

}

func post() {
	resp, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("MethodPost：%s", content)
}

func put() {
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(request) // 相当于浏览器按下enter
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("MethodPut：%s", content)

}

func delete() {
	request, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	content, _ := io.ReadAll(resp.Body)

	fmt.Printf("MethodDelete：%s", content)

}
func main() {
	get()
	fmt.Println("----------------------------------------------------")
	post()
	fmt.Println("----------------------------------------------------")
	put()
	fmt.Println("----------------------------------------------------")
	delete()
}
