package timedTask

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	"errors"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
)

// 初始化数据库的定时任务
func InitCron() {

	//初始定时任务 初始化时 将任务总数记为0
	global.Tasks.CronCount = 0
	//定义 最终返回的gc
	//初始化一个错误

	err := errors.New("")
	err = nil

	msgFromAs, err := QueryMsgFromCron(global.DB)
	if err != nil {
		log.Println("在初始化定时任务的时候 获取A表数据失败")
		return
	}
	for _, msgFromCron := range msgFromAs {

		//添加变量 来解决循环 拷贝问题  就是会重复替换值 导致每次加载值有问题
		tmpMsgFromCron := msgFromCron
		//添加发送微信任务
		err := AddTask(tmpMsgFromCron)
		if err != nil {
			return
		}
	}
	//fmt.Println(c.Entries())
}

// 重新初始化任务
func ResetTask() error {

	//删除原来所有的任务
	for _, entry := range global.Tasks.CronTask.Entries() {
		global.Tasks.CronTask.RemoveJob(entry.Name)
	}

	//初始化
	InitCron()
	return nil
}

// 添加task任务
func AddTask(msgFromCron model.MsgFromCron) error {
	global.Tasks.CronTask.AddFunc(msgFromCron.CronTime, func() {
		if msgFromCron.Type == 3 {
			err := encapsulationTask(msgFromCron)
			if err == nil {
				// 删除数据库
				global.DB.Delete(&msgFromCron)
				// 删除定时任务
				global.Tasks.CronTask.RemoveJob(msgFromCron.Name)
				return
			}
		}
		//添加发送任务
		utils.SendTimeNotice(msgFromCron)
	}, msgFromCron.Name)
	global.Tasks.CronCount += 1
	return nil
}

func encapsulationTask(cron model.MsgFromCron) error {
	var pointPosition model.PointPosition
	pointPosition.PointPositionId = cron.PointPositionId
	_ = global.DB.Find(&pointPosition)
	if *pointPosition.Status%10 != 0 {
		return nil
	}
	var tmp model.MsgFromCron
	tmp.Id = uuid.NewV4().String()
	tmp.Receive = cron.Receive
	tmp.WxName = cron.WxName
	tmp.ScheduledTime = cron.ScheduledTime
	tmp.Name = strings.Replace(cron.Name, "-3", "-2-", -1) + global.AssignmentNotStartedTag
	tmp.Type = 2
	tmp.CronTime = "0/30 * * * * *"
	tmp.PointPositionId = cron.PointPositionId
	global.DB.Create(tmp)
	err := AddTask(tmp)
	return err
}

////测试发送一次 不触发重置定时任务 A表的单个
//func (task Task) SentOneCronTask(msgFromA connectToDatabase.MsgFromA) error {
//	if msgFromA.IsSend == SendTureInA {
//	}
//	err := notice.SendMsg(init.Tasks.Base, msgFromA)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
////测试全部的任务
//func (task Task) SendAllCronTask() error {
//	msgFromAs, err := connectToDatabase.QueryMsgFromA(init.Tasks.Base)
//	if err != nil {
//		return err
//	}
//	for _, msgFromA := range msgFromAs {
//
//		//添加变量 来解决循环 拷贝问题  就是会重复替换值 导致每次加载值有问题
//		tmpMsgFromA := msgFromA
//		//添加任务
//		init.Tasks.SentOneCronTask(tmpMsgFromA)
//	}
//	return nil
//}
//
////重置所有任务
//func (task Task) ResetAllRestTask() error {
//	msgFromAs, err := connectToDatabase.QueryMsgFromA(init.Tasks.Base)
//	if err != nil {
//		return err
//	}
//	for _, msgFromA := range msgFromAs {
//
//		//添加变量 来解决循环 拷贝问题  就是会重复替换值 导致每次加载值有问题
//		tmpMsgFromA := msgFromA
//		//添加任务
//		init.Tasks.ResetOneResetTask(tmpMsgFromA)
//	}
//	return nil
//}
//
////重置单个任务
//func (task Task) ResetOneResetTask(msgFromA connectToDatabase.MsgFromA) error {
//
//	if msgFromA.IsSend == SendTureInA {
//		//当需要发送的时候 设置重置
//		if msgFromA.ResetTime != ResetTrueInA {
//
//			//重置方法
//			err := notice.ResetCompleteStatus(init.Tasks.Base, msgFromA)
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
