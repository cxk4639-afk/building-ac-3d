package service

import (
	"errors"
	"time"

	"building-ac-3d/backend/internal/dto"
	"building-ac-3d/backend/internal/model"
	"building-ac-3d/backend/internal/repository"
)

type TenantService struct {
	repo *repository.TenantRepository
}

func NewTenantService(repo *repository.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (s *TenantService) List() ([]dto.TenantDTO, error) {
	tenants, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	result := make([]dto.TenantDTO, 0, len(tenants))
	for _, tenant := range tenants {
		result = append(result, tenantToDTO(tenant))
	}
	return result, nil
}

func (s *TenantService) Create(req dto.TenantDTO) (uint, error) {
	if req.Name == "" || req.Code == "" {
		return 0, errors.New("租户名称和编码不能为空")
	}
	status := req.Status
	if status == 0 {
		status = 1
	}
	tenant := model.Tenant{
		Name:   req.Name,
		Code:   req.Code,
		Status: status,
		Remark: req.Remark,
	}
	if err := s.repo.Create(&tenant); err != nil {
		return 0, err
	}
	return tenant.ID, nil
}

func (s *TenantService) Update(req dto.TenantDTO) error {
	if req.ID == 0 {
		return errors.New("租户ID不能为空")
	}
	if req.Name == "" || req.Code == "" {
		return errors.New("租户名称和编码不能为空")
	}
	tenant, err := s.repo.FindByID(req.ID)
	if err != nil {
		return err
	}
	tenant.Name = req.Name
	tenant.Code = req.Code
	tenant.Status = req.Status
	tenant.Remark = req.Remark
	if tenant.Status == 0 {
		tenant.Status = 1
	}
	return s.repo.Update(tenant)
}

func (s *TenantService) Delete(id uint) error {
	if id == 0 {
		return errors.New("租户ID不能为空")
	}
	return s.repo.Delete(id)
}

func tenantToDTO(tenant model.Tenant) dto.TenantDTO {
	return dto.TenantDTO{
		ID:        tenant.ID,
		Name:      tenant.Name,
		Code:      tenant.Code,
		Status:    tenant.Status,
		Remark:    tenant.Remark,
		CreatedAt: formatTime(tenant.CreatedAt),
	}
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}
