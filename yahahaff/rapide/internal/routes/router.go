package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/middlewares"
	"github.com/yahahaff/rapide/internal/routes/ssl"
	"github.com/yahahaff/rapide/internal/routes/sys"
)

// RegisterAPIRoutes 注册分支路由
func RegisterAPIRoutes(Router *gin.Engine) {
	// ==================== 不需要 JWT 认证的路由 ====================
	
	// 1. 内部公开路由
	internalGroup := Router.Group("")
	{
		// 内部路由由sys包提供
		sys.InternalRouter(internalGroup)
	}
	
	// 2. 验证码路由
	captchaGroup := Router.Group("/api/captcha")
	{
		sys.CaptchaRouter(captchaGroup)
	}
	
	// 3. 认证相关路由 (/api/auth)
	authGroup := Router.Group("/api/auth")
	{
		sys.LoginRouter(authGroup)      // 登录
		sys.SingupRouter(authGroup)     // 注册
	}
	
	// ==================== 需要 JWT 认证的路由 ====================
	
	// 4. SSL相关路由 (/api/ssl)
	sslGroup := Router.Group("/api/ssl")
	sslGroup.Use(middlewares.AuthJWT()) // JWT认证
	{
		ssl.SSLCertRouter(sslGroup)
	}
	
	// 5. 系统管理路由 (/api/sys)
	sysGroup := Router.Group("/api/sys")
	sysGroup.Use(middlewares.AuthJWT()) // JWT认证
	{
		sys.MenuRouter(sysGroup)            // 菜单管理
		sys.CasbinRouter(sysGroup)          // 权限管理
		sys.UserRouter(sysGroup)            // 用户管理
		sys.OperationLogRouter(sysGroup)    // 操作日志
	}
}
