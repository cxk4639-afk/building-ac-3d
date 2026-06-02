package repository

import (
	"building-ac-3d/backend/internal/model"

	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) DB() *gorm.DB {
	return r.db
}

func (r *TenantRepository) FindAll() ([]model.Tenant, error) {
	var tenants []model.Tenant
	err := r.db.Order("id asc").Find(&tenants).Error
	return tenants, err
}

func (r *TenantRepository) FindByID(id uint) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := r.db.First(&tenant, id).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *TenantRepository) FindByCode(code string) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := r.db.Where("code = ?", code).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *TenantRepository) Create(tenant *model.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *TenantRepository) Update(tenant *model.Tenant) error {
	return r.db.Save(tenant).Error
}

func (r *TenantRepository) Delete(id uint) error {
	return r.db.Delete(&model.Tenant{}, id).Error
}
