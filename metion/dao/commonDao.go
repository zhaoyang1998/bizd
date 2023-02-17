package dao

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	"fmt"
	"log"
	"strconv"
)

func GetTotalDataByClientDao(clientId string) ([]model.EchartsPie, error) {
	var pies []model.EchartsPie
	result := global.DB.Model(model.PointPosition{}).Select("status as name, count(*) as value").
		Where("client_id = ? and type = 1", clientId).Group("status").Find(&pies)
	if result.Error != nil {
		log.Print(result.Error)
		return nil, result.Error
	}
	for i, pie := range pies {
		tmp, _ := strconv.Atoi(pie.Name)
		switch tmp % 10 {
		case 0:
			pies[i].Name = "未实施"
			break
		case 1:
			pies[i].Name = "实施中"
			break
		case 2:
			pies[i].Name = "已实施"
			break
		}
	}
	return pies, nil
}

func GetCurDataByClientDao(clientId string) ([]model.EchartsPie, error) {
	var pies []model.EchartsPie
	fmt.Print(utils.GetCurDayTime())
	fmt.Print(utils.GetNextDayTime())
	result := global.DB.Model(model.PointPosition{}).Select("status as name, count(*) as value").
		Where("client_id = ? and type = 1 and scheduled_time >= ? and scheduled_time < ?", clientId, utils.GetCurDayTime(), utils.GetNextDayTime()).Group("status").Find(&pies)
	if result.Error != nil {
		log.Print(result.Error)
		return nil, result.Error
	}
	for i, pie := range pies {
		tmp, _ := strconv.Atoi(pie.Name)
		switch tmp % 10 {
		case 0:
			pies[i].Name = "未实施"
			break
		case 1:
			pies[i].Name = "实施中"
			break
		case 2:
			pies[i].Name = "已实施"
			break
		}
	}
	return pies, nil
}
