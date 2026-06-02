package handler

import (
	"building-ac-3d/backend/internal/dto"
	"building-ac-3d/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	service *service.RoleService
}

func NewRoleHandler(service *service.RoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

func (h *RoleHandler) List(c *gin.Context) {
	data, err := h.service.List()
	if err != nil {
		fail(c, 400, err.Error())
		return
	}
	ok(c, data)
}

func (h *RoleHandler) Create(c *gin.Context) {
	var req dto.RoleDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, 400, "请求参数无效")
		return
	}
	id, err := h.service.Create(req)
	if err != nil {
		fail(c, 400, err.Error())
		return
	}
	ok(c, map[string]uint{"id": id})
}

func (h *RoleHandler) Update(c *gin.Context) {
	var req dto.RoleDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, 400, "请求参数无效")
		return
	}
	if err := h.service.Update(req); err != nil {
		fail(c, 400, err.Error())
		return
	}
	ok(c, true)
}

func (h *RoleHandler) Delete(c *gin.Context) {
	id, valid := queryID(c)
	if !valid {
		return
	}
	if err := h.service.Delete(id); err != nil {
		fail(c, 400, err.Error())
		return
	}
	ok(c, true)
}
