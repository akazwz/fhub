package user

import (
	"os"
	"strconv"
	"time"

	v1 "github.com/akazwz/gin/api/v1"
	"github.com/akazwz/gin/global"
	"github.com/akazwz/gin/middleware"
	"github.com/akazwz/gin/model"
	"github.com/akazwz/gin/model/request"
	"github.com/akazwz/gin/model/response"
	"github.com/akazwz/gin/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CreateTokenByUsernamePwd 账号密码登录获取token
func CreateTokenByUsernamePwd(c *gin.Context) {
	var userLogin request.LoginByUsernamePwd

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		response.BadRequest(v1.CodeErrorBindJson, "参数错误", c)
		return
	}

	err, user := service.LoginByUsernamePwdService(userLogin)
	if err != nil {
		response.BadRequest(v1.CodeErrorLogin, "登录失败", c)
		return
	}
	TokenNext(c, *user)
}

// TokenNext 生成并返回token
func TokenNext(c *gin.Context, user model.User) {
	j := middleware.NewJWT()

	bufferTime := global.CONF.JWT.BufferTime
	/* 环境变量覆盖 缓冲时长 */
	if len(os.Getenv("JWT_BUFFER_TIME")) > 0 {
		bufferTimeInt, err := strconv.ParseInt(os.Getenv("JWT_BUFFER_TIME"), 10, 64)
		if err != nil {
			return
		}
		bufferTime = bufferTimeInt
	}

	expiresTime := global.CONF.JWT.ExpiresTime
	/* 环境变量覆盖 过期时长 */
	if len(os.Getenv("JWT_EXPIRES_TIME")) > 0 {
		expiresTimeInt, err := strconv.ParseInt(os.Getenv("JWT_EXPIRES_TIME"), 10, 64)
		if err != nil {
			return
		}
		expiresTime = expiresTimeInt
	}

	claims := model.MyCustomClaims{
		UID:        user.UID,
		Username:   user.Username,
		BufferTime: bufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + expiresTime,
			Issuer:    "zwz",
		},
	}

	token, err := j.NewToken(claims)
	if err != nil {
		response.BadRequest(v1.CodeErrorCreateToken, "获取token失败", c)
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

	response.Created(v1.CodeSuccessLogin, response.LoginResponse{
		User:      userRes,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt,
	}, "登录成功", c)
}
