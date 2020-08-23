package product

import (
	"tdez/database.go"
	"tdez/models"
	"tdez/requests"
	"tdez/utils"

	"github.com/gin-gonic/gin"
)

func IssueActivation(c *gin.Context) {
	var request requests.ResProduct

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}

	if err := utils.Valid(request); err != nil {
		c.JSON(400, gin.H{"errors": err})
		return
	}

	useCode := c.MustGet("use_code").(int) //externalapp user code

	db, err := database.SetupDB()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}
	tx := db.Begin()

	var product models.ResProduct

	product.CodeIntUse = &useCode
	product.ResProductFill(request)

	if err := tx.Create(&product).Error; err != nil {
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
