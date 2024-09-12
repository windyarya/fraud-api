package middleware

import (
	"bitbucket.org/windyarya/backend-final/utils/token"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Middleware struct {
	DB *gorm.DB
}

type MiddlewareHandler func(echo.HandlerFunc) echo.HandlerFunc

func (m *Middleware) Authorisation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			response := map[string]interface{}{
				"message": "Authorisation header missing",
			}
			return c.JSON(http.StatusUnauthorized, response)
		}

		if !strings.HasPrefix(authHeader, "Bearer") {
			response := map[string]interface{}{
				"message": "Authorisation header invalid",
			}
			return c.JSON(http.StatusUnauthorized, response)
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == "" {
			response := map[string]interface{}{
				"message": "Authorisation token missing",
			}
			return c.JSON(http.StatusUnauthorized, response)
		}

		claims, err := token.VerifyToken(tokenStr)
		if err != nil {
			response := map[string]interface{}{
				"message": "Authorisation token invalid",
			}
			return c.JSON(http.StatusUnauthorized, response)
		}

		c.Set("role", claims.UserGroupID)
		c.Set("userID", claims.ID)

		return next(c)
	}
}

func (m *Middleware) RoleBased(next echo.HandlerFunc, requiredRole uint) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role").(uint)

		// RoleHierarchy := map[uint]string{
		// 	1: "Superadmin",
		// 	2: "Admin",
		// 	3: "Analyst",
		// 	4: "User",
		// }

		if role > requiredRole {
			response := map[string]interface{}{
				"message": "You don't have permission to access this resource",
			}
			return c.JSON(http.StatusForbidden, response)
		}

		// id := c.Param("id")
		// if id != "" {
		// 	access := m.AccessControl(id, c);
		// 	// fmt.Println(access)
		// 	if !access {
		// 		response := map[string]interface{}{
		// 			"message": "You don't have permission to access this resource",
		// 		}
		// 		return c.JSON(http.StatusForbidden, response)
		// 	}
		// }

		return next(c)
	}
}

func (m *Middleware) AccessControl(id string, c echo.Context) bool {
	userID := c.Get("userID").(uint)
	strID := strconv.Itoa(int(userID))
	// fmt.Println(strID, id)

	return strID == id
}
