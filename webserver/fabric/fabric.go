package fabric

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	InnerKey          = "$#$"
	FabricServerHost  = "8.140.121.31"
	FabricSerrverPort = "6789"
	HTTPTimeout       = 5
)

type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Response struct {
	Code    string `json:"code"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

func GenerateUniqueKey(username string, filename string) string {
	return username + InnerKey + filename
}

func GetFileValue(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	defer f.Close()
	if err != nil {
		logger.Warn("GetFileValue error:",err)
		return "", err
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		logger.Warn("GetFileValue error:",err)
		return "", err
	}
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum, nil
}

func GetPathFileValue(path string)(string,error){
	f ,err := os.Open(path)
	defer f.Close()
	if err != nil {
		logger.Warn("GetFileValue error:",err)
		return "", err
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		logger.Warn("GetFileValue error:",err)
		return "", err
	}
	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum, nil
}

func Post(url string, data interface{}) ([]byte,error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	logger.Info("http req is :",req)
	req.Header.Add("content-type", "application/json")
	if err != nil {
		logger.Warn("http post error",err)
		return nil,err
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: HTTPTimeout * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Warn("http post error",err)
		return nil,err
	}
	defer resp.Body.Close()

	result, err:= ioutil.ReadAll(resp.Body)
	if err!=nil{
		logger.Warn("http post error",err)
		return nil,err
	}
	return result,nil
}


func SetKV(key string, value string) error {
	url := fmt.Sprintf("http://%s:%s/fabric/set_kv",FabricServerHost,FabricSerrverPort)
	kv := KV{
		Key:   key,
		Value: value,
	}
	respBody, err := Post(url, kv)
	if err != nil {
		logger.Warn("SetKv failed:", err.Error())
		return err
	}
	resp := Response{}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		logger.Warn("SetKv failed:", err.Error())
		return err
	}
	logger.Info("SetKv success,resp is:", string(respBody))
	return nil
}

type GetValueReq struct {
	Key string `json:"key"`
}

func GetValue(key string) (string, error) {
	url := fmt.Sprintf("http://%s:%s/fabric/get_value",FabricServerHost,FabricSerrverPort)
	req := GetValueReq{
		Key: key,
	}
	respBody, err := Post(url, req)
	if err != nil {
		logger.Warn("GetValue failed:", err.Error())
		return "", err
	}
	resp := Response{}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		logger.Warn("GetValue failed:", err.Error())
		return "", err
	}
	logger.Info("GetValue success,resp is:", string(respBody))
	return resp.Value, nil
}
