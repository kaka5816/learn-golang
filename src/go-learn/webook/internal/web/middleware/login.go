package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddlewareBuilder struct {
}

func (m *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//注册和登录不要校验
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			return //无需登录校验
		}
		sess := sessions.Default(ctx)
		//sees的userid为nil时
		if sess.Get("UserId") == nil {
			//中断，不要继续往下执行，也就是不要执行后的业务逻辑
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
