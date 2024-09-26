package config

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/aergia-utils/constants"
	"github.com/ribeirosaimon/aergia-utils/domain/entities/sql"
	"github.com/ribeirosaimon/aergia-utils/properties"
	"github.com/ribeirosaimon/aergia/internal/controller"
)

type AergiaServer struct {
	ginEngine *gin.Engine
	config    *AergiaConfig
}

type AergiaConfig struct {
	ApiPort string
}

func NewAergiaServer(config *AergiaConfig) *AergiaServer {
	engine := gin.New()
	server := &AergiaServer{ginEngine: engine}
	server.config = config
	controller.StartControllers()
	controller.AddController(server.ginEngine)

	sql.CreateTableDatabase(properties.GetEnvironmentValue(constants.PostgressUrl))

	server.startServer()
	return server
}

func (a *AergiaServer) startServer() {
	a.ginEngine.Run(a.config.ApiPort)
}
