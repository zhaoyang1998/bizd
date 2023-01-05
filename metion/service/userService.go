package service

import (
	"bizd/metion/db"
	"bizd/metion/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var response model.Response
	userType := c.Query("type")
	if userType == "" {
		response.Code = http.StatusCreated
		response.Message = "用户类型不能为空"
		c.JSON(http.StatusOK, response)
		return
	}
	var users []model.User
	result := db.DB.Select("user_id,user_name").Where("type = ?", userType).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	data, _ := json.Marshal(users)
	response.Data = string(data)
	response.Code = http.StatusOK
	c.JSON(http.StatusOK, response)
}

func AddUser(c *gin.Context) {
	var user model.User
	c.Bind(&user)
	// 参数验证
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"code": 201, "msg": err.Error()})
		return
	}
	user.UserId = uuid.NewV4().String()
	result := db.DB.Create(user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	var response model.Response
	response.Message = "用户添加成功"
	response.Code = http.StatusOK
	c.JSON(http.StatusOK, response)
}

func DelUser(c *gin.Context) {
	var response model.Response
	var user model.User
	c.Bind(&user)
	if user.UserId == "" {
		response.Code = http.StatusCreated
		response.Message = "用户ID不能为空"
		c.JSON(http.StatusOK, response)
		return
	}
	result := db.DB.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	response.Code = http.StatusOK
	response.Message = "用户删除成功"
	c.JSON(http.StatusOK, response)
}
func UpdateUser(c *gin.Context) {
	var response model.Response
	var user model.User
	c.Bind(&user)
	result := db.DB.Model(&user).Updates(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	response.Code = http.StatusOK
	response.Message = "用户信息更新成功"
	c.JSON(http.StatusOK, response)
}
