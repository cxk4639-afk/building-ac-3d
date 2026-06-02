package dto

type TenantDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Status    int    `json:"status"`
	Remark    string `json:"remark"`
	CreatedAt string `json:"createdAt"`
}

type UserDTO struct {
	ID        uint   `json:"id"`
	TenantID  uint   `json:"tenantId"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	Nickname  string `json:"nickname"`
	RoleIDs   []uint `json:"roleIds"`
	Status    int    `json:"status"`
	CreatedAt string `json:"createdAt"`
}

type RoleDTO struct {
	ID        uint   `json:"id"`
	TenantID  uint   `json:"tenantId"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	MenuIDs   []uint `json:"menuIds"`
	Status    int    `json:"status"`
	CreatedAt string `json:"createdAt"`
}

type MenuDTO struct {
	ID         uint   `json:"id"`
	ParentID   uint   `json:"parentId"`
	Title      string `json:"title"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	Permission string `json:"permission"`
	Type       string `json:"type"`
	Sort       int    `json:"sort"`
	Visible    bool   `json:"visible"`
}
