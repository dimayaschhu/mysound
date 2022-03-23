package security

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dimayaschhu/mysound/src/entity"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type SecurityAdminService struct {
}

func NewSecurityAdminService() SecurityAdminService {
	return SecurityAdminService{}
}

func (s SecurityAdminService) CreateAuthMiddleware() *jwt.GinJWTMiddleware {
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

func (s SecurityAdminService) GetAuthorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*entity.User); ok && v.UserName == "admin" {
			return true
		}

		return false
	}
}

func (s SecurityAdminService) GetPayloadFunc() func(data interface{}) jwt.MapClaims {
	return GetDefaultPayloadFunc()
}

func (s SecurityAdminService) GetIdentityHandler() func(c *gin.Context) interface{} {
	return GetDefaultIdentityHandler()
}

func (s SecurityAdminService) GetAuthenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := loginVals.Username
		password := loginVals.Password

		if userID == "admin" && password == "admin" {
			return &entity.User{
				UserName:  userID,
				LastName:  "Bo-Yi",
				FirstName: "Wu",
			}, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
}

func (s SecurityAdminService) GetUnauthorized() func(c *gin.Context, code int, message string) {
	return GetDefaultUnauthorized()
}
