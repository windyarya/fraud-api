package routes

import (
	"bitbucket.org/windyarya/backend-final/controllers"
	"bitbucket.org/windyarya/backend-final/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Router(e *echo.Echo, db *gorm.DB) {
	user := controllers.UserHandler{DB: db}
	unit := controllers.UnitHandler{DB: db}
	group := controllers.GroupHandler{DB: db}

	// config := echojwt.Config{
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(token.JwtCustomClaims)
	// 	},
	// 	SigningKey: []byte(os.Getenv("TOKEN_SALT")),
	// }

	mw := middleware.Middleware{DB: db}

	AuthRoutes(e, user)
	UserRoutes(e, user, mw)
	UnitRoutes(e, unit, mw)
	GroupRoutes(e, group, mw)
}

func UserRoutes(e *echo.Echo, user controllers.UserHandler, mw middleware.Middleware) {
	u := e.Group("/api/v1")

	// u.Use(echojwt.WithConfig(config))

	u.GET("/users", mw.Authorisation(mw.RoleBased(user.GetUsers, uint(3))))
	u.GET("/users/:id", mw.Authorisation(mw.RoleBased(user.GetUser, uint(6))))
	u.PUT("/users/:id", mw.Authorisation(mw.RoleBased(user.Update, uint(6))))
	u.DELETE("/users/:id", mw.Authorisation(mw.RoleBased(user.Delete, uint(3))))
	// u.POST("/logout", user.Logout)
}

func AuthRoutes(e *echo.Echo, user controllers.UserHandler) {
	u := e.Group("/api/v1")

	u.POST("/login", user.Login)
	u.POST("/register", user.Register)
}

func UnitRoutes(e *echo.Echo, unit controllers.UnitHandler, mw middleware.Middleware) {
	u := e.Group("/api/v1")


	u.GET("/units", mw.Authorisation(mw.RoleBased(unit.GetAll, uint(6))))
	u.GET("/units/:id", mw.Authorisation(mw.RoleBased(unit.GetByID, uint(6))))
	u.POST("/units", mw.Authorisation(mw.RoleBased(unit.Create, uint(3))))
	u.PUT("/units/:id", mw.Authorisation(mw.RoleBased(unit.Update, uint(3))))
	u.DELETE("/units/:id", mw.Authorisation(mw.RoleBased(unit.Delete, uint(3))))
}

func GroupRoutes(e *echo.Echo, group controllers.GroupHandler, mw middleware.Middleware) {
	u := e.Group("/api/v1")

	u.GET("/groups", mw.Authorisation(mw.RoleBased(group.GetAll, uint(3))))
	u.GET("/groups/:id", mw.Authorisation(mw.RoleBased(group.GetByID, uint(6))))
	u.POST("/groups", mw.Authorisation(mw.RoleBased(group.Create, uint(3))))
	u.PUT("/groups/:id", mw.Authorisation(mw.RoleBased(group.Update, uint(3))))
	u.DELETE("/groups/:id", mw.Authorisation(mw.RoleBased(group.Delete, uint(3))))
}