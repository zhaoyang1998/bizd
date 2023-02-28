package service

import (
	"bizd/metion"
	"bizd/metion/dao"
	"bizd/metion/global"
	"bizd/metion/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

func GetUserByType(c *gin.Context) {
	var response model.Response
	userType := c.Query("type")
	if userType == "" {
		response.Code = http.StatusCreated
		response.Message = "用户类型不能为空"
		c.JSON(http.StatusOK, response)
		return
	}
	var users []model.User
	result := global.DB.Select("user_id,user_name").Where("type = ?", userType).Find(&users)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	data, _ := json.Marshal(users)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func GetUsers(c *gin.Context) {
	var user model.User
	var response model.Response
	var users []model.User
	var err error
	var pagination model.ResponsePagination
	_ = c.ShouldBindJSON(&user)
	pagination, users, err = dao.GetUsersDao(user)
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	data, _ := json.Marshal(users)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	response.ResponsePagination = pagination
	c.JSON(http.StatusOK, response)
}

func GetUsersByKeyword(c *gin.Context) {
	var search model.Search
	var response model.Response
	var users []model.User
	var err error
	var pagination model.ResponsePagination
	_ = c.ShouldBindJSON(&search)
	pagination, users, err = dao.GetUsersByKeywordDao(search)
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	data, _ := json.Marshal(users)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	response.ResponsePagination = pagination
	c.JSON(http.StatusOK, response)
}

func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	// 参数验证
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	user.UserId = uuid.NewV4().String()
	result := global.DB.Create(&user)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": result.Error,
		})
		return
	}
	var response model.Response
	response.Message = "请求成功"
	response.Code = http.StatusOK
	c.JSON(http.StatusOK, response)
}

func DelUser(c *gin.Context) {
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
	err = dao.DelUserByKeys(del)
	if err != nil {
		log.Print(err)
		c.JSON(200, gin.H{"code": 400, "message": err.Error()})
		return
	}
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func UpdateUser(c *gin.Context) {
	var response model.Response
	var user model.User
	_ = c.ShouldBindJSON(&user)
	result := global.DB.Model(&user).Updates(&user)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(200, gin.H{"code": 400, "message": result.Error})
		return
	}
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func Login(c *gin.Context) {
	var user model.User
	var err error
	_ = c.ShouldBindJSON(&user)
	user, err = dao.GetUserDao(user.UserAccount, user.UserPwd)
	msg := metion.ResponseMsg{
		Code:    200,
		Message: "成功！",
		Status:  "success",
	}
	if err != nil {
		msg.Code = 600
		msg.Message = "账号或密码错误"
		msg.Status = "fail"
	} else {
		token, _ := metion.CreateToken(user.UserId)
		msg.Data.Token = token
		msg.Message = "成功"
	}
	c.JSON(http.StatusOK, msg)
}
