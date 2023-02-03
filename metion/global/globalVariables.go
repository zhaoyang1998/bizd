package global

import (
	"bizd/metion/model"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Tasks *model.Task
var MySigningKey = []byte("Key of BIZD")
var ExpiresTime int64 = 60 * 60 * 24
var Issuer = "BIZD" // token签发人
// 数据库信息

var User = "root"
var Pwd = "baishan123"
var Ip = "172.18.89.86"
var Port = "3306"
var DbName = "bizd"

const (
	TimeFormat              = "2006-01-02 15:04:05"
	WxUrl                   = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=3d9df143-f9fa-4353-9685-002f11f82d52"
	AllocatingAssignmentTag = "分配提醒"
	AssignmentStartTag      = "开始提醒"
	AssignmentNotStartedTag = "超时未开始提醒"
)
