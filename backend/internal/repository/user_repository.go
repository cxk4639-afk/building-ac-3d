package repository

import (
	"building-ac-3d/backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Order("id asc").Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByTenantAndUsername(tenantID uint, username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("tenant_id = ? AND username = ?", tenantID, username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateWithRoles(user *model.User, roleIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return replaceUserRoles(tx, user.ID, roleIDs)
	})
}

func (r *UserRepository) UpdateWithRoles(user *model.User, roleIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			return err
		}
		return replaceUserRoles(tx, user.ID, roleIDs)
	})
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", id).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.User{}, id).Error
	})
}

func (r *UserRepository) RoleIDs(userID uint) ([]uint, error) {
	var rows []model.UserRole
	if err := r.db.Where("user_id = ?", userID).Find(&rows).Error; err != nil {
		return nil, err
	}
	ids := make([]uint, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.RoleID)
	}
	return ids, nil
}

func (r *UserRepository) Roles(userID uint) ([]model.Role, error) {
	var roles []model.Role
	err := r.db.Table("roles").
		Select("roles.*").
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Order("roles.id asc").
		Find(&roles).Error
	return roles, err
}

func replaceUserRoles(tx *gorm.DB, userID uint, roleIDs []uint) error {
	if err := tx.Where("user_id = ?", userID).Delete(&model.UserRole{}).Error; err != nil {
		return err
	}
	for _, roleID := range roleIDs {
		if roleID == 0 {
			continue
		}
		if err := tx.Create(&model.UserRole{UserID: userID, RoleID: roleID}).Error; err != nil {
			return err
		}
	}
	return nil
}
