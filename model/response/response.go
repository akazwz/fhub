package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

func Ok(code int, data interface{}, msg string, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Created(code int, data interface{}, msg string, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusCreated, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func BadRequest(code int, data interface{}, msg string, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Unauthorized(code int, msg string, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
		Code: code,
		Msg:  msg,
	})
}

func Forbidden(code int, msg string, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, Response{
		Code: code,
		Msg:  msg,
	})
}

func NotFound(code int, msg string, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, Response{
		Code: code,
		Msg:  msg,
	})
}
