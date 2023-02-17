package service

import (
	"bizd/metion/dao"
	"bizd/metion/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func GetHomePageData(c *gin.Context) {
	clientId := c.Param("clientId")
	var msg model.Response
	var echarts model.Echarts
	totalData, err := dao.GetTotalDataByClientDao(clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": "参数请求错误",
		})
		return
	}
	echarts.TotalData = totalData
	curData, err := dao.GetCurDataByClientDao(clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": "参数请求错误",
		})
		return
	}
	echarts.CurData = curData
	msg.Code = 200
	msg.Message = "请求成功"
	tmp, _ := json.Marshal(echarts)
	msg.Data = string(tmp)
	c.JSON(http.StatusOK, msg)
}
func GetMenuJson(c *gin.Context) {
	var response model.Response
	dataByte, err := ioutil.ReadFile("./metion/static/menu.json")
	if err != nil {
		log.Print("读取menu.json文件失败")
		c.JSON(200, gin.H{"code": 500, "message": "文件读取失败"})
	}
	var menus model.MenuJson
	err = json.Unmarshal(dataByte, &menus)
	if err == nil {
		response.Code = 0
		response.Message = "请求成功"
		tmp, _ := json.Marshal(menus)
		response.Data = string(tmp)
	} else {
		response.Code = 500
		response.Message = "文件错误"
	}
	c.JSON(http.StatusOK, response)
}
