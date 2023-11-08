package dao

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	"container/list"
	"log"
	"strconv"
)

func GetTotalDataByClientDao(clientId string) (model.EchartsPie, error) {
	pie := utils.InitPie()
	var data []model.EchartsPieData
	result := global.DB.Model(model.PointPosition{}).Select("status as name, count(*) as value").
		Where("client_id = ? ", clientId).Group("status").Find(&data)
	if result.Error != nil {
		log.Print(result.Error)
		return pie, result.Error
	}
	for i, item := range data {
		tmp, _ := strconv.Atoi(item.Name)
		switch tmp % 10 {
		case 0:
			data[i].Name = "未实施"
			break
		case 1:
			data[i].Name = "实施中"
			break
		case 2:
			data[i].Name = "已实施"
			break
		}
	}
	pie.Series[0].Data = data
	pie.Title.Text = "总实施数据"
	return pie, nil
}

func GetCurDataByClientDao(clientId string) (model.EchartsPie, error) {
	pie := utils.InitPie()
	var data []model.EchartsPieData
	result := global.DB.Model(model.PointPosition{}).Select("status as name, count(*) as value").
		Where("client_id = ? and scheduled_time >= ? and scheduled_time < ?", clientId, utils.GetCurDayTime(), utils.GetNextDayTime()).Group("status").Find(&data)
	if result.Error != nil {
		log.Print(result.Error)
		return pie, result.Error
	}
	for i, item := range data {
		tmp, _ := strconv.Atoi(item.Name)
		switch tmp % 10 {
		case 0:
			data[i].Name = "未实施"
			break
		case 1:
			data[i].Name = "实施中"
			break
		case 2:
			data[i].Name = "已实施"
			break
		}
	}
	pie.Series[0].Data = data
	pie.Title.Text = "当日实施数据"
	return pie, nil
}

func GetEfficiencyDataDao(clientId string) (model.EchartsLine, error) {
	line := utils.InitAxis()
	userIds := list.New()
	var pointPositionIds []string
	var pps []model.PointPosition
	var userNames []string
	result := global.DB.Model(&model.PointPosition{}).Joins("left join t_user user on user.user_id = t_point_position.implementer_id").
		Select("t_point_position.*, user.user_name as implementer_name").
		Where("client_id = ? and YEARWEEK(NOW()) - YEARWEEK(DATE_FORMAT(scheduled_time,'%Y-%m-%d')) < ?", clientId, 4).Group("implementer_id").Find(&pps)
	if (result.Error != nil) || (len(pps) == 0) {
		return line, result.Error
	}
	var series []model.EchartsLineSeries
	for _, item := range pps {
		userNames = append(userNames, item.ImplementerName)
		userIds.PushBack(item.ImplementerId)
		tmp := model.EchartsLineSeries{
			Type: "line",
			Data: []int{0, 0, 0, 0},
			Name: item.ImplementerName,
		}
		series = append(series, tmp)
		pointPositionIds = append(pointPositionIds, item.PointPositionId)
	}
	line.Legend.Data = userNames
	var i = 0
	for item := userIds.Front(); item != nil; item = item.Next() {
		var tmp []model.Result
		result = global.DB.Model(&model.PointPosition{}).Select("YEARWEEK(NOW()) - YEARWEEK(DATE_FORMAT(scheduled_time,'%Y-%m-%d')) as weeks,AVG(total_time) as times").
			Where("client_id = ? and YEARWEEK(NOW()) - YEARWEEK(DATE_FORMAT(scheduled_time,'%Y-%m-%d')) < ? and status %10 = 2 and implementer_id = ?", clientId, 4, item.Value).
			Group("YEARWEEK(NOW()) - YEARWEEK(DATE_FORMAT(scheduled_time,'%Y-%m-%d'))").Order("YEARWEEK(NOW()) - YEARWEEK(DATE_FORMAT(scheduled_time,'%Y-%m-%d')) DESC").Find(&tmp)
		if result.Error != nil {
			return line, result.Error
		}
		for _, item := range tmp {
			series[i].Data[3-item.Weeks] = int(item.Times / 60)
		}
		for j := 1; j < len(series[i].Data); j++ {
			if series[i].Data[j] == 0 {
				series[i].Data[j] = series[i].Data[j-1]
			}
		}
		i++
	}
	line.Series = series
	return line, nil
}
