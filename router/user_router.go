package router

import (
	"myapp/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {

	userController := controllers.NewUserController(db)

	r.GET("/users", userController.GetUsers)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	return r

}
