package controllers

import (
	"tdez/database.go"
	"tdez/models"
	"tdez/requests"
	"tdez/utils"

	"github.com/gin-gonic/gin"
)

func ExternalUserStore(c *gin.Context) {
	var request requests.EntUsersStore

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}

	if request.Cnpj == nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"O campo CNPJ é obrigatório ao cadastrar uma empresa parceira"}})
		c.Abort()
		return
	}

	if err := utils.Valid(request); err != nil {
		c.JSON(400, gin.H{"errors": err})
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

	request.Type = 1 //external user

	err = user.EntUsersFill(request)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})

		c.Abort()
		return
	}

	tx.Commit()

	c.JSON(200, gin.H{"messages": []string{"O usuário foi salvo com sucesso"}})
	c.Abort()
	return
}
