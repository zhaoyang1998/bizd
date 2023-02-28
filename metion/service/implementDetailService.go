package service

import (
	"bizd/metion/dao"
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

func GetPPDetail(c *gin.Context) {
	var details model.ImplementDetails
	_ = c.ShouldBindJSON(&details)
	var msg model.Response
	details, err := dao.GetPPDetailDao(details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": "参数请求错误",
		})
		return
	}
	msg.Code = 200
	msg.Message = "请求成功"
	tmp, _ := json.Marshal(details)
	msg.Data = string(tmp)
	c.JSON(http.StatusOK, msg)
}

//func GetNextPPDetail(c *gin.Context) {
//	var details model.ImplementDetails
//	_ = c.ShouldBindJSON(&details)
//	var msg model.Response
//	details, err := dao.GetNextPPDetailDao(details)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"code":    400,
//			"message": "参数请求错误",
//		})
//		return
//	}
//	msg.Code = 200
//	msg.Message = "请求成功"
//	tmp, _ := json.Marshal(details)
//	msg.Data = string(tmp)
//	c.JSON(http.StatusOK, msg)
//}

func GetNextPPDetail(c *gin.Context) {
	var msg model.Response
	var details model.ImplementDetails
	_ = c.ShouldBindJSON(&details)
	utils.Try(func() {
		details = dao.GetNextPPDetailDao(details)
		msg.Code = 200
		msg.Message = "请求成功"
		tmp, _ := json.Marshal(details)
		msg.Data = string(tmp)
	}, func(err interface{}) {
		res, _ := err.(model.MyError)
		msg.Message = res.Message
		msg.Code = res.Code
	})
	c.JSON(http.StatusOK, msg)
}

func GetPrevPPDetail(c *gin.Context) {
	var msg model.Response
	var details model.ImplementDetails
	_ = c.ShouldBindJSON(&details)
	utils.Try(func() {
		details = dao.GetPrevPPDetailDao(details)
		msg.Code = 200
		msg.Message = "请求成功"
		tmp, _ := json.Marshal(details)
		msg.Data = string(tmp)
	}, func(err interface{}) {
		res, _ := err.(model.MyError)
		msg.Message = res.Message
		msg.Code = res.Code
	})
	c.JSON(http.StatusOK, msg)
}

func DelPPDetail(c *gin.Context) {
	var msg model.Response
	utils.Try(func() {
		id := c.Param("id")
		if id == "" {
			panic(errors.New("请求参数错误"))
		}
		details := dao.DelPPDetailDao(id)
		msg.Code = 200
		msg.Message = "请求成功"
		tmp, _ := json.Marshal(details)
		msg.Data = string(tmp)
	}, func(err interface{}) {
		res, _ := err.(model.MyError)
		msg.Message = res.Message
		msg.Code = res.Code
	})
	c.JSON(http.StatusOK, msg)
}

func UpdatePPDetail(c *gin.Context) {
	var msg model.Response
	var details model.ImplementDetails
	_ = c.ShouldBindJSON(&details)
	utils.Try(func() {
		dao.UpdateCurDetail(details, global.DB)
		msg.Code = 200
		msg.Message = "请求成功"
		tmp, _ := json.Marshal(details)
		msg.Data = string(tmp)
	}, func(err interface{}) {
		res, _ := err.(model.MyError)
		msg.Message = res.Message
		msg.Code = res.Code
	})
	c.JSON(http.StatusOK, msg)
}

func GetAllDetail(c *gin.Context) {
	var msg model.Response
	utils.Try(func() {
		id := c.Param("id")
		if id == "" {
			panic(errors.New("请求参数错误"))
		}
		details := dao.GetAllDetail(id)
		msg.Code = 200
		msg.Message = "请求成功"
		tmp, _ := json.Marshal(details)
		msg.Data = string(tmp)
	}, func(err interface{}) {
		res, _ := err.(model.MyError)
		msg.Message = res.Message
		msg.Code = res.Code
	})
	c.JSON(http.StatusOK, msg)
}
