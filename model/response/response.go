package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"` //omitempty nil or default
	Msg  string      `json:"msg,omitempty"`
}

const (
	SUCCESS = 2000
	ERROR   = 4000
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		Code: ERROR,
		Msg:  "404 not found",
	})
}
