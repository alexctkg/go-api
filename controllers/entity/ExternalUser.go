package entity

import (
	"tdez/database"
	"tdez/models"
	"tdez/requests"
	"tdez/utils"

	"golang.org/x/crypto/bcrypt"

	cnpjValid "github.com/Nhanderu/brdoc"
	"github.com/gin-gonic/gin"
)

// ExternalUserStore godoc
// @Tags User
// @Summary Create a exteraluser
// @Description Create a external user, no athentication
// @Accept json
// @Produce json
// @Param Request body requests.EntExternalUserStore true "Request body"
// @Success 200 {object} models.DefaultSuccess
// @Failure 400 {object} models.DefaultError
// @Router /externalapp [post]
func ExternalUserStore(c *gin.Context) {
	var request requests.EntExternalUserStore

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
		c.Abort()
		return
	}

	if err := utils.Valid(request); err != nil {
		c.JSON(400, gin.H{"errors": err})
		return
	}

	if request.Cnpj == nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"O campo CNPJ é obrigatório ao cadastrar uma empresa parceira"}})
		c.Abort()
		return
	}

	cnpjValidate := cnpjValid.IsCNPJ(*request.Cnpj)

	if cnpjValidate == false {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{"CNPJ inválido"}})
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

	request.Type = 1 //external user

	pass, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(pass)
	request.ConfirmPassword = nil

	err = user.ExternaUserFill(request)
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
