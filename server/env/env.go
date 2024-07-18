package env

// Env contains all application singletons and lazy loads the dependancy tree

type Env struct {
	*loggerConfig
	*dbConfig
	*repositoriesConfig
	*servicesConfig
	*routerConfig
	*integrationsConfig
}

// Cleanup all environment resources that must perform cleanup actions on exit
// There is a cleaner way to do this by registering close functions lazily, but
// this works for now
func (e *Env) Cleanup() {

	if e.db != nil {
		e.db.Close()
	}
	if e.longRunningOperationsService != nil {
		e.longRunningOperationsService.Close()
	}
	if e.router != nil {
		e.router.Close()
	}
}

var app *Env

// NewEnv typically used to create different env based on config (prd, dev, stg).j
// For this project, will only create use the one type
func NewEnv() *Env {
	return app
}

func init() {

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
	routerConfig.init(&loggerConfig, &servicesConfig)
	integrationsConfig.init(&loggerConfig)
	servicesConfig.init(&loggerConfig, &repositoriesConfig, &integrationsConfig, &dbConfig)

	app = &Env{&loggerConfig, &dbConfig, &repositoriesConfig, &servicesConfig, &routerConfig, &integrationsConfig}
}
