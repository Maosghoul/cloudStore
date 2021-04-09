package db

import (
	"fmt"
	"github.com/cloudStore/fabricserver/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Db struct {
	conn *gorm.DB
}

var Dao Db

type KV struct {
	Key   string
	Value string
}

func Init() error {
	var err error
	username := config.Conf.MysqlConfig.UserName
	password := config.Conf.MysqlConfig.PassWord
	host := config.Conf.MysqlConfig.Host
	port := config.Conf.MysqlConfig.Port
	dbName := config.Conf.MysqlConfig.DBName
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbName, timeout)
	log.Printf("dsn:%s", dsn)

	Dao.conn, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	log.Printf("db init success")
	Dao.conn.LogMode(true)
	return nil
}
