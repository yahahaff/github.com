package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/routes/sys"
	"github.com/yahahaff/rapide/internal/routes/ssl"
	"github.com/yahahaff/rapide/internal/middlewares"
)

// RegisterAPIRoutes 注册分支路由
func RegisterAPIRoutes(Router *gin.Engine) {
	sys.RouterGroup(Router)
	
	// SSL相关需要jwt认证路由
	sslGroup := Router.Group("/api/ssl")
	//JWT认证 接口权限校验
	sslGroup.Use(middlewares.AuthJWT())
	ssl.SSLCertRouter(sslGroup)
}
