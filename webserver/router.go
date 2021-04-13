package main

import (
	"github.com/cloudStore/webserver/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(r *gin.Engine) {
	customizeRegister(r)
	r.GET("/ping", handler.Ping)
}

func customizeRegister(r *gin.Engine) {
	v := r.Group("/cs")
	// people manage
	v.POST("/login", handler.Login)
	v.POST("/register", handler.Register)
	v.POST("/modify", handler.Modify)

	// file manage
	v.POST("/upload_file", handler.UploadFile)
	v.POST("/delete_file", handler.DeleteFile)
	v.POST("/list_file", handler.ListFile)
	v.POST("/adult_file",handler.AdultFile)
	v.GET("/download_file",handler.DownloadFile)
}

func loadHTML(r *gin.Engine) {
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "云盘",
		})
	})
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{
			"title": "登录",
		})
	})
	r.GET("/reg", func(context *gin.Context) {
		context.HTML(http.StatusOK, "reg.html", gin.H{
			"title": "注册",
		})
	})
}
