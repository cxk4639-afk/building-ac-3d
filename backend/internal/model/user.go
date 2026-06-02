package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TenantID  uint      `gorm:"not null;uniqueIndex:uk_tenant_username" json:"tenantId"`
	Username  string    `gorm:"size:64;not null;uniqueIndex:uk_tenant_username" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Nickname  string    `gorm:"size:64;not null;default:''" json:"nickname"`
	Status    int       `gorm:"not null;default:1" json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}
