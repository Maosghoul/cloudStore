package db

import (
	"errors"
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

func (d *Db) AddKV(kv KV) error {
	err := d.conn.Exec("INSERT INTO `kv` (`id`,`key`,`value`) VALUES(0,?,?)", kv.Key, kv.Value).Error
	return err
}

func (d *Db) ModifyKV(kv KV) error {
	err := d.conn.Exec("UPDATE `kv` SET `value` = ? WHERE `key` = ?",
		kv.Value, kv.Key).Error
	return err
}

func (d *Db) DeleteKV(kv KV) error {
	err := d.conn.Exec("DELETE FROM `kv` WHERE `key` = ? AND `value` = ?", kv.Key, kv.Value).Error
	return err
}

func (d *Db) GetAllKV() ([]KV, error) {
	infos := make([]KV, 0)
	err := d.conn.Raw("SELECT * FROM `kv`").Scan(&infos).Error
	if err != nil {
		return nil, err
	}
	return infos, nil
}

func (d *Db) GetValueByKey(key string) (string, error) {
	infos := make([]KV, 0)
	err := d.conn.Raw("SELECT * FROM `kv` WHERE `key` = ?", key).Scan(&infos).Error
	if err != nil {
		return "", err
	}
	if len(infos) == 0 {
		return "", errors.New("not have this key")
	}
	if len(infos) != 1 {
		return "", errors.New("to many values")
	}
	return infos[0].Value, nil
}
