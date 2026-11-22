package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthResponse struct {
	Status string `json:"status" example:"ok"`
}

// GetHealth godoc
// @Summary      Health check
// @Description  Reports whether the API is up
// @Tags         health
// @Produce      json
// @Success      200  {object}  healthResponse
// @Router       /health [get]
func GetHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, healthResponse{Status: "ok"})
}
