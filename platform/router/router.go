package router

import (
	"encoding/gob"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"01-Login/platform/authenticator"
	"01-Login/platform/database"
	"01-Login/platform/middleware"
	"01-Login/web/app/callback"
	"01-Login/web/app/home"
	"01-Login/web/app/login"
	"01-Login/web/app/logout"
	"01-Login/web/app/user"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator, db *database.Database) *gin.Engine {

	router := gin.Default()
	home.SetDatabase(db)
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
		// ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", middleware.IsAuthenticated, home.Handler)
	router.POST("/create", home.Create)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	// router.POST("/create", middleware.IsAuthenticated, home.Create)
	// router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
