package dto

type LoginRequest struct {
	TenantCode string `json:"tenantCode"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type TenantInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type UserInfo struct {
	ID          uint       `json:"id"`
	Username    string     `json:"username"`
	Nickname    string     `json:"nickname"`
	Tenant      TenantInfo `json:"tenant"`
	Roles       []string   `json:"roles"`
	Permissions []string   `json:"permissions"`
}
