package security

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dimayaschhu/mysound/src/entity"
	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func GetDefaultPayloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*entity.User); ok {
			return jwt.MapClaims{
				"id":        v.UserName,
				"firstName": v.FirstName,
				"lastName":  v.LastName,
			}
		}
		return jwt.MapClaims{}
	}
}

func GetDefaultIdentityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &entity.User{
			UserName:  claims["id"].(string),
			FirstName: claims["firstName"].(string),
			LastName:  claims["lastName"].(string),
		}
	}
}

func GetDefaultAuthenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := loginVals.Username
		password := loginVals.Password

		if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
			return &entity.User{
				UserName:  userID,
				LastName:  "Bo-Yi",
				FirstName: "Wu",
			}, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
}

func GetDefaultAuthorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if _, ok := data.(*entity.User); ok {
			return true
		}

		return false
	}
}

func GetDefaultUnauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}
