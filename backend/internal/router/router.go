package router

import (
	"building-ac-3d/backend/internal/config"
	"building-ac-3d/backend/internal/handler"
	"building-ac-3d/backend/internal/middleware"
	"building-ac-3d/backend/internal/repository"
	"building-ac-3d/backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Options struct {
	Config *config.Config
	DB     *gorm.DB
}

func New(opts Options) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery(), middleware.CORS(opts.Config.CORS), middleware.Auth())

	tenantRepo := repository.NewTenantRepository(opts.DB)
	userRepo := repository.NewUserRepository(opts.DB)
	roleRepo := repository.NewRoleRepository(opts.DB)
	menuRepo := repository.NewMenuRepository(opts.DB)

	authHandler := handler.NewAuthHandler(service.NewAuthService(tenantRepo, userRepo))
	healthHandler := handler.NewHealthHandler(opts.DB)
	tenantHandler := handler.NewTenantHandler(service.NewTenantService(tenantRepo))
	userHandler := handler.NewUserHandler(service.NewUserService(userRepo))
	roleHandler := handler.NewRoleHandler(service.NewRoleService(roleRepo))
	menuHandler := handler.NewMenuHandler(service.NewMenuService(menuRepo))

	api := engine.Group("/api")
	{
		api.GET("/db/health", healthHandler.DBHealth)

		api.POST("/auth/login", authHandler.Login)
		api.GET("/auth/me", authHandler.Me)
		api.POST("/auth/logout", authHandler.Logout)

		system := api.Group("/system")
		{
			system.GET("/tenants", tenantHandler.List)
			system.POST("/tenants", tenantHandler.Create)
			system.PUT("/tenants", tenantHandler.Update)
			system.DELETE("/tenants", tenantHandler.Delete)

			system.GET("/users", userHandler.List)
			system.POST("/users", userHandler.Create)
			system.PUT("/users", userHandler.Update)
			system.DELETE("/users", userHandler.Delete)

			system.GET("/roles", roleHandler.List)
			system.POST("/roles", roleHandler.Create)
			system.PUT("/roles", roleHandler.Update)
			system.DELETE("/roles", roleHandler.Delete)

			system.GET("/menus", menuHandler.List)
			system.POST("/menus", menuHandler.Create)
			system.PUT("/menus", menuHandler.Update)
			system.DELETE("/menus", menuHandler.Delete)
		}
	}

	return engine
}
