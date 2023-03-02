package service

import (
	"bizd/metion/dao"
	"bizd/metion/model"
	"bizd/metion/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

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

func SaveDetail(c *gin.Context) {
	var msg model.Response
	var details []model.ImplementDetails
	_ = c.ShouldBindJSON(&details)
	utils.Try(func() {
		dao.SaveDetailDao(details)
		msg.Code = 200
		msg.Message = "请求成功"
		tmp, _ := json.Marshal(details)
		msg.Data = string(tmp)
	}, func(err interface{}) {
		fmt.Println(err)
		res, _ := err.(model.MyError)
		msg.Message = res.Message
		msg.Code = res.Code
	})
	c.JSON(http.StatusOK, msg)
}
