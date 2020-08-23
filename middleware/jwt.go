package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"tdez/models"

	"tdez/database.go"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Jwt ...
func Jwt(accessType int) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString != "" {
			bearerToken := strings.Split(tokenString, " ")

			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(os.Getenv("JWT_SECRET")), nil
				})

				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"authErrors": []string{"Invalid authorization token!!"}})
					c.Abort()
					return
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

					if claims["access"] == nil {
						c.JSON(http.StatusBadRequest, gin.H{"authErrors": []string{"Invalid payload"}})
						c.Abort()
						return
					}

					accessClaim := claims["access"].(int)
					if accessType != accessClaim {
						c.JSON(http.StatusBadRequest, gin.H{"authErrors": []string{"Access denied"}})
						c.Abort()
						return
					}

					db, er := database.SetupDB()
					if er != nil {
						c.JSON(400, gin.H{"authErrors": []string{"Error on database connection"}})
						c.Abort()
						return
					}
					defer db.Close()

					var user models.EntUsers
					query := db.Where("use_token = ?", bearerToken[1]).First(&user)
					if query.RowsAffected == 0 {
						c.JSON(400, gin.H{"authErrors": []string{"Invalid authorization token"}})
						c.Abort()
						return
					}

					c.Set("use_code", user.Code)

					return
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"authErrors": []string{err.Error()}})
					c.Abort()
					return
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"authErrors": []string{"Invalid authorization token"}})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"authErrors": []string{"An authorization header is required"}})
			c.Abort()
			return
		}
	}
}
