package handler

import (
	"fmt"
	"github.com/cloudStore/fabricserver/peer"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"log"
	"net/http"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "tong",
	})
}

type SetKVRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetValuesRequest struct {
	Key string `json:"key"`
}

type DELETEKVRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func SetKV(ctx *gin.Context) {
	req := SetKVRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		msg := fmt.Sprintf("SetKV param error:%v", err)
		logger.Info(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	err := peer.SetKV(req.Key, req.Value)
	if err != nil {
		msg := fmt.Sprintf("setkv  error:%v", err)
		logger.Warn("peer setkv error:", msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	msg := fmt.Sprintf("SetKV param success")
	log.Printf(msg)
	ctx.JSON(http.StatusOK, map[string]string{
		"message": msg,
	})

}

func GetValue(ctx *gin.Context) {
	req := GetValuesRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		msg := fmt.Sprintf("GetValue param error:%v", err)
		logger.Info(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	value, err := peer.GetValue(req.Key)
	if err != nil {
		msg := fmt.Sprintf("get value  error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
	}
	ctx.JSON(http.StatusOK, value)
}

func DeleteKV(ctx *gin.Context) {
	req := DELETEKVRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		msg := fmt.Sprintf("DeleteKV param error:%v", err)
		logger.Info(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	err := peer.SetKV(req.Key, "delete")
	if err != nil {
		msg := fmt.Sprintf("delete value  error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	msg := fmt.Sprintf("delete kv param success")
	logger.Info(msg)
	ctx.JSON(http.StatusOK, map[string]string{
		"message": msg,
	})

}
