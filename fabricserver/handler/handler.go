package handler

import (
	"fmt"
	"github.com/cloudStore/fabricserver/db"
	"github.com/gin-gonic/gin"
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
		log.Printf(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	go func(req SetKVRequest){
		err := db.Dao.ModifyKV(db.KV{
			Key:   req.Key,
			Value: req.Value,
		})
		if err != nil {
			msg := fmt.Sprintf("SetKV  error:%v", err)
			log.Printf(msg)
			return
		}
	}(req)


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
		log.Printf(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	value, err := db.Dao.GetValueByKey(req.Key)
	if err != nil {
		msg := fmt.Sprintf("GetValue  error:%v", err)
		log.Printf(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	ctx.JSON(http.StatusOK, value)
}

func DeleteKV(ctx *gin.Context) {
	req := DELETEKVRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		msg := fmt.Sprintf("DeleteKV param error:%v", err)
		log.Printf(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	go func(req DELETEKVRequest){
		err := db.Dao.ModifyKV(db.KV{
			Key:   req.Key,
			Value: "deleted",
		})
		if err != nil {
			msg := fmt.Sprintf("DeleteKV  error:%v", err)
			log.Printf(msg)
			return
		}
	}(req)

	msg := fmt.Sprintf("delete kv param success")
	log.Printf(msg)
	ctx.JSON(http.StatusOK, map[string]string{
		"message": msg,
	})

}
