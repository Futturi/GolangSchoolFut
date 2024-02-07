package handler

import (
	"fmt"
	"net/http"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var teacher models.Teacher
	err := c.BindJSON(&teacher)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	res, err := h.service.SignUp(teacher)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": res,
	})
}

func (h *Handler) SingIn(c *gin.Context) {
	var teacher models.SignInTeacher
	if err := c.BindJSON(&teacher); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	refresh, token, err := h.service.SignIn(teacher)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":         token,
		"refresh_token": refresh,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	var refresh models.Refresh
	if err := c.BindJSON(&refresh); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	fmt.Println(refresh.Token)

	token, err := h.service.RefreshToken(refresh.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
