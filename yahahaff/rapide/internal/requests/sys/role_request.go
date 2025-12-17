package sys

type RoleAddRequest struct {
	Name        string   `json:"name"  validate:"required"`
	Value       int      `json:"sort"`
	Desc        string   `json:"tips"`
	Permissions []string `json:"permissions" `
	Status      int      `json:"status"`
	Remark      string   `json:"remark"`
}

type RoleDeleteRequest struct {
	Id int `json:"id" validate:"required"`
}

type RoleMenuRequest struct {
	RoleID  int   `json:"role_id" validate:"required"`
	MenuIDs []int `json:"menu_ids" validate:"required"`
}
