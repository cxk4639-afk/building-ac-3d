package model

type RoleMenu struct {
	RoleID uint `gorm:"primaryKey" json:"roleId"`
	MenuID uint `gorm:"primaryKey" json:"menuId"`
}

func (RoleMenu) TableName() string {
	return "role_menus"
}
