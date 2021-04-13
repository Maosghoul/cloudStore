package handler

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/cloudStore/webserver/db"
	"github.com/cloudStore/webserver/fabric"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	DefaultFilePath = "/home/ubuntu/filecloud"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "tong",
	})
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"Email"`
}

type ModifyRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"Email"`
}

type DeleteFileRequest struct {
	Username string   `json:"username"`
	FileName []string `json:"filename"`
}

type AdultFileRequest struct{
	Username string `json:"username"`
	FileName []string `json:"filename"`
}

type DownloadFileRequest struct{
	Username string `json:"username"`
	FileName string `json:"filename"`
}


type ListFileRequest struct {
	Username string `json:"username"`
}

type ListFileResponse struct {
	FileInfo []FileInfo `json:"fileinfo"`
}

type FileInfo struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

func Login(ctx *gin.Context) {

	req := LoginRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		msg := fmt.Sprintf("Login param error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	user, err := db.Dao.FindUserByUsername(req.Username)
	if err != nil || user == nil {
		msg := fmt.Sprintf("Login  error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": msg,
		})
		return
	}
	pass := fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password)))
	if user.Password != pass {
		msg := fmt.Sprintf("Login  error:%v", errors.New("password is not true"))
		logger.Warn(msg)
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": msg,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Login success",
	})

}

func Register(ctx *gin.Context) {
	req := RegisterRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		msg := fmt.Sprintf("Register param error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"Register error": msg,
		})
		return
	}
	user := db.User{
		Username: req.Username,
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password))),
		Email:    req.Email,
	}
	if err := db.Dao.AddUser(user); err != nil {
		msg := fmt.Sprintf("Register  error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": msg,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Register success",
	})

}

func Modify(ctx *gin.Context) {
	req := ModifyRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		msg := fmt.Sprintf("Modify param error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"Modify error": msg,
		})
		return
	}
	user := db.User{
		Username: req.Username,
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password))),
		Email:    req.Email,
	}
	if err := db.Dao.ModifyUser(user); err != nil {
		msg := fmt.Sprintf("Modify  error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": msg,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Modify success",
	})

}

func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		msg := fmt.Sprintf("Upload param error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	username := ctx.PostForm("username")
	if username == "" {
		msg := fmt.Sprint("username or id is nil")
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	uniqueKey := fabric.GenerateUniqueKey(username, file.Filename)
	uniqueValue, err := fabric.GetFileValue(file)
	logger.Info("file key and value:",uniqueKey,uniqueValue)
	err = fabric.SetKV(uniqueKey, uniqueValue)
	if err!=nil{
		msg := fmt.Sprint("upload fabric network failed")
		logger.Warn(msg)
		ctx.JSON(http.StatusInternalServerError,map[string]string{
			"message":msg,
		})
		return
	}

	fileDir := fmt.Sprintf("%s/%s", DefaultFilePath, username)
	fmt.Println(fileDir)
	os.MkdirAll(fileDir, os.ModePerm)

	dst := fmt.Sprintf("%s/%s", fileDir, file.Filename)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		msg := fmt.Sprintf("UploadFile  error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": msg,
		})
		return
	}
	fileInfo := db.FileInfo{
		Username:   username,
		Filename:   file.Filename,
		UpdateTime: time.Now().Format("2006/01/02 15:04:05"),
	}
	err = db.Dao.AddFile(fileInfo)
	if err != nil {
		msg := fmt.Sprintf("UploadFile  error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": msg,
		})
		return
	}
	msg := fmt.Sprintf("upload file success dst:%v", dst)
	ctx.JSON(http.StatusOK, map[string]string{
		"message": msg,
	})
}

func DeleteFile(ctx *gin.Context) {
	req := DeleteFileRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		msg := fmt.Sprintf("DeleteFile param error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	fileInfos := make([]db.FileInfo, 0)
	for _, v := range req.FileName {
		fileInfos = append(fileInfos, db.FileInfo{
			Username: req.Username,
			Filename: v,
		})
	}
	fileDir := fmt.Sprintf("%s/%s", DefaultFilePath, req.Username)
	fmt.Println(fileDir)
	for _, info := range fileInfos {
		dst := fmt.Sprintf("%s/%s", fileDir, info.Filename)
		if err := os.Remove(dst); err != nil {
			msg := fmt.Sprintf("DeleteFile  error:%v", err)
			logger.Warn(msg)
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": msg,
			})
			return
		}
	}

	err := db.Dao.DeleteFile(fileInfos)
	if err != nil {
		msg := fmt.Sprintf("DeleteFile  error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": msg,
		})
		return
	}
	msg := fmt.Sprintf("delete file success dst:%v", fileDir)
	ctx.JSON(http.StatusOK, map[string]string{
		"message": msg,
	})

}

