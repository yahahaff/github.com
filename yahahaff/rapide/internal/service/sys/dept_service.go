package sys

import (
	"strconv"

	"github.com/yahahaff/rapide/internal/models/sys"
)

// DeptService 部门服务
type DeptService struct{}

// CreateDept 创建部门
func (ds *DeptService) CreateDept(dept sys.Dept) error {
	return dept.Create()
}

// GetDeptByID 根据ID获取部门
func (ds *DeptService) GetDeptByID(id string) (sys.Dept, error) {
	return sys.GetDeptByID(id)
}

// GetDeptList 获取部门列表
func (ds *DeptService) GetDeptList() ([]sys.Dept, error) {
	return sys.GetDeptList()
}

// UpdateDept 更新部门
func (ds *DeptService) UpdateDept(dept sys.Dept) error {
	return dept.Update()
}

// DeleteDept 删除部门
func (ds *DeptService) DeleteDept(id string) error {
	return sys.DeleteDept(id)
}

// BuildDeptTree 构建部门树形结构
func (ds *DeptService) BuildDeptTree(deptList []sys.Dept) []sys.Dept {
	// 创建部门ID到部门的映射
	deptMap := make(map[string]*sys.Dept)
	for i := range deptList {
		deptMap[deptList[i].ID] = &deptList[i]
	}

	// 构建树形结构
	var tree []sys.Dept
	for i := range deptList {
		dept := &deptList[i]
		if dept.Pid == 0 {
			// 根节点
			tree = append(tree, *dept)
		} else {
			// 非根节点，添加到父节点的Children中
			parentID := strconv.FormatUint(dept.Pid, 10)
			if parent, exists := deptMap[parentID]; exists {
				parent.Children = append(parent.Children, dept)
			}
		}
	}

	return tree
}

// GetDeptTree 获取部门树形结构
func (ds *DeptService) GetDeptTree() ([]sys.Dept, error) {
	deptList, err := sys.GetDeptList()
	if err != nil {
		return nil, err
	}
	return ds.BuildDeptTree(deptList), nil
}
