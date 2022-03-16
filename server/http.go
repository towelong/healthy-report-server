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
		ctx.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var u biz.User
		if err := ctx.ShouldBindJSON(&u); err != nil {
			if u.Username == "" || u.Password == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": 400,
					"msg":  "用户名或密码未填写",
				})
			}
			return
		}
		token, err := biz.Login(u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})

	r.POST("/register", func(ctx *gin.Context) {
		var u biz.User
		if err := ctx.ShouldBindJSON(&u); err != nil {
			if u.Username == "" || u.Password == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": 400,
					"msg":  "用户名或密码未填写",
				})
			}
			return
		}
		err := biz.Register(u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "注册成功",
		})
	})

	r.POST("/information", middleware.JWT, func(ctx *gin.Context) {
		var t = &biz.Task{}
		if err := ctx.ShouldBindJSON(t); err != nil {
			if t.UserID == 0 || t.SchoolID == "" || t.StudentID == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusBadRequest,
					"msg":  "参数非法",
				})
			}
			return
		}
		userId, ok := ctx.Get("uid")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "用户未获取到",
			})
			return
		}
		t.UserID = userId.(int32)
		err := biz.UploadInformation(t)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "信息上传成功",
		})
	})

	r.GET("/information", middleware.JWT, func(ctx *gin.Context) {
		userId, ok := ctx.Get("uid")
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "用户未获取到",
			})
			return
		}
		uid := userId.(int32)
		info, err := biz.FindUserInformation(uid)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, info)
	})
	r.Run(":8016")
}
