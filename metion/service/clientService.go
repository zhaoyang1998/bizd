package service

import (
	"bizd/metion/db"
	"bizd/metion/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/satori/go.uuid"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func AddClient(c *gin.Context) {
	var client model.Client
	c.Bind(&client)
	// 参数验证
	validate := validator.New()
	err := validate.Struct(client)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"code": 201, "msg": err.Error()})
		return
	}
	client.ClientId = uuid.NewV4().String()
	result := db.DB.Create(client)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	var response model.Response
	response.Code = http.StatusOK
	response.Message = "客户添加成功"
	c.JSON(http.StatusOK, response)
}

func GetClients(c *gin.Context) {
	var client model.Client
	c.Bind(&client)
	var clients []model.Client
	result := db.DB.Select("client_id,client_name,client_abbreviation,data_link,principal_id,status").Where(client).Find(&clients)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	var response model.Response
	data, _ := json.Marshal(clients)
	response.Data = string(data)
	response.Code = http.StatusOK
	c.JSON(http.StatusOK, response)
}

func DelClient(c *gin.Context) {
	var response model.Response
	var client model.Client
	c.Bind(&client)
	if client.ClientId == "" {
		response.Code = http.StatusCreated
		response.Message = "客户ID不能为空"
		c.JSON(http.StatusOK, response)
		return
	}
	result := db.DB.Delete(&client)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	response.Code = http.StatusOK
	response.Message = "客户删除成功"
	c.JSON(http.StatusOK, response)
}

func UpdateClient(c *gin.Context) {
	var response model.Response
	var client model.Client
	c.Bind(&client)
	result := db.DB.Model(&client).Updates(&client)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	response.Code = http.StatusOK
	response.Message = "客户信息更新成功"
	c.JSON(http.StatusOK, response)
}
