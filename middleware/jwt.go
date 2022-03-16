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
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "未携带令牌访问",
		})
		return
	}
	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)
	claims, e := module.VerifyToken(tokenStr)
	if e != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "令牌不合法",
		})
		return
	}
	if claims == nil || claims.UserID == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "令牌验证失败",
		})
		return
	}
	user, err := biz.FindUserById(claims.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "非法访问",
		})
		return
	}
	ctx.Set("uid", user.ID)
	ctx.Next()
}
