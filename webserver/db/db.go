package db

import (
	"fmt"
	"github.com/cloudStore/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Db struct {
	conn *gorm.DB
}

var Dao Db

type User struct {
	Username string
	Password string
	Email    string
}

type FileInfo struct {
	Username   string
	Filename   string
	UpdateTime string
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

func (d *Db) AddUser(user User) error {
	err := d.conn.Exec("INSERT INTO `user` (`id`,`username`,`password`,`email`) VALUES(0,?,?,?)",
		user.Username, user.Password, user.Email).Error
	return err
}

func (d *Db) FindUserByUsername(username string) (*User, error) {
	user := User{}
	err := d.conn.Raw("SELECT * FROM `user` WHERE `username` = ?", username).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *Db) ModifyUser(user User) error {
	err := d.conn.Exec("UPDATE `user` SET `password` = ? , `email` = ? WHERE `username` = ?",
		user.Password, user.Email, user.Username).Error
	return err
}

func (d *Db) AddFile(fileInfo FileInfo) error {
	err := d.conn.Exec("INSERT INTO `file` (`id`,`username`,`filename`,`update_time`) VALUES(0,?,?,?)",
		fileInfo.Username, fileInfo.Filename, fileInfo.UpdateTime).Error
	return err
}

func (d *Db) ModifyFile(info FileInfo) error {
	err := d.conn.Exec("UPDATE `file` SET `filename` = ?,`update_time` = ? WHERE `username` = ?", info.Filename, info.UpdateTime, info.Username).Error
	return err
}

func (d *Db) DeleteFile(infos []FileInfo) error {
	var err error
	for _, v := range infos {
		err = d.conn.Exec("DELETE FROM `file` WHERE `username` = ? AND `filename` = ?", v.Username, v.Filename).Error
		if err != nil {
			return err
		}
	}
	return err
}

func (d *Db) ListFile(info FileInfo) ([]FileInfo, error) {
	infos := make([]FileInfo, 0)
	err := d.conn.Raw("SELECT * FROM `file` WHERE `username` = ?", info.Username).Scan(&infos).Error
	if err != nil {
		return nil, err
	}
	return infos, nil
}
