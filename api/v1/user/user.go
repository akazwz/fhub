package user

import (
	"log"

	"github.com/akazwz/fhub/api/v1"
	"github.com/akazwz/fhub/model"
	"github.com/akazwz/fhub/model/request"
	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/service"
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

func GetUserProfileByToken(c *gin.Context) {
	claims, _ := c.Get("claims")
	customClaims := claims.(*model.MyCustomClaims)
	userUID := customClaims.UID

	err, user := service.GetUserProfileByUID(userUID.String())
	if err != nil {
		response.BadRequest(v1.CodeErrorNoSuchUser, "没有此用户", c)
		return
	}
	userRes := response.UserResponseProfile{
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		Gender:    user.Gender,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
	}
	response.Ok(v1.CodeSuccessGetUserProfile, userRes, "获取用户资料成功", c)
}
