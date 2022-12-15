package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitGormDB() (err error) {
	dsn := "root:baishan123@tcp(172.18.89.86:3306)/bizd?charset=utf8mb3&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Printf("数据库连接失败：%v\n", err)
	} else {
		fmt.Printf("数据库连接成功\n")
		DB = db
	}
	return err
}
