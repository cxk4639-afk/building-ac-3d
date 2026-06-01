package handler

import (
    "database/sql"
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "strings"
    "time"

    "building-ac-3d/backend/internal/interface/http/response"
)

type AuthHandler struct {
    db        *sql.DB
    jwtSecret string
}

type LoginRequest struct {
    TenantCode string `json:"tenantCode"`
    Username   string `json:"username"`
    Password   string `json:"password"`
}

type LoginResponse struct {
    Token        string `json:"token"`
    RefreshToken string `json:"refreshToken"`
}

type LoginUser struct {
    UserID     int64
    TenantID   int64
    Username   string
    Password   string
    Nickname   string
    TenantName string
    TenantCode string
}

func NewAuthHandler(db *sql.DB, jwtSecret string) *AuthHandler {
    return &AuthHandler{db: db, jwtSecret: jwtSecret}
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

    req.TenantCode = strings.TrimSpace(req.TenantCode)
    req.Username = strings.TrimSpace(req.Username)
    if req.TenantCode == "" {
        req.TenantCode = "default"
    }
    if req.Username == "" || req.Password == "" {
        response.Error(w, http.StatusBadRequest, "租户编码、用户名和密码不能为空")
        return
    }

    user, err := handler.findLoginUser(req.TenantCode, req.Username)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            response.Error(w, http.StatusUnauthorized, "用户名或密码错误")
            return
        }
        response.Error(w, http.StatusInternalServerError, err.Error())
        return
    }

    if user.Password != req.Password {
        response.Error(w, http.StatusUnauthorized, "用户名或密码错误")
        return
    }

    token := fmt.Sprintf("dev-token-%d-%d", user.TenantID, user.UserID)
    refreshToken := fmt.Sprintf("dev-refresh-token-%d-%d-%d", user.TenantID, user.UserID, time.Now().Unix())

    response.OK(w, LoginResponse{
        Token:        token,
        RefreshToken: refreshToken,
    })
}

func (handler *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Authorization") == "" {
        response.Error(w, http.StatusUnauthorized, "未登录或登录已过期")
        return
    }

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
    if r.Method != http.MethodPost {
        response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
        return
    }
    response.OK(w, nil)
}

func (handler *AuthHandler) findLoginUser(tenantCode string, username string) (*LoginUser, error) {
    row := handler.db.QueryRow(`
SELECT
    u.id,
    u.tenant_id,
    u.username,
    u.password,
    u.nickname,
    t.name,
    t.code
FROM users u
JOIN tenants t ON t.id = u.tenant_id
WHERE t.code = ?
  AND t.status = 1
  AND u.username = ?
  AND u.status = 1
LIMIT 1
`, tenantCode, username)

    var user LoginUser
    if err := row.Scan(
        &user.UserID,
        &user.TenantID,
        &user.Username,
        &user.Password,
        &user.Nickname,
        &user.TenantName,
        &user.TenantCode,
    ); err != nil {
        return nil, err
    }

    return &user, nil
}
