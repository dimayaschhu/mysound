package rourers

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dimayaschhu/mysound/src/entity"
	"github.com/gin-gonic/gin"
)

func helloAdminHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	res, _ := c.Get("id")
	fmt.Println(res)
	user := res.(*entity.User)
	c.JSON(200, gin.H{
		"userID":    claims["id"],
		"userName":  user.UserName,
		"firstName": user.FirstName,
		"LastName":  user.LastName,
		"text":      "Hello Admin.",
	})
}
