package app

import (
	"github.com/pleimer/ticketer/server/lib/once"
)

// Env contains all application singletons and lazy loads the dependancy tree
var App func() *Env

type Env struct {
	*loggerConfig
	*dbConfig
	*repositoriesConfig
	*servicesConfig
	*routerConfig
	*integrationsConfig
}

func init() {
	var e *Env

	loggerConfig := loggerConfig{}
	dbConfig := dbConfig{}
	repositoriesConfig := repositoriesConfig{}
	servicesConfig := servicesConfig{}
	routerConfig := routerConfig{}
	integrationsConfig := integrationsConfig{}

	// setup the singleton dependancy tree
	loggerConfig.init()
	dbConfig.init(&loggerConfig)
	repositoriesConfig.init(&dbConfig)
	servicesConfig.init(&loggerConfig, &repositoriesConfig)
	routerConfig.init(&loggerConfig, &servicesConfig)
	integrationsConfig.init(&loggerConfig)

	App = func() *Env {
		once.Once(func() {
			e = &Env{&loggerConfig, &dbConfig, &repositoriesConfig, &servicesConfig, &routerConfig, &integrationsConfig}
		})
		return e
	}
}
