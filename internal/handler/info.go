package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Info(c *gin.Context) {
	result, err := h.service.Info()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}
