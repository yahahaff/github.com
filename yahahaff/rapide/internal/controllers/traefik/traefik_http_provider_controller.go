package traefik

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/service/traefik"
	"github.com/yahahaff/rapide/pkg/response"
)

// TraefikHTTPProviderController Traefik HTTP自动发现控制器
type TraefikHTTPProviderController struct {
	controllers.BaseAPIController
	traefikHTTPProviderService *traefik.TraefikHTTPProviderService
}

// NewTraefikHTTPProviderController 创建TraefikHTTPProviderController实例
func NewTraefikHTTPProviderController(traefikHTTPProviderService *traefik.TraefikHTTPProviderService) *TraefikHTTPProviderController {
	return &TraefikHTTPProviderController{
		traefikHTTPProviderService: traefikHTTPProviderService,
	}
}

// GetHTTPProviderConfig 获取Traefik HTTP Provider配置
func (ctrl *TraefikHTTPProviderController) GetHTTPProviderConfig(c *gin.Context) {
	// 获取配置
	config, err := ctrl.traefikHTTPProviderService.GetHTTPProviderConfig()
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}

	// 返回配置，Traefik HTTP Provider期望的格式
	response.OK(c, config)
}
