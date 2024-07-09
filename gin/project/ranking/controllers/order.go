/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-09 11:27:33
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-09 16:02:31
 * @FilePath: /go-learn/gin/project/ranking/controllers/order.go
 * @Description:
 *
 */

package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

type Search struct {
	Cid  int    `json:"cid"`
	Name string `json:"name"`
}

func (o OrderController) GetList(c *gin.Context) {
	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)

	search := &Search{}

	err := c.BindJSON(&search)

	if err != nil {
		ReturnError(c, 4001, gin.H{"err": err})
		return
	}

	// ReturnSuccess(c, 0, param["cid"], param["name"], 1)
	ReturnSuccess(c, 0, search.Cid, search.Name, 1)

}
