package handler

import (
	"net/http"
	"strconv"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllLessonsTeacher(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	result, err := h.service.GetAllLessonsTeacher(userid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"lessons": result,
	})
}

func (h *Handler) CreateLesson(c *gin.Context) {
	var lesson models.Lesson
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	err = c.BindJSON(&lesson)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	result, err := h.service.CreateLesson(userId, lesson)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": result,
	})
}

func (h *Handler) DeleteLesson(c *gin.Context) {
	user, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	lesson_id, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	err = h.service.DeleteLesson(user, lesson_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
