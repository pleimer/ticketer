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
	*routerConfig
}

func init() {
	var a *app

	loggerConfig := loggerConfig{}
	dbConfig := dbConfig{}
	repositoriesConfig := repositoriesConfig{}
	routerConfig := routerConfig{}

	// setup the singleton dependancy tree
	loggerConfig.init()
	dbConfig.init(&loggerConfig)
	repositoriesConfig.init(&dbConfig)
	routerConfig.init(&loggerConfig)

	App = func() *app {
		once.Once(func() {
			a = &app{&loggerConfig, &dbConfig, &repositoriesConfig, &routerConfig}
		})
		return a
	}
}
