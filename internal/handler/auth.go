package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var teacher models.Teacher
	err := c.BindJSON(&teacher)
	if err != nil {
		log.Fatalf("error with signing up %s", err.Error())
	}
	res, err := h.service.SignUp(teacher)
	if err != nil {
		log.Fatalf("errors while creating teacher %s", err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": res,
	})
}

func (h *Handler) SingIn(c *gin.Context) {
	var teacher models.SignInTeacher
	if err := c.BindJSON(&teacher); err != nil {
		log.Fatalf("error with json data %s", err.Error())
	}
	refresh, token, err := h.service.SignIn(teacher)
	if err != nil {
		log.Fatalf("error with data %s", err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":         token,
		"refresh_token": refresh,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	var refresh models.Refresh
	if err := c.BindJSON(&refresh); err != nil {
		log.Fatalf("error while binding refresh to json %s", err.Error())
	}
	fmt.Println(refresh.Token)

	token, err := h.service.RefreshToken(refresh.Token)
	if err != nil {
		log.Fatalf("error with creating refresh token %s", err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
