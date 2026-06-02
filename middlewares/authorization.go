package middlewares

import (
	"fmt"
	"net/http"

	"example.com/utils"
	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	token := c.GetHeader("authorization")

	fmt.Println(token)

	claims, err := utils.ParseJwt(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	c.Set("userId", claims.UserId)
}
