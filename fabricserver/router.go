package main

import (
	"github.com/cloudStore/fabricserver/handler"
	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	customizeRegister(r)
	r.GET("/ping", handler.Ping)
}

func customizeRegister(r *gin.Engine) {
	v := r.Group("/fabric")

	v.POST("/set_kv",handler.SetKV)
	v.POST("/get_value",handler.GetValue)
	v.POST("/delete_kv",handler.DeleteKV)
}