package handler

import (
    "database/sql"
    "net/http"
    "time"

    "building-ac-3d/backend/internal/interface/http/response"
)

type SystemHandler struct {
    db *sql.DB
}

func NewSystemHandler(db *sql.DB) *SystemHandler {
    return &SystemHandler{db: db}
}

func (handler *SystemHandler) Tenants(w http.ResponseWriter, r *http.Request) {
    response.OK(w, []map[string]interface{}{
        {"id": 1, "name": "默认租户", "code": "default", "status": 1, "remark": "系统初始化租户", "createdAt": time.Now()},
    })
}

func (handler *SystemHandler) Users(w http.ResponseWriter, r *http.Request) {
    response.OK(w, []map[string]interface{}{
        {"id": 1, "tenantId": 1, "username": "admin", "nickname": "系统管理员", "roleIds": []int{1}, "status": 1, "createdAt": time.Now()},
    })
}

func (handler *SystemHandler) Roles(w http.ResponseWriter, r *http.Request) {
    response.OK(w, []map[string]interface{}{
        {"id": 1, "tenantId": 1, "name": "超级管理员", "code": "super_admin", "menuIds": []int{1, 2, 3, 4}, "status": 1, "createdAt": time.Now()},
    })
}

func (handler *SystemHandler) Menus(w http.ResponseWriter, r *http.Request) {
    response.OK(w, []map[string]interface{}{
        {"id": 1, "parentId": 0, "title": "系统管理", "path": "/system", "component": "Layout", "permission": "", "type": "catalog", "sort": 1, "visible": true},
        {"id": 2, "parentId": 1, "title": "租户管理", "path": "/system/tenant", "component": "system/tenant/index", "permission": "system:tenant:list", "type": "menu", "sort": 1, "visible": true},
        {"id": 3, "parentId": 1, "title": "用户管理", "path": "/system/user", "component": "system/user/index", "permission": "system:user:list", "type": "menu", "sort": 2, "visible": true},
        {"id": 4, "parentId": 1, "title": "角色管理", "path": "/system/role", "component": "system/role/index", "permission": "system:role:list", "type": "menu", "sort": 3, "visible": true},
    })
}
