package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/towelong/healthy-report-server/biz"
	"github.com/towelong/healthy-report-server/module"
)

func JWT(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("authorization")
	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "未携带令牌访问")
	}
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)
	claims, e := module.VerifyToken(tokenStr)
	if e != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "令牌不合法")
	}
	if claims == nil || claims.UserID == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "令牌验证失败")
	}
	user, err := biz.FindUserById(claims.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "非法访问")
	}
	ctx.Set("uid", user.ID)
	ctx.Next()
}
