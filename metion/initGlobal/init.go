package initGlobal

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/timedTask"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jakecoffman/cron"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func GormDB() (err error) {
	dsn := global.User + ":" + global.Pwd + "@tcp(" + global.Ip + ":" + global.Port + ")/" + global.DbName + "?charset=utf8mb3&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("数据库连接失败：%v\n", err)
	} else {
		fmt.Printf("数据库连接成功\n")
		global.DB = db
	}
	return err
}

func Inits() {
	_ = GormDB()
	initCron()
	initSystemParameters()
	initRedis()
	return
}
func initSystemParameters() {
	var params []model.SystemParameters
	_ = global.DB.Find(&params)
	for _, item := range params {
		global.SystemParameters[item.Key] = item.Value
	}
}
func initCron() {
	global.Tasks = &model.Task{}
	global.Tasks.CronTask = cron.New()
	timedTask.InitCron()
	go global.Tasks.CronTask.Start()
}

func initRedis() {
	// 创建Redis客户端
	global.RedisCli = redis.NewClient(&redis.Options{
		Addr:     global.RedisIp + ":" + global.RedisPort,
		Password: global.RedisPwd, // Redis无密码的话，这里为空
		DB:       global.RedisDb,
	})

	// 测试连接是否成功
	_, err := global.RedisCli.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("redis连接失败：%v\n", err)
	} else {
		fmt.Println("redis数据库连接成功")
	}

}
