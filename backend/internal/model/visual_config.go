package model

import "time"

type VisualConfig struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TenantID    uint      `gorm:"not null;default:1" json:"tenantId"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `gorm:"size:255;not null;default:''" json:"description"`
	ConfigData  string    `gorm:"type:json;not null" json:"configData"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (VisualConfig) TableName() string {
	return "visual_configs"
}
