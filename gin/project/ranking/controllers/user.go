/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-09 10:44:24
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-10 16:29:55
 * @FilePath: /go-learn/gin/project/ranking/controllers/user.go
 * @Description:
 *
 */
package controllers

import (
	"fmt"
	"ranking/pkg/logger"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSuccess(c, 0, name, id, 1)
}

func (u UserController) GetList(c *gin.Context) {
	logger.Write("日志信息", "user")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常", err)
		}
	}()
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	ReturnError(c, 4004, num3)
}
