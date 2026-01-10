package traefik

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/traefik"
	"github.com/yahahaff/rapide/internal/service"
)

// TraefikRouter 注册Traefik相关路由
func TraefikRouter(Router *gin.RouterGroup) {
	traefikGroup := Router.Group("/traefik")
	{
		tc := new(traefik.TraefikController)
		// 获取Traefik路由信息
		traefikGroup.GET("/routes", tc.GetRoutes)
		// 获取Traefik路由详情
		traefikGroup.GET("/routers/:name", tc.GetRouteDetail)
		// 获取Traefik中间件信息
		traefikGroup.GET("/middlewares", tc.GetMiddlewares)
		// 获取Traefik中间件详情
		traefikGroup.GET("/middlewares/:name", tc.GetMiddlewareDetail)
		// 获取Traefik服务信息
		traefikGroup.GET("/services", tc.GetServices)
		// 获取Traefik服务详情
		traefikGroup.GET("/services/:name", tc.GetServiceDetail)
		// 获取Traefik概览信息
		traefikGroup.GET("/overview", tc.GetOverview)
	}
}

// TraefikHTTPProviderRouter 注册Traefik HTTP自动发现路由，不需要认证
func TraefikHTTPProviderRouter(engine *gin.Engine) {
	// Traefik HTTP Provider配置路由，公开访问，不需要认证
	engine.GET("/api/traefik/provider", func(c *gin.Context) {
		// 从服务层获取配置
		config, err := service.Entrance.TraefikService.TraefikHTTPProviderService.GetHTTPProviderConfig()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// 直接返回配置，不添加任何包装，符合Traefik HTTP Provider期望的格式
		c.JSON(200, config)
	})
}
