package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一返回体 {code, data, msg}
type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"` //omitempty nil or default
	Msg  string      `json:"msg,omitempty"`
}

// NotFound 404
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		Code: 4040,
		Msg:  "404 not found",
	})
}

// Ok 成功 200
func Ok(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// Created 创建成功 201
func Created(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusCreated, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// BadRequest 请求失败 400
func BadRequest(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		Code: code,
		Msg:  msg,
	})
}

// Unauthorized 未授权 401
func Unauthorized(code int, message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code: code,
		Msg:  message,
	})
}

// Forbidden 禁止访问， 权限不足  403
func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{
		Msg: "Permission Denied",
	})
}
