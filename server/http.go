package server

import (
	"github.com/gin-gonic/gin"
	"github.com/towelong/healthy-report-server/biz"
)

func Run() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		err := biz.Register()
		if err != nil {
			return
		}
		ctx.JSON(200, gin.H{
			"ping": "pong",
		})
	})
	r.Run(":8016")
}
