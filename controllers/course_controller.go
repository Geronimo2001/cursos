package controllers

import (
	"myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CourseController struct {
	DB *gorm.DB
}

func NewCourseController(db *gorm.DB) *CourseController {
	return &CourseController{DB: db}
}

func (ctrl *CourseController) GetCourses(c *gin.Context) {
	var courses []models.Course
	ctrl.DB.Find(&courses)
	c.JSON(http.StatusOK, courses)
}

func (ctrl *CourseController) CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.DB.Create(&course)
	c.JSON(http.StatusOK, course)
}
