package sys

// DeptCreateRequest 创建部门请求
type DeptCreateRequest struct {
	Pid     string `json:"pid" binding:"required"`
	Name    string `json:"name" binding:"required,max=255"`
	Status  int    `json:"status" binding:"omitempty,oneof=0 1"`
	Remark  string `json:"remark" binding:"omitempty"`
}

// DeptUpdateRequest 更新部门请求
type DeptUpdateRequest struct {
	ID      string `json:"id" binding:"required"`
	Pid     string `json:"pid" binding:"required"`
	Name    string `json:"name" binding:"required,max=255"`
	Status  int    `json:"status" binding:"omitempty,oneof=0 1"`
	Remark  string `json:"remark" binding:"omitempty"`
}

// DeptIDRequest 部门ID请求
type DeptIDRequest struct {
	ID string `json:"id" binding:"required" uri:"id"`
}
