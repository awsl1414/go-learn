/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-09 11:27:33
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-09 13:31:05
 * @FilePath: /go-learn/gin/project/ranking/controllers/order.go
 * @Description:
 *
 */

package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetList(c *gin.Context) {
	cid := c.PostForm("cid")
	name := c.DefaultPostForm("name", "zhangsan")
	ReturnSuccess(c, 0, cid, name, 1)
}
