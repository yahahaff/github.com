package traefik

import (
	"github.com/yahahaff/rapide/internal/models"
	"github.com/yahahaff/rapide/pkg/types"
)

// TraefikRouter Traefik路由模型
type TraefikRouter struct {
	models.BaseModel
	models.CommonTimestampsField
	Name        string        `json:"name" gorm:"uniqueIndex:idx_router_name_protocol;not null"`
	EntryPoints types.JSONSlice `json:"entryPoints" gorm:"type:json"`
	Service     string        `json:"service" gorm:"not null"`
	Rule        string        `json:"rule" gorm:"not null"`
	RuleSyntax  string        `json:"ruleSyntax" gorm:"default:'default'"`
	Priority    int64         `json:"priority" gorm:"default:0"`
	Middlewares types.JSONSlice `json:"middlewares" gorm:"type:json"`
	TLS         types.JSONMap `json:"tls" gorm:"type:json"` // TLS配置
	Protocol    string        `json:"protocol" gorm:"type:varchar(10);default:'http';uniqueIndex:idx_router_name_protocol"` // http, tcp, udp
	Status      string        `json:"status" gorm:"default:'enabled'"`
	Provider    string        `json:"provider" gorm:"default:'http'"`
}

// TableName 指定表名
func (TraefikRouter) TableName() string {
	return "traefik_routers"
}
