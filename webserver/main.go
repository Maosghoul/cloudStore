package main

import (
	"github.com/cloudStore/config"
	"github.com/cloudStore/db"
	"github.com/gin-gonic/gin"
	"log"
)

func Init() {
	var err error
	defer func() {
		if e := recover(); e != nil {
			log.Printf("panic :%v\n", e)
		}
	}()
	err = config.Init()
	if err != nil {
		panic(err)
	}
	err = db.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	Init()
	r := gin.Default()
	defer func() {
		if err := r.Run(":6789"); err != nil {
			log.Printf("run gin error:%v\n", err)
		}
	}()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("./static/html/*")
	loadHTML(r)
	register(r)
}
