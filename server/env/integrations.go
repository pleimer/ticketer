package env

import (
	"github.com/pleimer/ticketer/server/integration/nylas"
	"github.com/pleimer/ticketer/server/lib/once"
)

type integrationsConfig struct {
	NylasClientConfig nylas.NylasClientConfig

	nylasClient *nylas.NylasClient
	NylasClient func() *nylas.NylasClient
}

func (i *integrationsConfig) init(loggerConfig *loggerConfig) {
	i.NylasClient = func() *nylas.NylasClient {
		once.Once(func() {
			i.nylasClient = nylas.NewNylasClient(i.NylasClientConfig)
		})
		return i.nylasClient
	}

}
