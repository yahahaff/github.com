package traefik

import (
	traefikDAO "github.com/yahahaff/rapide/internal/dao/traefik"
	traefikModel "github.com/yahahaff/rapide/internal/models/traefik"
)

// TraefikHTTPProviderService Traefik HTTP自动发现服务
type TraefikHTTPProviderService struct {
	traefikDAO *traefikDAO.TraefikDAO
}

// NewTraefikHTTPProviderService 创建TraefikHTTPProviderService实例
func NewTraefikHTTPProviderService() *TraefikHTTPProviderService {
	return &TraefikHTTPProviderService{
		traefikDAO: traefikDAO.NewTraefikDAO(),
	}
}

// GetHTTPProviderConfig 获取Traefik HTTP Provider配置
func (svc *TraefikHTTPProviderService) GetHTTPProviderConfig() (map[string]interface{}, error) {
	// 获取所有启用的路由
	routers, err := svc.traefikDAO.GetAllRouters()
	if err != nil {
		return nil, err
	}

	// 获取所有启用的服务
	services, err := svc.traefikDAO.GetAllServices()
	if err != nil {
		return nil, err
	}

	// 获取所有启用的中间件
	middlewares, err := svc.traefikDAO.GetAllMiddlewares()
	if err != nil {
		return nil, err
	}

	// 构建HTTP Provider配置，使用名称作为键，完全符合Traefik HTTP Provider格式
	config := map[string]interface{}{
		"http": map[string]interface{}{
			"routers":     buildRoutersConfig(routers),
			"services":    buildServicesConfig(services),
			"middlewares": buildMiddlewaresConfig(middlewares),
		},
	}

	return config, nil
}

// buildRoutersConfig 构建路由配置，以名称为键
func buildRoutersConfig(routers []traefikModel.TraefikRouter) map[string]interface{} {
	routerConfig := make(map[string]interface{})

	for _, router := range routers {
		routerData := map[string]interface{}{
			"entryPoints": router.EntryPoints,
			"service":     router.Service,
			"rule":        router.Rule,
		}

		// 添加可选字段
		if router.RuleSyntax != "" {
			routerData["ruleSyntax"] = router.RuleSyntax
		}

		if router.Priority > 0 {
			routerData["priority"] = router.Priority
		}

		if len(router.Middlewares) > 0 {
			routerData["middlewares"] = router.Middlewares
		}

		if len(router.TLS) > 0 {
			routerData["tls"] = router.TLS
		}

		// 使用名称作为键
		routerConfig[router.Name] = routerData
	}

	return routerConfig
}

// buildServicesConfig 构建服务配置，以名称为键
func buildServicesConfig(services []traefikModel.TraefikService) map[string]interface{} {
	serviceConfig := make(map[string]interface{})

	for _, service := range services {
		serviceData := make(map[string]interface{})

		// 根据服务类型添加不同的配置
		switch service.Type {
		case "loadbalancer":
			if len(service.LoadBalancer) > 0 {
				serviceData["loadBalancer"] = service.LoadBalancer
			}
		case "weighted":
			if len(service.Weighted) > 0 {
				serviceData["weighted"] = service.Weighted
			}
		case "mirror":
			if len(service.Mirror) > 0 {
				serviceData["mirror"] = service.Mirror
			}
		}

		// 为TCP/UDP服务添加特定配置
		if service.Protocol == "tcp" && len(service.TCP) > 0 {
			serviceData["tcp"] = service.TCP
		}

		if service.Protocol == "udp" && len(service.UDP) > 0 {
			serviceData["udp"] = service.UDP
		}

		// 使用名称作为键
		serviceConfig[service.Name] = serviceData
	}

	return serviceConfig
}

// buildMiddlewaresConfig 构建中间件配置，以名称为键
func buildMiddlewaresConfig(middlewares []traefikModel.TraefikMiddleware) map[string]interface{} {
	middlewareConfig := make(map[string]interface{})

	for _, middleware := range middlewares {
		middlewareData := map[string]interface{}{
			middleware.Type: middleware.Config,
		}

		// 使用名称作为键
		middlewareConfig[middleware.Name] = middlewareData
	}

	return middlewareConfig
}
