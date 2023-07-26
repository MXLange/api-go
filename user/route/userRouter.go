package userRoutes

import (
	userController "api/user/controller"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup) {
	router.POST("/", userController.CreateUser)
	// router.GET("/", controllers.ShowBooks)
	// router.GET("/:id", controllers.ShowBook)
	// router.PUT("/", controllers.UpdateBook)
	// router.DELETE("/:id", controllers.DeleteBook)
}
