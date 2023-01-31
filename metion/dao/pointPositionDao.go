package dao

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	"log"
)

func GetPointPosition(position model.PointPosition) (model.ResponsePagination, []model.PointPosition, error) {
	if position.PageSize == 0 {
		position.PageSize = 10
	}
	var pointPositions []model.PointPosition
	var pagination model.ResponsePagination
	var totalTmp int64
	result := global.DB.Joins("left join t_user on t_user.user_id = t_point_position.implementer_id").
		Joins("left join t_user t1 on t1.user_id = t_point_position.user_id").
		Joins("left join t_client client on client.client_id = t_point_position.client_id").
		Select("t_point_position.*, t_user.user_name as implementer_name, t1.user_name as user_name, client.client_abbreviation as client_abbreviation").Offset((position.PageNumber - 1) * position.PageSize).Limit(position.PageSize).Where(position).Find(&pointPositions)
	if result.Error != nil {
		log.Print(result.Error)
		return pagination, nil, result.Error
	}
	result = global.DB.Model(model.PointPosition{}).Joins("left join t_user on t_user.user_id = t_point_position.implementer_id").
		Joins("left join t_user t1 on t1.user_id = t_point_position.user_id").
		Joins("left join t_client client on client.client_id = t_point_position.client_id").
		Select("t_point_position.*, t_user.user_name as implementer_name, t1.user_name as user_name, client.client_abbreviation as client_abbreviation").Where(position).Count(&totalTmp)
	pagination = utils.EncapsulationPage(position.PageNumber, position.PageSize, totalTmp)
	return pagination, pointPositions, nil
}
