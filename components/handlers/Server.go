package handlers

import (
	_ "NAGger/components/docs"
	"fmt"
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
	RecordHandler      RecordHandler      `inject:"inline"`
}

// @title NAGger - Swagger
// @version 1.0
// @description This is the swagger API definition doc
// @contact.email anurags2909@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1

func (self *NagServer) initRoutes() {
	fmt.Println(self.Ginengine, self.Server)
	v1 := self.Ginengine.Group("/v1")

	healthGroup := v1.Group("/health")
	healthGroup.GET("/", self.HealthCheckHandler.HealthCheck)

	recordGroup := v1.Group("/record")
	recordGroup.POST("/create", self.RecordHandler.CreateRecord)

	self.Ginengine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (self *NagServer) Start() (err error) {
	self.initRoutes()

	self.Server.ListenAndServe()

	return
}
