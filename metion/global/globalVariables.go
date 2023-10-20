package global

import (
	"bizd/metion/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Tasks *model.Task
var MySigningKey = []byte("Key of BIZD")
var ExpiresTime int64 = 60 * 60 * 24
var Issuer = "BIZD" // token签发人
// User 数据库信息
var User = "root"
var Pwd = "baishan123"
var Ip = "172.18.89.54"
var Port = "3306"
var DbName = "test"

var RedisPwd = "baishan123"
var RedisIp = "172.18.89.54"
var RedisPort = "6379"
var RedisDb = 0
var RedisCli = &redis.Client{}

// DataDir 数据目录
var DataDir = "./data/bizd/"

const (
	TimeFormat              = "2006-01-02 15:04:05"
	TimeMinuteFormat        = "2006-01-02 15:04"
	TimeDayFormat           = "2006-01-02"
	WxUrlDefault            = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=3d9df143-f9fa-4353-9685-002f11f82d52"
	AllocatingAssignmentTag = "分配提醒"
	AssignmentStartTag      = "开始提醒"
	AssignmentNotStartedTag = "超时未开始提醒"
	WxUrlKey                = "WxUrl"
	DefaultTime             = "1991-01-01 00:00"
)
const (
	SheetName = "实施详情"
	ExcelName = "实施详情"
)

type ClientAndPointPositionStatus int

const (
	UnResearched = iota + 10
	InResearched
	EndOfResearched
	UnImplemented = iota + 17
	InImplemented
	EndOfImplementation
	UnPoc = iota + 24
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

var PointPositionToText = map[string]string{
	"clientAbbreviation": "客户",
	"pointPositionName":  "点位名称",
	"address":            "地址",
	"ip":                 "IP网段",
	"dataLink":           "资料连接",
	"cpeName":            "设备别名",
	"statusName":         "状态",
	"scheduledTime":      "预计实施时间",
	"startTime":          "开始时间",
	"endTime":            "结束时间",
	"remark":             "备注",
	"userName":           "负责人",
	"implementerName":    "实施人",
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
