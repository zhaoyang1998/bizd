package timedTask

import (
	"bizd/metion/db"
	"bizd/metion/model"
	"bizd/metion/utils"
	"errors"
	"github.com/jakecoffman/cron"
	"github.com/spf13/cast"
	"log"
)

const (
	SendTureInCron  = 0
	SendFalseInCron = 1
)

type Task struct {
	CronTask  *cron.Cron
	CronCount int
}

// 初始化数据库的定时任务
func (task *Task) InitCron() {

	//初始定时任务 初始化时 将任务总数记为0
	task.CronCount = 0
	//定义 最终返回的gc
	//初始化一个错误

	err := errors.New("")
	err = nil

	msgFromAs, err := QueryMsgFromCron(db.DB)
	if err != nil {
		log.Println("在初始化定时任务的时候 获取A表数据失败")
		return
	}
	for _, msgFromCron := range msgFromAs {

		//添加变量 来解决循环 拷贝问题  就是会重复替换值 导致每次加载值有问题
		tmpMsgFromCron := msgFromCron
		//添加发送微信任务
		task.AddTask(tmpMsgFromCron)
	}
	//fmt.Println(c.Entries())
}

// 重新初始化任务
func (task *Task) ResetTask() error {

	//删除原来所有的任务
	for _, entry := range task.CronTask.Entries() {
		task.CronTask.RemoveJob(entry.Name)
	}

	//初始化
	task.InitCron()
	return nil
}

// 添加task任务
func (task *Task) AddTask(MsgFromCron model.MsgFromCron) error {
	if MsgFromCron.IsSend == SendTureInCron {
		task.CronTask.AddFunc(MsgFromCron.CronTime, func() {
			//添加发送任务
			utils.SendTimeNotice(MsgFromCron)
		}, cast.ToString(MsgFromCron.Id)+"-cron")
		task.CronCount += 1
	}

	return nil

}

////测试发送一次 不触发重置定时任务 A表的单个
//func (task Task) SentOneCronTask(msgFromA connectToDatabase.MsgFromA) error {
//	if msgFromA.IsSend == SendTureInA {
//	}
//	err := notice.SendMsg(task.Base, msgFromA)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
////测试全部的任务
//func (task Task) SendAllCronTask() error {
//	msgFromAs, err := connectToDatabase.QueryMsgFromA(task.Base)
//	if err != nil {
//		return err
//	}
//	for _, msgFromA := range msgFromAs {
//
//		//添加变量 来解决循环 拷贝问题  就是会重复替换值 导致每次加载值有问题
//		tmpMsgFromA := msgFromA
//		//添加任务
//		task.SentOneCronTask(tmpMsgFromA)
//	}
//	return nil
//}
//
////重置所有任务
//func (task Task) ResetAllRestTask() error {
//	msgFromAs, err := connectToDatabase.QueryMsgFromA(task.Base)
//	if err != nil {
//		return err
//	}
//	for _, msgFromA := range msgFromAs {
//
//		//添加变量 来解决循环 拷贝问题  就是会重复替换值 导致每次加载值有问题
//		tmpMsgFromA := msgFromA
//		//添加任务
//		task.ResetOneResetTask(tmpMsgFromA)
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
//			err := notice.ResetCompleteStatus(task.Base, msgFromA)
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
