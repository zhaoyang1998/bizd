package initGlobal

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/timedTask"
	"fmt"
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
