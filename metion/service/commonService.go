package service

import (
	"bizd/metion/dao"
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sort"
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
	efficiencyData, err := dao.GetEfficiencyDataDao(clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": "参数请求错误",
		})
		return
	}
	if efficiencyData.Series == nil {
		echarts.NullEcharts = utils.GetNullLine()
		echarts.EfficiencyData.Flag = true
	} else {
		echarts.EfficiencyData = efficiencyData
	}
	msg.Code = 200
	msg.Message = "请求成功"
	tmp, _ := json.Marshal(echarts)
	msg.Data = string(tmp)
	c.JSON(http.StatusOK, msg)
}

func GetMenuJson(c *gin.Context) {
	var response model.Response
	dataByte, err := ioutil.ReadFile("./static/menu.json")
	if err != nil {
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

func GetAllStatus(c *gin.Context) {
	var msg model.Response
	utils.Try(func() {
		var statusList []model.StatusText
		for key, value := range global.ClientAndPointPositionStatusText {
			statusList = append(statusList, model.StatusText{Name: value, Value: int(key)})
		}
		sort.Slice(statusList, func(i, j int) bool {
			return statusList[i].Value < statusList[j].Value
		})
		msg.Code = 200
		msg.Message = "请求成功"
		tmp, _ := json.Marshal(statusList)
		msg.Data = string(tmp)
	}, func(err interface{}) {
		res, _ := err.(model.MyError)
		msg.Message = res.Message
		msg.Code = res.Code
	})
	c.JSON(http.StatusOK, msg)
}
