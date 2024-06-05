package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myapp/controllers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	courseController := controllers.NewCourseController(db)

	r.GET("/courses", courseController.GetCourses)
	r.POST("/courses", courseController.CreateCourse)
	r.PUT("/courses/:id", courseController.UpdateCourse)
	r.DELETE("/courses/:id", courseController.DeleteCourse)
	r.GET("/courses/name/:name", courseController.FindCourseByName)

	return r
}
