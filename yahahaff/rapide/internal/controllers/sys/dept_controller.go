package sys

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yahahaff/rapide/internal/controllers"
	"github.com/yahahaff/rapide/internal/models/sys"
	sysReq "github.com/yahahaff/rapide/internal/requests/sys"
	"github.com/yahahaff/rapide/internal/requests/validators"
	"github.com/yahahaff/rapide/internal/service"
	"github.com/yahahaff/rapide/pkg/response"
)

// DeptController 部门控制器
type DeptController struct {
	controllers.BaseAPIController
}

// GetDeptTree 获取部门树形结构
func (dc *DeptController) GetDeptTree(c *gin.Context) {
	deptTree, err := service.Entrance.SysService.DeptService.GetDeptTree()
	if err != nil {
		response.Abort500(c, "获取部门树形结构失败")
		return
	}
	response.OK(c, deptTree)
}

// CreateDept 创建部门
func (dc *DeptController) CreateDept(c *gin.Context) {
	var request sysReq.DeptCreateRequest
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 生成UUID作为部门ID
	uuid := uuid.New().String()

	// 转换请求为部门模型
	dept := sys.Dept{
		ID:         uuid,
		Pid:        0, // 默认为根部门，可根据实际情况调整
		Name:       request.Name,
		Status:     request.Status,
		Remark:     request.Remark,
		CreateTime: time.Now(),
	}

	// 调用服务创建部门
	err := service.Entrance.SysService.DeptService.CreateDept(dept)
	if err != nil {
		response.Abort500(c, "创建部门失败")
		return
	}

	response.OK(c, gin.H{"id": uuid})
}

// UpdateDept 更新部门
func (dc *DeptController) UpdateDept(c *gin.Context) {
	var request sysReq.DeptUpdateRequest

	// 1. 检查URL路径中是否有id参数，如果有则设置为请求体的id
	if idStr := c.Param("id"); idStr != "" {
		request.ID = idStr
	}

	// 2. 解析和验证请求体
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 3. 获取要更新的部门
	dept, err := service.Entrance.SysService.DeptService.GetDeptByID(request.ID)
	if err != nil {
		response.Abort500(c, "部门不存在")
		return
	}

	// 4. 只更新请求中包含的字段
	// 更新部门名称
	if request.Name != "" {
		dept.Name = request.Name
	}
	// 更新部门状态
	dept.Status = request.Status
	// 更新部门备注
	if request.Remark != "" {
		dept.Remark = request.Remark
	}

	// 5. 调用服务更新部门
	err = service.Entrance.SysService.DeptService.UpdateDept(dept)
	if err != nil {
		response.Abort500(c, "更新部门失败")
		return
	}

	response.OK(c, gin.H{"id": request.ID})
}

// DeleteDept 删除部门
func (dc *DeptController) DeleteDept(c *gin.Context) {
	var request sysReq.DeptIDRequest
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 调用服务删除部门
	err := service.Entrance.SysService.DeptService.DeleteDept(request.ID)
	if err != nil {
		response.Abort500(c, "删除部门失败")
		return
	}

	response.OK(c, gin.H{"id": request.ID})
}

// GetDept 获取单个部门信息
func (dc *DeptController) GetDept(c *gin.Context) {
	var request sysReq.DeptIDRequest
	if ok := validators.Validate(c, &request); !ok {
		return
	}

	// 获取部门信息
	dept, err := service.Entrance.SysService.DeptService.GetDeptByID(request.ID)
	if err != nil {
		response.Abort500(c, "部门不存在")
		return
	}

	response.OK(c, dept)
}
