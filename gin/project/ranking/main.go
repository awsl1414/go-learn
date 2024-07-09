/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-07 21:16:07
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-07 22:00:13
 * @FilePath: /go-learn/gin/project/ranking/main.go
 * @Description:
 *
 */

package main

import (
	"ranking/router"
)

func main() {
	r := router.Router()

	r.Run("127.0.0.1:9999")
}
