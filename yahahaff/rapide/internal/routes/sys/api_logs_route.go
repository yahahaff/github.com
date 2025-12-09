package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/sys"
)

func OperationLogRouter(Router *gin.RouterGroup) {
	// api-logs 路由
	olc := new(sys.OperationLogController)
	Router.GET("/api-logs", olc.GetOperationLog)
}
