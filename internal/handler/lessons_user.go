package handler

import (
	"net/http"
	"strconv"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllLessonsUser(c *gin.Context) {
	user_id, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	result, err := h.service.GetAllLessonsuser(user_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) GetLessonUser(c *gin.Context) {
	user_id, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	lesson_id, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	result, err := h.service.GetLessonUser(user_id, lesson_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) SolveHomework(c *gin.Context) {
	var hw models.HomeworkUser
	user_id, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	err = c.BindJSON(&hw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	lesson_id, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	err = h.service.SolveHomework(user_id, lesson_id, hw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
