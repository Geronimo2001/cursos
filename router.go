package main

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

	return r
}
