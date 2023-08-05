package middleware

import (
	"backend/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginForm struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	CaptchaId string `json:"captcha_id" form:"captcha_id"`
	Captcha   string `json:"captcha" form:"captcha"`
}

type UserClaim struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Id       int    `json:"id"`
}

func JWTAuthMiddleWare(key string) *jwt.GinJWTMiddleware {
	var m, _ = jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "versus_realm_s3cr3t_ybpxdgh^21",
		Key:             []byte(key),
		Timeout:         time.Hour * 168,
		MaxRefresh:      time.Hour,
		Authenticator:   LoginAuthenticator,
		Authorizator:    LoginAuthorizator,
		PayloadFunc:     PayloadFunc,
		Unauthorized:    Unauthorized,
		IdentityHandler: IdentityHandler,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
	return m
}

func LoginAuthenticator(ctx *gin.Context) (interface{}, error) {
	var user LoginForm

	if err := ctx.ShouldBindJSON(&user); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	if _, err := models.CheckLogin(user.Username, user.Password); err == nil {
		u := models.GetUserInfoByUsername(user.Username)
		return UserClaim{
			Username: u.Username,
			Nickname: u.Nickname,
			Id:       int(u.ID),
		}, nil
	} else {
		return nil, err
	}
}

func LoginAuthorizator(data interface{}, ctx *gin.Context) bool {
	if v, ok := data.(*models.Users); ok {
		ctx.Set("nickname", v.Nickname)
		ctx.Set("userId", v.ID)
		ctx.Set("userName", v.Username)
		return true
	}
	return false
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(UserClaim); ok {
		return jwt.MapClaims{
			"ID":       v.Id,
			"Nickname": v.Nickname,
			"Username": v.Username,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(ctx *gin.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return &UserClaim{
		Username: claims["Username"].(string),
		Id:       int(claims["ID"].(float64)),
		Nickname: claims["Nickname"].(string),
	}
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusForbidden, gin.H{
		"code": code,
		"msg":  message,
	})
}
