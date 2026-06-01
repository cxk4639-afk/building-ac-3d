package handler

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
)

type SystemHandler struct {
    db *sql.DB
}

func NewSystemHandler(db *sql.DB) *SystemHandler {
    return &SystemHandler{db: db}
}

type TenantDTO struct {
    ID        int64  `json:"id"`
    Name      string `json:"name"`
    Code      string `json:"code"`
    Status    int    `json:"status"`
    Remark    string `json:"remark"`
    CreatedAt string `json:"createdAt"`
}

type UserDTO struct {
    ID        int64   `json:"id"`
    TenantID  int64   `json:"tenantId"`
    Username  string  `json:"username"`
    Password  string  `json:"password,omitempty"`
    Nickname  string  `json:"nickname"`
    RoleIDs   []int64 `json:"roleIds"`
    Status    int     `json:"status"`
    CreatedAt string  `json:"createdAt"`
}

type RoleDTO struct {
    ID        int64   `json:"id"`
    TenantID  int64   `json:"tenantId"`
    Name      string  `json:"name"`
    Code      string  `json:"code"`
    MenuIDs   []int64 `json:"menuIds"`
    Status    int     `json:"status"`
    CreatedAt string  `json:"createdAt"`
}

type MenuDTO struct {
    ID         int64  `json:"id"`
    ParentID   int64  `json:"parentId"`
    Title      string `json:"title"`
    Path       string `json:"path"`
    Component  string `json:"component"`
    Permission string `json:"permission"`
    Type       string `json:"type"`
    Sort       int    `json:"sort"`
    Visible    bool   `json:"visible"`
}

func (handler *SystemHandler) Tenants(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handler.listTenants(w)
    case http.MethodPost:
        handler.createTenant(w, r)
    case http.MethodPut:
        handler.updateTenant(w, r)
    case http.MethodDelete:
        handler.deleteByID(w, r, "tenants")
    default:
        writeAPIError(w, http.StatusMethodNotAllowed, "method not allowed")
    }
}

func (handler *SystemHandler) Users(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handler.listUsers(w)
    case http.MethodPost:
        handler.createUser(w, r)
    case http.MethodPut:
        handler.updateUser(w, r)
    case http.MethodDelete:
        handler.deleteUser(w, r)
    default:
        writeAPIError(w, http.StatusMethodNotAllowed, "method not allowed")
    }
}

func (handler *SystemHandler) Roles(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handler.listRoles(w)
    case http.MethodPost:
        handler.createRole(w, r)
    case http.MethodPut:
        handler.updateRole(w, r)
    case http.MethodDelete:
        handler.deleteRole(w, r)
    default:
        writeAPIError(w, http.StatusMethodNotAllowed, "method not allowed")
    }
}

func (handler *SystemHandler) Menus(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handler.listMenus(w)
    case http.MethodPost:
        handler.createMenu(w, r)
    case http.MethodPut:
        handler.updateMenu(w, r)
    case http.MethodDelete:
        handler.deleteRoleMenuSafe(w, r)
    default:
        writeAPIError(w, http.StatusMethodNotAllowed, "method not allowed")
    }
}

func (handler *SystemHandler) listTenants(w http.ResponseWriter) {
    rows, err := handler.db.Query(`SELECT id, name, code, status, remark, created_at FROM tenants ORDER BY id DESC`)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    defer rows.Close()
    list := make([]TenantDTO, 0)
    for rows.Next() {
        var item TenantDTO
        if err := rows.Scan(&item.ID, &item.Name, &item.Code, &item.Status, &item.Remark, &item.CreatedAt); err != nil { writeAPIError(w, 500, err.Error()); return }
        list = append(list, item)
    }
    writeAPIOK(w, list)
}

func (handler *SystemHandler) createTenant(w http.ResponseWriter, r *http.Request) {
    var req TenantDTO
    if !decodeJSON(w, r, &req) { return }
    if req.Name == "" || req.Code == "" { writeAPIError(w, 400, "租户名称和编码不能为空"); return }
    if req.Status == 0 { req.Status = 1 }
    result, err := handler.db.Exec(`INSERT INTO tenants (name, code, status, remark) VALUES (?, ?, ?, ?)`, req.Name, req.Code, req.Status, req.Remark)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    id, _ := result.LastInsertId()
    writeAPIOK(w, map[string]int64{"id": id})
}

func (handler *SystemHandler) updateTenant(w http.ResponseWriter, r *http.Request) {
    var req TenantDTO
    if !decodeJSON(w, r, &req) { return }
    if req.ID == 0 { writeAPIError(w, 400, "id不能为空"); return }
    _, err := handler.db.Exec(`UPDATE tenants SET name=?, code=?, status=?, remark=? WHERE id=?`, req.Name, req.Code, req.Status, req.Remark, req.ID)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    writeAPIOK(w, true)
}

