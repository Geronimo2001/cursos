package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myapp/controllers"
)

func SetpRouterUser(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userController := controllers.NewUserController(db)

	r.GET("/users", userController.GetUsers)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)

	return r

}
