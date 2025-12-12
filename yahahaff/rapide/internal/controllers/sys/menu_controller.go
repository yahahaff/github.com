package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/models/sys"
	sysReq "github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/internal/utils"
	"github.com/yahahaff/rapide/pkg/database"
	"github.com/yahahaff/rapide/pkg/response"
)

// MenuController 菜单控制器
type MenuController struct {
	controllers.BaseAPIController
}

// GetUserMenus 获取用户菜单
func (mc *MenuController) GetUserMenus(c *gin.Context) {
	// 获取当前用户角色ID
	userRoleId := c.GetUint64("current_user_role_id")
	menus, err := service.Entrance.SysService.MenuService.GetUserMenus(userRoleId)
	if err != nil {
		response.Abort500(c, "获取用户菜单失败")
		return
	}
	response.OK(c, menus)
}

// GetMenuList 获取菜单列表
func (mc *MenuController) GetMenuList(c *gin.Context) {
	request := sysReq.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 处理分页参数，设置默认值
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 20 // 设置默认值
	}

	// 处理页码参数，确保页码大于0
	page := request.Page
	if page <= 0 {
		page = 1
	}

	// 查询一级菜单总数
	var total int64
	database.DB.Model(&sys.Menu{}).Where("parent_id IS NULL").Count(&total)

	// 一次性查询所有菜单，不分页
	var allMenus []*sys.Menu
	database.DB.Find(&allMenus)

	// 构建菜单树
	menuTree := utils.BuildMenuTree(allMenus)

	// 构造返回数据
	responseData := map[string]interface{}{
		"result": menuTree,
		"total":  total,
	}

	response.OK(c, responseData)
}
