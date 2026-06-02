package model

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TenantID  uint      `gorm:"not null;uniqueIndex:uk_tenant_role_code" json:"tenantId"`
	Name      string    `gorm:"size:64;not null" json:"name"`
	Code      string    `gorm:"size:64;not null;uniqueIndex:uk_tenant_role_code" json:"code"`
	Status    int       `gorm:"not null;default:1" json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (Role) TableName() string {
	return "roles"
}
