package userController

import (
	"api/database"
	"api/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	db := database.GetDB()
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't bind JSON. " + err.Error(),
		})
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't create user. " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}
