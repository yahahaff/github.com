package traefik

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/response"
)

// TraefikController Traefik控制器
type TraefikController struct {
	controllers.BaseAPIController
}

// GetRoutes 获取Traefik路由信息
func (tc *TraefikController) GetRoutes(c *gin.Context) {
	// 调用服务获取Traefik路由信息
	routes, err := service.Entrance.TraefikService.TraefikService.GetRoutes()
	if err != nil {
		// 记录错误日志
		log.Printf("Failed to get Traefik routes: %v", err)
		// 失败，显示错误提示
		response.Abort500(c, err.Error())
		return
	}

	// 获取分页参数，支持两种格式：page[currentPage] 和 page
	page := c.DefaultQuery("page[currentPage]", c.DefaultQuery("page", "1"))
	pageSize := c.DefaultQuery("page[pageSize]", c.DefaultQuery("pageSize", "10"))

	// 返回结果，包含列表和总数
	result := gin.H{
		"result":   routes,
		"total":    len(routes),
		"page":     page,
		"pageSize": pageSize,
	}

	// 成功，返回路由信息
	response.OK(c, result)
}

// GetMiddlewares 获取Traefik中间件信息
func (tc *TraefikController) GetMiddlewares(c *gin.Context) {
	// 调用服务获取Traefik中间件信息
	middlewares, err := service.Entrance.TraefikService.TraefikService.GetMiddlewares()
	if err != nil {
		// 记录错误日志
		log.Printf("Failed to get Traefik middlewares: %v", err)
		// 失败，显示错误提示
		response.Abort500(c, err.Error())
		return
	}

	// 获取分页参数，支持两种格式：page[currentPage] 和 page
	page := c.DefaultQuery("page[currentPage]", c.DefaultQuery("page", "1"))
	pageSize := c.DefaultQuery("page[pageSize]", c.DefaultQuery("pageSize", "10"))

	// 返回结果，包含列表和总数
	result := gin.H{
		"result":   middlewares,
		"total":    len(middlewares),
		"page":     page,
		"pageSize": pageSize,
	}

	// 成功，返回中间件信息
	response.OK(c, result)
}

// GetServices 获取Traefik服务信息
func (tc *TraefikController) GetServices(c *gin.Context) {
	// 调用服务获取Traefik服务信息
	services, err := service.Entrance.TraefikService.TraefikService.GetServices()
	if err != nil {
		// 记录错误日志
		log.Printf("Failed to get Traefik services: %v", err)
		// 失败，显示错误提示
		response.Abort500(c, err.Error())
		return
	}

	// 获取分页参数，支持两种格式：page[currentPage] 和 page
	page := c.DefaultQuery("page[currentPage]", c.DefaultQuery("page", "1"))
	pageSize := c.DefaultQuery("page[pageSize]", c.DefaultQuery("pageSize", "10"))

	// 返回结果，包含列表和总数
	result := gin.H{
		"result":   services,
		"total":    len(services),
		"page":     page,
		"pageSize": pageSize,
	}

	// 成功，返回服务信息
	response.OK(c, result)
}

// GetOverview 获取Traefik概览信息
func (tc *TraefikController) GetOverview(c *gin.Context) {
	// 调用服务获取Traefik概览信息
	overview, err := service.Entrance.TraefikService.TraefikService.GetOverview()
	if err != nil {
		// 记录错误日志
		log.Printf("Failed to get Traefik overview: %v", err)
		// 失败，显示错误提示
		response.Abort500(c, err.Error())
		return
	}

	// 成功，返回概览信息
	response.OK(c, overview)
}

// GetRouteDetail 获取Traefik HTTP路由详情
func (tc *TraefikController) GetRouteDetail(c *gin.Context) {
	// 获取路由名称
	routeName := c.Param("name")
	if routeName == "" {
		response.Abort400(c, "Route name is required")
		return
	}

	// 调用服务获取Traefik路由详情
	routeDetail, err := service.Entrance.TraefikService.TraefikService.GetRouteDetail(routeName)
	if err != nil {
		// 记录错误日志
		log.Printf("Failed to get Traefik route detail: %v", err)
		// 失败，显示错误提示
		response.Abort500(c, err.Error())
		return
	}

	// 成功，返回路由详情
	response.OK(c, routeDetail)
}

// GetServiceDetail 获取Traefik HTTP服务详情
func (tc *TraefikController) GetServiceDetail(c *gin.Context) {
	// 获取服务名称
	serviceName := c.Param("name")
	if serviceName == "" {
		response.Abort400(c, "Service name is required")
		return
	}

	// 调用服务获取Traefik服务详情
	serviceDetail, err := service.Entrance.TraefikService.TraefikService.GetServiceDetail(serviceName)
	if err != nil {
		// 记录错误日志
		log.Printf("Failed to get Traefik service detail: %v", err)
		// 失败，显示错误提示
		response.Abort500(c, err.Error())
		return
	}

	// 成功，返回服务详情
	response.OK(c, serviceDetail)
}

// GetMiddlewareDetail 获取Traefik HTTP中间件详情
func (tc *TraefikController) GetMiddlewareDetail(c *gin.Context) {
	// 获取中间件名称
	middlewareName := c.Param("name")
	if middlewareName == "" {
		response.Abort400(c, "Middleware name is required")
		return
	}

	// 调用服务获取Traefik中间件详情
	middlewareDetail, err := service.Entrance.TraefikService.TraefikService.GetMiddlewareDetail(middlewareName)
	if err != nil {
		// 记录错误日志
		log.Printf("Failed to get Traefik middleware detail: %v", err)
		// 失败，显示错误提示
		response.Abort500(c, err.Error())
		return
	}

	// 成功，返回中间件详情
	response.OK(c, middlewareDetail)
}
