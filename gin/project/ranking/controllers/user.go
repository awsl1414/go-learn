/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-09 10:44:24
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-09 13:21:25
 * @FilePath: /go-learn/gin/project/ranking/controllers/user.go
 * @Description:
 *
 */
package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSuccess(c, 0, name, id, 1)
}

func (u UserController) GetList(c *gin.Context) {
	ReturnError(c, 4004, "没有相关信息")
}