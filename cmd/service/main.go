package main

import (
	"github.com/hse-revizor/auth-service/internal/di"
	"github.com/hse-revizor/auth-service/internal/utils/config"
	"github.com/hse-revizor/auth-service/internal/utils/flags"
	"github.com/slipneff/gogger"
	"github.com/slipneff/gogger/log"
)

// @title           auth Service API
// @version         1.0
// @description     This is a Swagger documentation.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1
func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	container := di.New(cfg)
	gogger.ConfigureZeroLogger()
	log.Info("Service started")
	container.GetAuthService()
}
