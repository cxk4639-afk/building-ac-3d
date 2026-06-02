package model

import "time"

type Tenant struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Code      string    `gorm:"size:64;uniqueIndex;not null" json:"code"`
	Status    int       `gorm:"not null;default:1" json:"status"`
	Remark    string    `gorm:"size:255;not null;default:''" json:"remark"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (Tenant) TableName() string {
	return "tenants"
}
