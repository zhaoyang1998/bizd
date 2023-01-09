package main

import (
	"bizd/metion/db"
	"bizd/metion/timedTask"
	"bizd/metion/utils"
	"bizd/router"
	"fmt"
	"github.com/jakecoffman/cron"
)

func main() {
	// 开启日志
	utils.SetupLogger()
	// 初始化数据库连接
	_ = db.InitGormDB()
	// 开启定时任务
	task := &timedTask.Task{}
	task.CronTask = cron.New()
	task.InitCron()
	go task.CronTask.Start()

	router := router.SetupRouter()
	if err := router.Run(":8888"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
