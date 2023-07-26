package bookRoutes

import (
	controllers "api/book/controller"

	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.ShowBooks)
	router.GET("/:id", controllers.ShowBook)
	router.POST("/", controllers.CreateBook)
	router.PUT("/", controllers.UpdateBook)
	router.DELETE("/:id", controllers.DeleteBook)
}
