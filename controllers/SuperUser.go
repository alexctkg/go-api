package controllers

import (
	"tdez/database.go"
	"tdez/models"
	"tdez/requests"

	"github.com/gin-gonic/gin"
)

func SuperUserStore(c *gin.Context) {
	var request requests.EntUsers

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
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

	request.Type = 0 //superuser

	if err := user.EntUsersFill(request).Error; err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}

	if err := tx.Save(&user).Error; err != nil {
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
