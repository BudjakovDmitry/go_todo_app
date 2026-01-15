package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) me(c *gin.Context) {
	id, _ := c.Get("userId")
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}
