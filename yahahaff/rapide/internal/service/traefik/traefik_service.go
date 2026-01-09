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

// httpError 自定义HTTP错误类型
type httpError struct {
	statusCode int
	message    string
}

// Error 实现error接口
func (e *httpError) Error() string {
	return e.message
}