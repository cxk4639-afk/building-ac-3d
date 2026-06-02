package repository

import (
	"building-ac-3d/backend/internal/model"

	"gorm.io/gorm"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) FindAll() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Order("sort asc, id asc").Find(&menus).Error
	return menus, err
}

func (r *MenuRepository) FindByID(id uint) (*model.Menu, error) {
	var menu model.Menu
	if err := r.db.First(&menu, id).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *MenuRepository) Create(menu *model.Menu) error {
	return r.db.Create(menu).Error
}

func (r *MenuRepository) Update(menu *model.Menu) error {
	return r.db.Save(menu).Error
}

func (r *MenuRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("menu_id = ?", id).Delete(&model.RoleMenu{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.Menu{}, id).Error
	})
}
