package main

import (
	"bizd/metion/db"
	"bizd/metion/timedTask"
	"bizd/router"
	"fmt"
)

func main() {
	// 初始化数据库连接
	db.InitGormDB()
	// 开启定时任务
	timedTask.StartTimedTask()
	router := router.SetupRouter()
	if err := router.Run(":8888"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
