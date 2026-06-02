package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthUserIDKey   = "auth_user_id"
	AuthTenantIDKey = "auth_tenant_id"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if strings.HasPrefix(token, "dev-token-") {
			parts := strings.Split(token, "-")
			if len(parts) >= 5 {
				userID, userErr := strconv.ParseUint(parts[2], 10, 64)
				tenantID, tenantErr := strconv.ParseUint(parts[3], 10, 64)
				if userErr == nil && tenantErr == nil {
					c.Set(AuthUserIDKey, uint(userID))
					c.Set(AuthTenantIDKey, uint(tenantID))
				}
			}
		}
		c.Next()
	}
}
