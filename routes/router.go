package routes

import (
	bookRoutes "api/book/route"
	userRoutes "api/user/route"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")

	bookRoute := main.Group("/book")
	bookRoutes.SetupBookRoutes(bookRoute)
	userRoute := main.Group("/user")
	userRoutes.SetupUserRoutes(userRoute)
	return router
}
