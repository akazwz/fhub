package user

import (
	"log"

	"github.com/akazwz/gin/api/v1"
	"github.com/akazwz/gin/model/request"
	"github.com/akazwz/gin/model/response"
	"github.com/akazwz/gin/service"
	"github.com/gin-gonic/gin"
)

// CreateUserByUsernamePwd 通过用户名密码注册
func CreateUserByUsernamePwd(c *gin.Context) {
	/* 获取参数 */
	var register request.RegisterByUsernamePwd
	err := c.ShouldBindJSON(&register)
	if err != nil {
		log.Println("bind json error")
		response.BadRequest(v1.CodeErrorBindJson, "参数错误", c)
		return
	}
	/* 数据库新增用户 */
	err = service.RegisterByUsernamePwdService(register)
	if err != nil {
		log.Println("create user error")
		response.BadRequest(v1.CodeErrorCreateUser, "注册失败", c)
		return
	}
	/* 注册成功， 返回用户名 */
	userRes := response.UserResponseCreatedByUsernamePwd{Username: register.Username}
	response.Created(v1.CodeSuccessCreateUser, userRes, "注册成功", c)
}
