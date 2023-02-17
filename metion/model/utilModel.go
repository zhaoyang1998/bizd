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
	PageSize   int `json:"pageSize" form:"pageSize" gorm:"-"`
	PageNumber int `json:"pageNumber" form:"pageNumber" gorm:"-"`
}

type ResponsePagination struct {
	Total int `json:"total" form:"total"`
	Cur   int `json:"cur" form:"cur"`
	Next  int `json:"next" form:"next"`
	Prev  int `json:"prev" form:"prev"`
}

type Search struct {
	Keyword string `json:"keyword"`
	Pagination
}
type DelModel struct {
	Keys []string `json:"keys"`
}

type Echarts struct {
	TotalData []EchartsPie `json:"totalData"`
	CurData   []EchartsPie `json:"curData"`
}
type EchartsPie struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

type MenuJson struct {
	Permissions []string `json:"permissions"`
	Menus       []Menu   `json:"menus"`
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
