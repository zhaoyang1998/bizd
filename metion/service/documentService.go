package service

import (
	"bizd/metion/global"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type FormData struct {
	Content  string `json:"content"`
	Title    string `json:"title"`
	Customer string `json:"customer"`
}

func SaveDocumentFile(c *gin.Context) {
	var formData FormData
	err := c.BindJSON(&formData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 在这里可以将 formData.Content 和 formData.TableData 保存到数据库或者文件中
	// ...
	err = ioutil.WriteFile("test.doc", []byte(formData.Content), 0644) // 写入文件
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = postFileToRedis(formData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": "保存成功"})
}

//redis key的命名方式采用 客户名/实施点位名

//把文件内容放在redis中
func postFileToRedis(formData FormData) error {
	err := global.RedisCli.Set(context.Background(), formData.Customer+"/"+formData.Title, formData.Content, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

//把初始化界面把redis内容读出返回前端  获取doc内容
func GetDoc(c *gin.Context) {
	var formData FormData
	err := c.BindJSON(&formData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	get := global.RedisCli.Get(context.Background(), formData.Customer+"/"+formData.Title).Val()

	fmt.Println(formData.Customer + "/" + formData.Title)
	fmt.Println(get)
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": get})

}
