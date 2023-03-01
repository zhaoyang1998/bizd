package dao

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func GetPPDetailDao(details model.ImplementDetails) (model.ImplementDetails, error) {
	var tmp model.ImplementDetails
	result := global.DB.Model(model.ImplementDetails{}).Where("point_position_id = ? and seq = ?", details.PointPositionId, details.Seq).Find(&tmp)
	if result.Error == nil && result.RowsAffected == 0 {
		details.ImplementDetailId = uuid.NewV4().String()
		global.DB.Create(&details)
		details.Total = 1
		return details, nil
	} else if result.Error != nil {
		return tmp, result.Error
	}
	result = global.DB.Model(model.ImplementDetails{}).Where("point_position_id = ?", details.PointPositionId).Count(&tmp.Total)
	return tmp, nil
}

//func GetextPPDetailDao(details model.ImplementDetails) (model.ImplementDetails, error) {
//	tx := global.DB.Begin()
//	var tmp model.ImplementDetails
//	result := tx.Model(&details).Updates(&details)
//	if result.Error != nil {
//		tx.Rollback()
//		return tmp, result.Error
//	}
//	result = tx.Select("seq").Where("point_position_id = ?", details.PointPositionId).Order("seq desc").Find(&tmp)
//	if result.Error != nil {
//		tx.Rollback()
//		return tmp, result.Error
//	}
//	if tmp.Seq > details.Seq {
//		details.Seq++
//		tmp, err := GetPPDetailDao(details)
//		if err == nil {
//			return tmp, tx.Commit().Error
//		} else {
//			tx.Rollback()
//			return tmp, err
//		}
//	}
//	tmp.PointPositionId = details.PointPositionId
//	tmp.ImplementDetailId = uuid.NewV4().String()
//	tmp.Seq++
//	result = tx.Create(&tmp)
//	if result.Error != nil {
//		tx.Rollback()
//		return tmp, result.Error
//	}
//	tmp.Total = int64(tmp.Seq)
//	return tmp, tx.Commit().Error
//}

func GetNextPPDetailDao(details model.ImplementDetails) model.ImplementDetails {
	tx := global.DB.Begin()
	var tmp model.ImplementDetails
	UpdateCurDetail(details, tx)
	flag := JudgeCurLatest(details, tx)
	if !flag {
		details.Seq++
		tmp, _ := GetPPDetailDao(details)
		utils.JudgeCommit(tx)
		return tmp
	}
	tmp.PointPositionId = details.PointPositionId
	tmp.ImplementDetailId = uuid.NewV4().String()
	tmp.Seq = details.Seq + 1
	tmp = CreateDetail(tmp, tx)
	utils.JudgeCommit(tx)
	return tmp
}

func GetPrevPPDetailDao(details model.ImplementDetails) model.ImplementDetails {
	tx := global.DB.Begin()
	var tmp model.ImplementDetails
	UpdateCurDetail(details, tx)
	details.Seq--
	tmp, _ = GetPPDetailDao(details)
	utils.JudgeCommit(tx)
	return tmp
}

func CreateDetail(details model.ImplementDetails, tx *gorm.DB) model.ImplementDetails {
	result := tx.Create(&details)
	if result.Error != nil {
		tx.Rollback()
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
	details.Total = int64(details.Seq)
	return details
}

func UpdateCurDetail(details model.ImplementDetails, tx *gorm.DB) {
	details.TotalTime = utils.TimeFormatToUnix(details.EndTime) - utils.TimeFormatToUnix(details.StartTime)
	result := tx.Model(&details).Updates(&details)
	if result.Error != nil {
		tx.Rollback()
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
}

func JudgeCurLatest(details model.ImplementDetails, tx *gorm.DB) bool {
	var tmp model.ImplementDetails
	result := tx.Select("seq").Where("point_position_id = ?", details.PointPositionId).Order("seq desc").Find(&tmp)
	if result.Error != nil {
		tx.Rollback()
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
	return details.Seq == tmp.Seq
}

func DelPPDetailDao(id string) model.ImplementDetails {
	tx := global.DB.Begin()
	tmp := GetCurPPDetailDao(id, tx)
	result := tx.Model(model.ImplementDetails{}).Where("seq > ? ", tmp.Seq+1).
		Updates(map[string]interface{}{"seq": gorm.Expr("seq - 1")})
	if result.Error != nil {
		tx.Rollback()
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
	result = tx.Where("implement_detail_id = ?", id).Delete(model.ImplementDetails{})
	if result.Error != nil {
		tx.Rollback()
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
	utils.JudgeCommit(tx)
	return tmp
}

func GetAllDetail(id string) []model.ImplementDetails {
	var tmp []model.ImplementDetails
	result := global.DB.Where("point_position_id = ?", id).Order("seq asc").Find(&tmp)
	if result.Error != nil {
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
	return tmp
}

func GetCurPPDetailDao(id string, tx *gorm.DB) model.ImplementDetails {
	var tmp model.ImplementDetails
	result := tx.Model(model.ImplementDetails{}).Where("seq > (?) - 1 ", tx.Model(model.ImplementDetails{}).Select("seq").Where("implement_detail_id = ?", id)).
		Find(&tmp)
	if result.Error != nil {
		tx.Rollback()
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
	result = global.DB.Model(model.ImplementDetails{}).Where("point_position_id = ?", tmp.PointPositionId).Count(&tmp.Total)
	if result.Error != nil {
		tx.Rollback()
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
	return tmp
}
