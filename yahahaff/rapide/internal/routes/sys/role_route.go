package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers/sys"
)

func RoleRouter(Router *gin.RouterGroup) {
	// role
	roleGroup := Router.Group("/role")
	roc := new(sys.RoleController)
	roleGroup.GET("getRole", roc.GetRole)
	roleGroup.GET("list", roc.GetRole) // 添加list路由，指向相同的处理函数
	roleGroup.POST("addRole", roc.AddRole)
	roleGroup.DELETE("deleteRole", roc.DeleteRoleById)
	// 添加根据ID删除角色的路由
	roleGroup.DELETE("delete/:id", roc.DeleteRoleById)
}
