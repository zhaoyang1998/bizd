package service

import (
	"bizd/metion/db"
	"bizd/metion/model"
	"bizd/metion/utils"
)

func GetCurrentConductor() {
	var tmp model.Conductor
	var total int
	// 总人数
	db.DB.Select("count(*) as total").Table("t_conductor").Pluck("total", &total)
	// 获取当前主持人
	db.DB.Where("current_able = ? ", true).First(&tmp)
	// 更新当前主持人
	db.DB.Model(&model.Conductor{}).Where("current_able = ? ", true).
		Update("current_able", false)
	// 选择下次主持人
	db.DB.Model(&model.Conductor{}).Where("seq = ? ", (tmp.Seq+1)%total).
		Update("current_able", true)
	// 获取下一个主持人
	utils.SendConductorMessage(tmp)

}
