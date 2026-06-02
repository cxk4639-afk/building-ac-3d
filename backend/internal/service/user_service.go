package service

import (
	"errors"

	"building-ac-3d/backend/internal/dto"
	"building-ac-3d/backend/internal/model"
	"building-ac-3d/backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) List() ([]dto.UserDTO, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	result := make([]dto.UserDTO, 0, len(users))
	for _, user := range users {
		roleIDs, err := s.repo.RoleIDs(user.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, userToDTO(user, roleIDs))
	}
	return result, nil
}

func (s *UserService) Create(req dto.UserDTO) (uint, error) {
	if req.TenantID == 0 || req.Username == "" {
		return 0, errors.New("租户和用户名不能为空")
	}
	password := req.Password
	if password == "" {
		password = "123456"
	}
	status := req.Status
	if status == 0 {
		status = 1
	}
	user := model.User{
		TenantID: req.TenantID,
		Username: req.Username,
		Password: password,
		Nickname: req.Nickname,
		Status:   status,
	}
	if err := s.repo.CreateWithRoles(&user, req.RoleIDs); err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (s *UserService) Update(req dto.UserDTO) error {
	if req.ID == 0 {
		return errors.New("用户ID不能为空")
	}
	if req.TenantID == 0 || req.Username == "" {
		return errors.New("租户和用户名不能为空")
	}
	user, err := s.repo.FindByID(req.ID)
	if err != nil {
		return err
	}
	user.TenantID = req.TenantID
	user.Username = req.Username
	if req.Password != "" {
		user.Password = req.Password
	}
	user.Nickname = req.Nickname
	user.Status = req.Status
	if user.Status == 0 {
		user.Status = 1
	}
	return s.repo.UpdateWithRoles(user, req.RoleIDs)
}

func (s *UserService) Delete(id uint) error {
	if id == 0 {
		return errors.New("用户ID不能为空")
	}
	return s.repo.Delete(id)
}

func userToDTO(user model.User, roleIDs []uint) dto.UserDTO {
	return dto.UserDTO{
		ID:        user.ID,
		TenantID:  user.TenantID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		RoleIDs:   roleIDs,
		Status:    user.Status,
		CreatedAt: formatTime(user.CreatedAt),
	}
}
