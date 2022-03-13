package server

import "github.com/gin-gonic/gin"

func Run() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"ping": "pong",
		})
	})
}
