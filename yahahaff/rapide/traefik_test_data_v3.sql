-- Traefik v3.6测试数据
-- 创建时间: 2026-01-10

-- ===================================
-- 1. 中间件测试数据
-- ===================================
INSERT INTO traefik_middlewares (name, type, config, status, provider, protocol, created_at, updated_at) VALUES
('strip-app', 'stripPrefix', '{"prefixes": ["/app"]}', 'enabled', 'http', 'http', datetime('now'), datetime('now')),
('redirect-https', 'redirectScheme', '{"scheme": "https", "permanent": true}', 'enabled', 'http', 'http', datetime('now'), datetime('now')),
('add-headers', 'headers', '{"customResponseHeaders": {"X-Content-Type-Options": ["nosniff"], "X-Frame-Options": ["DENY"]}}', 'enabled', 'http', 'http', datetime('now'), datetime('now')),
('rate-limit', 'rateLimit', '{"average": 100, "burst": 200}', 'enabled', 'http', 'http', datetime('now'), datetime('now'));

-- ===================================
-- 2. 服务测试数据 (Traefik v3.6 格式: healthCheck 移到 loadBalancer 内部)
-- ===================================
INSERT INTO traefik_services (name, status, provider, protocol, type, load_balancer, tcp, created_at, updated_at) VALUES
-- my-service: HTTP负载均衡服务，包含2个后端服务器和健康检查
('my-service', 'enabled', 'http', 'http', 'loadbalancer', '{"servers": [{"url": "http://192.168.1.100:8080"}, {"url": "http://192.168.1.101:8080"}], "healthCheck": {"path": "/health", "interval": "30s", "timeout": "10s"}}', NULL, datetime('now'), datetime('now')),
-- api-service: API服务，包含1个后端服务器和健康检查
('api-service', 'enabled', 'http', 'http', 'loadbalancer', '{"servers": [{"url": "http://192.168.1.200:3000"}], "healthCheck": {"path": "/api/health", "interval": "15s", "timeout": "5s"}}', NULL, datetime('now'), datetime('now')),
-- tcp-service: TCP服务，用于数据库连接，包含健康检查
('tcp-service', 'enabled', 'http', 'tcp', 'loadbalancer', '{"servers": [{"url": "tcp://192.168.1.300:3306"}], "healthCheck": {"tcpSocket": {"timeout": "5s"}, "interval": "30s"}}', '{"idleTimeout": "360s"}', datetime('now'), datetime('now'));

-- ===================================
-- 3. 路由测试数据
-- ===================================
INSERT INTO traefik_routers (name, entry_points, service, rule, rule_syntax, priority, middlewares, tls, status, provider, protocol, created_at, updated_at) VALUES
('my-router', '["web"]', 'my-service', 'PathPrefix(`/app`)', 'default', 0, '["strip-app", "rate-limit"]', NULL, 'enabled', 'http', 'http', datetime('now'), datetime('now')),
('api-router', '["web", "websecure"]', 'api-service', 'Host(`api.example.com`)', 'default', 10, '["add-headers"]', '{"certResolver": "letsencrypt"}', 'enabled', 'http', 'http', datetime('now'), datetime('now')),
('redirect-router', '["web"]', 'api-service', 'Host(`www.example.com`)', 'default', 20, '["redirect-https"]', NULL, 'enabled', 'http', 'http', datetime('now'), datetime('now')),
('tcp-router', '["tcp"]', 'tcp-service', 'HostSNI(`*`)', 'default', 0, NULL, '{"passthrough": true}', 'enabled', 'http', 'tcp', datetime('now'), datetime('now'));
