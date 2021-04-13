package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Conf Config

type Config struct {
	MysqlConfig MysqlConfig `yaml:"MysqlConfig"`
}

type MysqlConfig struct {
	DBName   string `yaml:"DBName"`
	UserName string `yaml:"UserName"`
	PassWord string `yaml:"PassWord"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
}

func Init() error {
	yamlFile, err := ioutil.ReadFile("conf/fabricServer.yaml")
	log.Printf("config init success:%s\n", string(yamlFile))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		return err
	}
	return nil
}
