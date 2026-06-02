package handler

import (
	"net/http"
	"strconv"

	"building-ac-3d/backend/internal/dto"

	"github.com/gin-gonic/gin"
)

func ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, dto.Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func fail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, dto.Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func queryID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil || id == 0 {
		fail(c, 400, "ID参数无效")
		return 0, false
	}
	return uint(id), true
}
