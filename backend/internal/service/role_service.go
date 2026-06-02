package service

import (
	"errors"

	"building-ac-3d/backend/internal/dto"
	"building-ac-3d/backend/internal/model"
	"building-ac-3d/backend/internal/repository"
)

type RoleService struct {
	repo *repository.RoleRepository
}

func NewRoleService(repo *repository.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) List() ([]dto.RoleDTO, error) {
	roles, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	result := make([]dto.RoleDTO, 0, len(roles))
	for _, role := range roles {
		menuIDs, err := s.repo.MenuIDs(role.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, roleToDTO(role, menuIDs))
	}
	return result, nil
}

func (s *RoleService) Create(req dto.RoleDTO) (uint, error) {
	if req.TenantID == 0 || req.Name == "" || req.Code == "" {
		return 0, errors.New("租户、角色名称和角色编码不能为空")
	}
	status := req.Status
	if status == 0 {
		status = 1
	}
	role := model.Role{
		TenantID: req.TenantID,
		Name:     req.Name,
		Code:     req.Code,
		Status:   status,
	}
	if err := s.repo.CreateWithMenus(&role, req.MenuIDs); err != nil {
		return 0, err
	}
	return role.ID, nil
}

func (s *RoleService) Update(req dto.RoleDTO) error {
	if req.ID == 0 {
		return errors.New("角色ID不能为空")
	}
	if req.TenantID == 0 || req.Name == "" || req.Code == "" {
		return errors.New("租户、角色名称和角色编码不能为空")
	}
	role, err := s.repo.FindByID(req.ID)
	if err != nil {
		return err
	}
	role.TenantID = req.TenantID
	role.Name = req.Name
	role.Code = req.Code
	role.Status = req.Status
	if role.Status == 0 {
		role.Status = 1
	}
	return s.repo.UpdateWithMenus(role, req.MenuIDs)
}

func (s *RoleService) Delete(id uint) error {
	if id == 0 {
		return errors.New("角色ID不能为空")
	}
	return s.repo.Delete(id)
}

func roleToDTO(role model.Role, menuIDs []uint) dto.RoleDTO {
	return dto.RoleDTO{
		ID:        role.ID,
		TenantID:  role.TenantID,
		Name:      role.Name,
		Code:      role.Code,
		MenuIDs:   menuIDs,
		Status:    role.Status,
		CreatedAt: formatTime(role.CreatedAt),
	}
}
