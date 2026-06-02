package database

import (
	"building-ac-3d/backend/internal/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.Tenant{},
		&model.User{},
		&model.Role{},
		&model.Menu{},
		&model.UserRole{},
		&model.RoleMenu{},
		&model.VisualConfig{},
	); err != nil {
		return err
	}

	return seedDefaults(db)
}

func seedDefaults(db *gorm.DB) error {
	tenant := model.Tenant{
		ID:     1,
		Name:   "默认租户",
		Code:   "default",
		Status: 1,
		Remark: "系统初始化租户",
	}
	if err := db.FirstOrCreate(&tenant, model.Tenant{Code: "default"}).Error; err != nil {
		return err
	}

	user := model.User{
		ID:       1,
		TenantID: tenant.ID,
		Username: "admin",
		Password: "123456",
		Nickname: "系统管理员",
		Status:   1,
	}
	if err := db.FirstOrCreate(&user, model.User{TenantID: tenant.ID, Username: "admin"}).Error; err != nil {
		return err
	}

	role := model.Role{
		ID:       1,
		TenantID: tenant.ID,
		Name:     "超级管理员",
		Code:     "super_admin",
		Status:   1,
	}
	if err := db.FirstOrCreate(&role, model.Role{TenantID: tenant.ID, Code: "super_admin"}).Error; err != nil {
		return err
	}

	menus := []model.Menu{
		{ID: 1, ParentID: 0, Title: "系统管理", Path: "/system", Component: "Layout", Permission: "", Type: "catalog", Sort: 1, Visible: true},
		{ID: 2, ParentID: 1, Title: "租户管理", Path: "/system/tenant", Component: "system/tenant/index", Permission: "system:tenant:list", Type: "menu", Sort: 1, Visible: true},
		{ID: 3, ParentID: 1, Title: "用户管理", Path: "/system/user", Component: "system/user/index", Permission: "system:user:list", Type: "menu", Sort: 2, Visible: true},
		{ID: 4, ParentID: 1, Title: "角色管理", Path: "/system/role", Component: "system/role/index", Permission: "system:role:list", Type: "menu", Sort: 3, Visible: true},
		{ID: 5, ParentID: 1, Title: "菜单管理", Path: "/system/menu", Component: "system/menu/index", Permission: "system:menu:list", Type: "menu", Sort: 4, Visible: true},
		{ID: 6, ParentID: 0, Title: "3D配置", Path: "/visual-config", Component: "visual-config/index", Permission: "visual:config:list", Type: "menu", Sort: 2, Visible: true},
	}
	for _, menu := range menus {
		if err := db.FirstOrCreate(&menu, model.Menu{ID: menu.ID}).Error; err != nil {
			return err
		}
	}

	if err := db.FirstOrCreate(&model.UserRole{UserID: user.ID, RoleID: role.ID}).Error; err != nil {
		return err
	}
	for _, menu := range menus {
		if err := db.FirstOrCreate(&model.RoleMenu{RoleID: role.ID, MenuID: menu.ID}).Error; err != nil {
			return err
		}
	}

	return nil
}
