package handler

import (
	"net/http"

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
