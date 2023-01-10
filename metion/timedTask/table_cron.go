package timedTask

import (
	"bizd/metion/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

//QueryMsgFromCron 从Cron表中获取数据
func QueryMsgFromCron(base *gorm.DB) ([]model.MsgFromCron, error) {
	var res []model.MsgFromCron

	query := base.Find(&res)
	if query.Error != nil {
		return nil, errors.New("读取数据库Cron信息错误")
	}

	fmt.Printf("cron:%#v\n", res)
	return res, nil
}

//
////获取A表id最大值
//func QueryMaxIdFromCron(base *sql.DB) (int, error) {
//	var maxId int
//	query, err := base.Query("select  id from Cron order by id DESC limit 1")
//	if err != nil {
//		return 0, errors.New("读取数据库A信息错误")
//	}
//
//	//逐个添加
//	for query.Next() {
//		err := query.Scan(&maxId)
//		if err != nil {
//			return 0, errors.New("拿到数据库信息后 解析数据库有问题")
//		}
//	}
//	return maxId, nil
//}
//
////更新A表数据
//func UpdateCronTable(base *sql.DB, MsgFromCron MsgFromCron) error {
//
//	fmt.Println(MsgFromCron)
//	//判断是不是需要新增
//	if MsgFromCron.Id == 999999 {
//		_, err := base.Exec("insert  into Cron  (name,cronTime,resetTime,receive,receiveType,tags,isSend) values (?,?,?,?,?,?,?)",
//			MsgFromCron.Name, MsgFromCron.CronTime, MsgFromCron.ResetTime, MsgFromCron.Receive, MsgFromCron.ReceiveType, MsgFromCron.Tags,
//			MsgFromCron.IsSend)
//
//		return err
//	} else {
//		_, err := base.Exec("update Cron set name=?,cronTime=?,resetTime=?,receive=?,receiveType=?,tags=?,isSend=? where id=?",
//			MsgFromCron.Name, MsgFromCron.CronTime, MsgFromCron.ResetTime, MsgFromCron.Receive, MsgFromCron.ReceiveType, MsgFromCron.Tags,
//			MsgFromCron.IsSend, MsgFromCron.Id)
//		return err
//	}
//
//}
//
////删除A表数据
//func DeleteDataFromCron(base *sql.DB, id int) error {
//	_, err := base.Exec("DELETE FROM A WHERE id=?", id)
//
//	return err
//}
