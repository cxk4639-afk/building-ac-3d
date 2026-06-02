package model

import "time"

type Tenant struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"size:100;not null" json:"name"`
    Code      string    `gorm:"size:64;not null;uniqueIndex" json:"code"`
    Status    int       `gorm:"not null;default:1" json:"status"`
    Remark    string    `gorm:"size:255;not null;default:''" json:"remark"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    TenantID  uint      `gorm:"not null;index" json:"tenantId"`
    Username  string    `gorm:"size:64;not null;uniqueIndex:uk_tenant_username" json:"username"`
    Password  string    `gorm:"size:255;not null" json:"password,omitempty"`
    Nickname  string    `gorm:"size:64;not null;default:''" json:"nickname"`
    Status    int       `gorm:"not null;default:1" json:"status"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    Tenant    Tenant    `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
    Roles     []Role    `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}

type Role struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    TenantID  uint      `gorm:"not null;index;uniqueIndex:uk_tenant_role_code" json:"tenantId"`
    Name      string    `gorm:"size:64;not null" json:"name"`
    Code      string    `gorm:"size:64;not null;uniqueIndex:uk_tenant_role_code" json:"code"`
    Status    int       `gorm:"not null;default:1" json:"status"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    Tenant    Tenant    `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
    Menus     []Menu    `gorm:"many2many:role_menus;" json:"menus,omitempty"`
}

type Menu struct {
    ID         uint   `gorm:"primaryKey" json:"id"`
    ParentID   uint   `gorm:"not null;default:0;index" json:"parentId"`
    Title      string `gorm:"size:64;not null" json:"title"`
    Path       string `gorm:"size:255;not null;default:''" json:"path"`
    Component  string `gorm:"size:255;not null;default:''" json:"component"`
    Permission string `gorm:"size:128;not null;default:''" json:"permission"`
    Type       string `gorm:"size:16;not null" json:"type"`
    Sort       int    `gorm:"not null;default:0" json:"sort"`
    Visible    bool   `gorm:"not null;default:true" json:"visible"`
}

type VisualConfig struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    TenantID    uint      `gorm:"not null;default:1;index" json:"tenantId"`
    Name        string    `gorm:"size:100;not null" json:"name"`
    Description string    `gorm:"size:255;not null;default:''" json:"description"`
    ConfigData  string    `gorm:"type:json;not null" json:"configData"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}
