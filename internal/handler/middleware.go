package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	userHeader = "userId"
)

func (h *Handler) CheckIdentity(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	headerParts := strings.Split(auth, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid header"})
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Set(userHeader, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userHeader)
	if !ok {
		return 0, errors.New("error while geting auth header")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("error while converting header to int")
	}
	return idInt, nil
}

func (h *Handler) CheckIdentityUser(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty token"})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	spl := strings.Split(header, " ")
	if len(spl) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	userid, err := h.service.ParseTokenUser(spl[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Set(userHeader, userid)
}

func (h *Handler) CheckHealth(c *gin.Context) {
	user_id, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	result := h.service.CheckHealth(user_id)
	if result <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "your health <0, u cant study here"})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
