package traefik

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/types"
)

// TraefikService Traefik服务模型
type TraefikService struct {
	models.BaseModel
	models.CommonTimestampsField
	Name         string        `json:"name" gorm:"uniqueIndex:idx_service_name_protocol;not null"`
	Status       string        `json:"status" gorm:"default:'enabled'"`
	Provider     string        `json:"provider" gorm:"default:'http'"`
	Protocol     string        `json:"protocol" gorm:"type:varchar(10);default:'http';uniqueIndex:idx_service_name_protocol"` // http, tcp, udp
	Type         string        `json:"type" gorm:"not null"`                                                                  // loadbalancer, weighted, mirror, etc.
	LoadBalancer types.JSONMap `json:"loadBalancer" gorm:"type:json"`                                                         // 负载均衡器配置，包含healthCheck子字段
	Weighted     types.JSONMap `json:"weighted" gorm:"type:json"`                                                             // 加权服务配置
	Mirror       types.JSONMap `json:"mirror" gorm:"type:json"`                                                               // 镜像服务配置
	TCP          types.JSONMap `json:"tcp" gorm:"type:json"`                                                                  // TCP特定配置
	UDP          types.JSONMap `json:"udp" gorm:"type:json"`                                                                  // UDP特定配置
}

// TableName 指定表名
func (TraefikService) TableName() string {
	return "traefik_services"
}
