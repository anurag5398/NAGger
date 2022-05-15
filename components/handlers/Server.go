package handlers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

type NagServer struct {
	Ginengine *gin.Engine  `inject:""`
	Server    *http.Server `inject:""`

	/*Handlers*/
	HealthCheckHandler HealthCheckHandler `inject:"inline"`
}

// @title NAGger - Swagger
// @contact.email anurags2909@gmail.com

// @BasePath /v1

func (self *NagServer) initRoutes() {
	v1 := self.Ginengine.Group("/v1")

	healthGroup := v1.Group("/health")
	healthGroup.GET("/", self.HealthCheckHandler.HealthCheck)

	self.Ginengine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (self *NagServer) Start() (err error) {
	self.initRoutes()
	return
}
