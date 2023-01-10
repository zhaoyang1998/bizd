package model

type Conductor struct {
	Id          int `gorm:"primaryKey"`
	Username    string
	CurrentAble bool
	WxName      string
	Seq         int
}

type Client struct {
	// 客户唯一编码
	ClientId string `gorm:"primaryKey" json:"clientId" form:"clientId"`
	// 客户名称
	ClientName string `json:"clientName" form:"clientName" validate:"required"`
	// 客户简称
	ClientAbbreviation string `json:"clientAbbreviation" form:"clientAbbreviation" validate:"required"`
	// canvas账号
	CanvasAccount string `json:"canvasAccount" form:"canvasAccount"`
	// canvas密码
	CanvasPwd string `json:"canvasPwd" form:"canvasPwd"`
	// 客户资料连接
	DataLink string `json:"dataLink" form:"dataLink"`
	// 负责人ID
	PrincipalId string `json:"principalId" form:"principalId" validate:"required"`
	// 客户状态： 实施未开始、进行中、已完成:10,11,12 POC未开始、进行中、已完成:20,21,22
	Status int `json:"status" form:"status" validate:"required"`
	// 分页参数
	Pagination
}

type PointPosition struct {
	// 单位ID
	PointPositionId string `gorm:"primaryKey" json:"pointPositionId" form:"pointPositionId"`
	// 单位名称
	PointPositionName string `json:"pointPositionName" form:"pointPositionName" validate:"required"`
	// 所属客户ID
	ClientId string `json:"clientId" form:"clientId" validate:"required"`
	// 客户名称
	ClientAbbreviation string `json:"clientAbbreviation" form:"clientAbbreviation" gorm:"-"`
	// 创建者id
	UserId string `json:"userId" form:"userId"`
	// 创建者名称
	UserName string `json:"userName" form:"userName" gorm:"-"`
	// 单位地址
	Address string `json:"address" form:"address" validate:"required"`
	// ip规划
	Ip string `json:"ip" form:"ip"`
	// 实施类型 0：调研 1：正式实施 2：POC
	Type int `json:"type" form:"type" validate:"required"`
	// 人数
	PeopleNumbers int `json:"peopleNumbers" form:"peopleNumbers" validate:"required"`
	// 预计实施时间
	ScheduledTime string `json:"scheduledTime" form:"scheduledTime"`
	// 人员
	ImplementerId string `json:"implementerId" form:"implementerId"`
	// 人员名称
	ImplementerName string `json:"implementerName" form:"implementerName" gorm:"-"`
	// CpeName
	CpeName string `json:"cpeName" form:"cpeName"`
	// 状态，调研未开始、进行中、已完成:0,1,2 实施未开始、进行中、已完成:10,11,12 POC未开始、进行中、已完成:20,21,22
	Status int `json:"status" form:"status"`
	// 实施资链接
	DataLink string `json:"dataLink" form:"dataLink"`
	// 备注
	Remark string `json:"remark" form:"remark"`
	// 开始时间
	StartTime string `json:"startTime" form:"startTime"`
	// 结束时间
	EndTime string `json:"endTime" form:"endTime"`
	// 分页参数
	Pagination
}

type User struct {
	// 用户id
	UserId string `gorm:"primaryKey" json:"userId" form:"userId"`
	// 用户名
	UserName string `json:"userName" form:"userName" validate:"required"`
	// 密码
	UserPwd string `json:"userPwd" form:"userPwd" validate:"required"`
	// 微信名称
	WxName string `json:"wxName" form:"wxName" validate:"required"`
	// 人员类型 0：交付 1：项目管理
	Type int `json:"type" form:"type" validate:"required"`
	// 优先级
	Priority int `json:"priority" form:"priority"`
	// 当前工作量
	CurrentWorkload int `json:"currentWorkload" form:"currentWorkload"`
	// 分页参数
	Pagination
}

type Response struct {
	// 状态码
	Code int `json:"code"`
	// 数据
	Data string `json:"data"`
	// 提示消息
	Message string `json:"message"`
}
type Pagination struct {
	PageSize   int `json:"pageSize" form:"pageSize,default=10" gorm:"-"`
	PageNumber int `json:"pageNumber" form:"pageNumber,default=1" gorm:"-"`
}

// MsgFromCron 来获取Cron库的数据
type MsgFromCron struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CronTime    string `json:"cronTime"`
	Receive     string `json:"receive"`
	ReceiveType string `json:"receiveType"`
	Tags        string `json:"tags"`
	IsSend      int    `json:"isSend"`
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

func (MsgFromCron) TableName() string {
	return "t_cron"
}
