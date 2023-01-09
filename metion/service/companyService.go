package service

import (
	"bizd/metion/db"
	"bizd/metion/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func AddCompany(c *gin.Context) {
	var company model.Company
	_ = c.Bind(&company)
	// 参数验证
	validate := validator.New()
	err := validate.Struct(company)
	if err != nil {
		log.Print(err.Error())
		c.JSON(200, gin.H{"code": 201, "msg": err.Error()})
		return
	}
	company.CompanyId = uuid.NewV4().String()
	company.Status = company.Type * 10
	result := db.DB.Create(company)
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

func GetCompanies(c *gin.Context) {
	var company model.Company
	_ = c.Bind(&company)
	var companies []model.Company
	result := db.DB.Offset((company.PageNumber - 1) * company.PageSize).Limit(company.PageSize).Where(company).Find(&companies)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}
	var response model.Response
	data, _ := json.Marshal(companies)
	response.Data = string(data)
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}

func DelCompany(c *gin.Context) {
	var response model.Response
	var company model.Company
	_ = c.Bind(&company)
	if company.CompanyId == "" {
		response.Code = http.StatusCreated
		response.Message = "单位ID不能为空"
		c.JSON(http.StatusOK, response)
		return
	}
	result := db.DB.Delete(&company)
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

func UpdateCompany(c *gin.Context) {
	var response model.Response
	var company model.Company
	_ = c.Bind(&company)
	result := db.DB.Model(&company).Updates(&company)
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

func StartAssignment(c *gin.Context) {
	var response model.Response
	var company model.Company
	company.CompanyId = c.Query("companyId")
	if company.CompanyId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数不能为空",
		})
	}
	result := db.DB.Model(&company).Where(gorm.Expr("status % 10 = 0")).Updates(map[string]interface{}{"status": gorm.Expr("status + 1"), "start_time": time.Now()})
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

func FinishAssignment(c *gin.Context) {
	var response model.Response
	var company model.Company
	company.CompanyId = c.Query("companyId")
	if company.CompanyId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数不能为空",
		})
	}
	result := db.DB.Model(&company).Where(gorm.Expr("status % 10 = 1")).Updates(map[string]interface{}{"status": gorm.Expr("status + 1"), "end_time": time.Now()})
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

func AllocatingAssignment(c *gin.Context) {
	var response model.Response
	var company model.Company
	_ = c.Bind(&company)
	result := db.DB.Model(&company).Updates(&company)
	if result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": result.Error,
		})
		return
	}
	response.Code = http.StatusOK
	response.Message = "请求成功"
	c.JSON(http.StatusOK, response)
}
