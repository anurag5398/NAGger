package handlers

import (
	"NAGger/models/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckHandler struct {
	Config config.AppConfig `inject:"inline"`
}

// @Summary Health Check
// @Description Gives OK if server is up and inRotation
// @Tags Health
// @Produce json
// @Success 200 {object} string
// @Router /health/ [get]
func (self *HealthCheckHandler) HealthCheck(c *gin.Context) {
	if self.Config.Server.InRotation {
		c.JSON(http.StatusOK, "OK")
	} else {
		c.JSON(http.StatusServiceUnavailable, "OOR-OOPS")
	}
}
