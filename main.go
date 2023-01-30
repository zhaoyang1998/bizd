package main

import (
	"bizd/metion/initGlobal"
	"bizd/metion/utils"
	"bizd/router"
	"fmt"
)

func main() {
	// 开启日志
	utils.SetupLogger()
	// 初始化
	initGlobal.Inits()
	router := router.SetupRouter()
	if err := router.Run(":8888"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
