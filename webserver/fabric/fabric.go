package fabric

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
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

func HTTPPost(url string, content []byte) ([]byte, error) {
	logger.Info("http post info", "url is:"+url)
	client := &http.Client{
		Timeout: HTTPTimeout * time.Second,
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(content))
	if err != nil {
		return nil, errors.New("init http post error")
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ioutil read body error")
	}
	logger.Info("resp is ", string(body))
	return body, nil
}

func SetKV(key string, value string) error {
	url := fmt.Sprintf("http://%s:%s/fabric/set_kv",FabricServerHost,FabricSerrverPort)
	kv := KV{
		Key:   key,
		Value: value,
	}
	reqBody, err := json.Marshal(kv)
	if err != nil {
		logger.Warn("SetKv failed:", err.Error())
		return err
	}
	respBody, err := HTTPPost(url, reqBody)
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
	url := fmt.Sprintf("http://%s:%s/fabric/set_kv",FabricServerHost,FabricSerrverPort)
	req := GetValueReq{
		Key: key,
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		logger.Warn("GetValue failed:", err.Error())
		return "", err
	}
	respBody, err := HTTPPost(url, reqBody)
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
