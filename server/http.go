package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/towelong/healthy-report-server/biz"
	"github.com/towelong/healthy-report-server/middleware"
)

func Run() {
	r := gin.Default()
	r.Use(middleware.CORS)
	r.GET("/", func(ctx *gin.Context) {
		err := biz.Register()
		if err != nil {
			return
		}
		ctx.JSON(200, gin.H{
			"ping": "pong",
		})
	})
	r.POST("/user", middleware.JWT, func(ctx *gin.Context) {
		biz.Register()
		ctx.JSON(http.StatusOK, gin.H{
			"token": "",
		})
	})
	r.POST("/infomation", middleware.JWT, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "上传信息成功",
		})
	})
	r.Run(":8016")
}
