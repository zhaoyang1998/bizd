package service

import (
	"bizd/metion/dao"
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func AddPointPosition(c *gin.Context) {
	var pointPosition model.PointPosition
	_ = c.ShouldBindJSON(&pointPosition)
	// 参数验证
	validate := validator.New()
	err := validate.Struct(pointPosition)
	if err != nil {
		log.Print(err.Error())
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	pointPosition.PointPositionId = uuid.NewV4().String()
	pointPosition.UserId = utils.GetCurrentUserId(c)
	var tmp = *pointPosition.Type * 10
	pointPosition.Status = &tmp
	result := global.DB.Create(&pointPosition)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
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
	pagination, pointPositions, err = dao.GetPointPositionDao(pointPosition)
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	utils.TranPointPositionStatus(pointPositions)
	var response model.Response
	data, _ := json.Marshal(pointPositions)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	response.ResponsePagination = pagination
	c.JSON(http.StatusOK, response)
}

func GetPointPositionByKeyword(c *gin.Context) {
	var search model.Search
	_ = c.ShouldBindJSON(&search)
	var pointPositions []model.PointPosition
	var pagination model.ResponsePagination
	var err error
	pagination, pointPositions, err = dao.GetPointPositionByKeywordDao(search)
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	utils.TranPointPositionStatus(pointPositions)
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
	var del = model.DelModel{}
	var err error
	_ = c.ShouldBindJSON(&del)
	if len(del.Keys) == 0 {
		response.Code = http.StatusCreated
		response.Message = "参数不能为空"
		c.JSON(http.StatusOK, response)
		return
	}
	err = dao.DelPointPositionByKeys(del)
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
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
		c.JSON(200, gin.H{"code": 400, "message": result.Error})
		return
	}
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func StartAssignment(c *gin.Context) {
	var response model.Response
	var pointPosition model.PointPosition
	err := c.Bind(&pointPosition)
	if err != nil || pointPosition.PointPositionId == "" {
		c.JSON(200, gin.H{"code": 400, "message": "请求参数错误"})
		return
	}
	err = dao.StartAssignmentDao(pointPosition)
	if err == nil {
		response.Code = 200
		response.Message = "请求成功"
	} else {
		response.Code = 400
		response.Message = err.Error()
	}
	c.JSON(http.StatusOK, response)
}

func CancelAssignment(c *gin.Context) {
	var response model.Response
	var pointPosition model.PointPosition
	err := c.Bind(&pointPosition)
	if err != nil || pointPosition.PointPositionId == "" {
		c.JSON(200, gin.H{"code": 400, "message": "请求参数错误"})
		return
	}
	err = dao.CancelAssignmentDao(pointPosition)
	if err == nil {
		response.Code = 200
		response.Message = "请求成功"
	} else {
		response.Code = 400
		response.Message = err.Error()
	}
	c.JSON(http.StatusOK, response)
}

func FinishAssignment(c *gin.Context) {
	var response model.Response
	var pointPosition model.PointPosition
	var err error
	pointPosition.PointPositionId = c.Param("pointPositionId")
	if pointPosition.PointPositionId == "" {
		c.JSON(200, gin.H{"code": 400, "message": "请求参数错误"})
	}
	err = dao.FinishAssignmentDao(pointPosition)
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err})
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
	if pointPosition.ImplementerId == "" {
		c.JSON(200, gin.H{"code": 400, "message": "请求参数错误"})
		return
	}
	err := dao.AllocatingAssignmentDao(pointPosition)
	if err != nil {
		response.Code = 400
		response.Message = err.Error()
	} else {
		response.Code = 200
		response.Message = "请求成功"
	}
	c.JSON(http.StatusOK, response)
}
