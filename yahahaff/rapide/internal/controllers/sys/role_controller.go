package sys

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yahahaff/rapide/internal/controllers"
	sysDao "github.com/yahahaff/rapide/internal/dao/sys"
	"github.com/yahahaff/rapide/internal/models/sys"
	sysReq "github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/pkg/response"
)

type RoleController struct {
	controllers.BaseAPIController
}

func (rc *RoleController) GetRole(c *gin.Context) {
	// 1. 验证表单
	request := sysReq.PaginationRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 处理分页参数，设置默认值
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 10 // 设置默认值
	}

	// 处理页码参数
	page := request.Page
	if page <= 0 {
		page = 1
	}

	// 获取角色列表和总数
	roles, total := sysDao.GetRoles(page, pageSize)

	// 构造返回数据
	responseData := map[string]interface{}{
		"result":   roles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}

	response.OK(c, responseData)
}

func (rc *RoleController) AddRole(c *gin.Context) {

	// 1. 验证表单
	request := sysReq.RoleAddRequest{}
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 2. 验证成功，创建数据
	RoleModel := sys.Role{
		RoleName: request.Name,
		Sort:     request.Value,  // 注意：request.Value 对应 Sort 字段
		Remark:   request.Remark, // 注意：request.Remark 对应 Remark 字段
		Status:   request.Status,
		// RoleValue 和 RoleCode 需要根据业务逻辑生成，这里暂时使用默认值
		RoleValue: request.Name,
		RoleCode:  request.Name,
	}

	// 3. 创建角色
	RoleModel.Create()
	if RoleModel.ID == 0 {
		response.Abort500(c, "创建角色失败，请稍后尝试~")
		return
	}

	// 4. 处理权限关联
	if len(request.Permissions) > 0 {
		// 将字符串数组转换为uint64数组
		menuIDs := make([]uint64, 0, len(request.Permissions))
		for _, perm := range request.Permissions {
			var menuID uint64
			if _, err := fmt.Sscan(perm, &menuID); err == nil {
				menuIDs = append(menuIDs, menuID)
			}
		}

		// 分配菜单权限
		if len(menuIDs) > 0 {
			if err := RoleModel.AssignMenus(menuIDs); err != nil {
				response.Abort500(c, "分配菜单权限失败，请稍后尝试~")
				return
			}
		}
	}

	response.OK(c, RoleModel)
}

func (rc *RoleController) DeleteRoleById(c *gin.Context) {

	// 1. 从URL路径中获取id参数
	idStr := c.Param("id")
	var id int
	if _, err := fmt.Sscan(idStr, &id); err != nil || id <= 0 {
		response.Abort500(c, "无效的角色ID")
		return
	}

	// 2. 删除数据
	sysDao.RoleDeletelById(id)
	response.Success(c)

}
