package handler

import (
	"net/http"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUpUser(c *gin.Context) {
	var user models.Student

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	result, err := h.service.SignUpStudent(user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"id": result,
	})
}

func (h *Handler) SignInUser(c *gin.Context) {
	var userLog models.SignInStudent
	err := c.BindJSON(&userLog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"errors": err.Error(),
		})
	}
	token, err := h.service.SignInStudent(userLog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"errors": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": token,
	})
}
