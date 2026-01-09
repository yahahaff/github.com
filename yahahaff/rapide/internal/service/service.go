package service

import (
	"github.com/yahahaff/rapide/internal/service/sys"
	"github.com/yahahaff/rapide/internal/service/ssl"
	"github.com/yahahaff/rapide/internal/service/traefik"
)

var Entrance = ServiceGroup{}

type ServiceGroup struct {
	SysService    sys.SysGroup
	SSLService    ssl.SSLGroup
	TraefikService traefik.TraefikGroup
}
