package handlers

import (
	"NAGger/models/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckHandler struct {
	Config config.Server `inject:"serverConfig"`
}

// @Summary Health Check
// @Description Gives OK if server is up and inRotation
// @Tags Health
// @Produce json
// @Success 200 {object} string
// @Router /v1/health/ [get]
func (self *HealthCheckHandler) HealthCheck(c *gin.Context) {
	fmt.Println(self.Config)
	if self.Config.InRotation {
		c.JSON(http.StatusOK, "OK")
	} else {
		c.JSON(http.StatusServiceUnavailable, "OOR-OOPS")
	}
}