func (handler *SystemHandler) listUsers(w http.ResponseWriter) {
    rows, err := handler.db.Query(`SELECT id, tenant_id, username, nickname, status, created_at FROM users ORDER BY id DESC`)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    defer rows.Close()
    list := make([]UserDTO, 0)
    for rows.Next() {
        var item UserDTO
        if err := rows.Scan(&item.ID, &item.TenantID, &item.Username, &item.Nickname, &item.Status, &item.CreatedAt); err != nil { writeAPIError(w, 500, err.Error()); return }
        item.RoleIDs = handler.findIDs(`SELECT role_id FROM user_roles WHERE user_id=?`, item.ID)
        list = append(list, item)
    }
    writeAPIOK(w, list)
}

func (handler *SystemHandler) createUser(w http.ResponseWriter, r *http.Request) {
    var req UserDTO
    if !decodeJSON(w, r, &req) { return }
    if req.TenantID == 0 || req.Username == "" { writeAPIError(w, 400, "租户和用户名不能为空"); return }
    if req.Password == "" { req.Password = "123456" }
    if req.Status == 0 { req.Status = 1 }
    result, err := handler.db.Exec(`INSERT INTO users (tenant_id, username, password, nickname, status) VALUES (?, ?, ?, ?, ?)`, req.TenantID, req.Username, req.Password, req.Nickname, req.Status)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    id, _ := result.LastInsertId()
    handler.replaceRelations("user_roles", "user_id", "role_id", id, req.RoleIDs)
    writeAPIOK(w, map[string]int64{"id": id})
}

func (handler *SystemHandler) updateUser(w http.ResponseWriter, r *http.Request) {
    var req UserDTO
    if !decodeJSON(w, r, &req) { return }
    if req.ID == 0 { writeAPIError(w, 400, "id不能为空"); return }
    if req.Password != "" {
        _, _ = handler.db.Exec(`UPDATE users SET password=? WHERE id=?`, req.Password, req.ID)
    }
    _, err := handler.db.Exec(`UPDATE users SET tenant_id=?, username=?, nickname=?, status=? WHERE id=?`, req.TenantID, req.Username, req.Nickname, req.Status, req.ID)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    handler.replaceRelations("user_roles", "user_id", "role_id", req.ID, req.RoleIDs)
    writeAPIOK(w, true)
}

func (handler *SystemHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
    id := queryID(r)
    if id == 0 { writeAPIError(w, 400, "id不能为空"); return }
    _, _ = handler.db.Exec(`DELETE FROM user_roles WHERE user_id=?`, id)
    _, err := handler.db.Exec(`DELETE FROM users WHERE id=?`, id)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    writeAPIOK(w, true)
}

func (handler *SystemHandler) listRoles(w http.ResponseWriter) {
    rows, err := handler.db.Query(`SELECT id, tenant_id, name, code, status, created_at FROM roles ORDER BY id DESC`)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    defer rows.Close()
    list := make([]RoleDTO, 0)
    for rows.Next() {
        var item RoleDTO
        if err := rows.Scan(&item.ID, &item.TenantID, &item.Name, &item.Code, &item.Status, &item.CreatedAt); err != nil { writeAPIError(w, 500, err.Error()); return }
        item.MenuIDs = handler.findIDs(`SELECT menu_id FROM role_menus WHERE role_id=?`, item.ID)
        list = append(list, item)
    }
    writeAPIOK(w, list)
}

func (handler *SystemHandler) createRole(w http.ResponseWriter, r *http.Request) {
    var req RoleDTO
    if !decodeJSON(w, r, &req) { return }
    if req.TenantID == 0 || req.Name == "" || req.Code == "" { writeAPIError(w, 400, "租户、角色名称和角色编码不能为空"); return }
    if req.Status == 0 { req.Status = 1 }
    result, err := handler.db.Exec(`INSERT INTO roles (tenant_id, name, code, status) VALUES (?, ?, ?, ?)`, req.TenantID, req.Name, req.Code, req.Status)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    id, _ := result.LastInsertId()
    handler.replaceRelations("role_menus", "role_id", "menu_id", id, req.MenuIDs)
    writeAPIOK(w, map[string]int64{"id": id})
}

func (handler *SystemHandler) updateRole(w http.ResponseWriter, r *http.Request) {
    var req RoleDTO
    if !decodeJSON(w, r, &req) { return }
    if req.ID == 0 { writeAPIError(w, 400, "id不能为空"); return }
    _, err := handler.db.Exec(`UPDATE roles SET tenant_id=?, name=?, code=?, status=? WHERE id=?`, req.TenantID, req.Name, req.Code, req.Status, req.ID)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    handler.replaceRelations("role_menus", "role_id", "menu_id", req.ID, req.MenuIDs)
    writeAPIOK(w, true)
}

