package auth

import (
	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/service"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	authService := service.AuthService{}

	uidAny, _ := c.Get("uid")
	uid := uidAny.(string)

	user := authService.FindUserByID(uid)
	/* 用户不存在 */
	if user == nil {
		response.BadRequest(400, nil, "账户不存在", c)
		return
	}
	capacity := authService.FindCapacityByUID(user.ID)
	/* 用户不存在 */
	if capacity == nil {
		response.BadRequest(400, nil, "账户容量不存在", c)
		return
	}
	response.Ok(400, gin.H{
		"user":     user,
		"capacity": capacity,
	}, "success", c)
}
