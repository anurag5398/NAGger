package components

import (
	"NAGger/components/handlers"
	"NAGger/models/config"
	"fmt"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AppInjector interface {
	BuildApp() (app interface{}, err error)
}

type Ninjector struct {
	graph  inject.Graph
	config config.AppConfig
}

func (self *Ninjector) BuildApp() (app interface{}, err error) {
	router := handlers.NagServer{}
	app = &router

	engine := gin.Default()

	db := initDB(self.config.DB)

	err = self.graph.Provide(
		&inject.Object{Value: engine},
		&inject.Object{Value: &router},
		&inject.Object{Value: &http.Server{
			Addr:    fmt.Sprintf(":%v", self.config.Server.Port),
			Handler: engine,
		}},
		//Configs
		&inject.Object{Value: self.config.Server, Name: "serverConfig"},
		&inject.Object{Value: self.config.Database, Name: "databaseConfig"},

		//Handler
		&inject.Object{Value: &handlers.HealthCheckHandler{}},
	)

	if err == nil {
		err = self.graph.Populate()
	}

	return
}

func initDB(db config.DB) (db *gorm.DB) {
	var err error

}

func NewNinjector(config config.AppConfig) AppInjector {
	return &Ninjector{
		graph:  inject.Graph{},
		config: config,
	}
}
