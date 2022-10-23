package controllers

import (
	"net/http"

	"github.com/S3B4SZ17/Email_service/services"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {

	token := services.ExtractToken(c)
	user, _ := services.ExtractUser(&token)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
