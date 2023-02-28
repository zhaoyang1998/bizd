package service

import (
	"bizd/metion/dao"
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
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
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	client.ClientId = uuid.NewV4().String()
	client.PrincipalId = utils.GetCurrentUserId(c)
	result := global.DB.Create(&client)
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

func GetClients(c *gin.Context) {
	var client model.Client
	_ = c.ShouldBindJSON(&client)
	var clients []model.Client
	var err error
	var pagination model.ResponsePagination
	pagination, clients, err = dao.GetClientsDao(client)
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	var response model.Response
	data, _ := json.Marshal(clients)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	response.ResponsePagination = pagination
	c.JSON(http.StatusOK, response)
}

func GetAllClients(c *gin.Context) {
	var clients []model.Client
	var err error
	clients, err = dao.GetAllClientsDao()
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	var response model.Response
	data, _ := json.Marshal(clients)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func GetClientsByKeyword(c *gin.Context) {
	var search model.Search
	_ = c.ShouldBindJSON(&search)
	var clients []model.Client
	var err error
	var pagination model.ResponsePagination
	pagination, clients, err = dao.GetClientsByKeywordDao(search)
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	var response model.Response
	data, _ := json.Marshal(clients)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	response.ResponsePagination = pagination
	c.JSON(http.StatusOK, response)
}

func DelClient(c *gin.Context) {
	var response model.Response
	var del = model.DelModel{}
	var err error
	_ = c.ShouldBindJSON(&del)
	if len(del.Keys) == 0 {
		response.Code = http.StatusCreated
		response.Message = "客户ID不能为空"
		c.JSON(http.StatusOK, response)
		return
	}
	err = dao.DelClientByKeys(del)
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
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
		c.JSON(200, gin.H{"code": 400, "message": result.Error})
		return
	}
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}
