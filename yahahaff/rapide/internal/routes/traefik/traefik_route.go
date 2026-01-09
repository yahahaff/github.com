package traefik

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/traefik"
)

// TraefikRouter 注册Traefik相关路由
func TraefikRouter(Router *gin.RouterGroup) {
	traefikGroup := Router.Group("/traefik")
	{
		tc := new(traefik.TraefikController)
		// 获取Traefik路由信息
		traefikGroup.GET("/routes", tc.GetRoutes)
		// 获取Traefik中间件信息
		traefikGroup.GET("/middlewares", tc.GetMiddlewares)
		// 获取Traefik服务信息
		traefikGroup.GET("/services", tc.GetServices)
		// 获取Traefik概览信息
		traefikGroup.GET("/overview", tc.GetOverview)
	}
}
