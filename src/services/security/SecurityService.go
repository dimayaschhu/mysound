package security

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type SecurityService struct {
}

func NewSecurityService() SecurityService {
	return SecurityService{}
}

func (s SecurityService) CreateAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     "id",
		PayloadFunc:     s.GetPayloadFunc(),
		IdentityHandler: s.GetIdentityHandler(),
		Authenticator:   s.GetAuthenticator(),
		Authorizator:    s.GetAuthorizator(),
		Unauthorized:    s.GetUnauthorized(),
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	if err := authMiddleware.MiddlewareInit(); err != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + err.Error())
	}

	return authMiddleware
}

func (s SecurityService) GetPayloadFunc() func(data interface{}) jwt.MapClaims {
	return GetDefaultPayloadFunc()
}

func (s SecurityService) GetIdentityHandler() func(c *gin.Context) interface{} {
	return GetDefaultIdentityHandler()
}

func (s SecurityService) GetAuthenticator() func(c *gin.Context) (interface{}, error) {
	return GetDefaultAuthenticator()
}

func (s SecurityService) GetAuthorizator() func(data interface{}, c *gin.Context) bool {
	return GetDefaultAuthorizator()
}

func (s SecurityService) GetUnauthorized() func(c *gin.Context, code int, message string) {
	return GetDefaultUnauthorized()
}
