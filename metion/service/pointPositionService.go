package service

import (
	"bizd/metion/db"
	"bizd/metion/model"
	"bizd/metion/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func AddPointPosition(c *gin.Context) {
	var pointPosition model.PointPosition

	er := c.ShouldBindJSON(&pointPosition)
	fmt.Println(er)
	fmt.Printf("%+v", pointPosition)
	// 参数验证
	validate := validator.New()
	err := validate.Struct(pointPosition)
	if err != nil {
		log.Print(err.Error())
		c.JSON(200, gin.H{"code": 201, "msg": err.Error()})
		return
	}
	pointPosition.PointPositionId = uuid.NewV4().String()
	pointPosition.Status = pointPosition.Type * 10
	result := db.DB.Create(pointPosition)
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
	_ = c.Bind(&pointPosition)
	var pointPositions []model.PointPosition
	result := db.DB.Joins("left join t_user on t_user.user_id = t_point_position.implementer_id").
		Joins("left join t_user t1 on t1.user_id = t_point_position.user_id").
		Joins("left join t_client client on client.client_id = t_point_position.client_id").
		Select("t_point_position.*, t_user.user_name as implementer_name, t1.user_name as user_name, client.client_abbreviation as client_abbreviation").Offset((pointPosition.PageNumber - 1) * pointPosition.PageSize).Limit(pointPosition.PageSize).Where(pointPosition).Find(&pointPositions)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	var response model.Response
	data, _ := json.Marshal(pointPositions)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func DelPointPosition(c *gin.Context) {
	var response model.Response
	var pointPosition model.PointPosition
	_ = c.Bind(&pointPosition)
	if pointPosition.PointPositionId == "" {
		response.Code = http.StatusCreated
		response.Message = "单位ID不能为空"
		c.JSON(http.StatusOK, response)
		return
	}
	result := db.DB.Delete(&pointPosition)
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
	_ = c.Bind(&pointPosition)
	result := db.DB.Model(&pointPosition).Updates(&pointPosition)
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
	pointPosition.PointPositionId = c.Query("pointPositionId")
	if pointPosition.PointPositionId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数不能为空",
		})
	}
	result := db.DB.Model(&pointPosition).Where(gorm.Expr("status % 10 = 0")).Updates(map[string]interface{}{"status": gorm.Expr("status + 1"),
		"start_time": utils.GetNowTime(), "current_workload": gorm.Expr("current_workload + 1")})
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

func FinishAssignment(c *gin.Context) {
	var response model.Response
	var pointPosition model.PointPosition
	pointPosition.PointPositionId = c.Query("pointPositionId")
	if pointPosition.PointPositionId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数不能为空",
		})
	}
	result := db.DB.Model(&pointPosition).Where(gorm.Expr("status % 10 = 1")).Updates(map[string]interface{}{"status": gorm.Expr("status + 1"), "end_time": utils.GetNowTime(),
		"current_workload": gorm.Expr("current_workload - 1")})
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
	_ = c.Bind(&pointPosition)
	result := db.DB.Model(&pointPosition).Updates(&pointPosition)
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
