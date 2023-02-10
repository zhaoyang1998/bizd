package service

import (
	"bizd/metion/dao"
	"bizd/metion/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHomePageData(c *gin.Context) {
	clientId := c.Param("clientId")
	var msg model.Response
	data, err := dao.GetTotalDataByClientDao(clientId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    400,
			"message": "参数请求错误",
		})
		return
	}
	msg.Code = 200
	msg.Message = "请求成功"
	tmp, _ := json.Marshal(data)
	msg.Data = string(tmp)
	c.JSON(http.StatusOK, msg)
}
