package main

import (
	"NAGger/components"
	"NAGger/components/handlers"
	"NAGger/models/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	var err error
	var AppConfig config.AppConfig
	var bytes []byte
	var app interface{}

	if bytes, err = ioutil.ReadFile("D:\\Goland\\Projects\\Nagger\\components\\config.yaml"); err == nil {
		if err = yaml.Unmarshal(bytes, &AppConfig); err == nil {
			injector := components.NewNinjector(AppConfig)
			if app, err = injector.BuildApp(); err == nil {
				router := app.(*handlers.NagServer)
				err = router.Start()
			}
		}
	}

	if err != nil {
		panic(err)
	}
}
