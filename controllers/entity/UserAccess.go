package entity

import (
	"fmt"
	"tdez/database.go"
	"tdez/models"
	"tdez/requests"
	"tdez/resources"
	"tdez/utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var request requests.EntUsersLogin

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"errors": []string{err.Error()}})
		return
	}

	if err := utils.Valid(request); err != nil {

		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"Falha no login"}})
		c.Abort()
		return
	}

	db, err := database.SetupDB()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}
	tx := db.Begin()

	var user models.EntUsers

	query := tx.Where("use_email = ?", request.Email).Find(&user)
	if query.RowsAffected == 0 {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"Falha no login"}})
		c.Abort()
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	fmt.Println(err)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"E-mail/senha incorretos"}})
		c.Abort()
		return
	}

	token, err := resources.CreateTokenJWT(user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"Falha no login"}})
		c.Abort()
		return
	}

	spew.Dump("porcodio")

	user.Token = &token

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"Falha no login"}})
		c.Abort()
		return
	}

	tx.Commit()
	c.JSON(200, gin.H{"messages": user.Token})
	c.Abort()
	return
}
