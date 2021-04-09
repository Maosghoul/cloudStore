package handler

import "github.com/gin-gonic/gin"

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "tong",
	})
}

type SetKVRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func SetKV(ctx *gin.Context) {

}

func GetValue(ctx *gin.Context) {

}
