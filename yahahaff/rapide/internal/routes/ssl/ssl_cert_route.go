// Package ssl
package ssl

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/ssl"
)

// SSLCertRouter SSL证书路由
func SSLCertRouter(Router *gin.RouterGroup) {
	sslCertGroup := Router.Group("/ssl")
	{
		sslCertController := new(ssl.SSLCertController)
		// 获取SSL证书列表
		sslCertGroup.GET("/list", sslCertController.GetSSLCertList)
		// 创建SSL证书
		sslCertGroup.POST("/create", sslCertController.CreateSSLCert)
	}
}
