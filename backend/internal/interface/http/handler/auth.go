package handler

import (
    "database/sql"
    "encoding/json"
    "net/http"

    "building-ac-3d/backend/internal/interface/http/response"
)

type AuthHandler struct {
    db *sql.DB
}

type LoginRequest struct {
    TenantCode string `json:"tenantCode"`
    Username   string `json:"username"`
    Password   string `json:"password"`
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
    return &AuthHandler{db: db}
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
        return
    }

    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        response.Error(w, http.StatusBadRequest, "invalid json body")
        return
    }

    if req.TenantCode == "" {
        req.TenantCode = "default"
    }

    if req.Username != "admin" || req.Password != "123456" {
        response.Error(w, http.StatusUnauthorized, "用户名或密码错误")
        return
    }

    response.OK(w, map[string]string{
        "token":        "dev-token",
        "refreshToken": "dev-refresh-token",
    })
}

func (handler *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
    response.OK(w, map[string]interface{}{
        "id":       1,
        "username": "admin",
        "nickname": "系统管理员",
        "tenant": map[string]interface{}{
            "id":   1,
            "name": "默认租户",
            "code": "default",
        },
        "roles":       []string{"super_admin"},
        "permissions": []string{"*:*:*"},
    })
}

func (handler *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
    response.OK(w, nil)
}
