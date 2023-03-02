package dao

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	uuid "github.com/satori/go.uuid"
)

func GetAllDetail(id string) []model.ImplementDetails {
	var tmp []model.ImplementDetails
	result := global.DB.Where("point_position_id = ?", id).Order("seq asc").Find(&tmp)
	if result.Error != nil {
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
	return tmp
}

func SaveDetailDao(details []model.ImplementDetails) {
	tx := global.DB.Begin()
	var ids []string
	for i, item := range details {
		item.Seq = i + 1
		if item.ImplementDetailId == "" {
			item.ImplementDetailId = uuid.NewV4().String()
			result := tx.Create(&item)
			if result.Error != nil {
				tx.Rollback()
				panic(model.MyError{Code: 400, Message: result.Error.Error()})
			}
		} else {
			result := tx.Model(&item).Updates(&item)
			if result.Error != nil {
				tx.Rollback()
				panic(model.MyError{Code: 400, Message: result.Error.Error()})
			}
		}
		ids = append(ids, item.ImplementDetailId)
	}
	result := tx.Where("point_position_id = ? and implement_detail_id not in ?", details[0].PointPositionId, ids).
		Delete(model.ImplementDetails{})
	if result.Error != nil {
		tx.Rollback()
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
	utils.JudgeCommit(tx)
}
