package main

import (
	"fmt"

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

// @host     localhost:8383
// @BasePath /api
func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode)
	container := di.New(cfg)
	gogger.ConfigureZeroLogger()
	log.Info(fmt.Sprintf("Server starting on %s:%d", cfg.Host, cfg.Port))

	err := container.GetHttpServer().ListenAndServe()
	if err != nil {
		log.Panic(err, "Fail serve HTTP")
	}
}
