package bookController

import (
	"api/database"
	"api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	var book models.Book

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDB()
	err = db.First(&book, id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't find book with id " + strconv.Itoa(id) + ". " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	db := database.GetDB()
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't bind JSON. " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't create book. " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	var books []models.Book

	db := database.GetDB()

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't find books. " + err.Error(),
		})
		return
	}
	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't bind JSON",
		})
		return
	}

	db := database.GetDB()
	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't update book.",
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDB()
	err = db.Delete(&models.Book{}, id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Can't find book with id " + strconv.Itoa(id) + ". " + err.Error(),
		})
		return
	}
	c.Status(204)
}
