package service

import (
	"errors"
	"fmt"
	"time"

	"building-ac-3d/backend/internal/dto"
	"building-ac-3d/backend/internal/repository"
)

type AuthService struct {
	tenantRepo *repository.TenantRepository
	userRepo   *repository.UserRepository
}

func NewAuthService(tenantRepo *repository.TenantRepository, userRepo *repository.UserRepository) *AuthService {
	return &AuthService{tenantRepo: tenantRepo, userRepo: userRepo}
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	if req.TenantCode == "" {
		req.TenantCode = "default"
	}
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名和密码不能为空")
	}

	tenant, err := s.tenantRepo.FindByCode(req.TenantCode)
	if err != nil {
		return nil, errors.New("租户不存在")
	}
	if tenant.Status != 1 {
		return nil, errors.New("租户已停用")
	}

	user, err := s.userRepo.FindByTenantAndUsername(tenant.ID, req.Username)
	if err != nil {
		return nil, errors.New("用户不存在或密码错误")
	}
	if user.Status != 1 || user.Password != req.Password {
		return nil, errors.New("用户不存在或密码错误")
	}

	return &dto.LoginResponse{
		Token:        buildDevToken(user.ID, tenant.ID),
		RefreshToken: fmt.Sprintf("dev-refresh-%d-%d-%d", user.ID, tenant.ID, time.Now().Unix()),
	}, nil
}

func (s *AuthService) Me(userID uint) (*dto.UserInfo, error) {
	if userID == 0 {
		userID = 1
	}

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("登录状态无效")
	}
	tenant, err := s.tenantRepo.FindByID(user.TenantID)
	if err != nil {
		return nil, errors.New("租户不存在")
	}
	roles, err := s.userRepo.Roles(user.ID)
	if err != nil {
		return nil, err
	}
	roleCodes := make([]string, 0, len(roles))
	for _, role := range roles {
		roleCodes = append(roleCodes, role.Code)
	}

	return &dto.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Tenant: dto.TenantInfo{
			ID:   tenant.ID,
			Name: tenant.Name,
			Code: tenant.Code,
		},
		Roles:       roleCodes,
		Permissions: []string{"*:*:*"},
	}, nil
}

func buildDevToken(userID uint, tenantID uint) string {
	return fmt.Sprintf("dev-token-%d-%d-%d", userID, tenantID, time.Now().Unix())
}