func (handler *SystemHandler) deleteRole(w http.ResponseWriter, r *http.Request) {
    id := queryID(r)
    if id == 0 { writeAPIError(w, 400, "id不能为空"); return }
    _, _ = handler.db.Exec(`DELETE FROM user_roles WHERE role_id=?`, id)
    _, _ = handler.db.Exec(`DELETE FROM role_menus WHERE role_id=?`, id)
    _, err := handler.db.Exec(`DELETE FROM roles WHERE id=?`, id)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    writeAPIOK(w, true)
}

func (handler *SystemHandler) listMenus(w http.ResponseWriter) {
    rows, err := handler.db.Query(`SELECT id, parent_id, title, path, component, permission, type, sort, visible FROM menus ORDER BY sort ASC, id ASC`)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    defer rows.Close()
    list := make([]MenuDTO, 0)
    for rows.Next() {
        var item MenuDTO
        var visible int
        if err := rows.Scan(&item.ID, &item.ParentID, &item.Title, &item.Path, &item.Component, &item.Permission, &item.Type, &item.Sort, &visible); err != nil { writeAPIError(w, 500, err.Error()); return }
        item.Visible = visible == 1
        list = append(list, item)
    }
    writeAPIOK(w, list)
}

func (handler *SystemHandler) createMenu(w http.ResponseWriter, r *http.Request) {
    var req MenuDTO
    if !decodeJSON(w, r, &req) { return }
    if req.Title == "" || req.Type == "" { writeAPIError(w, 400, "菜单名称和类型不能为空"); return }
    visible := 0
    if req.Visible { visible = 1 }
    result, err := handler.db.Exec(`INSERT INTO menus (parent_id, title, path, component, permission, type, sort, visible) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, req.ParentID, req.Title, req.Path, req.Component, req.Permission, req.Type, req.Sort, visible)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    id, _ := result.LastInsertId()
    writeAPIOK(w, map[string]int64{"id": id})
}

func (handler *SystemHandler) updateMenu(w http.ResponseWriter, r *http.Request) {
    var req MenuDTO
    if !decodeJSON(w, r, &req) { return }
    if req.ID == 0 { writeAPIError(w, 400, "id不能为空"); return }
    visible := 0
    if req.Visible { visible = 1 }
    _, err := handler.db.Exec(`UPDATE menus SET parent_id=?, title=?, path=?, component=?, permission=?, type=?, sort=?, visible=? WHERE id=?`, req.ParentID, req.Title, req.Path, req.Component, req.Permission, req.Type, req.Sort, visible, req.ID)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    writeAPIOK(w, true)
}

func (handler *SystemHandler) deleteRoleMenuSafe(w http.ResponseWriter, r *http.Request) {
    id := queryID(r)
    if id == 0 { writeAPIError(w, 400, "id不能为空"); return }
    _, _ = handler.db.Exec(`DELETE FROM role_menus WHERE menu_id=?`, id)
    _, err := handler.db.Exec(`DELETE FROM menus WHERE id=?`, id)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    writeAPIOK(w, true)
}

func (handler *SystemHandler) deleteByID(w http.ResponseWriter, r *http.Request, table string) {
    id := queryID(r)
    if id == 0 { writeAPIError(w, 400, "id不能为空"); return }
    _, err := handler.db.Exec("DELETE FROM "+table+" WHERE id=?", id)
    if err != nil { writeAPIError(w, 500, err.Error()); return }
    writeAPIOK(w, true)
}

func (handler *SystemHandler) findIDs(query string, id int64) []int64 {
    rows, err := handler.db.Query(query, id)
    if err != nil { return []int64{} }
    defer rows.Close()
    ids := make([]int64, 0)
    for rows.Next() {
        var value int64
        if err := rows.Scan(&value); err == nil { ids = append(ids, value) }
    }
    return ids
}

func (handler *SystemHandler) replaceRelations(table, leftColumn, rightColumn string, leftID int64, rightIDs []int64) {
    _, _ = handler.db.Exec("DELETE FROM "+table+" WHERE "+leftColumn+"=?", leftID)
    for _, rightID := range rightIDs {
        _, _ = handler.db.Exec("INSERT IGNORE INTO "+table+" ("+leftColumn+", "+rightColumn+") VALUES (?, ?)", leftID, rightID)
    }
}

func decodeJSON(w http.ResponseWriter, r *http.Request, target interface{}) bool {
    defer r.Body.Close()
    if err := json.NewDecoder(r.Body).Decode(target); err != nil {
        writeAPIError(w, http.StatusBadRequest, "invalid json body")
        return false
    }
    return true
}

func queryID(r *http.Request) int64 {
    id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
    return id
}

func writeAPIOK(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    _ = json.NewEncoder(w).Encode(map[string]interface{}{"code": 0, "message": "success", "data": data})
}

func writeAPIError(w http.ResponseWriter, status int, message string) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(status)
    _ = json.NewEncoder(w).Encode(map[string]interface{}{"code": status, "message": message, "data": nil})
}
