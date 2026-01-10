package traefik

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/types"
)

// TraefikMiddleware Traefik中间件模型
type TraefikMiddleware struct {
	models.BaseModel
	models.CommonTimestampsField
	Name     string        `json:"name" gorm:"uniqueIndex:idx_middleware_name_protocol;not null"`
	Type     string        `json:"type" gorm:"not null"` // stripPrefix, redirectRegex, headers, etc.
	Config   types.JSONMap `json:"config" gorm:"type:json;not null"`
	Status   string        `json:"status" gorm:"default:'enabled'"`
	Provider string        `json:"provider" gorm:"default:'http'"`
	Protocol string        `json:"protocol" gorm:"type:varchar(10);default:'http';uniqueIndex:idx_middleware_name_protocol"` // http, tcp, tls
}

// TableName 指定表名
func (TraefikMiddleware) TableName() string {
	return "traefik_middlewares"
}
