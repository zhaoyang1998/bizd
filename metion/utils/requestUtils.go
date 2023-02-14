package utils

import (
	"bizd/metion"
	"github.com/gin-gonic/gin"
)

func GetCurrentUserId(c *gin.Context) string {
	token := c.Request.Header.Get("token")
	claim, _ := metion.ParseToken(token)
	return claim.UserId
}
