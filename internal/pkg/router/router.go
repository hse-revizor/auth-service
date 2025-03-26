package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/hse-revizor/auth-service/docs"
	"github.com/hse-revizor/auth-service/internal/pkg/service/auth"
	"github.com/hse-revizor/auth-service/internal/utils/config"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	cfg  *config.Config
	auth *AuthHandler
}

func NewRouter(cfg *config.Config, authService *auth.Service) *Handler {
	return &Handler{
		cfg:  cfg,
		auth: NewAuthHandler(cfg, authService),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	api := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"
	router := api.Group("/api/v1")
	{
		router.GET("/", h.auth.HandleHome)
		router.GET("/login", h.auth.HandleLogin)
		router.GET("/auth/github/callback", h.auth.HandleCallback)
	}
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return api
}
