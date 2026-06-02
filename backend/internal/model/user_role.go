package model

type UserRole struct {
	UserID uint `gorm:"primaryKey" json:"userId"`
	RoleID uint `gorm:"primaryKey" json:"roleId"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
