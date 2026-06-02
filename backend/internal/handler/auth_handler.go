package handler

import (
	"strings"

	"building-ac-3d/backend/internal/dto"
	"building-ac-3d/backend/internal/middleware"
	"building-ac-3d/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, 400, "请求参数无效")
		return
	}
	result, err := h.service.Login(req)
	if err != nil {
		fail(c, 400, err.Error())
		return
	}
	ok(c, result)
}

func (h *AuthHandler) Me(c *gin.Context) {
	if strings.TrimSpace(c.GetHeader("Authorization")) == "" {
		fail(c, 400, "未登录")
		return
	}
	userID, _ := c.Get(middleware.AuthUserIDKey)
	id, _ := userID.(uint)
	if id == 0 {
		fail(c, 400, "未登录")
		return
	}
	result, err := h.service.Me(id)
	if err != nil {
		fail(c, 400, err.Error())
		return
	}
	ok(c, result)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	ok(c, nil)
}
