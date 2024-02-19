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

	result, err := h.service.SignUpStudent(user, h.cfg)

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
	token, refer, err := h.service.SignInStudent(userLog)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"errors": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  token,
		"refresh_token": refer,
	})
}

func (h *Handler) RefreshUser(c *gin.Context) {
	var refresh models.Refresh
	err := c.BindJSON(&refresh)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	access, err := h.service.RefreshUser(refresh)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": access,
	})
}

func (h *Handler) CheckToken(c *gin.Context) {
	token := c.Param("email_token")
	err := h.service.CheckToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "your email is verified",
	})
}
