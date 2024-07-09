/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-09 10:47:53
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-09 13:36:19
 * @FilePath: /go-learn/gin/project/ranking/router/routers.go
 * @Description:
 *
 */
package router

import (
	"ranking/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/info/:id/:name", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetList)

	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}
