package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/sys"
)

func OperationLogRouter(Router *gin.RouterGroup) {

	{
		// api-logs 路由
		OperationLogGroup := Router.Group("/sys/api-logs")
		olc := new(sys.OperationLogController)
		OperationLogGroup.GET("getOperationLog", olc.GetOperationLog)

	}
}