func ListFile(ctx *gin.Context) {
	req := ListFileRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		msg := fmt.Sprintf("List file  error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	info := db.FileInfo{}
	info.Username = req.Username
	output, err := db.Dao.ListFile(info)
	if err != nil {
		msg := fmt.Sprintf("List file  error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": msg,
		})
		return
	}
	resp := ListFileResponse{}
	resp.FileInfo = make([]FileInfo, 0)
	for _, v := range output {
		resp.FileInfo = append(resp.FileInfo, FileInfo{
			Name: v.Filename,
			Time: v.UpdateTime,
		})
	}
	ctx.JSON(http.StatusOK, resp)

}

func AdultFile(ctx *gin.Context){
	req := AdultFileRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		msg := fmt.Sprintf("AdultFile param error:%v", err)
		logger.Warn(msg)
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"code":"1",
			"message": msg,
		})
		return
	}
	fileInfos := make([]db.FileInfo, 0)
	for _, v := range req.FileName {
		fileInfos = append(fileInfos, db.FileInfo{
			Username: req.Username,
			Filename: v,
		})
	}
	fileKV :=make([]fabric.KV,0)
	fileDir := fmt.Sprintf("%s/%s", DefaultFilePath, req.Username)
	for _, info := range fileInfos {
		uniqueKey := fabric.GenerateUniqueKey(info.Username,info.Filename)
		dst := fmt.Sprintf("%s/%s", fileDir, info.Filename)
		uniqueValue ,err := fabric.GetPathFileValue(dst)
		if err!=nil{
			msg := fmt.Sprintf("AdultFile  error:%v", err)
			logger.Warn(msg)
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"code":"1",
				"message": msg,
			})
			return
		}
		fileKV = append(fileKV ,fabric.KV{
			Key: uniqueKey,
			Value: uniqueValue,
		})
	}
	notPass :=make([]string,0)
	for _,kv:=range fileKV{
		value ,_:= fabric.GetValue(kv.Key)
		if value != kv.Value{
			idx := strings.LastIndex(kv.Key,"$")
			tmpName := kv.Key[idx+1:]
			notPass = append(notPass,tmpName)
		}
	}
	if len(notPass)==0{
		msg :=fmt.Sprintf("All file pass the adult")
		ctx.JSON(http.StatusOK,map[string]string{
			"code":"0",
			"message":msg,
		})
	}else{
		msg := fmt.Sprintf("%v",notPass)
		ctx.JSON(http.StatusOK,map[string]string{
			"code":"-1",
			"message":msg,
		})
	}
}

func DownloadFile(ctx *gin.Context){
	req := DownloadFileRequest{}
	req.Username,_ = ctx.GetQuery("username")
	req.FileName ,_= ctx.GetQuery("filename")
	fileInfos := make([]db.FileInfo, 0)
		fileInfos = append(fileInfos, db.FileInfo{
			Username: req.Username,
			Filename: req.FileName,
		})

	fileDir := fmt.Sprintf("%s/%s", DefaultFilePath, req.Username)
	fmt.Println(fileDir)
	for _, info := range fileInfos {
		dst := fmt.Sprintf("%s/%s", fileDir, info.Filename)
		logger.Info("dst is :",dst)
		f,err := os.Open(dst)
		if err!=nil{
			msg := fmt.Sprintf("DownloadFile  error:%v", err)
			logger.Warn(msg)
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"code":"1",
				"message": msg,
			})
			return
		}
		fileContent ,err :=ioutil.ReadAll(f)
		if err!=nil || fileContent==nil{
			msg := fmt.Sprintf("DownloadFile  error:%v", err)
			logger.Warn(msg)
			ctx.JSON(http.StatusInternalServerError, map[string]string{
				"code":"1",
				"message": msg,
			})
			return
		}
		logger.Info("fileContent",string(fileContent))
		ctx.Writer.Header().Add("content-type","application/octet-stream;charset=utf-8")
		ctx.Writer.Header().Add("content-disposition",fmt.Sprintf("attachment; filename=\"%s\"", info.Filename))
		ctx.Writer.Write(fileContent)
		f.Close()
	}
	logger.Info("download file success")
}