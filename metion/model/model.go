package model

import (
	"github.com/jakecoffman/cron"
)

type Conductor struct {
	Id          int `gorm:"primaryKey"`
	Username    string
	CurrentAble bool
	WxName      string
	Seq         int
}

type Client struct {
	// 客户唯一编码
	ClientId string `gorm:"primaryKey" json:"clientId,omitempty" form:"clientId"`
	// 客户名称
	ClientName string `json:"clientName,omitempty" form:"clientName" validate:"required"`
	// 客户简称
	ClientAbbreviation string `json:"clientAbbreviation,omitempty" form:"clientAbbreviation" validate:"required"`
	// canvas账号
	CanvasAccount string `json:"canvasAccount,omitempty" form:"canvasAccount"`
	// canvas密码
	CanvasPwd string `json:"canvasPwd,omitempty" form:"canvasPwd"`
	// 客户资料连接
	DataLink string `json:"dataLink,omitempty" form:"dataLink"`
	// 负责人ID
	PrincipalId   string `json:"principalId,omitempty" form:"principalId"`
	PrincipalName string `json:"principalName,omitempty" form:"principalName" gorm:"->"`
	// 客户状态： 实施未开始、进行中、已完成:10,11,12 POC未开始、进行中、已完成:20,21,22
	Status     int    `json:"status,omitempty" form:"status" validate:"required"`
	StatusName string `json:"statusName,omitempty" gorm:"-"`
	UpdatedAt  int    `json:"updatedAt,-"`
	// 分页参数
	Pagination
}

type PointPosition struct {
	// 单位ID
	PointPositionId string `gorm:"primaryKey" json:"pointPositionId,omitempty" form:"pointPositionId"`
	// 单位名称
	PointPositionName string `json:"pointPositionName,omitempty" form:"pointPositionName" validate:"required"`
	// 所属客户ID
	ClientId string `json:"clientId,omitempty" form:"clientId" validate:"required"`
	// 客户名称
	ClientAbbreviation string `json:"clientAbbreviation,omitempty" form:"clientAbbreviation" gorm:"->"`
	// 创建者id
	UserId string `json:"userId,omitempty" form:"userId"`
	// 创建者名称
	UserName string `json:"userName,omitempty" form:"userName" gorm:"->"`
	// 单位地址
	Address string `json:"address,omitempty" form:"address" validate:"required"`
	// ip规划
	Ip string `json:"ip,omitempty" form:"ip"`
	// 实施类型 0：调研 1：正式实施 2：POC
	Type *int `json:"type,omitempty" form:"type" validate:"required"`
	// 人数
	PeopleNumbers *int `json:"peopleNumbers,omitempty" form:"peopleNumbers"`
	// 预计实施时间
	ScheduledTime string `json:"scheduledTime,omitempty" form:"scheduledTime"`
	// 人员
	ImplementerId string `json:"implementerId,omitempty" form:"implementerId"`
	// 人员名称
	ImplementerName string `json:"implementerName,omitempty" form:"implementerName" gorm:"->"`
	// CpeName
	CpeName string `json:"cpeName,omitempty" form:"cpeName"`
	// 状态，调研未开始、进行中、已完成:0,1,2 实施未开始、进行中、已完成:10,11,12 POC未开始、进行中、已完成:20,21,22
	Status     *int   `gorm:"FORCE" json:"status,omitempty" form:"status"`
	StatusName string `gorm:"-" json:"statusName,omitempty" form:"statusName"`
	// 实施资链接
	DataLink string `json:"dataLink,omitempty" form:"dataLink"`
	// 备注
	Remark string `json:"remark,omitempty" form:"remark"`
	// 开始时间
	StartTime string `json:"startTime,omitempty" form:"startTime"`
	// 结束时间
	EndTime   string `json:"endTime,omitempty" form:"endTime"`
	TotalTime int    `json:"totalTime,omitempty" form:"totalTime"`
	UpdatedAt int    `json:"updatedAt,-"`
	// 分页参数
	Pagination
}

type User struct {
	// 用户id
	UserId string `gorm:"primaryKey" json:"userId,omitempty" form:"userId"`
	// 账号
	UserAccount string `json:"userAccount,omitempty" form:"userAccount" validate:"required"`
	// 用户名
	UserName string `json:"userName,omitempty" form:"userName" validate:"required"`
	// 密码
	UserPwd string `json:"userPwd,omitempty" form:"userPwd" validate:"required"`
	// 微信名称
	WxName string `json:"wxName,omitempty" form:"wxName" validate:"required"`
	// 人员类型 1：交付 2：项目管理
	Type     int    `json:"type,omitempty" form:"type" validate:"required"`
	TypeName string `json:"typeName,omitempty" gorm:"-"`
	// 优先级
	Priority int `json:"priority,omitempty" form:"priority"`
	// 当前工作量
	CurrentWorkload int `json:"currentWorkload,omitempty" form:"currentWorkload"`
	UpdatedAt       int `json:"updatedAt,-"`
	// 分页参数
	Pagination
}

type ImplementDetails struct {
	ImplementDetailId string `json:"implementDetailId" gorm:"primaryKey"`
	PointPositionId   string `json:"pointPositionId,omitempty"`
	StepName          string `json:"stepName,omitempty"`
	Desc              string `json:"desc,omitempty"`
	Seq               int    `json:"seq,omitempty"`
	UpdatedAt         int    `json:"updatedAt,-"`
	Total             int64  `json:"total,omitempty"  gorm:"->"`
	TotalTime         int64  `json:"totalTime,omitempty"  `
	StartTime         string `json:"startTime,omitempty"`
	EndTime           string `json:"endTime,omitempty"`
}

// MsgFromCron 来获取Cron库的数据
type MsgFromCron struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CronTime    string `json:"cronTime"`
	Receive     string `json:"receive"`
	ReceiveType string `json:"receiveType"`
	// 0：任务分配提醒  1：任务开始提醒   2：任务超时未开始提醒
	Type            int    `json:"type"`
	IsSend          int    `json:"isSend"`
	WxName          string `json:"wxName"`
	PointPositionId string `json:"pointPositionId"`
	ScheduledTime   string `json:"scheduledTime" gorm:"-"`
}

type Task struct {
	CronTask  *cron.Cron
	CronCount int
}

type SystemParameters struct {
	Key   string
	Value string
}

// TableName 对应数据库中的表名

func (Conductor) TableName() string {
	return "t_conductor"
}

func (Client) TableName() string {
	return "t_client"
}

func (PointPosition) TableName() string {
	return "t_point_position"
}

func (User) TableName() string {
	return "t_user"
}

func (ImplementDetails) TableName() string {
	return "t_implement_details"
}

func (MsgFromCron) TableName() string {
	return "t_cron"
}

func (SystemParameters) TableName() string {
	return "t_system_parameters"
}
