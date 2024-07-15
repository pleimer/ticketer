package app

import (
	"github.com/pleimer/ticketer/server/lib/once"
)

// App contains all application singletons and lazy loads the dependancy tree
var App func() *app

type app struct {
	*loggerConfig
	*dbConfig
	*repositoriesConfig
	*servicesConfig
	*routerConfig
}

func init() {
	var a *app

	loggerConfig := loggerConfig{}
	dbConfig := dbConfig{}
	repositoriesConfig := repositoriesConfig{}
	servicesConfig := servicesConfig{}
	routerConfig := routerConfig{}

	// setup the singleton dependancy tree
	loggerConfig.init()
	dbConfig.init(&loggerConfig)
	repositoriesConfig.init(&dbConfig)
	servicesConfig.init(&loggerConfig, &repositoriesConfig)
	routerConfig.init(&loggerConfig, &servicesConfig)

	App = func() *app {
		once.Once(func() {
			a = &app{&loggerConfig, &dbConfig, &repositoriesConfig, &servicesConfig, &routerConfig}
		})
		return a
	}
}
