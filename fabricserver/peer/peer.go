package peer

import (
	"encoding/json"
	"github.com/wonderivan/logger"
	"os/exec"
	"strings"
)

const DefaultPeerPath = "/root/fabric-1.1.0-preview-demo/bin/peer"

type LastBin struct {
	JsonArgs []string `json:"Args"`
}

func GetValue(key string) (string, error) {
	Args := []string{"get", key}
	lastBin := LastBin{JsonArgs: Args}
	xx, _ := json.Marshal(lastBin)
	cmdArgs := []string{"chaincode", "invoke", "-o", "orderer.example.com:7050", "-C", "mychannel", "-n", "setkv", "-c", string(xx)}
	cmd := exec.Command(DefaultPeerPath, cmdArgs...)
	out, err := cmd.CombinedOutput()
	logger.Info("peer get value output:", string(out))
	if err != nil {
		logger.Warn("peer get value error:" + err.Error())
		return "", err
	}
	firstIndex := strings.Index(string(out), "\"")
	lastIndex := strings.LastIndex(string(out), "\"")
	logger.Info("value is :", string(out)[firstIndex+1:lastIndex])
	return string(out)[firstIndex+1 : lastIndex], nil
}

func SetKV(key string, value string) error {
	Args := []string{"set", key, value}
	lastBin := LastBin{JsonArgs: Args}
	xx, _ := json.Marshal(lastBin)
	cmdArgs := []string{"chaincode", "invoke", "-o", "orderer.example.com:7050", "-C", "mychannel", "-n", "setkv", "-c", string(xx)}
	cmd := exec.Command(DefaultPeerPath, cmdArgs...)
	out, err := cmd.CombinedOutput()
	logger.Info("peer set kv output:", string(out))
	if err != nil {
		logger.Warn("peer set kv error:" + err.Error())
		return err
	}
	logger.Info("peer output:", string(out))
	return nil
}
