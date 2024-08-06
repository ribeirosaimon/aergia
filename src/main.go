package main

import (
	"fmt"
	"os"

	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/logs"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia/internal/config"
)

func main() {
	logs.LOG.Message("Init file: " + os.Getenv(string(constants.AERGIA)))
	properties.NewPropertiesFile()
	config.NewAergiaServer(&config.AergiaConfig{
		ApiPort: fmt.Sprintf(":%s", properties.GetEnvironmentValue(constants.ApiPort)),
	})
}
