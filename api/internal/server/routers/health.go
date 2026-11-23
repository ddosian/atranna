package routers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthResponse struct {
	Status    string `json:"status" example:"ok"`
	DBVersion string `json:"db_version" example:"3.46.0"`
}

type healthUnavailableResponse struct {
	Status string `json:"status" example:"database unavailable"`
}

// GetHealth godoc
// @Summary      Health check
// @Description  Reports whether the API is up and the database version
// @Tags         health
// @Produce      json
// @Success      200  {object}  healthResponse
// @Failure      503  {object}  healthUnavailableResponse
// @Router       /health [get]
func GetHealth(conn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var version string
		err := conn.QueryRowContext(c.Request.Context(), "SELECT sqlite_version()").Scan(&version)

		if err != nil {
			c.IndentedJSON(http.StatusServiceUnavailable, healthUnavailableResponse{Status: "database unavailable"})
			return
		}
		c.IndentedJSON(http.StatusOK, healthResponse{Status: "ok", DBVersion: version})
	}
}
