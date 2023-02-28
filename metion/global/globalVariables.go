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
	TimeDayFormat           = "2006-01-02"
	WxUrlDefault            = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=3d9df143-f9fa-4353-9685-002f11f82d52"
	AllocatingAssignmentTag = "分配提醒"
	AssignmentStartTag      = "开始提醒"
	AssignmentNotStartedTag = "超时未开始提醒"
	WxUrlKey                = "WxUrl"
)

type ClientAndPointPositionStatus int

const (
	UnResearched = iota
	InResearched
	EndOfResearched
	UnImplemented = iota + 7
	InImplemented
	EndOfImplementation
	UnPoc = iota + 17
	InPoc
	EndPoc
)

var ClientAndPointPositionStatusText = map[ClientAndPointPositionStatus]string{
	UnResearched:        "未调研",
	InResearched:        "调研中",
	EndOfResearched:     "调研结束",
	UnImplemented:       "实施未开始",
	InImplemented:       "实施中",
	EndOfImplementation: "实施结束",
	UnPoc:               "POC未开始",
	InPoc:               "POC进行中",
	EndPoc:              "POC结束",
}

type UserType int

const (
	Delivery = iota + 1
	PM
	ADMIN
)

var UserTypeText = map[UserType]string{
	Delivery: "交付",
	PM:       "项目经理",
	ADMIN:    "系统管理员",
}

var SystemParameters = map[string]string{}
