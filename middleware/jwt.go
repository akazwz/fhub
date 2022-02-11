package middleware

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/akazwz/gin/global"
	"github.com/akazwz/gin/model"
	"github.com/akazwz/gin/model/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that`s not even a token")
	TokenInvalid     = errors.New("could`t handle this token")
)

//NewJWT 生成新的 JWT
func NewJWT() *JWT {
	/* 配置文件读取 signKey */
	signKey := global.CONF.JWT.SigningKey
	/* 环境变量覆盖 signKey */
	if len(os.Getenv("JWT_SIGNING_KEY")) > 0 {
		signKey = os.Getenv("JWT_SIGNING_KEY")
	}
	return &JWT{SigningKey: []byte(signKey)}
}

// NewToken 生成token
func (j *JWT) NewToken(claims model.MyCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenStr string) (*model.MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &model.MyCustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})

	log.Println(token)
	log.Println(err)

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*model.MyCustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

const (
	CodeTokenNUll        = 4010
	CodeTokenExpired     = 4011
	CodeTokenNotValidYet = 4012
	CodeTokenMalformed   = 4013
	CodeTokenInvalid     = 4014
	CodeNoSuchUser       = 4015
	CodePermissionDenied = 4016
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		/* 没有 token */
		if len(token) < 1 {
			response.Unauthorized(CodeTokenNUll, "no token", c)
			c.Abort()
			return
		}

		log.Println(token)

		j := NewJWT()
		claims, err := j.ParseToken(token)
		/* 解析 token 出错 */
		if err != nil {
			switch err {
			case TokenExpired:
				response.Unauthorized(CodeTokenExpired, "Token Expired", c)
				c.Abort()
				return
			case TokenNotValidYet:
				response.Unauthorized(CodeTokenNotValidYet, "Token Not Valid Yet", c)
				c.Abort()
				return
			case TokenMalformed:
				response.Unauthorized(CodeTokenMalformed, "Token Malformed", c)
				c.Abort()
				return
			case TokenInvalid:
				response.Unauthorized(CodeTokenInvalid, "Token Invalid", c)
				c.Abort()
				return
			default:
				c.Abort()
				return
			}
		}

		/* token 将要过期 */
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.CONF.JWT.ExpiresTime
			newToken, _ := j.NewToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("'new-token", newToken)
			c.Header("'new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
