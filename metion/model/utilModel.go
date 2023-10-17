package model

type Response struct {
	// 状态码
	Code int `json:"code"`
	// 数据
	Data string `json:"data"`
	// 提示消息
	Message string `json:"message"`
	Token   string `json:"token"`
	ResponsePagination
}

type Pagination struct {
	PageSize   int `json:"pageSize,omitempty" form:"pageSize" gorm:"-"`
	PageNumber int `json:"pageNumber,omitempty" form:"pageNumber" gorm:"-"`
}

type ResponsePagination struct {
	Total int `json:"total,omitempty" form:"total"`
	Cur   int `json:"cur,omitempty" form:"cur"`
	Next  int `json:"next,omitempty" form:"next"`
	Prev  int `json:"prev,omitempty" form:"prev"`
}

type Search struct {
	Keyword string `json:"keyword" gorm:"-"`
	ETime   string `json:"eTime" gorm:"-"`
	STime   string `json:"sTime" gorm:"-"`
	PointPosition
	Pagination
}
type DelModel struct {
	Keys []string `json:"keys"`
}

type Echarts struct {
	TotalData      EchartsPie  `json:"totalData"`
	CurData        EchartsPie  `json:"curData"`
	EfficiencyData EchartsLine `json:"efficiencyData,omitempty"`
	NullEcharts    NullEcharts `json:"nullEcharts,omitempty"`
}
type EchartsPieData struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}
type EchartsPie struct {
	Series  []EchartsPieSeries `json:"series,omitempty"`
	Legend  EchartsLegend      `json:"legend,omitempty"`
	Title   EchartsTitle       `json:"title,omitempty"`
	ToolTip EchartsToolTip     `json:"tooltip,omitempty"`
}
type EchartsPieSeries struct {
	Type  string           `json:"type,omitempty"`
	Data  []EchartsPieData `json:"data,omitempty"`
	Label EchartsLabel     `json:"label"`
}
type EchartsLabel struct {
	Show      bool   `json:"show,omitempty"`
	Formatter string `json:"formatter,omitempty"`
}
type EchartsLegend struct {
	Data   []string `json:"data,omitempty"`
	Left   string   `json:"left,omitempty"`
	Orient string   `json:"orient,omitempty"`
}
type EchartsLine struct {
	Legend  EchartsLegend       `json:"legend,omitempty"`
	XAxis   Axis                `json:"xAxis,omitempty"`
	YAxis   Axis                `json:"yAxis,omitempty"`
	Tooltip EchartsToolTip      `json:"tooltip,omitempty"`
	Series  []EchartsLineSeries `json:"series,omitempty"`
	Title   EchartsTitle        `json:"title,omitempty"`
	Flag    bool                `json:"flag"`
}
type EchartsToolTip struct {
	Trigger string `json:"trigger,omitempty"` // axis   item   none三个值
}
type NullEcharts struct {
	Title EchartsTitle `json:"title,omitempty"`
}
type EchartsTitle struct {
	Text string `json:"text,omitempty"`
	Left string `json:"left,omitempty"`
	X    string `json:"x,omitempty"`
	Y    string `json:"y,omitempty"`
}
type EchartsLineSeries struct {
	Name string `json:"name,omitempty"`
	Data []int  `json:"data,omitempty"`
	Type string `json:"type,omitempty"`
}
type Axis struct {
	Type          string        `json:"type,omitempty"`
	Data          []string      `json:"data,omitempty"`
	Name          string        `json:"name,omitempty"`
	NameTextStyle NameTextStyle `json:"nameTextStyle,omitempty"`
}
type NameTextStyle struct {
	FontWeight int `json:"fontWeight,omitempty"`
	FontSize   int `json:"fontSize,omitempty"`
}

type Result struct {
	Times float32
	Weeks int
}
type MenuJson struct {
	Permissions []string `json:"permissions,omitempty"`
	Menus       []Menu   `json:"menus,omitempty"`
}

type Menu struct {
	Id           int    `json:"permissions"`
	Path         string `json:"path"`
	Name         string `json:"name"`
	Url          string `json:"url"`
	Type         int    `json:"type"`
	Icon         string `json:"icon"`
	Show         int    `json:"show"`
	Tab          int    `json:"tab"`
	Multiple     int    `json:"multiple"`
	Keepalive    int    `json:"keepalive"`
	Sort         int    `json:"sort"`
	NameCn       string `json:"name_cn"`
	NameEn       string `json:"name_en"`
	MenuId       int    `json:"menu_id"`
	ParentId     int    `json:"parent_id"`
	EnterpriseId int    `json:"enterprise_id"`
	Children     []Menu `json:"children"`
}

type MyError struct {
	Code    int
	Message string
	Error   error
}

type StatusText struct {
	Name  string
	Value int
}
