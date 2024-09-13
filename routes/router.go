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
	account := controllers.AccountHandler{DB: db}
	activity := controllers.ActivityHandler{DB: db}
	alert := controllers.AlertHandler{DB: db}

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
	AccountRoutes(e, account, mw)
	ActivityRoutes(e, activity, mw)
	AlertRoutes(e, alert, mw)
}

func UserRoutes(e *echo.Echo, user controllers.UserHandler, mw middleware.Middleware) {
	u := e.Group("/api/v1")

	// u.Use(echojwt.WithConfig(config))

	u.GET("/users", mw.Authorisation(mw.RoleBased(user.GetUsers, uint(1))))
	u.GET("/users/:id", mw.Authorisation(mw.RoleBased(user.GetUser, uint(4))))
	u.PUT("/users/:id", mw.Authorisation(mw.RoleBased(user.Update, uint(4))))
	u.DELETE("/users/:id", mw.Authorisation(mw.RoleBased(user.Delete, uint(1))))
	// u.POST("/logout", user.Logout)
}

func AuthRoutes(e *echo.Echo, user controllers.UserHandler) {
	u := e.Group("/api/v1")

	u.POST("/login", user.Login)
	u.POST("/register", user.Register)
}

func UnitRoutes(e *echo.Echo, unit controllers.UnitHandler, mw middleware.Middleware) {
	u := e.Group("/api/v1")


	u.GET("/units", mw.Authorisation(mw.RoleBased(unit.GetAll, uint(4))))
	u.GET("/units/:id", mw.Authorisation(mw.RoleBased(unit.GetByID, uint(4))))
	u.POST("/units", mw.Authorisation(mw.RoleBased(unit.Create, uint(1))))
	u.PUT("/units/:id", mw.Authorisation(mw.RoleBased(unit.Update, uint(1))))
	u.DELETE("/units/:id", mw.Authorisation(mw.RoleBased(unit.Delete, uint(1))))
}

func GroupRoutes(e *echo.Echo, group controllers.GroupHandler, mw middleware.Middleware) {
	u := e.Group("/api/v1")

	u.GET("/groups", mw.Authorisation(mw.RoleBased(group.GetAll, uint(1))))
	u.GET("/groups/:id", mw.Authorisation(mw.RoleBased(group.GetByID, uint(4))))
	u.POST("/groups", mw.Authorisation(mw.RoleBased(group.Create, uint(1))))
	u.PUT("/groups/:id", mw.Authorisation(mw.RoleBased(group.Update, uint(1))))
	u.DELETE("/groups/:id", mw.Authorisation(mw.RoleBased(group.Delete, uint(1))))
}

func AccountRoutes(e *echo.Echo, account controllers.AccountHandler, mw middleware.Middleware) {
	u := e.Group("/api/v1")

	u.GET("/accounts", mw.Authorisation(mw.RoleBased(account.GetAll, uint(2))))
	u.GET("/accounts/:id", mw.Authorisation(mw.RoleBased(account.GetByID, uint(4))))
	u.POST("/accounts", mw.Authorisation(mw.RoleBased(account.Create, uint(1))))
	u.PUT("/accounts/:id", mw.Authorisation(mw.RoleBased(account.Update, uint(1))))
	u.DELETE("/accounts/:id", mw.Authorisation(mw.RoleBased(account.Delete, uint(1))))
}

func ActivityRoutes(e *echo.Echo, activity controllers.ActivityHandler, mw middleware.Middleware) {
	u := e.Group("/api/v1")

	u.GET("/activities", mw.Authorisation(mw.RoleBased(activity.GetAll, uint(4))))
	u.GET("/activities/:id", mw.Authorisation(mw.RoleBased(activity.GetByID, uint(4))))
	u.POST("/activities", mw.Authorisation(mw.RoleBased(activity.Create, uint(4))))
}

func AlertRoutes(e *echo.Echo, alert controllers.AlertHandler, mw middleware.Middleware) {
	u := e.Group("/api/v1")

	u.GET("/alerts", mw.Authorisation(mw.RoleBased(alert.GetAll, uint(4))))
	u.GET("/alerts/:id", mw.Authorisation(mw.RoleBased(alert.GetByID, uint(3))))
	u.POST("/alerts", mw.Authorisation(mw.RoleBased(alert.Create, uint(4))))
	u.PUT("/alerts/:id", mw.Authorisation(mw.RoleBased(alert.Update, uint(3))))
	u.DELETE("/alerts/:id", mw.Authorisation(mw.RoleBased(alert.Delete, uint(2))))
}