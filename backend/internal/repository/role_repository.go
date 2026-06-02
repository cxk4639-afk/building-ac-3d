package repository

import (
	"building-ac-3d/backend/internal/model"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) FindAll() ([]model.Role, error) {
	var roles []model.Role
	err := r.db.Order("id asc").Find(&roles).Error
	return roles, err
}

func (r *RoleRepository) FindByID(id uint) (*model.Role, error) {
	var role model.Role
	if err := r.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) CreateWithMenus(role *model.Role, menuIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(role).Error; err != nil {
			return err
		}
		return replaceRoleMenus(tx, role.ID, menuIDs)
	})
}

func (r *RoleRepository) UpdateWithMenus(role *model.Role, menuIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(role).Error; err != nil {
			return err
		}
		return replaceRoleMenus(tx, role.ID, menuIDs)
	})
}

func (r *RoleRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id = ?", id).Delete(&model.RoleMenu{}).Error; err != nil {
			return err
		}
		if err := tx.Where("role_id = ?", id).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.Role{}, id).Error
	})
}

func (r *RoleRepository) MenuIDs(roleID uint) ([]uint, error) {
	var rows []model.RoleMenu
	if err := r.db.Where("role_id = ?", roleID).Find(&rows).Error; err != nil {
		return nil, err
	}
	ids := make([]uint, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.MenuID)
	}
	return ids, nil
}

func replaceRoleMenus(tx *gorm.DB, roleID uint, menuIDs []uint) error {
	if err := tx.Where("role_id = ?", roleID).Delete(&model.RoleMenu{}).Error; err != nil {
		return err
	}
	for _, menuID := range menuIDs {
		if menuID == 0 {
			continue
		}
		if err := tx.Create(&model.RoleMenu{RoleID: roleID, MenuID: menuID}).Error; err != nil {
			return err
		}
	}
	return nil
}
