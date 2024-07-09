/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-07 23:30:16
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-09 11:24:29
 * @FilePath: /go-learn/gin/project/ranking/controllers/common.go
 * @Description:
 *
 */

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type JsonErrStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(http.StatusOK, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonErrStruct{Code: code, Msg: msg}
	c.JSON(http.StatusOK, json)
}
