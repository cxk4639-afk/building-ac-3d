package model

type Menu struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ParentID   uint   `gorm:"not null;default:0" json:"parentId"`
	Title      string `gorm:"size:64;not null" json:"title"`
	Path       string `gorm:"size:255;not null;default:''" json:"path"`
	Component  string `gorm:"size:255;not null;default:''" json:"component"`
	Permission string `gorm:"size:128;not null;default:''" json:"permission"`
	Type       string `gorm:"size:16;not null" json:"type"`
	Sort       int    `gorm:"not null;default:0" json:"sort"`
	Visible    bool   `gorm:"not null;default:true" json:"visible"`
}

func (Menu) TableName() string {
	return "menus"
}
