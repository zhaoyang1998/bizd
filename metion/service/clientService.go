package service

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/satori/go.uuid"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func AddClient(c *gin.Context) {
	var client model.Client
	_ = c.ShouldBindJSON(&client)
	// 参数验证
	validate := validator.New()
	err := validate.Struct(client)
	if err != nil {
		log.Print(err.Error())
		c.JSON(200, gin.H{"code": 201, "msg": err.Error()})
		return
	}
	client.ClientId = uuid.NewV4().String()
	result := global.DB.Create(client)
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

func GetClients(c *gin.Context) {
	var client model.Client
	_ = c.ShouldBindJSON(&client)
	var clients []model.Client
	result := global.DB.Select("client_id,client_name,client_abbreviation,data_link,principal_id,status").Where(client).Find(&clients)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	var response model.Response
	data, _ := json.Marshal(clients)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func DelClient(c *gin.Context) {
	var response model.Response
	var client model.Client
	_ = c.ShouldBindJSON(&client)
	if client.ClientId == "" {
		response.Code = http.StatusCreated
		response.Message = "客户ID不能为空"
		c.JSON(http.StatusOK, response)
		return
	}
	result := global.DB.Delete(&client)
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

func UpdateClient(c *gin.Context) {
	var response model.Response
	var client model.Client
	_ = c.ShouldBindJSON(&client)
	result := global.DB.Model(&client).Updates(&client)
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
