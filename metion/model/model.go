package model

type Conductor struct {
	Id          int `gorm:"primaryKey"`
	Username    string
	CurrentAble bool
	WxName      string
	Seq         int
}

// 对应数据库中的表名

func (Conductor) TableName() string {
	return "t_conductor"
}
