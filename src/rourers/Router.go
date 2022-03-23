package rourers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dimayaschhu/mysound/src/services/security"
	"github.com/gin-gonic/gin"
	"log"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter() Router {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	router := Router{engine: engine}

	return router
}

func (route Router) Engine() *gin.Engine {
	return route.engine
}

func (route *Router) SetRoutes() {
	route.engine.NoRoute(func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	route.setAuthGroup()
	route.setAdminGroup()
}

func (route *Router) setAuthGroup() {
	s := security.NewSecurityService()
	authMiddleware := s.CreateAuthMiddleware()

	route.engine.POST("/login", authMiddleware.LoginHandler)

	auth := route.engine.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	auth.GET("/hello", helloUserHandler)
}

func (route *Router) setAdminGroup() {
	s := security.NewSecurityAdminService()
	authMiddleware := s.CreateAuthMiddleware()

	route.engine.POST("/admin/login", authMiddleware.LoginHandler)

	auth := route.engine.Group("/admin")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	auth.GET("/hello", helloAdminHandler)
}
