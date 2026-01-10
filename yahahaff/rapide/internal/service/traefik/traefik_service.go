package traefik

import (
	"encoding/json"
	"io"
	"net/http"
)

// TraefikService 处理Traefik相关业务逻辑
type TraefikService struct{}

// TraefikGroup Traefik服务组
type TraefikGroup struct {
	TraefikService
	TraefikHTTPProviderService
}

// GetRoutes 获取Traefik路由信息
func (ts *TraefikService) GetRoutes() ([]map[string]interface{}, error) {
	// Traefik API地址
	url := "http://172.16.0.60:8080/api/http/routers"

	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, &httpError{statusCode: resp.StatusCode, message: resp.Status}
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var routes []map[string]interface{}
	if err := json.Unmarshal(body, &routes); err != nil {
		return nil, err
	}

	return routes, nil
}

// GetMiddlewares 获取Traefik中间件信息
func (ts *TraefikService) GetMiddlewares() ([]map[string]interface{}, error) {
	// Traefik API地址
	url := "http://172.16.0.60:8080/api/http/middlewares"

	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, &httpError{statusCode: resp.StatusCode, message: resp.Status}
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var middlewares []map[string]interface{}
	if err := json.Unmarshal(body, &middlewares); err != nil {
		return nil, err
	}

	return middlewares, nil
}

// GetServices 获取Traefik服务信息
func (ts *TraefikService) GetServices() ([]map[string]interface{}, error) {
	// Traefik API地址
	url := "http://172.16.0.60:8080/api/http/services"

	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, &httpError{statusCode: resp.StatusCode, message: resp.Status}
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var services []map[string]interface{}
	if err := json.Unmarshal(body, &services); err != nil {
		return nil, err
	}

	return services, nil
}

// GetOverview 获取Traefik概览信息
func (ts *TraefikService) GetOverview() (map[string]interface{}, error) {
	// 获取HTTP相关数据
	httpRouters, err := ts.GetRoutes()
	if err != nil {
		return nil, err
	}

	httpServices, err := ts.GetServices()
	if err != nil {
		return nil, err
	}

	httpMiddlewares, err := ts.GetMiddlewares()
	if err != nil {
		return nil, err
	}

	// 构建概览数据
	overview := map[string]interface{}{
		"http": map[string]interface{}{
			"routers": map[string]interface{}{
				"total":    len(httpRouters),
				"warnings": 0,
				"errors":   0,
			},
			"services": map[string]interface{}{
				"total":    len(httpServices),
				"warnings": 0,
				"errors":   0,
			},
			"middlewares": map[string]interface{}{
				"total":    len(httpMiddlewares),
				"warnings": 0,
				"errors":   0,
			},
		},
		"tcp": map[string]interface{}{
			"routers": map[string]interface{}{
				"total":    0,
				"warnings": 0,
				"errors":   0,
			},
			"services": map[string]interface{}{
				"total":    0,
				"warnings": 0,
				"errors":   0,
			},
			"middlewares": map[string]interface{}{
				"total":    0,
				"warnings": 0,
				"errors":   0,
			},
		},
		"udp": map[string]interface{}{
			"routers": map[string]interface{}{
				"total":    0,
				"warnings": 0,
				"errors":   0,
			},
			"services": map[string]interface{}{
				"total":    0,
				"warnings": 0,
				"errors":   0,
			},
		},
		"features": map[string]interface{}{
			"tracing":    "",
			"metrics":    "",
			"accessLog": false,
		},
		"providers": []string{"Docker"},
	}

	return overview, nil
}

// GetRouteDetail 获取Traefik HTTP路由详情
func (ts *TraefikService) GetRouteDetail(routeName string) (map[string]interface{}, error) {
	// Traefik API地址
	url := "http://172.16.0.60:8080/api/http/routers/" + routeName

	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, &httpError{statusCode: resp.StatusCode, message: resp.Status}
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var routeDetail map[string]interface{}
	if err := json.Unmarshal(body, &routeDetail); err != nil {
		return nil, err
	}

	return routeDetail, nil
}

// GetServiceDetail 获取Traefik HTTP服务详情
func (ts *TraefikService) GetServiceDetail(serviceName string) (map[string]interface{}, error) {
	// Traefik API地址
	url := "http://172.16.0.60:8080/api/http/services/" + serviceName

	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, &httpError{statusCode: resp.StatusCode, message: resp.Status}
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var serviceDetail map[string]interface{}
	if err := json.Unmarshal(body, &serviceDetail); err != nil {
		return nil, err
	}

	return serviceDetail, nil
}

// GetMiddlewareDetail 获取Traefik HTTP中间件详情
func (ts *TraefikService) GetMiddlewareDetail(middlewareName string) (map[string]interface{}, error) {
	// Traefik API地址
	url := "http://172.16.0.60:8080/api/http/middlewares/" + middlewareName

	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, &httpError{statusCode: resp.StatusCode, message: resp.Status}
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
	var middlewareDetail map[string]interface{}
	if err := json.Unmarshal(body, &middlewareDetail); err != nil {
		return nil, err
	}

	return middlewareDetail, nil
}

// httpError 自定义HTTP错误类型
type httpError struct {
	statusCode int
	message    string
}

// Error 实现error接口
func (e *httpError) Error() string {
	return e.message
}