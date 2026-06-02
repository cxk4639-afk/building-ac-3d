package service

import (
	"errors"

	"building-ac-3d/backend/internal/dto"
	"building-ac-3d/backend/internal/model"
	"building-ac-3d/backend/internal/repository"
)

type MenuService struct {
	repo *repository.MenuRepository
}

func NewMenuService(repo *repository.MenuRepository) *MenuService {
	return &MenuService{repo: repo}
}

func (s *MenuService) List() ([]dto.MenuDTO, error) {
	menus, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	result := make([]dto.MenuDTO, 0, len(menus))
	for _, menu := range menus {
		result = append(result, menuToDTO(menu))
	}
	return result, nil
}

func (s *MenuService) Create(req dto.MenuDTO) (uint, error) {
	if req.Title == "" || req.Type == "" {
		return 0, errors.New("菜单名称和类型不能为空")
	}
	menu := model.Menu{
		ParentID:   req.ParentID,
		Title:      req.Title,
		Path:       req.Path,
		Component:  req.Component,
		Permission: req.Permission,
		Type:       req.Type,
		Sort:       req.Sort,
		Visible:    req.Visible,
	}
	if err := s.repo.Create(&menu); err != nil {
		return 0, err
	}
	return menu.ID, nil
}

func (s *MenuService) Update(req dto.MenuDTO) error {
	if req.ID == 0 {
		return errors.New("菜单ID不能为空")
	}
	if req.Title == "" || req.Type == "" {
		return errors.New("菜单名称和类型不能为空")
	}
	menu, err := s.repo.FindByID(req.ID)
	if err != nil {
		return err
	}
	menu.ParentID = req.ParentID
	menu.Title = req.Title
	menu.Path = req.Path
	menu.Component = req.Component
	menu.Permission = req.Permission
	menu.Type = req.Type
	menu.Sort = req.Sort
	menu.Visible = req.Visible
	return s.repo.Update(menu)
}

func (s *MenuService) Delete(id uint) error {
	if id == 0 {
		return errors.New("菜单ID不能为空")
	}
	return s.repo.Delete(id)
}

func menuToDTO(menu model.Menu) dto.MenuDTO {
	return dto.MenuDTO{
		ID:         menu.ID,
		ParentID:   menu.ParentID,
		Title:      menu.Title,
		Path:       menu.Path,
		Component:  menu.Component,
		Permission: menu.Permission,
		Type:       menu.Type,
		Sort:       menu.Sort,
		Visible:    menu.Visible,
	}
}
