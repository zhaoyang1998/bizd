package service

import (
	"bizd/metion/dao"
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/timedTask"
	"bizd/metion/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func AddPointPosition(c *gin.Context) {
	var pointPosition model.PointPosition
	_ = c.ShouldBindJSON(&pointPosition)
	// 参数验证
	validate := validator.New()
	err := validate.Struct(pointPosition)
	if err != nil {
		log.Print(err.Error())
		c.JSON(200, gin.H{"code": 201, "msg": err.Error()})
		return
	}
	pointPosition.PointPositionId = uuid.NewV4().String()
	var tmp = *pointPosition.Type * 10
	pointPosition.Status = &tmp
	result := global.DB.Create(pointPosition)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	var response model.Response
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func GetPointPosition(c *gin.Context) {
	var pointPosition model.PointPosition
	_ = c.ShouldBindJSON(&pointPosition)
	var pointPositions []model.PointPosition
	var pagination model.ResponsePagination
	var err error
	pagination, pointPositions, err = dao.GetPointPosition(pointPosition)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	var response model.Response
	data, _ := json.Marshal(pointPositions)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	response.ResponsePagination = pagination
	c.JSON(http.StatusOK, response)
}

func DelPointPosition(c *gin.Context) {
	var response model.Response
	var pointPosition model.PointPosition
	_ = c.ShouldBindJSON(&pointPosition)
	if pointPosition.PointPositionId == "" {
		response.Code = http.StatusCreated
		response.Message = "单位ID不能为空"
		c.JSON(http.StatusOK, response)
		return
	}
	result := global.DB.Delete(&pointPosition)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func UpdatePointPosition(c *gin.Context) {
	var response model.Response
	var pointPosition model.PointPosition
	_ = c.ShouldBindJSON(&pointPosition)
	if pointPosition.Type != nil {
		fmt.Println(*pointPosition.Type * 10)
		var tmp = *pointPosition.Type * 10
		pointPosition.Status = &tmp
	}
	result := global.DB.Model(&pointPosition).Updates(&pointPosition).Update("status", pointPosition.Status)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func StartAssignment(c *gin.Context) {
	var response model.Response
	var pointPosition model.PointPosition
	var msgFormCron []model.MsgFromCron
	err := c.Bind(&pointPosition)
	if err != nil || pointPosition.PointPositionId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": "参数请求错误",
		})
		return
	}
	result := global.DB.Where("point_position_id = ?", pointPosition.PointPositionId).Find(&msgFormCron)
	for _, val := range msgFormCron {
		global.Tasks.CronTask.RemoveJob(val.Name)
	}
	result = global.DB.Where("point_position_id = ?", pointPosition.PointPositionId).Delete(model.MsgFromCron{})
	if pointPosition.PointPositionId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数不能为空",
		})
	}
	result = global.DB.Model(&pointPosition).Where(gorm.Expr("status % 10 = 0")).Updates(map[string]interface{}{"status": gorm.Expr("status + 1"),
		"start_time": utils.GetNowTime()})
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": result.Error,
		})
		return
	}
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func FinishAssignment(c *gin.Context) {
	var response model.Response
	var pointPosition model.PointPosition
	pointPosition.PointPositionId = c.Query("pointPositionId")
	if pointPosition.PointPositionId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数不能为空",
		})
	}
	result := global.DB.Model(&pointPosition).Where(gorm.Expr("status % 10 = 1")).Updates(map[string]interface{}{"status": gorm.Expr("status + 1"), "end_time": utils.GetNowTime()})
	result = global.DB.Model(model.User{UserId: pointPosition.ImplementerId}).Update("current_workload", gorm.Expr("current_workload + 1"))
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func AllocatingAssignment(c *gin.Context) {
	var response model.Response
	var pointPosition model.PointPosition
	var user model.User
	var tmp []model.MsgFromCron
	_ = c.Bind(&pointPosition)
	// 判断当前任务是否已分配
	result := global.DB.Where("point_position_id = ?", pointPosition.PointPositionId).Find(&tmp)
	if result.RowsAffected != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    200,
			"message": "当前实施地点已分配",
		})
		return
	}
	result = global.DB.Model(&pointPosition).Updates(&pointPosition)
	result = global.DB.Model(model.User{UserId: pointPosition.ImplementerId}).Update("current_workload", gorm.Expr("current_workload + 1"))
	user.UserId = pointPosition.ImplementerId
	result = global.DB.Find(&user)
	var msgFormCorns, err = encapsulationTasks(pointPosition, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": "参数请求错误",
		})
		return
	}
	global.DB.Create(msgFormCorns)
	for _, val := range msgFormCorns {
		err := timedTask.AddTask(val)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "添加定时任务失败",
			})
			return
		}
	}
	for _, val := range global.Tasks.CronTask.Entries() {
		fmt.Printf("%+v", val)
		fmt.Println()
	}
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": result.Error,
		})
		return
	}
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
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
	fmt.Println(scheduledTime)
	msgFormCorns[0].CronTime = utils.DateConversionCron(time.Now().Add(time.Second * 60 * 1))
	msgFormCorns[1].CronTime = utils.DateConversionCron(scheduledTime.Add(time.Minute * 2 * -1))
	msgFormCorns[2].CronTime = utils.DateConversionCron(scheduledTime)
	return msgFormCorns, err
}
