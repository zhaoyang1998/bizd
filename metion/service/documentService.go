package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type FormData struct {
	Content string `json:"content"`
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

	c.JSON(http.StatusOK, gin.H{"message": "保存成功"})
}
