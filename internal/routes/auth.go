package routes

import (
	"movie-reservation-system/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.SignUp)
	r.GET("/home", controllers.Home)
	r.GET("/admin", controllers.Admin)
	r.GET("/signout", controllers.SignOut)
}
