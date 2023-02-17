package dao

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/timedTask"
	"bizd/metion/utils"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
	"unsafe"
)

func GetPointPositionDao(position model.PointPosition) (model.ResponsePagination, []model.PointPosition, error) {
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

func GetPointPositionByKeywordDao(search model.Search) (model.ResponsePagination, []model.PointPosition, error) {
	if search.PageSize == 0 {
		search.PageSize = 10
	}
	var pointPositions []model.PointPosition
	var pagination model.ResponsePagination
	var totalTmp int64
	result := global.DB.Joins("left join t_user on t_user.user_id = t_point_position.implementer_id").
		Joins("left join t_user t1 on t1.user_id = t_point_position.user_id").
		Joins("left join t_client client on client.client_id = t_point_position.client_id").
		Select("t_point_position.*, t_user.user_name as implementer_name, t1.user_name as user_name, client.client_abbreviation as client_abbreviation").Offset((search.PageNumber-1)*search.PageSize).Limit(search.PageSize).
		Where("point_position_name LIKE ? or address LIKE ? or cpe_name LIKE ? or ip LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%", "%"+search.Keyword+"%", "%"+search.Keyword+"%").Find(&pointPositions)
	if result.Error != nil {
		log.Print(result.Error)
		return pagination, nil, result.Error
	}
	result = global.DB.Model(model.PointPosition{}).Joins("left join t_user on t_user.user_id = t_point_position.implementer_id").
		Joins("left join t_user t1 on t1.user_id = t_point_position.user_id").
		Joins("left join t_client client on client.client_id = t_point_position.client_id").
		Where("point_position_name LIKE ? or address LIKE ? or cpe_name LIKE ? or ip LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%", "%"+search.Keyword+"%", "%"+search.Keyword+"%").Count(&totalTmp)
	pagination = utils.EncapsulationPage(search.PageNumber, search.PageSize, totalTmp)
	return pagination, pointPositions, nil
}

func DelPointPositionByKeys(del model.DelModel) error {
	result := global.DB.Model(model.PointPosition{}).Delete("pointPositionId", del.Keys)
	if result.Error != nil {
		log.Print(result.Error)
		return result.Error
	} else {
		return nil
	}
}
func FinishAssignmentDao(pointPosition model.PointPosition) error {
	tx := global.DB.Begin()
	result := tx.Model(&pointPosition).Where(gorm.Expr("status % 10 = 1")).Updates(map[string]interface{}{"status": gorm.Expr("status + 1"), "end_time": utils.GetNowTime()})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Exec("SELECT count(*) FROM `t_point_position` WHERE t_point_position.status % 10 != 2 and t_point_position.client_id = (SELECT `client_id` FROM `t_point_position` WHERE `t_point_position`.`point_position_id` = ?)",
		pointPosition.PointPositionId)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Exec("UPDATE t_client client LEFT JOIN t_point_position pp ON client.client_id = pp.client_id SET client.status = client.status + 1 WHERE pp.point_position_id = ? and (client.status % 10 = 0 or "+strconv.Itoa(*(*int)(unsafe.Pointer(&result.RowsAffected)))+"=1)",
		pointPosition.PointPositionId)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Exec("UPDATE t_user user LEFT JOIN t_point_position pp ON user.user_id = pp.implementer_id SET user.current_workload = current_workload - 1 WHERE pp.point_position_id = ?",
		pointPosition.PointPositionId)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

func AllocatingAssignmentDao(pointPosition model.PointPosition) error {
	var user model.User
	var tmp []model.MsgFromCron
	tx := global.DB.Begin()
	// 判断当前任务是否已分配
	result := tx.Where("point_position_id = ?", pointPosition.PointPositionId).Find(&tmp)
	if result.RowsAffected != 0 {
		return errors.New("当前实施地点已分配")
	}
	result = tx.Model(&pointPosition).Updates(map[string]interface{}{"implementer_id": pointPosition.ImplementerId, "scheduled_time": pointPosition.ScheduledTime})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Model(model.User{UserId: pointPosition.ImplementerId}).Update("current_workload", gorm.Expr("current_workload + 1"))
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	user.UserId = pointPosition.ImplementerId
	result = tx.Find(&user)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	var msgFormCorns, err = encapsulationTasks(pointPosition, user)
	if err != nil {
		tx.Rollback()
		return errors.New("任务分配失败")
	}
	result = tx.Create(msgFormCorns)
	if result.Error != nil {
		tx.Rollback()
		return err
	}
	for _, val := range msgFormCorns {
		err := timedTask.AddTask(val)
		if err != nil {
			tx.Rollback()
			return errors.New("添加定时任务失败")
		}
	}
	return tx.Commit().Error
}

func StartAssignmentDao(pointPosition model.PointPosition) error {
	var msgFormCron []model.MsgFromCron
	tx := global.DB.Begin()
	result := tx.Where("point_position_id = ?", pointPosition.PointPositionId).Find(&msgFormCron)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	for _, val := range msgFormCron {
		global.Tasks.CronTask.RemoveJob(val.Name)
	}
	result = tx.Where("point_position_id = ?", pointPosition.PointPositionId).Delete(model.MsgFromCron{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Model(&pointPosition).Where(gorm.Expr("status % 10 = 0")).Updates(map[string]interface{}{"status": gorm.Expr("status + 1"),
		"start_time": utils.GetNowTime()})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

func CancelAssignmentDao(pointPosition model.PointPosition) error {
	var msgFormCron []model.MsgFromCron
	tx := global.DB.Begin()
	result := tx.Where("point_position_id = ?", pointPosition.PointPositionId).Find(&msgFormCron)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	for _, val := range msgFormCron {
		global.Tasks.CronTask.RemoveJob(val.Name)
	}
	result = tx.Where("point_position_id = ?", pointPosition.PointPositionId).Delete(model.MsgFromCron{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Exec("UPDATE t_user user LEFT JOIN t_point_position pp ON user.user_id = pp.implementer_id SET user.current_workload = current_workload - 1 WHERE pp.point_position_id = ?",
		pointPosition.PointPositionId)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	result = tx.Model(&pointPosition).Updates(map[string]interface{}{"status": gorm.Expr("(status / 10) * 10"),
		"start_time": nil, "implementer_id": nil, "scheduled_time": nil})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}

func encapsulationTasks(position model.PointPosition, user model.User) ([3]model.MsgFromCron, error) {
	var msgFormCorns [3]model.MsgFromCron
	msgFormCorns[0].Id = uuid.NewV4().String()
	msgFormCorns[1].Id = uuid.NewV4().String()
	msgFormCorns[2].Id = uuid.NewV4().String()
	msgFormCorns[0].Receive = global.WxUrl
	msgFormCorns[1].Receive = global.WxUrl
	msgFormCorns[2].Receive = global.WxUrl
	msgFormCorns[0].Name = position.PointPositionName + "-0-" + global.AllocatingAssignmentTag
	msgFormCorns[1].Name = position.PointPositionName + "-1-" + global.AssignmentStartTag
	msgFormCorns[2].Name = position.PointPositionName + "-3"
	msgFormCorns[0].Type = 0
	msgFormCorns[1].Type = 1
	msgFormCorns[2].Type = 3
	msgFormCorns[0].PointPositionId = position.PointPositionId
	msgFormCorns[1].PointPositionId = position.PointPositionId
	msgFormCorns[2].PointPositionId = position.PointPositionId
	msgFormCorns[0].WxName = user.WxName
	msgFormCorns[1].WxName = user.WxName
	msgFormCorns[2].WxName = user.WxName
	msgFormCorns[0].ScheduledTime = position.ScheduledTime
	msgFormCorns[1].ScheduledTime = position.ScheduledTime
	msgFormCorns[2].ScheduledTime = position.ScheduledTime
	var scheduledTime, err = time.Parse(global.TimeFormat, position.ScheduledTime)
	if err != nil {
		return msgFormCorns, err
	}
	msgFormCorns[0].CronTime = utils.DateConversionCron(time.Now().Add(time.Second * 60 * 1))
	msgFormCorns[1].CronTime = utils.DateConversionCron(scheduledTime.Add(time.Minute * -1))
	msgFormCorns[2].CronTime = utils.DateConversionCron(scheduledTime)
	return msgFormCorns, err
}
