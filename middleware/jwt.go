package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/towelong/healthy-report-server/module"
)

func JWT(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("authorization")
	if bearerToken == "" {
		ctx.JSON(http.StatusUnauthorized, "未携带令牌访问")
		ctx.Abort()
	}
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)
	claims, e := module.VerifyToken(tokenStr)
	if e != nil {
		ctx.JSON(http.StatusUnauthorized, "令牌不合法")
		ctx.Abort()
	}
	if claims == nil {
		ctx.JSON(http.StatusUnauthorized, "令牌验证失败")
		ctx.Abort()
	}
	// TODO: 身份验证
	ctx.Next()
}
