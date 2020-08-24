package entity

import (
	"tdez/database"
	"tdez/models"
	"tdez/requests"
	"tdez/utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

// SuperUserStore godoc
// @Tags User
// @Summary Create a superuser
// @Description Create a super user, no athentication
// @Accept json
// @Produce json
// @Param Request body requests.EntUsersStore true "Request body"
// @Success 200 {object} models.DefaultSuccess
// @Failure 400 {object} models.DefaultError
// @Router /superuser [post]
func SuperUserStore(c *gin.Context) {
	var request requests.EntSuperUsersStore

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": []string{err.Error()}})
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

	request.Type = 0 //superuser

	pass, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(pass)
	request.ConfirmPassword = nil

	err = user.SuperUserFill(request)
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
